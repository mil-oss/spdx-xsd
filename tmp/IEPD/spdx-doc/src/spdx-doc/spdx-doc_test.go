package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "testing"

    . "github.com/franela/goblin"
    . "github.com/onsi/gomega"
)

// TestSpdxDoc ...
func TestSpdxDoc(t *testing.T) {
    g := Goblin(t)
    RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

    xf, ferr := ioutil.ReadFile(testinstances["spdx-doc-test-instance.xml"])
    if ferr != nil {
        fmt.Printf(ferr.Error())
    }
    var spdxdoc = NewSpdxDocument()
    err := xml.Unmarshal([]byte(xf), &spdxdoc)
    if err != nil {
        fmt.Printf(err.Error())
    }
    g.Describe("SpdxDocument",func() {
        g.It("Must have Annotation",func() {
            Expect(spdxdoc.Annotation.Date).To(Equal("2018-04-12T13:20:00"))
            Expect(spdxdoc.Annotation.AnnotationTypeCode).To(Equal("Other"))
            Expect(spdxdoc.Annotation.Comment).To(Equal("Test string one"))
            Expect(spdxdoc.Annotation.Annotator).To(Equal("Test string one"))
        })
        g.It("Must have name",func() {
            Expect(spdxdoc.Name).To(Equal("Test string one"))
        })
        g.It("Must have comment",func() {
            Expect(spdxdoc.Comment).To(Equal("Test string one"))
        })
        g.It("Must have CreationInfo",func() {
            Expect(spdxdoc.CreationInfo.LicenseListVersion).To(Equal("Test string one"))
            Expect(spdxdoc.CreationInfo.Created).To(Equal("2018-04-12T13:20:00"))
            Expect(spdxdoc.CreationInfo.Comment).To(Equal("Test string one"))
            Expect(spdxdoc.CreationInfo.Creator).To(Equal("Test string one"))
        })
        g.It("Must have specVersion",func() {
            Expect(spdxdoc.SpecVersion).To(Equal("Test string one"))
        })
        g.It("Must have dataLicense",func() {
            Expect(spdxdoc.DataLicense).To(Equal("Test string one"))
        })
    })

}