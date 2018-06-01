package stub

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type myBlockcode2 struct {
}

func (b *myBlockcode2) OnInit() {
}

type myBlockcode3 struct {
}

func (b *myBlockcode3) OnInit() {
	On("func1", func() (interface{}, error) {
		return nil, fmt.Errorf("an error")
	})
}

type myBlockcode4 struct {
}

func (b *myBlockcode4) OnInit() {
	On("func1", func() (interface{}, error) {
		return 10, nil
	})
}

var _ = Describe("Service", func() {

	AfterEach(func() {
		reset()
	})

	Describe(".Invoke", func() {
		It("should return err if function does not exist", func() {
			bc := new(myBlockcode2)
			defaultStub.blockcode = bc
			service := newService(defaultStub)
			res := service.Invoke(Args{
				Func: "unknown",
			})
			Expect(res).ToNot(BeNil())
			Expect(res.Error).To(BeTrue())
			Expect(res.Body).To(Equal("unknown function `unknown`"))
		})

		It("should return error returned by the invoked function", func() {
			bc := new(myBlockcode3)
			defaultStub.blockcode = bc
			service := newService(defaultStub)
			res := service.Invoke(Args{
				Func: "func1",
			})
			Expect(res).ToNot(BeNil())
			Expect(res.Error).To(BeTrue())
			Expect(res.Body).To(Equal("an error"))
		})

		It("should return success value returned by the invoked function", func() {
			bc := new(myBlockcode4)
			defaultStub.blockcode = bc
			service := newService(defaultStub)
			res := service.Invoke(Args{
				Func: "func1",
			})
			Expect(res).ToNot(BeNil())
			Expect(res.Error).To(BeFalse())
			Expect(res.Body).To(Equal(10))
		})
	})
})
