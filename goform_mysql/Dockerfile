FROM ubuntu:16.04
MAINTAINER s-ari

RUN apt-get update && apt-get install -y mysql-client
RUN mkdir -p /usr/local/goform_sql/html
COPY goform_sql /usr/local/goform_sql/
COPY ./html/index.html.tpl /usr/local/goform_sql/html
COPY create_table.sql /usr/local/goform_sql/
RUN chmod 755 /usr/local/goform_sql/goform_sql

EXPOSE 8080
ENTRYPOINT ["/usr/local/goform_sql/goform_sql"]

