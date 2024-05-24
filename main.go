package main

import (
	"fmt"
	// "sort"
	// "strings"
	// "math"
)

// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v \n", n)
// }

// func sayBye(n string) {
// 	fmt.Printf("Goodbye %v \n", n)
// }

// func cycleNames(n []string, f func(string)) {
// 	for _, v := range n {
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }

// func getInitials(n string) (string, string) {
// 	s := strings.ToUpper(n)
// 	names := strings.Split(s, " ")

// 	var initials []string
// 	for _, v := range names {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], "_"
// }

var score = 99.5

func main() {
	// var student1 string = "John" //type is string
	// var student2 = "Jane"        //type is inferred
	// x := 2                       //type is inferred

	// fmt.Println(student1)
	// fmt.Println(student2)
	// fmt.Println(x)

	// In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type
	// var a string
	// var b int
	// var c bool

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// strings

	// var nameOne string = "Abdimalik"
	// var nameTwo = "Osman"
	// var nameThree string
	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "Abdullahi"
	// nameThree = "Hassan"
	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "Abdi"
	// fmt.Println(nameFour)

	// ints
	// var ageOne int = 24
	// var ageTwo = 32
	// ageThree := 33
	// fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint16 = 256

	// var scoreOne float32 = 29.98
	// var scoreTwo float64 = 7777777772882822.8
	// scoreThree := 1.5
	// fmt.Println(scoreOne, scoreTwo, scoreThree)

	// var ages = [3]int{20, 25, 30}

	// names := [4]string{"yoshi", "mario", "peach", "bowser"}
	// names[1] = "luigi"

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// // slices (use arrays under the hood)
	// var scores = []int{100, 50, 60}
	// scores[2] = 25
	// scores = append(scores, 85)

	// fmt.Println(scores, len(scores))

	// // slice ranges
	// rangeOne := names[1:4] // doesn't include pos 4 element
	// rangeTwo := names[2:]  //includes the last element
	// rangeThree := names[:3]

	// fmt.Println(rangeOne, rangeTwo, rangeThree)
	// fmt.Printf("the type of rangeOne is %T \n", rangeOne)

	// rangeOne = append(rangeOne, "koopa")
	// fmt.Println(rangeOne)

	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.Contains(greeting, "bye"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.Split(greeting, " "))

	// // the original value is unchanged
	// fmt.Println("original string value =", greeting)

	// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)

	// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "bowser"))

	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++ // infinite loop without this
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is:", i)
	// }

	// names := []string{"mario", "luigi", "yoshi", "peach"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, val := range names {
	// 	fmt.Printf("the value at pos %v is %v \n", index, val)
	// 	val = "new string"
	// }

	// for _, val := range names {
	// 	fmt.Print(val, ",")
	// 	val = "new string"
	// }

	// // changing val in a loop does not mutate the original slice
	// fmt.Println(names)

	// age := 45

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is not less than 40")
	// }

	// names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	// for index, val := range names {
	// 	if index == 1 {
	// 		fmt.Println("continuing at pos", index)
	// 		continue
	// 	}
	// 	if index > 2 {
	// 		fmt.Println("breaking at pos", index)
	// 		break
	// 	}
	// 	fmt.Printf("the value at pos %v is %v \n", index, val)
	// }

	// sayGreeting("Abdimalik")
	// sayGreeting("Osman")
	// sayBye("Abdimalik")

	// cycleNames([]string{"cloud", "barret", "tifa"}, sayGreeting)
	// cycleNames([]string{"cloud", "barret", "tifa"}, sayBye)

	// a1 := circleArea(10.5)
	// a2 := circleArea(15)

	// fmt.Println(a1, a2)
	// fmt.Printf("circle 1 area is %0.3f & circle 2 area is %0.3f \n", a1, a2)

	// fn1, sn1 := getInitials("tifa lockhart")
	// fmt.Println(fn1, sn1)

	// fn2, sn2 := getInitials("cloud strife")
	// fmt.Println(fn2, sn2)

	// fn3, sn3 := getInitials("barret")
	// fmt.Println(fn3, sn3)

	// accesing greetings file
	// sayHello("Abdimalik")
	// showScore()

	// for _, v := range points {
	// 	fmt.Println(v)
	// }

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[845775485])

	phonebook[984759373] = "bowser"
	fmt.Println(phonebook)

	phonebook[647583927] = "yoshi"
	fmt.Println(phonebook)
}
