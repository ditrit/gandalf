package log

import (
	"log"
	"os"
)

func OpenLogFile(path string) {
	file, err := os.OpenFile(path+"/gandalf.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//defer file.Close()

	log.SetOutput(file)
}
