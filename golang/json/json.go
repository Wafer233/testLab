package json

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

// struct -> json
func JsonMarshal() {
	p := Person{
		Name:  "wafer",
		Age:   20,
		Email: "wafer@mail.com",
	}
	b, _ := json.Marshal(p)
	//	b = []byte(`{"Name":"wafer","Age":20,"Email":"wafer@mail.com"}`)
	fmt.Printf("\n%v\n", string(b))
}

func JsonUnmarshal() {
	b := []byte(`{"Name":"wafer","Age":20,"Email":"wafer@mail.com"}`)
	var person Person
	json.Unmarshal(b, &person)
	fmt.Printf("\n%+v\n", person)

}
