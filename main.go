package main

import (
	"github.com/YashKumarVerma/bird-nest/modules/shell"
	"github.com/YashKumarVerma/bird-nest/modules/ui"
)

func main() {
	ui.Heading("Bird Nest")
	ui.Description("a handy repl shell to quickly build nest.js services")
	shell.Init()
}
