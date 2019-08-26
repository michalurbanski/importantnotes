package main

import (
	"fmt"
	"importantnotes/enums"
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
		return
	}

	note := models.NewNote(1, "text", enums.Important)
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
