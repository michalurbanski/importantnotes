package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting program")

	file, err := os.Open("./data/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// TODO: add line only if it's an important line
		lines = append(lines, scanner.Text())
	}

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
