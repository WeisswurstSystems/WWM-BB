package usecase

import (
	"log"
	"os"
)

var LOG = log.New(os.Stdout, "USER: ", log.Ldate|log.Ltime|log.Lshortfile)
