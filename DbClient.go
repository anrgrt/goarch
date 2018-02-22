package main

import(
"fmt"
"clientupdate"
"clientregister"
"clientviewall"
"clientviewmy"
)
func main(){
	fmt.Println("\n\t\tWhat do u want us to do:")
	fmt.Println("\t\t1: Register")
	fmt.Println("\t\t2:View All users")
	fmt.Println("\t\t3: view Ur account")
	fmt.Println("\t\t4:Update ur information")
	fmt.Println("\t\t5: Delete ur account")	
	var input int
	fmt.Scanln(&input)
	switch input{

		case 1: clientregister.Register()
		case 2: clientviewall.ViewAll()
		case 3: clientviewmy.ViewMy()
		case 4: clientupdate.Update()
		//case 5: Delete()
		default: fmt.Println("Please Enter a Valid Input")

	}








}

