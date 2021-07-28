package main

import "fmt"

var global_var = "testvariable"

func main() {

	fmt.Println("Hello, World!")
	var fname string
	var lname string
	fmt.Println("enter fname:")
	fmt.Scan(&fname)
	fmt.Println("enter lname:")
	fmt.Scan(&lname)
	customer := getData(fname, lname)
	fmt.Println(customer)
	fmt.Println(global_var)

}

func getData(fname string, lname string) (customer string) {
	if fname != "" && lname != "" {
		return fname + " " + lname
	} else {
		return ""
	}
}
