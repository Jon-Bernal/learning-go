package main

import "fmt"

func main() {
	fmt.Println("Methods")

	fmt.Println("----- Example 1 -----")
	p1 := person{"James", "Bond", 20}
	fmt.Println(p1.fullName())
	fmt.Println(p1.maxAge())

	fmt.Println("----- Example 2 -----")
	u1 := User{"James", "Bond"}
	fmt.Println(u1.userDetails())
}

// ========== Methods ========== //
// General Sytax for a method. THey're not object like in other languages. It's more like a function gets attached to a type.
// func (receiverName receiverType) methodName (parameter List)(returnTypes){
//  <em>  method body</em>
// }

// ----- Example 1 ----- //
type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + " " + p.last
}

func (p person) maxAge() int {
	return p.age
}

// ----- Example 2 ----- //
type User struct {
   Name  string
   Email string
}
 
func (u User) userDetails() string {
   return fmt.Sprintf("User name is: %s and email is: %s", u.Name, u.Email)
}
