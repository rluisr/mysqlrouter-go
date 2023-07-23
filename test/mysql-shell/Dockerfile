FROM oraclelinux:7-slim

ARG MYSQL_SHELL_PACKAGER_URL="https://dev.mysql.com/get/Downloads/MySQL-Shell/mysql-shell-8.1.0-1.el7.x86_64.rpm"

COPY run.sh /run.sh
RUN chmod +x /run.sh

COPY setupCluster.js /setupCluster.js

RUN yum install -y $MYSQL_SHELL_PACKAGER_URL \
    && yum clean all

ENTRYPOINT ["/run.sh"]
CMD ["mysqlsh"]
