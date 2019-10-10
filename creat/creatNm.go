package creat

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CreatNum() int{
	fmt.Println("input a num pls?")
	f := bufio.NewReader(os.Stdin)
                                          	num, _ := f.ReadString('\n')
	num = strings.TrimRight(num, "\n")
	i,err := strconv.ParseInt(num,10,32)
	if err!=nil {
		print(err)
	}
	return int(i)
}

func CreatRoutine() int{
	fmt.Println("how many routine would you like?")
	f := bufio.NewReader(os.Stdin)
	num, _ := f.ReadString('\n')
	num = strings.TrimRight(num, "\n")
	i,err := strconv.Atoi(num)
	if err!=nil {
		print(err)
	}
	return i
}

