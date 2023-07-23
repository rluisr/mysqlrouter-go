FROM oraclelinux:7-slim

ARG MYSQL_SHELL_PACKAGER_URL

COPY run.sh /run.sh
RUN chmod +x /run.sh

COPY setupCluster.js /setupCluster.js

RUN yum install -y $MYSQL_SHELL_PACKAGER_URL \
    && yum clean all

ENTRYPOINT ["/run.sh"]
CMD ["mysqlsh"]
