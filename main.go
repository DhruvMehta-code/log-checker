package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	Path := flag.String("path", "sample.log", "you can also change log file location")
	level := flag.String("level", "[error]", "you can change any level like error,alert...etc")
	flag.Parse()
	file, err := os.Open(*Path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	read := bufio.NewReader(file)
	for {
		s, err := read.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}

}
