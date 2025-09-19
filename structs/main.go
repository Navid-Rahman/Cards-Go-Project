package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// 1st way of defining struct
	navid := person{"Navid", "Rahman"}
	fmt.Println(navid)

	// 2nd way of defining struct
	seyam := person{firstName: "Sadman", lastName: "Seyam"}
	fmt.Println(seyam)

	// 3rd way of defining struct
	var fatin person
	fmt.Println(fatin)
	fmt.Printf("%+v", fatin)

	fatin.firstName = "Fatin"
	fatin.lastName = "Benyamin"
	fmt.Println(fatin)
}
