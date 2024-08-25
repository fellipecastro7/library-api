package main

import (
	"library-api/adapter/http"
	"library-api/container"
)

func main() {
	cont := container.NewContainer()

	r := http.SetupRouter(cont)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
