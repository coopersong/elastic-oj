package main

import (
	"os"

	"elastic-oj/router"
)

func main() {
	r := router.NewRouter()
	if err := r.Run(); err != nil {
		os.Exit(1)
	}
}
