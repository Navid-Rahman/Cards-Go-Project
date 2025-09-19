package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	navid := person{
		firstName: "Navid",
		lastName:  "Rahman",
		contact: contactInfo{
			email:   "navidrahman92@gmail.com",
			zipCode: 1234,
		},
	}

	fmt.Println(navid)

	fmt.Printf("%+v", navid)

}
