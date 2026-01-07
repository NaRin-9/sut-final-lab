package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositiveEmployee(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := Employee{
		Name:       "Thongnarin",
		Salary:     20000,
		EmployeeID: "HR-1234",
	}

	ok, err := govalidator.ValidateStruct(employee)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
