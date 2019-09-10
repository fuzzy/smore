FROM alpine:3.9
RUN apk update
RUN apk add go gcc libpthread-stubs util-linux musl-utils musl-dev musl git ca-certificates ca-certificates-cacert
RUN mkdir -p /config
RUN rm -rf /go/src/git.devfu.net/fuzzy/smore/
ADD . /go/src/git.devfu.net/fuzzy/smore/
RUN env GOPATH=/go go get -v git.devfu.net/fuzzy/smore
RUN apk del go gcc libpthread-stubs util-linux musl-utils musl-dev musl git
VOLUME /config
COPY smore.yml /config/site.yml
CMD /bin/sh -c '/go/bin/smore -config /config/site.yml'

