package main

import (
	"os"

	"bazil.org/bazil/cli"
)

//go:generate go run task/gen-imports.go -o commands.gen.go bazil.org/bazil/cli/...
//go:generate go get -v go.pedge.io/protoeasy/cmd/protoeasy
//go:generate protoeasy --go --grpc --context bazil.org/bazil .

func main() {
	code := cli.Main()
	os.Exit(code)
}
