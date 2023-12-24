package main

import "fmt"

//multiple named output
//if we name the output which we want return and then keep the return statement as it is then it will take that value accordingly
var bit int = 12
 func namedReturn(a , b int) (x , y int) {
	x = a
	y = b
	return
 }

// multiple output return type
func multiString(x string , y string) (string , string) {
	x = x + "hey"
	y = y + "new"
	return y , x
}

// function takes parameter as func (name type) <returnType> {}
func format(p int) int {	
	a , b , c := 4 , 5 , 6
	fmt.Printf("binary for a , b , c are  %b , %b , %b \n" , a , b , c)
	return p
}

func main()  {
	// fmt.Println("Hello gophers again.. 😊😊😊")
	str := "Aryan"
	var age int = 21 
	// ----Format printing -------
	fmt.Printf("My name is %s, and my age is %d \n"  , str,  age)
	/*
	----- ASSIGNING VALUE -----

	var f int = 92
	VARIABLE holds a VALUE of a specific TYPE

	for assinging any value easily
	a := "hey"
	a := 12
	*/

	var f int = 92
	var g int // gives an default value of 0(int) false(boolean)
	var check bool
	_ = 123 // blank idetifier
	a := 0 //also a var
	a = 7 // see
	b := "hey"
	
	fmt.Print("hey" , g   , a , b , "hey" ,f , check , "\n")

	format(27)
	

	fmt.Println(multiString("Aryan", "Anamika"))
	fmt.Println(namedReturn(23 , 45))

}