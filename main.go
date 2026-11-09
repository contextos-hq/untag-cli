package main

import (
	"github.com/contextos-hq/untag-cli/cmd"
	"github.com/contextos-hq/untag-cli/internal/config"
	"github.com/contextos-hq/untag-cli/internal/updater"
)

func main() {
	updater.Start(config.Version)
	cmd.Execute()
}
