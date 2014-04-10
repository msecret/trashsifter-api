FROM msecret/go
MAINTAINER Marco Secret, msegreto@miceover.com

ENV DEBIAN_FRONTEND noninteractive:w
ENV GOPATH /srv/go
ENV APP_PATH $GOPATH/src/github.com/msecret/trashsifter-api
ENV PATH /srv/go/bin:$PATH

ADD . $APP_PATH
## WORKDIR doesn't not expand env vars
## see https://github.com/dotcloud/docker/issues/2637 
WORKDIR /srv/go/src/github.com/msecret/trashsifter-api

VOLUME ["/srv/go/src/github.com/msecret/trashsifter-api"]

RUN go get github.com/githubnemo/CompileDaemon
RUN go get
RUN go build

EXPOSE 80

CMD ["CompileDaemon", "-directory=/srv/go/src/github.com/msecret/trashsifter-api", "-command=/srv/go/src/github.com/msecret/trashsifter-api/trashsifter-api"]
