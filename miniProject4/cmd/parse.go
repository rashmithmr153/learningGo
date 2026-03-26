package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

type ParseState int
const (
	DateParse ParseState =iota
	TimeParse
	Done
)


type Date struct{
	date int
	month int
	year int	
}

type Time struct{
	hour int
	min int
	sec int
}

type TimeStamp struct{
	day Date
	time Time
	State ParseState
}
func NewTimestamp() TimeStamp{
	return TimeStamp{}
}


func ParseTimeStamps(data string)(TimeStamp,error){
	T1:=NewTimestamp()
	strLen,err:=T1.parse(data)
	if err!=nil{
		return T1,err
	}
	if len(data)!=strLen{
		return T1,fmt.Errorf("Error while parsing:length mismatch")
	}
	return T1,nil
}


func (t *TimeStamp) parse(data string)(int,error){
	read:=0
	idx:=strings.Index(data,"T")
	if idx==-1{
		return 0,fmt.Errorf("Invalid time stamp")
	}
	for t.State!=Done{
		switch t.State{
		case DateParse:
			date,n,err:=ParseDate(data[read:idx])
			if err!=nil{
				return 0,err
			}
			if valid,err:=isDate(*date);!valid{
				return 0,err
			}
			read+=n+len("T")
			t.day=*date
			t.State=TimeParse
		case TimeParse:
			time,n,err:=ParseTime(data[read:])
			if err!=nil{
				return 0,err
			}
			if valid,err:=isTime(*time);!valid{
				return 0,err
			}
			read+=n
			t.time=*time
			t.State=Done
		default:
			return read,fmt.Errorf("Invalid state")
		}
	}
	return read,nil
}

func isLeapYear(year int) bool{
		if year%4==0{
		if year%100==0{
			if year%400==0{
				return true
			}
		}else{
			return true
		}
	}
	return false
}


func isDate(d Date) (bool,error){
	date:=d.date
	month:=d.month
	year:=d.year
	flag:=isLeapYear(year)


	switch(month){
	case 2:
		if !flag{	
			if date<1 ||date>28{
				return false,fmt.Errorf("invalid date for february in non-leap year")
			}else {return true,nil} 
		}
		if date<1 ||date>29{
				return false,fmt.Errorf("invalid date for february in leap year")
		}
	case 1,3,5,7,8,10,12:
		if date<1 || date>31{
		return false,fmt.Errorf("invalid date")
		}
	case 4,6,9,11:
		if date<1 || date>30{
		return false,fmt.Errorf("invalid date")
		}
	default:
		return false,fmt.Errorf("invalid month")
	}
return true,nil
}


func ParseDate(data string)(*Date,int,error){
	
	readlen:=len(data)
	dateParts:=strings.Split(data,"-")
	if len(dateParts)!=3{
		return nil,0,fmt.Errorf("Invalid date in time stamp")
	}
	y,_:=strconv.Atoi(dateParts[0])
	m,_:=strconv.Atoi(dateParts[1])
	d,_:=strconv.Atoi(dateParts[2])
	
	return &Date{
		date:d,
		month: m,
		year: y,
		},readlen,nil
}


func ParseTime(t string)(*Time,int,error){
	parts:=strings.Split(t,":")
	readlen:=len(t)
	if len(parts)!=3{
		return nil,0,fmt.Errorf("Invalid time")
	}
	hour,_:=strconv.Atoi(parts[0])
	minute,_:=strconv.Atoi(parts[1])
	second,_:=strconv.Atoi(strings.Trim(parts[2],"Z"))
	return &Time{hour:hour,min:minute,sec:second,},readlen,nil
}

func isTime(t Time) (bool,error){
	h,m,s:=t.hour,t.min,t.sec

	if h<0 || h>23 {
		return false,fmt.Errorf("invalid hour value")
		
	}
	if m<0 || m>59{
		return false,fmt.Errorf("invalid minute value")
		
	}
	if s<0 || s>59 {
		return false,fmt.Errorf("invalid second value")	
	}
	return true,nil
}