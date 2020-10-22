package main

import (
	"github.com/YashKumarVerma/bird-nest/modules/shell"
	"github.com/YashKumarVerma/bird-nest/modules/ui"
	"github.com/kyokomi/emoji"
)

func main() {
	ui.Heading("Bird Nest " + emoji.Sprint(":bird: :tiger:"))
	ui.Description("a friendly cli to quickly build nest.js modules")
	shell.Init()
}
