package test

import (
	"myapp/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestHee(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("test", func(t *testing.T) {
		approved := entity.Product{
			Name:        "hee",
			Price:       45,
			Quantity:    45,
			Description: "geekkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk",
		}

		ok, err := govalidator.ValidateStruct(approved)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

