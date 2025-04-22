package main

import (
	"context"
	"github.com/hiddenmarten/ogen-head-example/internal"
	"net/http"
	"os"
)

func main() {
	api := internal.NewHumaAPI(context.Background(), http.NewServeMux())
	yaml, err := api.OpenAPI().DowngradeYAML()
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("api/oas/openapi.yaml", yaml, 0o644)
	if err != nil {
		panic(err)
	}
}
