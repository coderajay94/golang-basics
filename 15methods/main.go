package main

import "fmt"

type Customer struct {
	Name     string
	Salary   float64
	IsActive bool
	Email    string
}

func main() {

	fmt.Println("calling methods in golang")

	ajay := Customer{
		Name:     "ajay Kumar",
		Salary:   12344.23,
		IsActive: true,
		Email:    "ajay@gmail.com",
	}

	fmt.Println(ajay)

	ajay.IsActiveUser()

	ajay.ChangeEmail()

	// it doesn't reflect the change happened as, we have just passed the copy
	fmt.Println(ajay)

	// in case you really want to change the email
	//pass the pointer

	ajay.ChangeEmailWithPointer()
	fmt.Println(ajay)

}

//now actual object pointer is passed so when you changed it , it will reflect the change
func (customer *Customer) ChangeEmailWithPointer() {
	customer.Email = "ajay@hotmail.com"
	fmt.Println("changeEmailWithPointer: New email for ajaykumar is:", customer.Email)
}

func (customer Customer) ChangeEmail() {
	customer.Email = "ajay@hotmail.com"
	fmt.Println("changeEmail:New email for ajaykumar is:", customer.Email)
}

//methods are going to be associated with the struct
func (customer Customer) IsActiveUser() {

	if customer.IsActive {
		fmt.Println("Customer is active")
	}
}
