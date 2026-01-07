package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeSalary(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := Employee{
		Name:       "Thongnarin",
		Salary:     5000,
		EmployeeID: "HR-1234",
	}

	ok, err := govalidator.ValidateStruct(employee)

	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
}
