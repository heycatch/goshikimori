App := 
Token := 

client_id := 
client_secret := 
refresh := 
auth_code :=

list:
	@grep '^[^#[:space:]].*:' Makefile

doc:
	godoc -http=:1337 -goroot=.

docker-build:
	docker build --no-cache -t shikimori-docs -f docker/Dockerfile .

docker-start:
	docker run -d -p 1337:1337 shikimori-docs

test:
	go test -v

run:
	go run cmd/main.go

request:
	curl -X GET https://shikimori.me/api/constants/anime \
	-H "User-Agent: $(App)" \
	-H "Authorization: Bearer $(Token)"

token:
	curl -X POST "https://shikimori.me/oauth/token" \
	-H "User-Agent: $(App)" \
	-F grant_type="refresh_token" \
	-F client_id="$(client_id)" \
	-F client_secret="$(client_secret)" \
	-F refresh_token="$(refresh)"

auth:
	curl -X POST "https://shikimori.me/oauth/token" \
	-H "User-Agent: $(App)" \
	-F grant_type="authorization_code" \
	-F client_id="$(client_id)" \
	-F client_secret="$(client_secret)" \
	-F code="$(auth_code)" \
	-F redirect_uri="urn:ietf:wg:oauth:2.0:oob"
