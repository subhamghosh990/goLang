package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) SetName(n string) {
	p.Name = n
}

func (p *Person) SetAge(a int) {
	p.Age = a
}
func (p *Person) GetName() string {
	return p.Name
}
func (p *Person) GetAge() int {
	return p.Age
}

func main() {
	jsonT := `{"name":"subham","age":28, "add":"alipurduar"}`
	var o Person
	err := json.Unmarshal([]byte(jsonT), &o)
	if err != nil {
		fmt.Println("err : ", err)
	} else {
		fmt.Println(o.GetAge(), o.GetName())
	}
	d, err := json.Marshal(o)
	if err != nil {
		fmt.Println("err : ", err)
	} else {
		fmt.Println(string(d))
	}
}
