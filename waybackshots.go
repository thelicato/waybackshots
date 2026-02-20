package main

import (
	"github.com/thelicato/waybackshots/cmd"
	"github.com/thelicato/waybackshots/pkg/utils"
)

var version = "0.1.1"

func main() {
	utils.Banner(version)
	cmd.Execute()
}
