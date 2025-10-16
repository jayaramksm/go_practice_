package main

import (
	"fmt"
)

type Adress struct {
	State   string
	Pincode int
}

type Person struct {
	Name string
	Age  int
	Adress
}

func main() {

	per := []Person{
		{Name: "jayaram", Age: 26, Adress: Adress{State: "ap", Pincode: 533212}},
		{Name: "ram", Age: 27, Adress: Adress{State: "ts", Pincode: 500055}},
	}

	fmt.Println("person===>", per, per[0].Name, per[1].Adress.Pincode)

	per1 := []map[string]interface{}{
		{"Name": "jayaram", "Age": 26, "Adress": map[string]interface{}{"State": "ap", "Pincode": 533212}},
		{"Name": "ram", "Age": 27, "Adress": map[string]interface{}{"State": "ts", "Pincode": 500055}},
	}
	fmt.Println("person_interface===>", per1, per1[1]["Adress"])

	if Adress, ok := per1[1]["Adress"].(map[string]interface{}); ok {
		fmt.Println("person_Adress_interface===>", Adress["Pincode"])

	}

}
