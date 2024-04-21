package main

import "github.com/greenblat17/digital_spb/internal/app"

const configPath = "config/config.yml"

func main() {
	app.Run(configPath)
}
