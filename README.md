<div align="center">
  <h1>goshik api for golang</h1>
</div>

## Foreword
- Work with API occurs only through OAuth2.

## Documentation
* API v1 https://shikimori.one/api/doc/1.0
* API v2 https://shikimori.one/api/doc/2.0 
* OAuth2 https://shikimori.one/oauth

## Makefile
- First step:
  * Add app "User-Agent: APPLICATION_NAME_HER"
  * Add access token "Authorization: Bearer ACCEESS_TOKEN_HER"
  * Add refresh token refresh_token="REFREHS_TOKEN_HER"
- Second step:
  * make test (check test connection)
  * make refresh (refresh access token)
