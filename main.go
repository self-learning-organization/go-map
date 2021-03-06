package main

import "fmt"

func main() {
	colors := map[string]string{ // One of the ways to declare a map. First string between the square brackets indicates that the keys are of type string. And the next occurrence of string says all the values are string as well
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}

	// var colors map[string]string // Another way of declaring a map. Since we didn't assign an actual value, Go will initialize it with its zero value. I.e. an empty map - has no key value pairs inside of it

	// colors := make(map[string]string) // Yet another way of declaring a map, using Go's built-in function. This line here and the previous way of declaring a map are pretty much equivalent for all intents and purposes

	// colors["white"] = "#ffffff" // To add a key value pair

	// delete(colors, "white") // To delete a key value pair, use the built in function 'delete'

	printMap(colors)
}

func printMap(c map[string]string) { // Breakdown of syntax in notes' point 2
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// Notes:
// 1. In Go, a map is a collection of key-value pairs
//    Map is Go is like - Hash in Ruby - Object in JavaScript - Dict in Python
//    - In maps, both the keys and the values are statically typed
//      So whenever we add some number of keys to a map in Go, they must all be of the same exact type.
//      Likewise, all the different values that we add as well must also be of the exact same type.
//      Just the keys and the values themselves don't have to be of the same type just all the different values have to.
//      I.e. key1, key2, key3 must be all of the same type
//           and value1, value2, value3 must be all of the same type
//
// 2. func printMap(c map[string]string) { // Breakdown of syntax in notes' point 2
//	     for color, hex := range c {
//
//	     }
//    }
//    The c in the function's parameter is the argument's name, whereas the map[string]string is the type of the map
//    color in the for loop is the key for this step through the loop, whereas the hex is the value for this step through the loop
//
// 3. Differences between Map and Struct
//    Map:
//    - All keys must be the same type
//    - All values must be the same type
//    - Keys are indexed - we can iterate over them
//    - Use to represent a collection of related properties <- Use case
//    - Don't need to know all the keys at compile time
//    - Reference type!
//    Struct:
//    - Values can be of different type
//    - Keys don't support indexing
//    - You need to know all the different fields at compile time
//    - Use to represent a "thing" with a lot of different properties <- Use case
//    - Value type!

// Quiz 9: Test Your Knowledge: Maps
// 1. Can some of the keys in a single map be of type int and others of type string? No
// 2. Can some of the values in a single map be of type int and others of type string? No
// 3. Take a look at the following code. What would the print statement log out? map[dog: bark cat: purr]
/* 	package main
	import "fmt"

	func main() {
 		m := map[string]string{
   			"dog": "bark",
 		}

		changeMap(m)

 		fmt.Println(m)
	}

	func changeMap(m map[string]string) {
 	m["cat"] = "purr"
	} */
// 4. What would happen if we tried to run the following program? The compiler would throw an error because the variable key was created, but never used
/* 	package main
	import "fmt"

	func main() {
		m := map[string]string{
   			"dog": "bark",
   			"cat": "purr",
 		}

 		for key, value := range m {
   			fmt.Println(value)
 		}
	} */
