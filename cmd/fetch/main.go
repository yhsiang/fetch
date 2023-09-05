package main

import (
	"github.com/yhsiang/fetch/pkg/cmd"
	"github.com/yhsiang/fetch/pkg/cmd/fetch"
)

func main() {
	err := fetch.NewCommand().Execute()
	cmd.CheckError(err)
}
