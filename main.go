package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	var hello string
	hello = "hello"

	println(hello)

	//-----------------------------------

	world := "world"

	println(world)

	//-------------------------------------

	strSl := []string{"this", "is", "a", "slice"}

	for _, str := range strSl {
		print(str + " ")
	}

	//------------------------------------

	print("\n")

	rsltStr := strings.Join(strSl, " | ")

	println(rsltStr)

	//------------------------------------

	err := ioutil.WriteFile("testWrite.txt", []byte(rsltStr), 0666)

	if err != nil {
		panic(err)
	}
}
