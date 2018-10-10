package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "testing"

    . "github.com/franela/goblin"
    . "github.com/onsi/gomega"
)

var testinstances = map[string]string{
    "test_instance.xml":      "xml/test_instance.xml",
}
func TestLicense(t *testing.T) {
    g := Goblin(t)
    RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

    xf, ferr := ioutil.ReadFile(testinstances["test_instance.xml"])
    if ferr != nil {
        fmt.Printf(ferr.Error())
    }
    var spdx = NewLicense()
    err := xml.Unmarshal([]byte(xf), &spdx)
    if err != nil {
        fmt.Printf(err.Error())
    }
    g.Describe("License",func() {
	g.It("Must have licenseId",func() {
		Expect(spdx.LicenseID).To(Equal("Test string one"))
        })
	g.It("Must have comment",func() {
		Expect(spdx.CommentText).To(Equal("Test string one"))
        })
	g.It("Must have seeAlso",func() {
		Expect(spdx.SeeAlsoURI).To(Equal("http://anyuri.org"))
        })
	g.It("Must have name",func() {
		Expect(spdx.Name).To(Equal("Test string one"))
        })
	g.It("Must have IsDeprecatedLicenseID",func() {
		Expect(spdx.IsDeprecatedLicenseID).To(Equal("true"))
        })
	g.It("Must have isOsiApproved",func() {
		Expect(spdx.IsOsiApprovedIndicator).To(Equal("true"))
        })
	g.It("Must have IsFsfLibre",func() {
		Expect(spdx.IsFsfLibreIndicator).To(Equal("true"))
        })
	g.It("Must have standardLicenseHeader",func() {
		Expect(spdx.StandardLicenseHeader).To(Equal("Test string one"))
        })
	g.It("Must have licenseText",func() {
		Expect(spdx.LicenseText).To(Equal("Test string one"))
        })
	g.It("Must have standardLicenseTemplate",func() {
		Expect(spdx.StandardLicenseTemplate).To(Equal("Test string one"))
        })
    })

}