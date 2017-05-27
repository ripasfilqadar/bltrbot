FROM gcr.io/google-appengine/golang
COPY . /go/src/github.com/ripasfilqadar/bltrbot/bltrbot/
COPY . /go/src/app
RUN go-wrapper install

ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone


