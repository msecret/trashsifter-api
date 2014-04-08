FROM msecret/go
MAINTAINER Marco Secret, msegreto@miceover.com

ENV DEBIAN_FRONTEND noninteractive:w
ENV GOPATH /srv/go
ENV APP_PATH $GOPATH/src/github.com/msecret/trashsifter-api

ADD . $APP_PATH
## WORKDIR doesn't not expand env vars
## see https://github.com/dotcloud/docker/issues/2637 
WORKDIR /srv/go/src/github.com/msecret/trashsifter-api

RUN go get
RUN go build

EXPOSE 80

ENTRYPOINT ["./trashsifter-api"]
