package main

import "fmt"

func main() {
	type MyString = string

	str := "BCD"
	mystr1 := MyString(str)
	//mystr2 := MyString("A" + str)
	fmt.Printf("%T(%q)==%T(%q):%v\n", str, str, mystr1, mystr1, str == mystr1)
	//res true

	strs := []string{"E", "F", "G"}
	mystrs := []MyString(strs)
	fmt.Printf("A value of type []MyString: %T(%q)\n",
		mystrs, mystrs)
	fmt.Printf("Type %T is the same as type %T.\n", mystrs, strs)
	fmt.Println()
	//这段代码主要说明是可以给type取别名的
}
