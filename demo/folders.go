package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("creating node_modules folder with sample data")
	os.Mkdir("./demo/node_modules", 0755)

	for i := 0; i < 32; i++ {

		// convert i to string
		s := fmt.Sprintf("%d", i)
		os.WriteFile("./demo/node_modules/"+s, []byte("hello"), 0644)
	}
}
