# ggz

[![GoDoc](https://godoc.org/github.com/go-ggz/ggz?status.svg)](https://godoc.org/github.com/go-ggz/ggz)
[![Build Status](https://cloud.drone.io/api/badges/go-ggz/ggz/status.svg)](https://cloud.drone.io/go-ggz/ggz)
[![Build status](https://ci.appveyor.com/api/projects/status/prjvsklt3io5nuhn/branch/master?svg=true)](https://ci.appveyor.com/project/appleboy/ggz/branch/master)
[![codecov](https://codecov.io/gh/go-ggz/ggz/branch/master/graph/badge.svg)](https://codecov.io/gh/go-ggz/ggz)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ggz/ggz)](https://goreportcard.com/report/github.com/go-ggz/ggz)
[![codebeat badge](https://codebeat.co/badges/6fc8d61a-17c1-446d-a895-2dc6a8d1c16c)](https://codebeat.co/projects/github-com-go-ggz-ggz-master)
[![Docker Pulls](https://img.shields.io/docker/pulls/goggz/ggz-server.svg)](https://hub.docker.com/r/goggz/ggz-server/)
[![Get your own image badge on microbadger.com](https://images.microbadger.com/badges/image/goggz/ggz-server.svg)](https://microbadger.com/images/goggz/ggz-server "Get your own image badge on microbadger.com")

An URL shortener service written in Golang.

## Features

* Support [MySQL](https://www.mysql.com/), [Postgres](https://www.postgresql.org/) or [SQLite](https://www.sqlite.org/) Database.
* Support [RESTful](https://en.wikipedia.org/wiki/Representational_state_transfer) or [GraphQL](http://graphql.org/) API.
* Support [Auth0](https://auth0.com/) or [Firebase](https://firebase.google.com/) Single Sign On (default is `auth0`).
* Support expose [prometheus](https://prometheus.io/) metrics and database data like count of registerd users.
* Support install TLS certificates from [Let's Encrypt](https://letsencrypt.org/) automatically.
* Support [QR Code](https://en.wikipedia.org/wiki/QR_code) Generator from shorten URL.
* Support local disk storage or [Minio Object Storage](https://minio.io/).
* Support linux and windows container, see [Docker Hub](https://hub.docker.com/r/goggz/ggz/tags/).
* Support integrate with [Grafana](https://grafana.com/) service.

## Requirement

Go version: `1.13`

## Start app using docker-compose

See the `docker-compose.yml`

```yml
version: '3'

services:
  ggz:
    image: goggz/ggz
    restart: always
    ports:
      - 8080:8080
      - 8081:8081
    environment:
      - GGZ_DB_DRIVER=sqlite3
      - GGZ_SERVER_HOST=http://localhost:8080
      - GGZ_SERVER_SHORTEN_HOST=http://localhost:8081
      - GGZ_AUTH0_PEM_PATH=test.pem
```

## Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/go-ggz/ggz.svg)](https://starcharts.herokuapp.com/go-ggz/ggz)
