package errors

import (
	"errors"
	"fmt"
	"log"
)

func check(str string) (string, error) {
	if str == "string" {
		err := errors.New("str equal")
		return "", err
	} else {
		return str, nil
	}
}

func Err() {
	str, err := check("string")
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Print("err: ", err)
		return
	}
	fmt.Println("str: ", str)
}
