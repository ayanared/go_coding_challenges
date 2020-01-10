package main

import (
	"fmt"
	"strings"
)

func main() {
	// Uppercase each string in this array as a new array called uppercaseLyrics:
	// const lyrics = ['never', 'gonna', 'give', 'you', 'up'];
	// const uppercaseLyrics = ???

	lyrics := []string{"never", "gonna", "give", "you", "up"}
	fmt.Println(lyrics)

	var uppercaseLyrics []string

	for _, v := range lyrics {
		uppercaseWord := strings.ToUpper(string(v[0])) + v[1:]
		uppercaseLyrics = append(uppercaseLyrics, uppercaseWord)
	}

	fmt.Println(uppercaseLyrics)

	// Return an array of names for each of these person objects:
	// const people = [
	// 	{
	// 		name: 'George Michael',
	// 		age: 14,
	// 		title: 'Mr. Manager'
	// 	},
	// 	{
	// 		name: 'T-Bone',
	// 		age: 34,
	// 		title: 'Arsonist'
	// 	},
	// 	{
	// 		name: 'George Oscar',
	// 		age: 32,
	// 		title: 'Illusionist'
	// 	}
	// ];

	// const names = ???
	type person struct {
		name  string
		age   int
		title string
	}
	people := []person{
		person{name: "George Micheal", age: 34, title: "Mr. Manager"},
		person{name: "T-Bone", age: 34, title: "Arsonist"},
		person{name: "George Oscar", age: 32, title: "Illusionist"},
	}
	var names []string
	fmt.Println(people)
	for _, p := range people {
		names = append(names, p.name)
	}
	fmt.Println(names)

	// Return the amount of tax to charge for each of these products, assuming a tax rate of 7%:
	// const products = [
	// 	{
	// 		name: 'iPad',
	// 		price: 549.99
	// 	},
	// 	{
	// 		name: 'iPhone',
	// 		price: 799.99
	// 	},
	// 	{
	// 		name: 'iPod',
	// 		price: 2.99
	// 	}
	// ];

	// const tax = ???

	type product struct {
		name  string
		price float32
	}

	products := []product{
		product{name: "iPad", price: 549.99},
		product{name: "iPhone", price: 799.99},
		product{name: "iPod", price: 2.99},
	}
	var tax float32 = 0

	for _, p := range products {
		tax = tax + p.price*0.07
	}
	taxInDollars := fmt.Sprintf("%.2f", tax)
	fmt.Println("The tax is: $" + taxInDollars)

}
