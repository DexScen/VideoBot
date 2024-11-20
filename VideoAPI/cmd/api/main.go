package main

import (
	"log/slog"

	"github.com/DexScen/VideoBot/VideoAPI/pkg/utils"
)

func main() {
	err := utils.NewSlog()
	if err != nil {
		panic("logger config failure")
	}
	slog.Info("Api started working")
}
