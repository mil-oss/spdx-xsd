package spdxsecism

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
    g.Describe("SoftwareEvidenceArchive",func() {
	g.It("Must have Software Information",func() {
            Expect(spdx.SoftwareInformation.ProductTitleText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.SoftwareInformation.SoftwareNameText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.SoftwareInformation.SoftwareOrgText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.SoftwareInformation.VersionInformation.SoftwareVersionText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionIndicator).To(Equal("true"))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionText).To(Equal("1.0.0-alpha.beta"))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionInformation.APIName).To(Equal("APIname"))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionInformation.MajorVersionNumeric).To(Equal("1"))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionInformation.MinorVersionNumeric).To(Equal("1"))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionInformation.PatchVersionText).To(Equal("1"))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionInformation.PreReleaseVersionText).To(Equal("-alpha"))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionInformation.BuildMetaText).To(Equal("+exp.sha.5114f85"))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionDate).To(Equal("2018-02-15T09:00:00"))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonSemanticVersionIndicator).To(Equal("true"))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonEpochSegmentText).To(Equal("1"))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonReleaseSegmentText).To(Equal("0"))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonPreReleaseSegmentText).To(Equal("3"))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonPostReleaseSegmentText).To(Equal("2"))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonDevelopmentReleaseSegmentText).To(Equal("3"))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonVersionText).To(Equal("1.0.3.2.3"))
            Expect(spdx.SoftwareInformation.GroupingText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.SoftwareInformation.ReleaseNotesText).To(Equal("Paragraph text max length 512 characters."))
        })
	g.It("Must have File Information",func() {
            Expect(spdx.FileInformation.ComputerFileNameText).To(Equal("filename"))
            Expect(spdx.FileInformation.FileExtensionText).To(Equal(".xml"))
        })
	g.It("Must have Authoritative Source Information",func() {
            Expect(spdx.AuthoritativeSourceInformation.SourceHashText).To(Equal("49FE985C79ACECDAC8AC147A88E872FF4E134650367A9D7FC1EFCBAD8C28B47C"))
            Expect(spdx.AuthoritativeSourceInformation.SourceURIText).To(Equal("good:URI"))
            Expect(spdx.AuthoritativeSourceInformation.AuthoritativeDigitalSignatureIndicator).To(Equal("true"))
        })
	g.It("Must have Ecosystem Information",func() {
            Expect(spdx.EcosystemInformation.Ecosystem.CommitterQuantity).To(Equal("100"))
            Expect(spdx.EcosystemInformation.Ecosystem.Languages.LanguageText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.EcosystemInformation.Ecosystem.CommitLogs.CommitLogText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.EcosystemInformation.CompanyInformation.NameText).To(Equal("Line text max length 48 characters."))
        })
	g.It("Must have Dependency Information",func() {
            Expect(spdx.DependencyInformation.Compiler.NameText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.DependencyInformation.Compiler.CompilerVersionText).To(Equal("Line text max length 48 characters."))
        })
	g.It("Must have License Information",func() {
            Expect(spdx.LicenseInformation.LicenseCategoryCode).To(Equal("FOSS"))
            Expect(spdx.LicenseInformation.LicenseCode[0]).To(Equal("LGPL"))
            Expect(spdx.LicenseInformation.EndOfLifeIndicator).To(Equal("true"))
        })
	g.It("Must have Vulnerability Information",func() {
            Expect(spdx.VulnerabilityInformation.VirusInformation.KnownVirusQuantity).To(Equal("100"))
            Expect(spdx.VulnerabilityInformation.VirusInformation.EngineVersionText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.VulnerabilityInformation.VirusInformation.VirusEngineText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.VulnerabilityInformation.VirusInformation.ScannedDirectoryQuantity).To(Equal("100"))
            Expect(spdx.VulnerabilityInformation.VirusInformation.ScannedFileQuantity).To(Equal("100"))
            Expect(spdx.VulnerabilityInformation.VirusInformation.InfectedFileQuantity).To(Equal("100"))
            Expect(spdx.VulnerabilityInformation.VirusInformation.DataScannedText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.VulnerabilityInformation.VirusInformation.DataReadText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.VulnerabilityInformation.VirusInformation.TimeText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.VulnerabilityInformation.VirusInformation.FileNotesText).To(Equal("Paragraph text max length 512 characters."))
            Expect(spdx.VulnerabilityInformation.VirusInformation.EngineDetails.EngineVersionText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.VulnerabilityInformation.VirusInformation.EngineDetails.DatabaseVersionText).To(Equal("Line text max length 48 characters."))
        })
	g.It("Must have Governance Risk Compliance",func() {
            Expect(spdx.GovernanceRiskCompliance.RiskCode).To(Equal("medium"))
            Expect(spdx.GovernanceRiskCompliance.StatementOfAssuranceText).To(Equal("Paragraph text max length 512 characters."))
        })
	g.It("Must have Delivery Information",func() {
            Expect(spdx.DeliveryInformation.DeliveringOrganizationText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.DeliveryInformation.DestinationText).To(Equal("Line text max length 48 characters."))
            Expect(spdx.DeliveryInformation.HashText).To(Equal("49FE985C79ACECDAC8AC147A88E872FF4E134650367A9D7FC1EFCBAD8C28B47C"))
            Expect(spdx.DeliveryInformation.LastRegistryUpdateDate).To(Equal("2018-02-15T09:00:00"))
        })
    })

}