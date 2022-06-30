mysqlrouter-go
==============
[![Build Status](https://cloud.drone.io/api/badges/rluisr/mysqlrouter-go/status.svg)](https://cloud.drone.io/rluisr/mysqlrouter-go)

client for getting mysql-router information.

Supported version
-----------------
- 20190715 (8.0.17 - 8.0.29)

Enable HTTP Server and REST API
-------------------------------
See [MySQL Router 8.0.17 and the REST API by lefred](https://lefred.be/content/mysqlrouter-8-0-17-and-the-rest-api/).

Usage
-----
```go
mysqlrouter.New("http://localhost:8080", "luis", "luis", nil)
```

See [example](example/main.go) and [client_test.go](client_test.go)

Supported endpoint
-------------------
### server
- [x] HTTPS with verify

### cluster
- [x] /metadata
- [x] /metadata/metadata_name/config
- [x] /metadata/metadata_name/status

### app
- [x] /router/status

### route
- [x] /routes
- [x] /routes/route_name/status
- [x] /routes/route_name/health
- [x] /routes/route_name/destinations
- [x] /routes/route_name/connections

Developer
---------
```shell
$ cd test && make local
```
