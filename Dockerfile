FROM golang
ADD . /go/src/github.com/vallard/trudiedo-api
WORKDIR /go/src/github.com/vallard/trudiedo-api
# grab all the deps
RUN go get
EXPOSE 8080
CMD ["go", "run", "main.go"]
