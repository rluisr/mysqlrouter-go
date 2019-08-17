mysqlrouter-go
==============
[![Build Status](https://cloud.drone.io/api/badges/rluisr/mysqlrouter-go/status.svg)](https://cloud.drone.io/rluisr/mysqlrouter-go)

client for getting mysql-router information.

Usage
-----
```go
mr, err := mysqlrouter.New("https://mysqlrouter-test.xzy.pw", "luis", "luis")
routes, err := mr.GetAllRoutes()
```

See [example](example/main.go)

mysqlrouter-test.xzy.pw
-----------------------
I'm providing MySQL Router 8.0.17 for testing.  
You can access this with basic `luis:luis`.  
But you can only GET method.

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
- [x] /routes/router_name/status
- [x] /routes/router_name/health
- [x] /routes/router_name/destinations
- [x] /routes/router_name/connections
