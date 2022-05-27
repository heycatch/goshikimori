App :=
Token :=

client_id :=
client_secret :=
refresh :=
auth_code :=

list:
	@grep '^[^#[:space:]].*:' Makefile

test:
	go test -v

request:
	curl -X GET https://shikimori.one/api/animes?search=initial+d \
	-H "User-Agent: $(App)" \
	-H "Authorization: Bearer $(Token)" \
	-F limit=5

token:
	curl -X POST "https://shikimori.one/oauth/token" \
	-H "User-Agent: $(App)" \
	-F grant_type="refresh_token" \
	-F client_id="$(client_id)" \
	-F client_secret="$(client_secret)" \
	-F refresh_token="$(refresh)"

auth:
	curl -X POST "https://shikimori.one/oauth/token" \
	-H "User-Agent: $(App)" \
	-F grant_type="authorization_code" \
	-F client_id="$(client_id)" \
	-F client_secret="$(client_secret)" \
	-F code="$(auth_code)" \
	-F redirect_uri="urn:ietf:wg:oauth:2.0:oob"
