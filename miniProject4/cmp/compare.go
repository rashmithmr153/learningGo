package cmp

import(
	"Date/cmd"
)

func TimeCmp(t1,t2 string)(int,error)  {
	T1,err1:=cmd.ParseTimeStamps(t1)
	T2,err2:=cmd.ParseTimeStamps(t2)
	if err1!=nil{
		return 0,err1
	}
	if err2!=nil{
		return 0,err2
	}
	return 0,nil
}