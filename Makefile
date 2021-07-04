all: build

build:
	go build .

test:
	go test -v

token:
	curl -X POST "https://shikimori.one/oauth/token" \
	-H "User-Agent: " \
	-F grant_type="refresh_token" \
	-F client_id="" \
	-F client_secret="" \
	-F refresh_token=""
