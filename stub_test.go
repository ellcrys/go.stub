package stub

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type myBlockcode struct {
}

func (b *myBlockcode) OnInit() {
	On("add", func() (interface{}, error) {
		return 2 + 2, nil
	})
}

var _ = Describe("Go.Stub", func() {

	AfterEach(func() {
		reset()
	})

	Describe(".On", func() {

		It("should not add func if nil is passed as a function", func() {
			On("func1", nil)
			Expect(defaultStub.funcs).ToNot(HaveKey("func1"))
		})

		It("should successfully add function", func() {
			f := func() (interface{}, error) { return nil, nil }
			On("func1", f)
			Expect(defaultStub.funcs).To(HaveKey("func1"))
			Expect(defaultStub.funcs["func1"]).ToNot(BeNil())
		})
	})

	Describe(".Run", func() {

		It("should return panic when nil is pass", func() {
			Expect(func() {
				Run(nil)
			}).To(Panic())
		})

		It("should set default block code on default stub", func() {
			bc := new(myBlockcode)
			Expect(defaultStub.blockcode).To(BeNil())
			go Run(bc)
			time.Sleep(100 * time.Millisecond)
			close(defaultStub.wait)
			Expect(defaultStub.blockcode).ToNot(BeNil())
		})

		// It("should set default block code on default stub", func() {
		// 	bc := new(myBlockcode)
		// 	Expect(defaultStub.blockcode).To(BeNil())
		// 	Run(bc)
		// 	Expect(defaultStub.blockcode).ToNot(BeNil())
		// })
	})

	Describe(".getFunc", func() {

		It("should return nil when function is not found", func() {
			Expect(getFunc("unknown")).To(BeNil())
		})

		It("should successfully return the function", func() {
			f := func() (interface{}, error) { return nil, nil }
			On("func1", f)
			Expect(getFunc("func1")).ToNot(BeNil())
		})
	})

	Describe(".stopService", func() {
		It("will close the wait channel", func() {
			closed := make(chan bool)
			go func() {
				<-defaultStub.wait
				closed <- true
			}()
			time.Sleep(50 * time.Millisecond)
			stopService()
			Expect(<-closed).To(BeTrue())
		})
	})
})
