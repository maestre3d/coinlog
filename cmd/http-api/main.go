package main

import (
	"github.com/maestre3d/coinlog/restapi"
)

func main() {
	app, clean, err := restapi.NewCoinlogHTTP()
	if err != nil {
		panic(err)
	}
	defer clean()
	app.Start()
}
