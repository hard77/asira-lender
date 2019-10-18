FROM golang:alpine  AS build-env

ARG APPNAME="asira_lender"
ARG ENV="dev"

#RUN adduser -D -g '' golang
#USER root

ADD . $GOPATH/src/"${APPNAME}"
WORKDIR $GOPATH/src/"${APPNAME}"

RUN apk add --update git gcc libc-dev;
RUN apk --no-cache add curl
#  tzdata wget gcc libc-dev make openssl py-pip;
RUN go get -u github.com/golang/dep/cmd/dep

RUN cd $GOPATH/src/"${APPNAME}"
RUN cp deploy/dev-config.yaml config.yaml
RUN dep ensure -v
RUN go build -v -o "${APPNAME}-res"

RUN ls -alh $GOPATH/src/
RUN ls -alh $GOPATH/src/"${APPNAME}"
RUN ls -alh $GOPATH/src/"${APPNAME}"/vendor
RUN pwd

FROM alpine

WORKDIR /go/src/
COPY --from=build-env /go/src/asira_lender/asira_lender-res /go/src/asira_lender
COPY --from=build-env /go/src/asira_lender/deploy/dev-config.yaml /go/src/config.yaml
COPY --from=build-env /go/src/asira_lender/permissions.yaml /go/src/permissions.yaml

RUN pwd
#ENTRYPOINT /app/asira_lender-res
CMD ["/go/src/asira_lender","run"]

EXPOSE 8000
