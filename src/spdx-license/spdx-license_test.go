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
        g.It("Must have licenseId",func() {
            Expect(spdxlic.LicenseID).To(Equal(""))
        })
        g.It("Must have comment",func() {
            Expect(spdxlic.Comment).To(Equal(""))
        })
        g.It("Must have seeAlso",func() {
            Expect(spdxlic.SeeAlso).To(Equal(""))
        })
        g.It("Must have name",func() {
            Expect(spdxlic.Name).To(Equal(""))
        })
        g.It("Must have IsDeprecatedLicenseID",func() {
            Expect(spdxlic.IsDeprecatedLicenseID).To(Equal(""))
        })
        g.It("Must have isOsiApproved",func() {
            Expect(spdxlic.IsOsiApproved).To(Equal(""))
        })
        g.It("Must have IsFsfLibre",func() {
            Expect(spdxlic.IsFsfLibre).To(Equal(""))
        })
        g.It("Must have standardLicenseHeader",func() {
            Expect(spdxlic.StandardLicenseHeader).To(Equal(""))
        })
        g.It("Must have licenseText",func() {
            Expect(spdxlic.LicenseText).To(Equal(""))
        })
        g.It("Must have standardLicenseTemplate",func() {
            Expect(spdxlic.StandardLicenseTemplate).To(Equal(""))
        })
    })

}