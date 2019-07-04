package main

import "fmt"

func main() {
	fmt.Println("Have fun for fucking thing inthis training")
	foo() //function foo
	fmt.Println("Make This More")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("this a ODD number =",i)
			//fmt.Printf("|%-6s|%-6s|\n",i,"this a ODD number")
		}
		if i%2 != 0 {
			fmt.Println("this Not ODD number =",i)
			//fmt.Printf("|%-6s|%-6s|\n",i,"this Not ODD number")
		}

	}
	bar() // bar function
}

func foo() {
	fmt.Println("Test the function Foo")
}

func bar() {
	fmt.Println("If Failure you will control flow")
}
//control flow:
// (1) sequence
// (2) loop interactive
// (3) condition
