package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeEmployeeCode(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := Employee{
		Name:       "Thongnarin",
		Salary:     20000,
		EmployeeID: "aa-1234",
	}

	ok, err := govalidator.ValidateStruct(employee)

	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by - and 4 digits (0-9)"))
}
