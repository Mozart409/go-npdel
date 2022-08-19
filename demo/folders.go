package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("creating node_modules folder with sample data")
	os.Mkdir("./demo/node_modules", 0755)

	for i := 0; i < 32; i++ {

		// convert i to string
		s := fmt.Sprintf("%d", i)
		file, err := os.Create("./demo/node_modules/" + s)
		if err != nil {
			log.Fatal(err)
		}

		if err := file.Truncate(1e7); err != nil {
			log.Fatal(err)
		}
	}
}
