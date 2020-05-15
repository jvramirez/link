package main

import (
	"fmt"
	"os"

	"github.com/jvramirez/link"
)

func main() {
	fmt.Println("Hello, GO!")

	tFiles := []string{"ex1.html", "ex2.html", "ex3.html", "ex4.html", "ex5.html", "ex6.html"}

	for _, file := range tFiles {
		displayLinks(file)
	}

}

func displayLinks(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	links, err := link.Parse(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", links)
}
