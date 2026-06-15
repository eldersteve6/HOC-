package main 

import(
	"fmt"
	"os"
)
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: please provide input")
		return
	}

	arg := os.Args[1]

	words, err := getLetters(arg)
	if err != nil {
		
		for _, w := range words {
			fmt.Println(w)
		}
		fmt.Println(err)
		return
	}

	for _, w := range words {
		fmt.Println(w)
	}
}