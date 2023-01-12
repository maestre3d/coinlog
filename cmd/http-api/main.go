package main

import "github.com/maestre3d/coinlog/di"

func main() {
	app, clean, err := di.NewCoinlogHTTP()
	if err != nil {
		panic(err)
	}
	defer clean()
	app.Start()
}
