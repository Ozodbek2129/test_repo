package main

import (
	"fmt"
	"strconv"

	// "github.com/Ozodbek2129/test_repo/greating"
	"github.com/Ozodbek2129/test_repo/user"
)

func main(){
	// fmt.Println(greating.Greet())
	// fmt.Println("Mening ismim Ozodbek.")
	u:=user.User{}
	u.ValidateUser()
	
	var ageStr string

	fmt.Printf("Ismni kiriting: ")
	fmt.Scan(&u.Name)

	fmt.Printf("Yoshni kiriting: ")
	fmt.Scan(&ageStr)
	age, ok := strconv.Atoi(ageStr)
	if ok != nil {
		fmt.Println("Yoshni noto'g'ri kiritdingiz.")
		return
	}
	u.Age=age

	fmt.Printf("Emailni kiriting: ")
	fmt.Scan(&u.Email)

	errors := u.ValidateUser()

	if len(errors) > 0 {
		fmt.Println("\nXatolar:")
		for _, e := range errors {
			fmt.Println(e)
		}
	} else {
		fmt.Println("\nFoydalanuvchi ma'lumotlari to'g'ri!")
	}
}
