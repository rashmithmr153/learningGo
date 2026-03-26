package main

import (
	"fmt"
	"Date/cmp"
)

func main(){
	timestamp1:="2024-02-03T10:15:30Z"
	timestamp2:="2024-02-03T10:15:30Z"
	valid,err:=cmp.TimeCmp(timestamp1,timestamp2)
	if err!=nil{
		panic(err)
	}
	switch(valid){
	case 0:
		fmt.Print("Both time stamps are from same")
		return
	case 1:
		fmt.Print("time stamp 1  latest")
		return
	case -1:
		fmt.Print("time stamp 2  latest")
		return
	default:
		panic("Somehow error shown")
	}
}