FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/go-getting-started /opt/bin/WWM-BB

CMD ["/opt/bin/WWM-BB"]
