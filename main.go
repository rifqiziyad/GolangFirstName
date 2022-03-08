package main

import "fmt"

func main()  {
	// === declare variabel ===

	// var firstName string = "Rifqi"
	// var lastName string
	// lastName = "Imtinan"

	// == declare variable without var (type inference) ==

	// midName := "Ziyad"

	// == declare multiple variable with comma ==

	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"

	// var fourth, fifth, sixth string = "empat", "lima", "enam"

	// seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// with type inference technique can declare multiple variable and different data type

	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// == use underscore (_) if the variable is not used (reserved variable) ==

	// _ = "belajar Golang"
	// _ = "Golang itu mudah"
	// name, _ := "john", "wick"

	// == declare variable using "new" ==
	// name := new (string)
	// fmt.Println(name) 	---> 0xc000010230
	// fmt.Println(*name) ---> ""

	// declare vaiable using "make"
	// for channel, slice, map

	fmt.Println("Golang First Time")
 }
