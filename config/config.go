package config

import (
	"log"
	"os"
	"strings"
)

var (
	SrcPath string
)

func LoadSrcPath() {
	srcPath, err := os.Getwd()
	if err != nil {
		log.Println("Unable to fetch current path:", err)
	}

	SrcPath = strings.TrimRight(srcPath, "config")
}
