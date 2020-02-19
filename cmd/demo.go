package main

import (
	"log"

	"github.com/mlavergn/goaudit"
)

func main() {
	platform := goaudit.NewClientPlatform("Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405")
	log.Println(platform)
}
