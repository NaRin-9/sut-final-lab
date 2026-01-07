package entity

import "gorm.io/gorm"

type Employee struct{
	gorm.Model
	Name string `valid:"stringlength(2|80)"`
	Salary float64 `valid:"range(15000|200000)~Salary must be between 15000 and 200000"`
	EmployeeID string `valid:"matches(^[A-Z]{2}-\\d{4})~EmployeeCode must be 2 uppercase English letters (A-Z) followed by - and 4 digits (0-9)"`

}
