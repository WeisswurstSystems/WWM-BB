package usecase

import (
	"log"
	"os"
)

var LOG = log.New(os.Stdout, "MEETING: ", log.Ldate|log.Ltime|log.Lshortfile)
