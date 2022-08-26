package partner

import (
	"fmt"
	_ "github.com/kononovn/partner-temp/tests/partner-tests/tests"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestVrf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Partner example tests suite")
}

var _ = BeforeSuite(func() {
	By(fmt.Sprintf("Partner's Before Suite"))

})

var _ = AfterSuite(func() {
	By(fmt.Sprintf("Partners AfterSuite"))
})
