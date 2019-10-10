package creat

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func DispatchByChannel(c chan string) {
	fmt.Println("who want chocolate?")
	member := []string{}
	f := bufio.NewReader(os.Stdin)
	for {
		input, _ := f.ReadString('\n')
		if len(input) == 1 {
			break
		}
		member = append(member, input)
	}
	for i:=0;i<len(member);i++ {
		go println(member[rand.Intn(len(member))], " takes ", <-c)
	}
}
