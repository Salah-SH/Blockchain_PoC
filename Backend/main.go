package main

import (
	"fmt"
	"os"

	"github.com/cagnotteApp/Backend/api"
	"github.com/cagnotteApp/Backend/listener"
)

func main() {
	fmt.Println("Starting the application...")
	go listener.New()

	err := api.New()

	if err != nil {
		os.Exit(1)
	}

}
