App :=
Key :=

client_id :=
client_secret :=
refresh :=

list:
	@grep '^[^#[:space:]].*:' Makefile

test:
	go test -v

request:
	curl -X GET https://shikimori.one/api/users/whoami \
	-H "User-Agent: $(App)" \
	-H "Authorization: Bearer $(Key)"

token:
	curl -X POST "https://shikimori.one/oauth/token" \
	-H "User-Agent: $(App)" \
	-F grant_type="refresh_token" \
	-F client_id="$(client_id)" \
	-F client_secret="$(client_secret)" \
	-F refresh_token="$(refresh)"
