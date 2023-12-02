//go:generate bash -c "mkdir -p codegen && go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4 -generate types,server,spec -package codegen api/openapi.yaml > codegen/api.go"

package main

import "github.com/CorrectRoadH/Likit/cmd"

func main() {
	cmd.Main()
}
