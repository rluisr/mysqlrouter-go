mysqlrouter-go
==============
client for getting mysql-router information.

Usage
-----
```go
mr, err := mysqlrouter.New("http://db-router.luis.local:8080", "luis", "luis")
if err != nil {
    panic(err)
}

routes, err := mr.GetAllRoutes()
```

See [example](example/main.go)

Supported endpoint
-------------------
### cluster
- [ ] /metadata
- [ ] /metadata/metadata_name/config
- [ ] /metadata/metadata_name/status

### app
- [ ] /router/status

### route
- [x] /routes
- [x] /routes/router_name/status
- [x] /routes/router_name/health
- [x] /routes/router_name/destinations
- [x] /routes/router_name/connections
