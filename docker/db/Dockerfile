FROM mysql:8.0.32

ARG MYSQL_DIR=/etc/mysql
ARG MYSQL_LOG_DIR=/var/log/mysql

RUN mkdir -p $MYSQL_LOG_DIR \
    && touch $MYSQL_LOG_DIR/mysql.log \
    && chown -R mysql:mysql $MYSQL_LOG_DIR

COPY my.cnf $MYSQL_DIR/conf.d/
