package main

import (
	"log"
	"os"

	"github.com/slamchillz/platnova/constant"
)

func init() {
	// Check if temp folder exists, if not create it
	if _, err := os.Stat(constant.DefaultOutPutDir); os.IsNotExist(err) {
		err = os.Mkdir(constant.DefaultOutPutDir, 0755)
		if err != nil {
			log.Println("Error creating output directory:", err)
		}
	}
}
