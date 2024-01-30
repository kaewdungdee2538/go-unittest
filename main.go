package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	c := CustomerRepositoryMock{}
	c.On("GetCustomer",1).Return("Bond",18,nil)
	c.On("GetCustomer",2).Return("",0,errors.New("not found"))

	
	name,age,err := c.GetCustomer(2)
	if err != nil{
		panic(err)
	}
	fmt.Printf("name: %v,age: %v\n" ,name,age)
}

type CustomerRepository interface{
	GetCustomer(id int) (name string, age int, err error)
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (r *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := r.Called(id)
	return args.String(0),args.Int(1),args.Error(2)
}