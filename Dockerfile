FROM alpine:latest

WORKDIR /root/docs

COPY . .

EXPOSE 1337

RUN apk add --no-cache go

RUN go install -v golang.org/x/tools/cmd/godoc@latest && ln -sf /root/go/bin/godoc /usr/local/bin/godoc

CMD ["godoc", "-http=:1337", "-goroot=."]
