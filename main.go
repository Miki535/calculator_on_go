package main

import "fmt"

func main() {
	for {
		var s1 float64
		var s2 float64
		var choise uint8
		fmt.Println("$ 1)+, 2) -, 3) *, 4) /, 5:")
		fmt.Scan(&choise)
		if choise == 1 {
			fmt.Println("$ Enter firts number:")
			fmt.Scan(&s1)

			fmt.Println("$ Enter second number:")
			fmt.Scan(&s2)

			fmt.Println("You result is: \n" + fmt.Sprint(s1+s2))

			var a string
			fmt.Scan(&a)
		} else if choise == 2 {
			fmt.Println("$ Enter first number:")
			fmt.Scan(&s1)

			fmt.Println("$ Enter second number:")
			fmt.Scan(&s2)

			fmt.Println("You result is: \n" + fmt.Sprint(s1-s2))

			var a string
			fmt.Scan(&a)
		} else if choise == 3 {
			fmt.Println("$ Enter first number:")
			fmt.Scan(&s1)

			fmt.Println("$ Enter second number:")
			fmt.Scan(&s2)

			fmt.Println("$ You result is: \n" + fmt.Sprint(s1*s2))

			var a string
			fmt.Scan(&a)
		} else if choise == 4 {
			fmt.Println("$ Enter first number:")
			fmt.Scan(&s1)

			fmt.Println("$ Enter second number:")
			fmt.Scan(&s2)

			fmt.Println("$ You result is: \n" + fmt.Sprint(s1/s2))

			var a string
			fmt.Scan(&a)
		} else if choise == 5 {
			fmt.Println("Bye!")
			break
		}
	}
}
