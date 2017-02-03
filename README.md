[![Build Status](https://travis-ci.org/adorsys/golang-chi-rest-db-oauth-sample.svg?branch=master)](https://travis-ci.org/adorsys/golang-chi-rest-db-oauth-sample)

# golang-chi-rest-db-oauth-sample
REST sample with all the stuff we need in our day jobs

- [REST](https://github.com/pressly/chi) with documentation
- [JWT](https://github.com/goware/jwtauth) with [central IDP](https://auth0.com)
- [DB migrations](https://github.com/mattes/migrate) on [postgres](https://github.com/lib/pq)
- [external configuration with TOML](https://github.com/BurntSushi/toml)

## Setup
- install Go 1.7
- start service: `go run main.go --conf data/conf/dev.toml`
- [get JWT](https://buildrunclick.eu.auth0.com/login?client=0beCklFKuabEpbQ2SJ34m6JmwxYDsn5H&protocol=oauth2&redirect_uri=https://adorsys.de/karriere.html&response_type=token&scope=openid roles) (admin@buildrun.click:admin)
  - copy token from URL after redirect
- try with curl:
```bash
curl --request GET \
  --url http://localhost:3333/articles \
  --header 'accept: application/json' \
  --header 'authorization: Bearer $JWT'
```
- generate route markdown docs: `go run main.go --routes`
