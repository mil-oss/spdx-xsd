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
func TestSpdxLicense(t *testing.T) {
    g := Goblin(t)
    RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

    xf, ferr := ioutil.ReadFile(testinstances["test_instance.xml"])
    if ferr != nil {
        fmt.Printf(ferr.Error())
    }
    var spdxlic = NewLicense()
    err := xml.Unmarshal([]byte(xf), &spdxlic)
    if err != nil {
        fmt.Printf(err.Error())
    }
    g.Describe("SPDX",func() {
        g.It("Must have isOsiApproved",func() {
            Expect(spdxlic.IsOsiApproved).To(Equal("true"))
        })
        g.It("Must have standardLicenseHeader",func() {
            Expect(spdxlic.StandardLicenseHeader).To(Equal("Test string one"))
        })
        g.It("Must have licenseText",func() {
            Expect(spdxlic.LicenseText).To(Equal("Test string one"))
        })
        g.It("Must have standardLicenseTemplate",func() {
            Expect(spdxlic.StandardLicenseTemplate).To(Equal("Test string one"))
        })
    })

}