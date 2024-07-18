package main

import (
	"github.com/hoangxt/go-ecommerce-ba-api/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")[default]
}
