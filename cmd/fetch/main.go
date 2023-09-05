package main

import (
	"github.com/yhsiang/autify/pkg/cmd"
	"github.com/yhsiang/autify/pkg/cmd/fetch"
)

func main() {
	err := fetch.NewCommand().Execute()
	cmd.CheckError(err)
}
