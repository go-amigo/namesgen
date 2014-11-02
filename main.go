package main

import (
	"fmt"
	"github.com/docker/docker/pkg/namesgenerator"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("%s unable to convert %s to an integer", err, os.Args[1])
		}
		fmt.Println(namesgenerator.GetRandomName(num))
	} else {
		fmt.Println(namesgenerator.GetRandomName(0))
	}
}
