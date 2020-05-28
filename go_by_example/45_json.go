package main

import (
	"encoding/json"
	"fmt"
)

type res1 struct {
	Page   int      `json:"page"` // Can't have a space
	Fruits []string `json:"fruits"`
}

func main() {

	// Generic decoding
	var dat map[string]interface{}
	byt := []byte(`{"num":6.13,"strs":["a","b"], "test": "test"}`)
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)

	fmt.Println(dat["test"])

	// Normal values
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Encoding
	a := res1{
		Page:   43,
		Fruits: []string{"zodiac", "tohru", "rat"}}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))

	// Decoding into a structure
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := res1{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}
