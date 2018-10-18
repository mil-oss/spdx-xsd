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
            Expect(spdx.SoftwareInformation.ProductTitleText).To(Equal(""))
            Expect(spdx.SoftwareInformation.SoftwareNameText).To(Equal(""))
            Expect(spdx.SoftwareInformation.SoftwareOrgText).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.SoftwareVersionText).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionIndicator).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionText).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionInformation.APIName).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionInformation.MajorVersionNumeric).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionInformation.MinorVersionNumeric).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionInformation.PatchVersionText).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionInformation.PreReleaseVersionText).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionInformation.BuildMetaText).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.SemanticVersionDate).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonSemanticVersionIndicator).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonEpochSegmentText).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonReleaseSegmentText).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonPreReleaseSegmentText).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonPostReleaseSegmentText).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonDevelopmentReleaseSegmentText).To(Equal(""))
            Expect(spdx.SoftwareInformation.VersionInformation.PythonVersionText).To(Equal(""))
            Expect(spdx.SoftwareInformation.GroupingText).To(Equal(""))
            Expect(spdx.SoftwareInformation.ReleaseNotesText).To(Equal(""))
        })
	g.It("Must have File Information",func() {
            Expect(spdx.FileInformation.ComputerFileNameText).To(Equal(""))
            Expect(spdx.FileInformation.FileExtensionText).To(Equal(""))
        })
	g.It("Must have Authoritative Source Information",func() {
            Expect(spdx.AuthoritativeSourceInformation.SourceHashText).To(Equal(""))
            Expect(spdx.AuthoritativeSourceInformation.SourceURIText).To(Equal(""))
            Expect(spdx.AuthoritativeSourceInformation.AuthoritativeDigitalSignatureIndicator).To(Equal(""))
        })
	g.It("Must have Ecosystem Information",func() {
            Expect(spdx.EcosystemInformation.Ecosystem.CommitterQuantity).To(Equal(""))
            Expect(spdx.EcosystemInformation.Ecosystem.Languages.LanguageText).To(Equal(""))
            Expect(spdx.EcosystemInformation.Ecosystem.CommitLogs.CommitLogText).To(Equal(""))
            Expect(spdx.EcosystemInformation.CompanyInformation.NameText).To(Equal(""))
        })
	g.It("Must have Dependency Information",func() {
            Expect(spdx.DependencyInformation.Compiler.NameText).To(Equal(""))
            Expect(spdx.DependencyInformation.Compiler.CompilerVersionText).To(Equal(""))
        })
	g.It("Must have License Information",func() {
            Expect(spdx.LicenseInformation.LicenseCategoryCode).To(Equal(""))
            Expect(spdx.LicenseInformation.LicenseCode[0]).To(Equal(""))
            Expect(spdx.LicenseInformation.EndOfLifeIndicator).To(Equal(""))
        })
	g.It("Must have Vulnerability Information",func() {
            Expect(spdx.VulnerabilityInformation.VirusInformation.KnownVirusQuantity).To(Equal(""))
            Expect(spdx.VulnerabilityInformation.VirusInformation.EngineVersionText).To(Equal(""))
            Expect(spdx.VulnerabilityInformation.VirusInformation.VirusEngineText).To(Equal(""))
            Expect(spdx.VulnerabilityInformation.VirusInformation.ScannedDirectoryQuantity).To(Equal(""))
            Expect(spdx.VulnerabilityInformation.VirusInformation.ScannedFileQuantity).To(Equal(""))
            Expect(spdx.VulnerabilityInformation.VirusInformation.InfectedFileQuantity).To(Equal(""))
            Expect(spdx.VulnerabilityInformation.VirusInformation.DataScannedText).To(Equal(""))
            Expect(spdx.VulnerabilityInformation.VirusInformation.DataReadText).To(Equal(""))
            Expect(spdx.VulnerabilityInformation.VirusInformation.TimeText).To(Equal(""))
            Expect(spdx.VulnerabilityInformation.VirusInformation.FileNotesText).To(Equal(""))
            Expect(spdx.VulnerabilityInformation.VirusInformation.EngineDetails.EngineVersionText).To(Equal(""))
            Expect(spdx.VulnerabilityInformation.VirusInformation.EngineDetails.DatabaseVersionText).To(Equal(""))
        })
	g.It("Must have Governance Risk Compliance",func() {
            Expect(spdx.GovernanceRiskCompliance.RiskCode).To(Equal(""))
            Expect(spdx.GovernanceRiskCompliance.StatementOfAssuranceText).To(Equal(""))
        })
	g.It("Must have Delivery Information",func() {
            Expect(spdx.DeliveryInformation.DeliveringOrganizationText).To(Equal(""))
            Expect(spdx.DeliveryInformation.DestinationText).To(Equal(""))
            Expect(spdx.DeliveryInformation.HashText).To(Equal(""))
            Expect(spdx.DeliveryInformation.LastRegistryUpdateDate).To(Equal(""))
        })
    })

}