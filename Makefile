test_go:
	go test -v

test_request:
	curl -X GET http://shikimori.one/api/users/whoami \
	-H "User-Agent: APPLICATION_NAME" \
	-H "Authorization: Bearer ACCESS_TOKEN"

token:
	curl -X POST "https://shikimori.one/oauth/token" \
	-H "User-Agent: APPLICATION_NAME" \
	-F grant_type="refresh_token" \
	-F client_id="CLIENT_ID" \
	-F client_secret="CLIENT_SECRET" \
	-F refresh_token="REFRESH_TOKEN"
