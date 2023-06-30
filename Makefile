App := Api Test
Token := XadcAlKJ32gF-Dpjuk4mSdcciPRqaxYEhxAo4-pFvoU

client_id := bce7ad35b631293ff006be882496b29171792c8839b5094115268da7a97ca34c
client_secret := 811459eada36b14ff0cf0cc353f8162e72a7d6e6c7930b647a5c587d1beffe68
refresh := P-HMUgXHWcrf-L6e2IZ531KxwXRGaFsJZc1bSPf-UKw
auth_code := QwtuhrfZblfamTauRQaigfK4ofbAVhbKccocumQmd4A

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
