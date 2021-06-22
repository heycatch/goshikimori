all: run

run:
	cd example && go build . && ./example

test:
	curl -X GET https://shikimori.one/api/users/whoami \
	-H "User-Agent: Api Test" \
	-H "Authorization: Bearer ACCESS_TOKEN"

refresh:
	curl -X POST "https://shikimori.one/oauth/token" \
	-H "User-Agent: Api Test" \
	-F grant_type="refresh_token" \
	-F client_id="bce7ad35b631293ff006be882496b29171792c8839b5094115268da7a97ca34c" \
	-F client_secret="811459eada36b14ff0cf0cc353f8162e72a7d6e6c7930b647a5c587d1beffe68" \
	-F refresh_token="REFRESH_TOKEN"
