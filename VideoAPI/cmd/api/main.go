package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/DexScen/VideoBot/VideoAPI/pkg/utils"
)

func init(){
	f, err := os.OpenFile("VideoAPI\\pkg\\utils\\videoapi.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic("file open failure")
	}
	defer f.Close()
	utils.NewSlog(f)
}

func main() {
	
	slog.Info("Api started working")
	i := 0
	for {
		slog.Info(
			"Api started working",
			"curr_i", i,
		)
		i++
		time.Sleep(5 * time.Second)
	}
}
