package main

import "fmt"

type adress struct {
	state  string
	distic string
}

type Person struct {
	name   string
	age    int
	adress adress
}

func main() {

	pers := map[string]interface{}{

		" per1": map[string]interface{}{
			"name":   "jayaram",
			"age":    26,
			"adress": map[string]interface{}{"state": "ap", "distic": "konaseema"},
		},
		"per2": map[string]interface{}{
			"name":   "jayaram",
			"age":    27,
			"adress": map[string]interface{}{"state": "ts", "distic": "hyderabad"},
		},
	}

	if per2, ok := pers["per2"].(map[string]interface{}); ok {

		if adress, ok := per2["adress"].(map[string]interface{}); ok {
			fmt.Println("adress of per2:", adress)

		}
		if name, ok := per2["name"].(string); ok {
			fmt.Println("Name of per2:", name)
		}

	}

	fmt.Println("per2==>", pers, pers["per1"], pers["per2"])

	per2 := pers["per2"]

	fmt.Println("per2==>", per2)

	pers1 := map[string]Person{
		"per1": {
			name:   "jayaram",
			adress: adress{state: "ap", distic: "konaseema"},
			age:    26,
		},
		"per2": {
			name:   "jayaram",
			age:    27,
			adress: adress{state: "ts", distic: "hyderabad"},
		},
	}
	fmt.Println("pers1===>", pers1, pers1["per1"], pers1["per2"].age)
}
