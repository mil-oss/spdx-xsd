package spdxsec

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

var testinstances = map[string]string{
	"test_instance.xml": "xml/test_instance.xml",
}

func TestSoftwareEvidenceArchive(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	xf, ferr := ioutil.ReadFile(testinstances["test_instance.xml"])
	if ferr != nil {
		fmt.Printf(ferr.Error())
	}
	var spdx = NewSoftwareEvidenceArchive()
	err := xml.Unmarshal([]byte(xf), &spdx)
	if err != nil {
		fmt.Printf(err.Error())
	}
	g.Describe("SoftwareEvidenceArchive", func() {
	})

}
