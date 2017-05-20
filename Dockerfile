FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/bltr_bot /opt/bin/bltr_bot
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/bltr_bot"]

ENV PORT 3000
