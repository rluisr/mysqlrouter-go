[WIP] mysqlrouter-go
==============
client for getting mysql-router information.

under the work in progress.

Usage
-----
```go
mr, err := mysqlrouter.New("http://db-router.luis.local:8080", "luis", "luis")
if err != nil {
    panic(err)
}

routes, err := mr.GetAllRoutes()
if err != nil {
    panic(err)
}

for _, route := range routes.Item {
    fmt.Printf("name: %s\n", route.Name)
}
```