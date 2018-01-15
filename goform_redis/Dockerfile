FROM ubuntu:16.04
MAINTAINER s-ari

RUN mkdir -p /usr/local/goform/html
COPY goform /usr/local/goform/
COPY ./html/index.html.tpl /usr/local/goform/html
RUN chmod 755 /usr/local/goform/goform

EXPOSE 8080
ENTRYPOINT ["/usr/local/goform/goform"]

