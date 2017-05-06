package gomegaconvey

import (
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/smartystreets/goconvey/convey/reporting"
)

//
// We hook into the Gomega library and provide
// the same functions as the DSL. Then by adding
// the appropriate Goconvey code, we can create
// the bridge between the two libraries.
//

func RegisterTestingT(t types.GomegaTestingT) {
	gomega.RegisterTestingT(t)
}

func Ω(actual interface{}, extra ...interface{}) gomega.GomegaAssertion {
	return gomega.Ω(actual, extra)
}

func Expect(actual interface{}, extra ...interface{}) gomega.GomegaAssertion {
	return ExpectWithOffset(0, actual, extra...)
}

func ExpectWithOffset(offset int, actual interface{}, extra ...interface{}) gomega.GomegaAssertion {
	return gomega.ExpectWithOffset(offset, actual, extra)
}

func Eventually(actual interface{}, intervals ...interface{}) gomega.GomegaAssertion {
	return EventuallyWithOffset(0, actual, intervals)
}

func EventuallyWithOffset(offset int, actual interface{}, intervals ...interface{}) gomega.GomegaAssertion {
	return EventuallyWithOffset(offset, actual, intervals)
}

func Consistently(actual interface{}, intervals ...interface{}) gomega.GomegaAsyncAssertion {
	return ConsistentlyWithOffset(0, actual, intervals)
}

func ConsistentlyWithOffset(offset int, actual interface{}, intervals ...interface{}) gomega.GomegaAsyncAssertion {
	return gomega.ConsistentlyWithOffset(offset, actual, intervals)
}

//
// Gomega matchers
//

func Equal(expected interface{}) types.GomegaMatcher {
	reporting.NewSuccessReport()
	return gomega.Equal(expected)
}
