#!/bin/bash
set -e

if [ "$1" = 'mysqlrouter' ]; then
    if [[ -z $MYSQL_HOST || -z $MYSQL_PORT || -z $MYSQL_USER || -z $MYSQL_PASSWORD ]]; then
	    echo "We require all of"
	    echo "    MYSQL_HOST"
	    echo "    MYSQL_PORT"
	    echo "    MYSQL_USER"
	    echo "    MYSQL_PASSWORD"
	    echo "to be set. Exiting."
	    exit 1
    fi

    PASSFILE=$(mktemp)
    echo "$MYSQL_PASSWORD" > "$PASSFILE"
    DEFAULTS_EXTRA_FILE=$(mktemp)
    cat >"$DEFAULTS_EXTRA_FILE" <<EOF
[client]
password="$MYSQL_PASSWORD"
EOF
    unset MYSQL_PASSWORD
    max_tries=12
    attempt_num=0
    until (echo > "/dev/tcp/$MYSQL_HOST/$MYSQL_PORT") >/dev/null 2>&1; do
      echo "Waiting for mysql server $MYSQL_HOST ($attempt_num/$max_tries)"
      sleep $(( attempt_num++ ))
      if (( attempt_num == max_tries )); then
        exit 1
      fi
    done
    echo "Successfully contacted mysql server at $MYSQL_HOST. Checking for cluster state."

    mysql --defaults-extra-file="$DEFAULTS_EXTRA_FILE" -u "$MYSQL_USER" -h "$MYSQL_HOST" -P "$MYSQL_PORT" -e "show status;"

    if ! [[ "$(mysql --defaults-extra-file="$DEFAULTS_EXTRA_FILE" -u "$MYSQL_USER" -h "$MYSQL_HOST" -P "$MYSQL_PORT" -e "show status;" 2> /dev/null)" ]]; then
	    echo "Can not connect to database. Exiting."
	    exit 1
    fi
    if [[ -n $MYSQL_INNODB_CLUSTER_MEMBERS ]]; then
      attempt_num=0
      echo $attempt_num
      echo $max_tries
      until [ "$(mysql --defaults-extra-file="$DEFAULTS_EXTRA_FILE" -u "$MYSQL_USER" -h "$MYSQL_HOST" -P "$MYSQL_PORT" -N performance_schema -e "select count(MEMBER_STATE) = $MYSQL_INNODB_CLUSTER_MEMBERS from replication_group_members where MEMBER_STATE = 'ONLINE';" 2> /dev/null)" -eq 1 ]; do
             echo "Waiting for $MYSQL_INNODB_CLUSTER_MEMBERS cluster instances to become available via $MYSQL_HOST ($attempt_num/$max_tries)"
             sleep $(( attempt_num++ ))
             if (( attempt_num == max_tries )); then
                     exit 1
             fi
      done
      echo "Successfully contacted cluster with $MYSQL_INNODB_CLUSTER_MEMBERS members. Bootstrapping."
    fi
    echo "Successfully contacted mysql server at $MYSQL_HOST. Trying to bootstrap."
    echo "Hello! I'm tired for creating this fuck containers!"

    chown mysqlrouter /mysqlrouter.pwd

    mysqlrouter --bootstrap "$MYSQL_USER@$MYSQL_HOST:$MYSQL_PORT" --user=mysqlrouter --directory /tmp/mysqlrouter --force < "$PASSFILE"
    sed -i -e 's/logging_folder=.*$/logging_folder=/' /tmp/mysqlrouter/mysqlrouter.conf
    # Remove sections related to rest api
    sed -i -e '60,83d' /tmp/mysqlrouter/mysqlrouter.conf
    # then add
    cat << EOF >> /tmp/mysqlrouter/mysqlrouter.conf
[http_server]
port=8080

[rest_api]

[rest_router]
require_realm=somerealm

[rest_routing]
require_realm=somerealm

[rest_metadata_cache]
require_realm=somerealm

[http_auth_realm:somerealm]
backend=somebackend
method=basic
name=Some Realm

[http_auth_backend:somebackend]
backend=file
filename=/mysqlrouter.pwd
EOF
    echo "Starting mysql-router."

    exec "$@" --config /tmp/mysqlrouter/mysqlrouter.conf
fi

rm -f "$PASSFILE"
rm -f "$DEFAULTS_EXTRA_FILE"
unset DEFAULTS_EXTRA_FILE

exec "$@"