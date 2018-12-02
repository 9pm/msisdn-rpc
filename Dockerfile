FROM golang:alpine
ADD . ./
RUN go build -o main
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT [“/main”]