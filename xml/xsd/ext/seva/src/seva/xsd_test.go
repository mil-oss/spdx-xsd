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
	"test_data.xml":     "xml/test_data.xml",
	"test_instance.xml": "xml/test_instance.xml",
}

func Testsecurity(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	xf, ferr := ioutil.ReadFile(testinstances["test_instance.xml"])
	if ferr != nil {
		fmt.Printf(ferr.Error())
	}
	var security = NewSoftwareEvidenceArchive()
	err := xml.Unmarshal([]byte(xf), &security)
	if err != nil {
		fmt.Printf(err.Error())
	}
	g.Describe("security", func() {
		g.It("Must have Software Information", func() {
			Expect(security.SoftwareInformation.ProductTitleText).To(Equal("Line text max length 48 characters."))
			Expect(security.SoftwareInformation.SoftwareNameText).To(Equal("Line text max length 48 characters."))
			Expect(security.SoftwareInformation.SoftwareOrgText).To(Equal("Line text max length 48 characters."))
			Expect(security.SoftwareInformation.VersionInformation.SoftwareVersionText).To(Equal("Line text max length 48 characters."))
			Expect(security.SoftwareInformation.VersionInformation.SemanticVersionIndicator).To(Equal("true"))
			Expect(security.SoftwareInformation.VersionInformation.SemanticVersionText).To(Equal("1.0.0-alpha.beta"))
			Expect(security.SoftwareInformation.VersionInformation.SemanticVersionInformation.APIName).To(Equal("APIname"))
			Expect(security.SoftwareInformation.VersionInformation.SemanticVersionInformation.MajorVersionNumeric).To(Equal("1"))
			Expect(security.SoftwareInformation.VersionInformation.SemanticVersionInformation.MinorVersionNumeric).To(Equal("1"))
			Expect(security.SoftwareInformation.VersionInformation.SemanticVersionInformation.PatchVersionText).To(Equal("1"))
			Expect(security.SoftwareInformation.VersionInformation.SemanticVersionInformation.PreReleaseVersionText).To(Equal("-alpha"))
			Expect(security.SoftwareInformation.VersionInformation.SemanticVersionInformation.BuildMetaText).To(Equal("+exp.sha.5114f85"))
			Expect(security.SoftwareInformation.VersionInformation.SemanticVersionDate).To(Equal("2018-02-15T09:00:00"))
			Expect(security.SoftwareInformation.VersionInformation.PythonSemanticVersionIndicator).To(Equal("true"))
			Expect(security.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonEpochSegmentText).To(Equal("1"))
			Expect(security.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonReleaseSegmentText).To(Equal("0"))
			Expect(security.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonPreReleaseSegmentText).To(Equal("3"))
			Expect(security.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonPostReleaseSegmentText).To(Equal("2"))
			Expect(security.SoftwareInformation.VersionInformation.PythonSemanticVersion.PythonDevelopmentReleaseSegmentText).To(Equal("3"))
			Expect(security.SoftwareInformation.VersionInformation.PythonVersionText).To(Equal("1.0.3.2.3"))
			Expect(security.SoftwareInformation.GroupingText).To(Equal("Line text max length 48 characters."))
			Expect(security.SoftwareInformation.ReleaseNotesText).To(Equal("Paragraph text max length 512 characters."))
		})
		g.It("Must have File Information", func() {
			Expect(security.FileInformation.FileNameText).To(Equal("FileNameNoExtension"))
			Expect(security.FileInformation.FileExtensionText).To(Equal(".ext"))
		})
		g.It("Must have Authoritative Source Information", func() {
			Expect(security.AuthoritativeSourceInformation.SourceHashText).To(Equal("49FE985C79ACECDAC8AC147A88E872FF4E134650367A9D7FC1EFCBAD8C28B47C"))
			Expect(security.AuthoritativeSourceInformation.SourceURIText).To(Equal("good:URI"))
			Expect(security.AuthoritativeSourceInformation.AuthoritativeDigitalSignatureIndicator).To(Equal("true"))
		})
		g.It("Must have Ecosystem Information", func() {
			Expect(security.EcosystemInformation.Ecosystem.CommitterQuantity).To(Equal("100"))
			Expect(security.EcosystemInformation.Ecosystem.Languages.LanguageText[0]).To(Equal("Line text max length 48 characters."))
			Expect(security.EcosystemInformation.Ecosystem.CommitLogs.CommitLogText[0]).To(Equal("Line text max length 48 characters."))
			Expect(security.EcosystemInformation.CompanyInformation.Name).To(Equal("Line text max length 48 characters."))
		})
		g.It("Must have Dependency Information", func() {
			Expect(security.DependencyInformation.Dependencies.Dependency[0].Name).To(Equal("Line text max length 48 characters."))
			Expect(security.DependencyInformation.Dependencies.Dependency[0].DependencyVersionText).To(Equal("Line text max length 48 characters."))
			Expect(security.DependencyInformation.Compiler.Name).To(Equal("Line text max length 48 characters."))
			Expect(security.DependencyInformation.Compiler.CompilerVersionText).To(Equal("Line text max length 48 characters."))
		})
		g.It("Must have License Information", func() {
			Expect(security.LicenseInformation.LicenseCategoryCode).To(Equal("FOSS"))
			Expect(security.LicenseInformation.LicenseCode[0]).To(Equal("LGPL"))
			Expect(security.LicenseInformation.EndOfLifeIndicator).To(Equal("true"))
		})
		g.It("Must have Vulnerability Information", func() {
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityIdentifierText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].SourceText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityTitleText).To(Equal("Paragraph text max length 512 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilitySummaryText).To(Equal("Paragraph text max length 512 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityScoreValue).To(Equal("123.456"))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityScoreVersionValue).To(Equal("123.456"))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityScoreSystemText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV2.VectorStringText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV2.AccessVectorText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV2.AccessComplexityText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV2.VulnerabilityAuthenticationText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV2.ConfidentialityImpactText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV2.IntegrityImpactText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV2.AvailabilityImpactText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV2.VulnerabilityBaseScoreValue).To(Equal("123.456"))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV3.VectorStringText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV3.AccessVectorText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV3.AccessComplexityText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV3.PrivilegesRequiredText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV3.UserInteractionText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV3.VulnerabilityScopeText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV3.ConfidentialityImpactText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV3.IntegrityImpactText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV3.AvailabilityImpactText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV3.VulnerabilityBaseScoreValue).To(Equal("123.456"))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScoreDetails.CommonVulnerabilityScoringSystemV3.BaseSeverityText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityVectorText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].AccessComplexityText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityAuthenticationText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ConfidentialityImpactText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].IntegrityImpactText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].AvailabilityImpactText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilitySourceText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].AssessmentCheckText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].ScannerText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].RecommendationText).To(Equal("Paragraph text max length 512 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityReference.TypeOfReferenceText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityReference.ReferenceSourceText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityReference.URIText).To(Equal("good:URI"))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityReference.ReferenceText).To(Equal("Paragraph text max length 512 characters."))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityModifiedDate).To(Equal("2018-02-15T09:00:00"))
			Expect(security.VulnerabilityInformation.Vulnerabilities.Vulnerability[0].VulnerabilityPublishedDate).To(Equal("2018-02-15T09:00:00"))
			Expect(security.VulnerabilityInformation.VirusInformation.KnownVirusQuantity).To(Equal("100"))
			Expect(security.VulnerabilityInformation.VirusInformation.EngineVersionText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.VirusInformation.VirusEngineText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.VirusInformation.ScannedDirectoryQuantity).To(Equal("100"))
			Expect(security.VulnerabilityInformation.VirusInformation.ScannedFileQuantity).To(Equal("100"))
			Expect(security.VulnerabilityInformation.VirusInformation.InfectedFileQuantity).To(Equal("100"))
			Expect(security.VulnerabilityInformation.VirusInformation.DataScannedText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.VirusInformation.DataReadText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.VirusInformation.TimeText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.VirusInformation.FileNotesText).To(Equal("Paragraph text max length 512 characters."))
			Expect(security.VulnerabilityInformation.VirusInformation.EngineDetails.EngineVersionText).To(Equal("Line text max length 48 characters."))
			Expect(security.VulnerabilityInformation.VirusInformation.EngineDetails.DatabaseVersionText).To(Equal("Line text max length 48 characters."))
		})
		g.It("Must have Governance Risk Compliance", func() {
			Expect(security.GovernanceRiskCompliance.RiskCode).To(Equal("medium"))
			Expect(security.GovernanceRiskCompliance.StatementOfAssuranceText).To(Equal("Paragraph text max length 512 characters."))
		})
		g.It("Must have Delivery Information", func() {
			Expect(security.DeliveryInformation.DeliveringOrganizationText).To(Equal("Line text max length 48 characters."))
			Expect(security.DeliveryInformation.DestinationText).To(Equal("Line text max length 48 characters."))
			Expect(security.DeliveryInformation.HashText).To(Equal("49FE985C79ACECDAC8AC147A88E872FF4E134650367A9D7FC1EFCBAD8C28B47C"))
			Expect(security.DeliveryInformation.LastRegistryUpdateDate).To(Equal("2018-02-15T09:00:00"))
		})
	})

}
