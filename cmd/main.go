package main

import (
	"fmt"
	"schedule/internal/config"
)

func main() {
	cfg := config.GetConfig()
	fmt.Println(cfg)
}
