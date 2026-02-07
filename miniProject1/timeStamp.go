package main
import (
	"fmt"
	"errors"
)



func strToInt(str string)(int,error){
	//function to covert string vlaues to integers
	//i/p: str sting
	//op num int
	if str==""{
		//checking for if empty values is given as i/p
		return 0,errors.New("cannot convert empty string")
	}
	num:=0
	for _,c:= range str{
		if c<'0' || c>'9'{
			//here checking done for only numeric values passed in i/p or not
			return 0,errors.New("function can only have integers as inputs in string")
		}
		num =num*10+int(c-'0')
	}
	return num,nil
}

func compValues(v1,v2 int)int{
	//helper functon for compareing values
	if v1<v2{
		return -1
	} else if v1>v2{
		return 1
	}
	return 0
}
func isLeapYear(year int)bool{
	//this function checkes wheather given year is leap year or not
	//input: year as int
	//o/p: true if leap year or false
	//for checking logic 
	// check:https://www.geeksforgeeks.org/dsa/program-check-given-year-leap-year/
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

func isDateValid(date,month,year int)(bool,error){
	//this funtions checkes wheteher given date is valis or not
	//input:date, month, year all in int
	//output// boolean(true if valid,flase if invalid) and error
	flag:=isLeapYear(year)

	//first checked for is given month is valid or not
	if month<1 || month>12{
		return false,errors.New("invalid month")
	}else{
		//here special condition checked for feb mont
		if month==2{
			if !flag{
				//for non leap yaers	
				if date>28{
					return false,errors.New("invalid date for february in non-leap year")
				}
			}
			if date>29{
				//for leap yaers
					return false,errors.New("invalid date for february in leap year")
			}
		}
	}
	if date<1 || date>31{
		//checking for vsalid dates
		return false,errors.New("invalid date")
	}

	return true,nil;
}



func totalDaysinmonth(date,month,year int)(int,error){
	//This logic is bit confusing
	// the formula to calulate is like this
	// total days= date + extra days in months before + 30*(months before)
	// extra days in months before is stored in arrays below for leap and non-leap years as sum
	// eg: for march in leap year extra days are 1(feb) +0(jan)=1
	// for may in non-leap year extra days are 0(jan)+0(feb)+1(mar)+0(apr)=1
	// so total days in may 15 in non-leap year = 15+1+30*4=136
	leapExtradays:=[13]int{0,1,0,1,1,2,2,3,4,4,5,5,6}
	extradays:=[13]int{0,1,-1,0,0,1,1,2,3,3,4,4,5}
	flag:=isLeapYear(year)
	validDate,err1:=isDateValid(date,month,year)
	if !validDate{
		return 0,err1
	}
	if flag{
		return date+ leapExtradays[month-1] +((month-1)*30),nil
	}
	return date+extradays[month-1]+((month-1)*30),nil
}
func dateCmp(t1,t2 string)(int,error){
	//this function takes 2 dates t1 and t2 as string and compares them
	//the restuen values are following
		//if date1 is earlier return -1,with error as nil
		//if date1 is later return 1,with error as nil
		//if equal return 0,with error as nil
		//if any errros found return 0, with errors


	//checking for valid length of date
	if len(t1)<8 || len(t2)<8{
		return 0,errors.New("invalid date format")
	}

	//errors checked intaially as if any alphabets or charctes passed to the function
	//dates is coverted to int for convince in calculations
	_,err1:=strToInt(t1)
	_,err2:=strToInt(t2)
	if err1!=nil || err2!=nil{
		if err1!=nil{
			return 0,err1
		}
		return 0,err2
	}
	y1:=t1[0:4]
	y2:=t2[0:4]
	year1,_:=strToInt(y1)
	year2,_:=strToInt(y2)
	m1,_:=strToInt(t1[4:6])
	m2,_:=strToInt(t2[4:6])
	d1,_:=strToInt(t1[6:8])
	d2,_:=strToInt(t2[6:8])
	//intial checking is done for years here with help of helper function
	if compValues(year1, year2)!=0{
		//checking for if date is valid or not
		validDate1,err1:=isDateValid(d1,m1,year1)
		validDate2,err2:=isDateValid(d2,m2,year2)
		if !validDate1{
			return 0,err1
		}
		if !validDate2{
			return 0,err2
		}
		return compValues(year1, year2),nil
	}else{
		//here date and moths are passed to the function to get total days present till date in that year
		dayCount1,err1:=totalDaysinmonth(d1,m1,year1)
		dayCount2,err2:=totalDaysinmonth(d2,m2,year2)
		if err1!=nil || err2!=nil{
			if err1!=nil{
				return 0,err1
			}
			return 0,err2
		}
		if compValues(dayCount1, dayCount2)!=0{
			return compValues(dayCount1,dayCount2),nil
		}else{
			return 0,nil
		}
	}
}
func timeComp(t1,t2 string){
	//this function takes 2 times(HH:MM:SS) t1 and t2 as string and compares them
	//the resturn values are following
		//if time1 is earlier return -1,with error as nil
		//if time1 is later return 1,with error as nil
		//if equal return 0,with error as nil
		//if any errros found return 0, with errors


	_,err1:=strToInt(t1)
	_,err2:=strToInt(t2)
	if err1!=nil || err2!=nil{
		if err1!=nil{
			fmt.Println("Error in TimeStamp1: ",err1)
			return	
		}
		fmt.Println("Error in TimeStamp2: ",err2)
		return
	}

	h1,_:=strToInt(t1[0:2])
	h2,_:=strToInt(t2[0:2])
	m1,_:=strToInt(t1[2:4])
	m2,_:=strToInt(t2[2:4])
	s1,_:=strToInt(t1[4:6])
	s2,_:=strToInt(t2[4:6])
	//follwoing block checkes  wheter its is valid time
	if h1<0 || h1>23 || h2<0 || h2>23{
		fmt.Println("invalid hour value")
		return
	}
	if m1<0 || m1>59 || m2<0 || m2>59{
		fmt.Println("invalid minute value")
		return
	}
	if s1<0 || s1>59 || s2<0 || s2>59{
		fmt.Println("invalid second value")
		return
	}
	//checking done as total seconds passed in that day till that time
	totalSec1:=h1*3600 + m1*60 + s1
	totalSec2:=h2*3600 + m2*60 + s2
	result:=compValues(totalSec1,totalSec2)
	switch result{
	case -1:
		fmt.Println("TimeStamp1 is earlier than TimeStamp2")
	case 1:
		fmt.Println("TimeStamp1 is later than TimeStamp2")
	default:
		fmt.Println("TimeStamp1 and TimeStamp2 are from same time")
	}
}



func timeStampcmp(t1,t2 string){
	
	// checking for valid type of timestamps, now checking for ISO standards
	if len(t1)<20 || len(t2)<20{
		fmt.Println("invalid timestamp format")
		return 
	}
	//stipping dates from timw stamps without the sperteros '-'
	DateNum1:=t1[0:4]+t1[5:7]+t1[8:10]
	DateNum2:=t2[0:4]+t2[5:7]+t2[8:10]
	//stripping times from time stamps without the spertors ':'
	timeNum1:=t1[11:13]+t1[14:16]+t1[17:19]
	timeNum2:=t2[11:13]+t2[14:16]+t2[17:19]
	// first passingg only the dates if they are same or not
	result,err:=dateCmp(DateNum1,DateNum2)
		if err != nil {
		fmt.Println("Error comparing timestamps")
		fmt.Println("Error: ",err)
		return
	}
	switch result{
	case -1:
		fmt.Println("TimeStamp1 is earlier than TimeStamp2")
	case 1:
		fmt.Println("TimeStamp1 is later than TimeStamp2")
	default:
		timeComp(timeNum1,timeNum2)
	}
}

func main(){
	TimeStamp1:="2024-02-03T10:15:30Z"
	TimeStamp2:="2024-02-03T10:15:30Z"
	//timestamps 1 &2 are passed to functions
	timeStampcmp(TimeStamp1,TimeStamp2)
}