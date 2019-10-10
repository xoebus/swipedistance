package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xoebus/swipedistance"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		word := scanner.Text()
		dist, ok := swipedistance.Sum(word)
		if !ok {
			continue
		}

		perChar := dist / float64(len(word))

		fmt.Printf("%f\t%f\t%s\n", dist, perChar, word)
	}
}
