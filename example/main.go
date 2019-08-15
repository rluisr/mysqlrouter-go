package main

import (
	"fmt"

	"github.com/rluisr/mysqlrouter-go"
)

func main() {
	mr, err := mysqlrouter.New("http://localhost:5901", "luis", "luis")
	if err != nil {
		panic(err)
	}

	routes, err := mr.GetAllRoutes()
	if err != nil {
		panic(err)
	}

	route := routes[1].Name
	fmt.Printf("route name: %s\n", route)

	routeStatus, err := mr.GetRouteStatus(route)
	if err != nil {
		panic(err)
	}
	fmt.Printf("route status: %+v\n", routeStatus)

	routeHealth, err := mr.GetRouteHealth(route)
	if err != nil {
		panic(err)
	}
	fmt.Printf("route health: %+v\n", routeHealth)

	routeDestinations, err := mr.GetRouteDestinations(route)
	if err != nil {
		panic(err)
	}
	fmt.Printf("route destinations: %+v\n", routeDestinations[0])

	routeConnections, err := mr.GetRouteConnections(route)
	if err != nil {
		panic(err)
	}
	fmt.Printf("route connections: %+v\n", routeConnections[0])

	metadatas, err := mr.GetAllMetadata()
	if err != nil {
		panic(err)
	}

	metadata := metadatas[0].Name
	fmt.Printf("metadata name: %s\n", metadata)

	metadataConfig, err := mr.GetMetadataConfig(metadata)
	if err != nil {
		panic(err)
	}
	fmt.Printf("metadata config: %+v\n", metadataConfig)

	metadataStatus, err := mr.GetMetadataStatus(metadata)
	if err != nil {
		panic(err)
	}
	fmt.Printf("metadata status: %+v\n", metadataStatus)

	routerStatus, err := mr.GetRouterStatus()
	if err != nil {
		panic(err)
	}
	fmt.Printf("router status: %+v\n", routerStatus)
}
