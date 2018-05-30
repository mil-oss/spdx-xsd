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
        g.It("Must have ",func() {
            Expect(spdxlic.IsDeprecatedLicenseID).To(Equal(""))
        })
        g.It("Must have StandardLicenseHeader",func() {
            Expect(spdxlic.StandardLicenseHeader).To(Equal("Test string one"))
        })
        g.It("Must have StandardLicenseTemplate",func() {
            Expect(spdxlic.StandardLicenseTemplate).To(Equal("Test string one"))
        })
        g.It("Must have LicenseText",func() {
            Expect(spdxlic.LicenseText).To(Equal("Test string one"))
        })
        g.It("Must have IsOsiApproved",func() {
            Expect(spdxlic.IsOsiApproved).To(Equal("true"))
        })
        g.It("Must have IsFsfLibre",func() {
            Expect(spdxlic.IsFsfLibre).To(Equal("true"))
        })
        g.It("Must have ",func() {
            Expect(spdxlic.LicenseID).To(Equal(""))
        })
        g.It("Must have Name",func() {
            Expect(spdxlic.Name).To(Equal("Test string one"))
        })
        g.It("Must have SeeAlso",func() {
            Expect(spdxlic.SeeAlso[0]).To(Equal("http://anyuri.org"))
        })
        g.It("Must have Comment",func() {
            Expect(spdxlic.Comment).To(Equal("Test string one"))
        })
    })

}