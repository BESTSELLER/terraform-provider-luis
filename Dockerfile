FROM golang:1.16 as build

WORKDIR $GOPATH/src/github.com/BESTSELLER/terraform-provider-luis
COPY . .
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init -g api/router.go
RUN GO111MODULE=on CGO_ENABLED=0 go install

FROM alpine

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/BESTSELLER/terraform-provider-luis"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

WORKDIR /
COPY --from=build /go/bin/terraform-provider-luis /

CMD /terraform-provider-luis
EXPOSE 3000
