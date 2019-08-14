package main

import (
	"fmt"

	"github.com/rluisr/mysqlrouter-go"
)

func main() {
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
}
