FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/bltrbot /opt/bin/bltrbot


CMD ["/opt/bin/bltrbot"]

