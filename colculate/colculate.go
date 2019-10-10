package colculate

func CoculateByMutiR(c chan int, num int,nRoutine int) int{
	flag := num / nRoutine //1000/5=200
	c<-0
	for i:=1;i<nRoutine;i++{
		//1-200
		//201-400
		go increase(c,flag*(nRoutine-1)+1,flag*nRoutine)
	}
	return <-c
}

func increase(c chan int, startnum int, stopnm int) chan int{
	temp := 0
	print("increase called,",startnum,"\n")
	for j:=startnum;j<stopnm;j++{
		temp = temp+j
	}
	temp = temp + <-c
	c<-temp
	return c
}

func AddNormal(startnum int, stopnm int) int{
	temp := 0
	for j:=startnum;j<stopnm;j++{
		temp = temp+j
	}
	return temp
}