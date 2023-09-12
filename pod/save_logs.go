package pod

import (
	"log"
	"os"
)

func SaveLogsFromPod(logs string) {
	f, err := os.Create("logs.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(logs)

	if err != nil {
		log.Fatal(err)
	}
}
