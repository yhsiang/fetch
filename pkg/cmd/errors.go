package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
)

func CheckError(err error) {
	if err != nil {
		if err != context.Canceled && err != io.EOF {
			fmt.Println(err.Error())
		}
		os.Exit(1)
	}
}
