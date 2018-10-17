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
	"test_data.xml":     "xml/instance/test_data.xml",
	"test_instance.xml": "xml/instance/test_instance.xml",
}

func TestSeva(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	xf, ferr := ioutil.ReadFile(testinstances["test_instance.xml"])
	if ferr != nil {
		fmt.Printf(ferr.Error())
	}
	var seva = NewSoftwareEvidenceArchiveISM()
	err := xml.Unmarshal([]byte(xf), &seva)
	if err != nil {
		fmt.Printf(err.Error())
	}
	g.Describe("SEVA", func() {
		g.It("Must have Software Information", func() {
			Expect(seva.SoftwareInformation.ProductTitleText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.SoftwareInformation.SoftwareNameText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.SoftwareInformation.SoftwareOrgText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.SoftwareInformation.VersionInformation.SoftwareVersionText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.SoftwareInformation.VersionInformation.SemanticVersionIndicator.Data).To(Equal("true"))
			Expect(seva.SoftwareInformation.VersionInformation.SemanticVersionText.Data).To(Equal("1.0.0-alpha.beta"))
			Expect(seva.SoftwareInformation.VersionInformation.SemanticVersionInformation.APIName.Data).To(Equal("APIname"))
			Expect(seva.SoftwareInformation.VersionInformation.SemanticVersionInformation.MajorVersionNumeric.Data).To(Equal("1"))
			Expect(seva.SoftwareInformation.VersionInformation.SemanticVersionInformation.MinorVersionNumeric.Data).To(Equal("1"))
			Expect(seva.SoftwareInformation.VersionInformation.SemanticVersionInformation.PatchVersionText.Data).To(Equal("1"))
			Expect(seva.SoftwareInformation.VersionInformation.SemanticVersionInformation.PreReleaseVersionText.Data).To(Equal("-alpha"))
			Expect(seva.SoftwareInformation.VersionInformation.SemanticVersionInformation.BuildMetaText.Data).To(Equal("+exp.sha.5114f85"))
			Expect(seva.SoftwareInformation.VersionInformation.SemanticVersionDate.Data).To(Equal("2018-02-15T09:00:00"))
			Expect(seva.SoftwareInformation.VersionInformation.PythonSemanticVersionIndicator.Data).To(Equal("true"))
			Expect(seva.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonEpochSegmentText.Data).To(Equal("1"))
			Expect(seva.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonReleaseSegmentText.Data).To(Equal("0"))
			Expect(seva.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonPreReleaseSegmentText.Data).To(Equal("3"))
			Expect(seva.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonPostReleaseSegmentText.Data).To(Equal("2"))
			Expect(seva.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonDevelopmentReleaseSegmentText.Data).To(Equal("3"))
			Expect(seva.SoftwareInformation.VersionInformation.PythonVersionText.Data).To(Equal("1.0.3.2.3"))
			Expect(seva.SoftwareInformation.GroupingText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.SoftwareInformation.ReleaseNotesText.Data).To(Equal("Paragraph text max length 512 characters."))
		})
		g.It("Must have File Information", func() {
			Expect(seva.FileInformation.FileNameText.Data).To(Equal("FileNameNoExtension"))
			Expect(seva.FileInformation.FileExtensionText.Data).To(Equal(".ext"))
		})
		g.It("Must have Authoritative Source Information", func() {
			Expect(seva.AuthoritativeSourceInformation.SourceHashText.Data).To(Equal("49FE985C79ACECDAC8AC147A88E872FF4E134650367A9D7FC1EFCBAD8C28B47C"))
			Expect(seva.AuthoritativeSourceInformation.SourceURIText.Data).To(Equal("good:URI"))
			Expect(seva.AuthoritativeSourceInformation.AuthoritativeDigitalSignatureIndicator.Data).To(Equal("true"))
		})
		g.It("Must have Ecosystem Information", func() {
			Expect(seva.EcosystemInformation.Ecosystem.CommitterQuantity.Data).To(Equal("100"))
			Expect(seva.EcosystemInformation.Ecosystem.Languages.LanguageText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.EcosystemInformation.Ecosystem.CommitLogs.CommitLogText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.EcosystemInformation.CompanyInformation.Name.Data).To(Equal("Line text max length 48 characters."))
		})
		g.It("Must have Dependency Information", func() {
			Expect(seva.DependencyInformation.Dependencies.Dependency.Name.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.DependencyInformation.Dependencies.Dependency.DependencyVersionText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.DependencyInformation.Compiler.Name.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.DependencyInformation.Compiler.CompilerVersionText.Data).To(Equal("Line text max length 48 characters."))
		})
		g.It("Must have License Information", func() {
			Expect(seva.LicenseInformation.LicenseCategoryCode.Data).To(Equal("FOSS"))
			Expect(seva.LicenseInformation.LicenseCode.Data).To(Equal("LGPL"))
			Expect(seva.LicenseInformation.EndOfLifeIndicator.Data).To(Equal("true"))
		})
		g.It("Must have Vulnerability Information", func() {
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityIdentifierText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.SourceText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityTitleText.Data).To(Equal("Paragraph text max length 512 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilitySummaryText.Data).To(Equal("Paragraph text max length 512 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityScoreValue.Data).To(Equal("123.456"))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityScoreVersionValue.Data).To(Equal("123.456"))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityScoreSystemText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV2.VectorStringText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV2.AccessVectorText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV2.AccessComplexityText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV2.VulnerabilityAuthenticationText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV2.ConfidentialityImpactText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV2.IntegrityImpactText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV2.AvailabilityImpactText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV2.VulnerabilityBaseScoreValue.Data).To(Equal("123.456"))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV3.VectorStringText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV3.AccessVectorText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV3.AccessComplexityText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV3.PrivilegesRequiredText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV3.UserInteractionText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV3.VulnerabilityScopeText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV3.ConfidentialityImpactText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV3.IntegrityImpactText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV3.AvailabilityImpactText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV3.VulnerabilityBaseScoreValue.Data).To(Equal("123.456"))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScoreDetails.CommonVulnerabilityScoringSystemV3.BaseSeverityText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityVectorText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.AccessComplexityText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityAuthenticationText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ConfidentialityImpactText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.IntegrityImpactText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.AvailabilityImpactText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilitySourceText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.AssessmentCheckText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.ScannerText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.RecommendationText.Data).To(Equal("Paragraph text max length 512 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityReference.TypeOfReferenceText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityReference.ReferenceSourceText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityReference.URIText.Data).To(Equal("good:URI"))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityReference.ReferenceText.Data).To(Equal("Paragraph text max length 512 characters."))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityModifiedDate.Data).To(Equal("2018-02-15T09:00:00"))
			Expect(seva.VulnerabilityInformation.Vulnerabilities.Vulnerability.VulnerabilityPublishedDate.Data).To(Equal("2018-02-15T09:00:00"))
			Expect(seva.VulnerabilityInformation.VirusInformation.KnownVirusQuantity.Data).To(Equal("100"))
			Expect(seva.VulnerabilityInformation.VirusInformation.EngineVersionText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.VirusInformation.VirusEngineText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.VirusInformation.ScannedDirectoryQuantity.Data).To(Equal("100"))
			Expect(seva.VulnerabilityInformation.VirusInformation.ScannedFileQuantity.Data).To(Equal("100"))
			Expect(seva.VulnerabilityInformation.VirusInformation.InfectedFileQuantity.Data).To(Equal("100"))
			Expect(seva.VulnerabilityInformation.VirusInformation.DataScannedText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.VirusInformation.DataReadText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.VirusInformation.TimeText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.VirusInformation.FileNotesText.Data).To(Equal("Paragraph text max length 512 characters."))
			Expect(seva.VulnerabilityInformation.VirusInformation.EngineDetails.EngineVersionText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.VulnerabilityInformation.VirusInformation.EngineDetails.DatabaseVersionText.Data).To(Equal("Line text max length 48 characters."))
		})
		g.It("Must have Governance Risk Compliance", func() {
			Expect(seva.GovernanceRiskCompliance.RiskCode.Data).To(Equal("medium"))
			Expect(seva.GovernanceRiskCompliance.StatementOfAssuranceText.Data).To(Equal("Paragraph text max length 512 characters."))
		})
		g.It("Must have Delivery Information", func() {
			Expect(seva.DeliveryInformation.DeliveringOrganizationText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.DeliveryInformation.DestinationText.Data).To(Equal("Line text max length 48 characters."))
			Expect(seva.DeliveryInformation.HashText.Data).To(Equal("49FE985C79ACECDAC8AC147A88E872FF4E134650367A9D7FC1EFCBAD8C28B47C"))
			Expect(seva.DeliveryInformation.LastRegistryUpdateDate.Data).To(Equal("2018-02-15T09:00:00"))
		})
	})

}
