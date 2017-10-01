package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string `json:"N"`
}

func (a A) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"name": %q}`, a.Name)), nil
}

type B struct {
	A
	Desc string `json:"D"`
}

type C struct {
	Desc string `json:"D"`
}

func (c C) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"desc": %q}`, c.Desc)), nil
}

type D struct {
	A
	C
}

func main() {
	a := A{Name: "Martin"}
	data, _ := json.Marshal(a)
	fmt.Println("a:", string(data))

	b := B{A: a, Desc: "The Sad Robot"}
	data, _ = json.Marshal(b)
	fmt.Println("b:", string(data), "where is D?")

	c := C{Desc: "The Sad Robot"}
	data, _ = json.Marshal(c)
	fmt.Println("c:", string(data))

	d := D{a, c}
	data, _ = json.Marshal(d)
	fmt.Println("d:", string(data), "where are name and desc?")
}
