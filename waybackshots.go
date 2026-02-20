package main

import (
	"github.com/thelicato/waybackshots/cmd"
	"github.com/thelicato/waybackshots/pkg/utils"
)

var version = "0.2.0"

func main() {
	utils.Banner(version)
	cmd.Execute()
}
