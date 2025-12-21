package main

import (
	"fmt"
	"os"

	models_v0 "gothub_release_dl/models/v0"
)

func main() {
	package_v0, err := models_v0.LoadJson("package.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", fmt.Sprintf("%s", err))
	}

	fmt.Println("%s", package_v0)
}
