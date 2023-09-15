FROM ubuntu:latest
LABEL maintainer="andreyisback@yandex.ru"

WORKDIR /root/docs
COPY . .

RUN apt update && apt install -y wget && \
wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz && \
tar -C /usr/local/ -zxf go1.21.1.linux-amd64.tar.gz

ENV PATH $PATH:/usr/local/go/bin

RUN go install -v golang.org/x/tools/cmd/godoc@latest && \
ln -sf /root/go/bin/godoc /usr/local/bin/godoc

EXPOSE 1337

CMD ["godoc", "-http=:1337", "-goroot=."]
