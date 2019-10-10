package creat

import (
	"bufio"
	"fmt"
	"os"
)

/**
*/
func InputByChannle(c chan string) {
	fmt.Println("pls input some chocolate")
	f := bufio.NewReader(os.Stdin)
	for {
		input, _ := f.ReadString('\n')
		if len(input) == 1 {
			break
		}
		c <- input
	}
}