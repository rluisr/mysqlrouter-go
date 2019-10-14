mysqlrouter-go
==============
[![Build Status](https://cloud.drone.io/api/badges/rluisr/mysqlrouter-go/status.svg)](https://cloud.drone.io/rluisr/mysqlrouter-go)
[![Coverage Status](https://coveralls.io/repos/github/rluisr/mysqlrouter-go/badge.svg?branch=test)](https://coveralls.io/github/rluisr/mysqlrouter-go?branch=test)

client for getting mysql-router information.

Supported version
-----------------
- 20190715 (8.0.17, 8.0.18)

Enable HTTP Server and REST API
-------------------------------
See [MySQL Router 8.0.17 and the REST API by lefred](https://lefred.be/content/mysqlrouter-8-0-17-and-the-rest-api/).

Usage
-----
```go
mr, err := mysqlrouter.New("https://mysqlrouter-test.xzy.pw", "luis", "luis")
routes, err := mr.GetAllRoutes()
```

See [example](example/main.go)

About mysqlrouter-test.xzy.pw
-----------------------
I provide MySQL Router 8.0.17 for testing.  
You can access this domain with basic `luis:luis` and you can request only GET method.  

Supported endpoint
-------------------
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
