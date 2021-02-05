package log

import (
	"bytes"
	"fmt"
	"log"
)

func Log() {
	// logger will write type of bytes.Buffer
	buf := bytes.Buffer{}

	// 2nd args is prefix, last args is opts
	// conf options can conbine with logic and symbol
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ldate)

	logger.Println("test")

	logger.SetPrefix("new logger: ")

	logger.Printf("you can also add args(%v) and use Fataln to log and crash", true)

	fmt.Println(buf.String())
}
