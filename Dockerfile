FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/bltr_bot /opt/bin/bltr_bot


CMD ["/opt/bin/bltr_bot"]

