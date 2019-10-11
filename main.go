package main

import (
	"fmt"
	"importantnotes/models"
	"importantnotes/readers/filereader"
	"log"
)

func main() {
	fmt.Println("Starting program")

	path := "./data/input.txt"
	lines, err := filereader.ReadLines(path)
	if err != nil {
		log.Fatal(err)
		//panic("Something went wrong") // panic message is not printed when
		// log.Fatal() is used
		return
	}

	note := models.NewNote(1, "text")
	fmt.Println(note)

	fmt.Println(lines)
	fmt.Println(len(lines))

	// lines := []string{}

	// for {
	// 	_, err := fmt.Fscanln(file, lines)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		} else {
	// 			fmt.Println("Scan error: ", err)
	// 			return
	// 		}
	// 	}
	// }

	fmt.Println("Program finished.")
}
