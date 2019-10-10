package main

import (
	"exercise/colculate"
	"exercise/creat"
	"time"
)

func main() {
	//choco := make(chan string ,10)
	//creat.InputByChannle(choco)
	//creat.DispatchByChannel(choco)


	num := creat.CreatNum()
	addch := make(chan int,num)
 	nRoutine := creat.CreatRoutine()
	strat := time.Now() // get current time
 	sum:=colculate.CoculateByMutiR(addch,num,nRoutine)
 	tm := time.Since(strat)
	stratn := time.Now()
	sumNormal:=colculate.AddNormal(1,num)
 	tn := time.Since(stratn)
 	print(sum," := sum value of mutiR, and time is :",tm,"\n")
 	print(sumNormal," := sum value of normal, and time is :", tn,"\n")
}
