package greet

import( "fmt"
"errors"
"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string,error) {
	if name==""{
		return "",errors.New("can not passs emplty name to hello function")
	}
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message,nil
}


func MutiName(names []string) (map[string] string ,error){
	messege:=make(map[string]string)
	for _,name:= range names{
		messge,err:=Hello(name)
		if err!=nil{
			return nil,err
		}
		messege[name]=messge
	}
	return messege,nil
}

func randomFormat() string{
	//var formats []string
	formats :=[]string{
		"Hi %v, welcome!",
		"Great to see you %v",
		"Hola %v!",
		"Namaste %v",
	}
	return formats[rand.Intn(len(formats))]
}
