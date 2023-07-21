# docker buildx build --platform linux/amd64 -t mysqlshell:8.0.29 . --load

FROM oraclelinux:7-slim

ARG KEY=https://repo.mysql.com/RPM-GPG-KEY-mysql
ARG MYSQL_SHELL_PACKAGE_URL="https://dev.mysql.com/get/mysql80-community-release-el7-8.noarch.rpm"

COPY run.sh /run.sh
RUN chmod +x /run.sh

COPY setupCluster.js /setupCluster.js

RUN rpmkeys --import $KEY \
    && yum install -y $MYSQL_SHELL_PACKAGE_URL \
    && yum clean all

ENTRYPOINT ["/run.sh"]
CMD ["mysqlsh"]
