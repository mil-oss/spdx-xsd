package main

import "encoding/xml"

//NewSoftwareEvidenceArchive ...
func NewSoftwareEvidenceArchive() *SoftwareEvidenceArchive {
	return &SoftwareEvidenceArchive{
		// Required for the proper namespacing
		AttrXmlnsXsi: "http://www.w3.org/2001/XMLSchema-instance",
		AttrXmlns:    "spdx:xsd::1.0",
	}
}

//SoftwareEvidenceArchive ... A data item for automated software supply chain metadata
type SoftwareEvidenceArchive struct {
	AttrXmlnsXsi                   string                          `xml:"xmlns:xsi,attr" json:"AttrXmlnsXsi,omitempty"`
	AttrXmlns                      string                          `xml:"xmlns,attr" json:"AttrXmlns,omitempty"`
	SoftwareInformation            *SoftwareInformation            `xml:"SoftwareInformation,omitempty"  json:"SoftwareInformation,omitempty"`
	FileInformation                *FileInformation                `xml:"FileInformation,omitempty"  json:"FileInformation,omitempty"`
	AuthoritativeSourceInformation *AuthoritativeSourceInformation `xml:"AuthoritativeSourceInformation,omitempty"  json:"AuthoritativeSourceInformation,omitempty"`
	EcosystemInformation           *EcosystemInformation           `xml:"EcosystemInformation,omitempty"  json:"EcosystemInformation,omitempty"`
	DependencyInformation          *DependencyInformation          `xml:"DependencyInformation,omitempty"  json:"DependencyInformation,omitempty"`
	LicenseInformation             *LicenseInformation             `xml:"LicenseInformation,omitempty"  json:"LicenseInformation,omitempty"`
	VulnerabilityInformation       *VulnerabilityInformation       `xml:"VulnerabilityInformation,omitempty"  json:"VulnerabilityInformation,omitempty"`
	GovernanceRiskCompliance       *GovernanceRiskCompliance       `xml:"GovernanceRiskCompliance,omitempty"  json:"GovernanceRiskCompliance,omitempty"`
	DeliveryInformation            *DeliveryInformation            `xml:"DeliveryInformation,omitempty"  json:"DeliveryInformation,omitempty"`
	XMLName                        xml.Name                        `xml:"SoftwareEvidenceArchive,omitempty"  json:"SoftwareEvidenceArchive,omitempty"`
}

//Annotation ...
type Annotation struct {
	Date               string   `xml:"Date,omitempty"  json:"Date,omitempty"`
	AnnotationTypeCode string   `xml:"AnnotationTypeCode,omitempty"  json:"AnnotationTypeCode,omitempty"`
	CommentText        string   `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	AnnotatorText      string   `xml:"AnnotatorText,omitempty"  json:"AnnotatorText,omitempty"`
	XMLName            xml.Name `xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
}

//AuthoritativeSourceInformation ... A data item for information about what is to be considered the authoritative source for a given artifact or source repository
type AuthoritativeSourceInformation struct {
	SourceHashText                         string   `xml:"SourceHashText,omitempty"  json:"SourceHashText,omitempty"`
	SourceURIText                          string   `xml:"SourceURIText,omitempty"  json:"SourceURIText,omitempty"`
	AuthoritativeDigitalSignatureIndicator string   `xml:"AuthoritativeDigitalSignatureIndicator,omitempty"  json:"AuthoritativeDigitalSignatureIndicator,omitempty"`
	XMLName                                xml.Name `xml:"AuthoritativeSourceInformation,omitempty"  json:"AuthoritativeSourceInformation,omitempty"`
}

//CommitLogs ...
type CommitLogs struct {
	CommitLogText []string `xml:"CommitLogText,omitempty"  json:"CommitLogText,omitempty"`
	XMLName       xml.Name `xml:"CommitLogs,omitempty"  json:"CommitLogs,omitempty"`
}

//CommonVulnerabilityScoringSystemV2 ...
type CommonVulnerabilityScoringSystemV2 struct {
	VectorStringText                string   `xml:"VectorStringText,omitempty"  json:"VectorStringText,omitempty"`
	AccessVectorText                string   `xml:"AccessVectorText,omitempty"  json:"AccessVectorText,omitempty"`
	AccessComplexityText            string   `xml:"AccessComplexityText,omitempty"  json:"AccessComplexityText,omitempty"`
	VulnerabilityAuthenticationText string   `xml:"VulnerabilityAuthenticationText,omitempty"  json:"VulnerabilityAuthenticationText,omitempty"`
	ConfidentialityImpactText       string   `xml:"ConfidentialityImpactText,omitempty"  json:"ConfidentialityImpactText,omitempty"`
	IntegrityImpactText             string   `xml:"IntegrityImpactText,omitempty"  json:"IntegrityImpactText,omitempty"`
	AvailabilityImpactText          string   `xml:"AvailabilityImpactText,omitempty"  json:"AvailabilityImpactText,omitempty"`
	VulnerabilityBaseScoreValue     string   `xml:"VulnerabilityBaseScoreValue,omitempty"  json:"VulnerabilityBaseScoreValue,omitempty"`
	XMLName                         xml.Name `xml:"CommonVulnerabilityScoringSystemV2,omitempty"  json:"CommonVulnerabilityScoringSystemV2,omitempty"`
}

//CommonVulnerabilityScoringSystemV3 ...
type CommonVulnerabilityScoringSystemV3 struct {
	VectorStringText            string   `xml:"VectorStringText,omitempty"  json:"VectorStringText,omitempty"`
	AccessVectorText            string   `xml:"AccessVectorText,omitempty"  json:"AccessVectorText,omitempty"`
	AccessComplexityText        string   `xml:"AccessComplexityText,omitempty"  json:"AccessComplexityText,omitempty"`
	PrivilegesRequiredText      string   `xml:"PrivilegesRequiredText,omitempty"  json:"PrivilegesRequiredText,omitempty"`
	UserInteractionText         string   `xml:"UserInteractionText,omitempty"  json:"UserInteractionText,omitempty"`
	VulnerabilityScopeText      string   `xml:"VulnerabilityScopeText,omitempty"  json:"VulnerabilityScopeText,omitempty"`
	ConfidentialityImpactText   string   `xml:"ConfidentialityImpactText,omitempty"  json:"ConfidentialityImpactText,omitempty"`
	IntegrityImpactText         string   `xml:"IntegrityImpactText,omitempty"  json:"IntegrityImpactText,omitempty"`
	AvailabilityImpactText      string   `xml:"AvailabilityImpactText,omitempty"  json:"AvailabilityImpactText,omitempty"`
	VulnerabilityBaseScoreValue string   `xml:"VulnerabilityBaseScoreValue,omitempty"  json:"VulnerabilityBaseScoreValue,omitempty"`
	BaseSeverityText            string   `xml:"BaseSeverityText,omitempty"  json:"BaseSeverityText,omitempty"`
	XMLName                     xml.Name `xml:"CommonVulnerabilityScoringSystemV3,omitempty"  json:"CommonVulnerabilityScoringSystemV3,omitempty"`
}

//CompanyInformation ...
type CompanyInformation struct {
	NameText string   `xml:"NameText,omitempty"  json:"NameText,omitempty"`
	XMLName  xml.Name `xml:"CompanyInformation,omitempty"  json:"CompanyInformation,omitempty"`
}

//Compiler ...
type Compiler struct {
	NameText            string   `xml:"NameText,omitempty"  json:"NameText,omitempty"`
	CompilerVersionText string   `xml:"CompilerVersionText,omitempty"  json:"CompilerVersionText,omitempty"`
	XMLName             xml.Name `xml:"Compiler,omitempty"  json:"Compiler,omitempty"`
}

//DeliveryInformation ... A data item for delivery information including the target url and last delivered date and time for a previously signed Software Evidence Archive
type DeliveryInformation struct {
	DeliveringOrganizationText string   `xml:"DeliveringOrganizationText,omitempty"  json:"DeliveringOrganizationText,omitempty"`
	DestinationText            string   `xml:"DestinationText,omitempty"  json:"DestinationText,omitempty"`
	HashText                   string   `xml:"HashText,omitempty"  json:"HashText,omitempty"`
	LastRegistryUpdateDate     string   `xml:"LastRegistryUpdateDate,omitempty"  json:"LastRegistryUpdateDate,omitempty"`
	XMLName                    xml.Name `xml:"DeliveryInformation,omitempty"  json:"DeliveryInformation,omitempty"`
}

//Dependencies ...
type Dependencies struct {
	Dependency []*Dependency `xml:"Dependency,omitempty"  json:"Dependency,omitempty"`
	XMLName    xml.Name      `xml:"Dependencies,omitempty"  json:"Dependencies,omitempty"`
}

//Dependency ...
type Dependency struct {
	NameText              string   `xml:"NameText,omitempty"  json:"NameText,omitempty"`
	DependencyVersionText string   `xml:"DependencyVersionText,omitempty"  json:"DependencyVersionText,omitempty"`
	XMLName               xml.Name `xml:"Dependency,omitempty"  json:"Dependency,omitempty"`
}

//DependencyInformation ... A data item for a list of dependencies for a given software project derived directly from the artifact or source dependency definition file
type DependencyInformation struct {
	Dependencies *Dependencies `xml:"Dependencies,omitempty"  json:"Dependencies,omitempty"`
	Compiler     *Compiler     `xml:"Compiler,omitempty"  json:"Compiler,omitempty"`
	XMLName      xml.Name      `xml:"DependencyInformation,omitempty"  json:"DependencyInformation,omitempty"`
}

//Ecosystem ...
type Ecosystem struct {
	CommitterQuantity string      `xml:"CommitterQuantity,omitempty"  json:"CommitterQuantity,omitempty"`
	Languages         *Languages  `xml:"Languages,omitempty"  json:"Languages,omitempty"`
	CommitLogs        *CommitLogs `xml:"CommitLogs,omitempty"  json:"CommitLogs,omitempty"`
	XMLName           xml.Name    `xml:"Ecosystem,omitempty"  json:"Ecosystem,omitempty"`
}

//EcosystemInformation ...
type EcosystemInformation struct {
	Ecosystem          *Ecosystem          `xml:"Ecosystem,omitempty"  json:"Ecosystem,omitempty"`
	CompanyInformation *CompanyInformation `xml:"CompanyInformation,omitempty"  json:"CompanyInformation,omitempty"`
	XMLName            xml.Name            `xml:"EcosystemInformation,omitempty"  json:"EcosystemInformation,omitempty"`
}

//EngineDetails ...
type EngineDetails struct {
	EngineVersionText   string   `xml:"EngineVersionText,omitempty"  json:"EngineVersionText,omitempty"`
	DatabaseVersionText string   `xml:"DatabaseVersionText,omitempty"  json:"DatabaseVersionText,omitempty"`
	XMLName             xml.Name `xml:"EngineDetails,omitempty"  json:"EngineDetails,omitempty"`
}

//FileInformation ...
type FileInformation struct {
	ComputerFileNameText string   `xml:"ComputerFileNameText,omitempty"  json:"ComputerFileNameText,omitempty"`
	FileExtensionText    string   `xml:"FileExtensionText,omitempty"  json:"FileExtensionText,omitempty"`
	XMLName              xml.Name `xml:"FileInformation,omitempty"  json:"FileInformation,omitempty"`
}

//GovernanceRiskCompliance ...
type GovernanceRiskCompliance struct {
	RiskCode                 string   `xml:"RiskCode,omitempty"  json:"RiskCode,omitempty"`
	StatementOfAssuranceText string   `xml:"StatementOfAssuranceText,omitempty"  json:"StatementOfAssuranceText,omitempty"`
	XMLName                  xml.Name `xml:"GovernanceRiskCompliance,omitempty"  json:"GovernanceRiskCompliance,omitempty"`
}

//Languages ...
type Languages struct {
	LanguageText []string `xml:"LanguageText,omitempty"  json:"LanguageText,omitempty"`
	XMLName      xml.Name `xml:"Languages,omitempty"  json:"Languages,omitempty"`
}

//LicenseInformation ...
type LicenseInformation struct {
	LicenseCategoryCode string   `xml:"LicenseCategoryCode,omitempty"  json:"LicenseCategoryCode,omitempty"`
	LicenseCode         []string `xml:"LicenseCode,omitempty"  json:"LicenseCode,omitempty"`
	EndOfLifeIndicator  string   `xml:"EndOfLifeIndicator,omitempty"  json:"EndOfLifeIndicator,omitempty"`
	XMLName             xml.Name `xml:"LicenseInformation,omitempty"  json:"LicenseInformation,omitempty"`
}

//PythonSemanticVersion ... A data item for a Python Distribution identified by a public version identifier which supports all defined version comparison operations
type PythonSemanticVersion struct {
	PythonEpochSegmentText              string   `xml:"PythonEpochSegmentText,omitempty"  json:"PythonEpochSegmentText,omitempty"`
	PythonReleaseSegmentText            string   `xml:"PythonReleaseSegmentText,omitempty"  json:"PythonReleaseSegmentText,omitempty"`
	PythonPreReleaseSegmentText         string   `xml:"PythonPreReleaseSegmentText,omitempty"  json:"PythonPreReleaseSegmentText,omitempty"`
	PythonPostReleaseSegmentText        string   `xml:"PythonPostReleaseSegmentText,omitempty"  json:"PythonPostReleaseSegmentText,omitempty"`
	PythonDevelopmentReleaseSegmentText string   `xml:"PythonDevelopmentReleaseSegmentText,omitempty"  json:"PythonDevelopmentReleaseSegmentText,omitempty"`
	XMLName                             xml.Name `xml:"PythonSemanticVersion,omitempty"  json:"PythonSemanticVersion,omitempty"`
}

//Relationship ...
type Relationship struct {
	RelationshipTypeCode string        `xml:"RelationshipTypeCode,omitempty"  json:"RelationshipTypeCode,omitempty"`
	CommentText          string        `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	RelatedSpdxElement   *Relationship `xml:"RelatedSpdxElement,omitempty"  json:"RelatedSpdxElement,omitempty"`
	XMLName              xml.Name      `xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
}

//ScoreDetails ...
type ScoreDetails struct {
	CommonVulnerabilityScoringSystemV2 *CommonVulnerabilityScoringSystemV2 `xml:"CommonVulnerabilityScoringSystemV2,omitempty"  json:"CommonVulnerabilityScoringSystemV2,omitempty"`
	CommonVulnerabilityScoringSystemV3 *CommonVulnerabilityScoringSystemV3 `xml:"CommonVulnerabilityScoringSystemV3,omitempty"  json:"CommonVulnerabilityScoringSystemV3,omitempty"`
	XMLName                            xml.Name                            `xml:"ScoreDetails,omitempty"  json:"ScoreDetails,omitempty"`
}

//SemanticVersionInformation ... A data type for semantic versioning information
type SemanticVersionInformation struct {
	APIName               string   `xml:"APIName,omitempty"  json:"APIName,omitempty"`
	MajorVersionNumeric   string   `xml:"MajorVersionNumeric,omitempty"  json:"MajorVersionNumeric,omitempty"`
	MinorVersionNumeric   string   `xml:"MinorVersionNumeric,omitempty"  json:"MinorVersionNumeric,omitempty"`
	PatchVersionText      string   `xml:"PatchVersionText,omitempty"  json:"PatchVersionText,omitempty"`
	PreReleaseVersionText string   `xml:"PreReleaseVersionText,omitempty"  json:"PreReleaseVersionText,omitempty"`
	BuildMetaText         string   `xml:"BuildMetaText,omitempty"  json:"BuildMetaText,omitempty"`
	XMLName               xml.Name `xml:"SemanticVersionInformation,omitempty"  json:"SemanticVersionInformation,omitempty"`
}

//SoftwareInformation ...
type SoftwareInformation struct {
	ProductTitleText   string              `xml:"ProductTitleText,omitempty"  json:"ProductTitleText,omitempty"`
	SoftwareNameText   string              `xml:"SoftwareNameText,omitempty"  json:"SoftwareNameText,omitempty"`
	SoftwareOrgText    string              `xml:"SoftwareOrgText,omitempty"  json:"SoftwareOrgText,omitempty"`
	VersionInformation *VersionInformation `xml:"VersionInformation,omitempty"  json:"VersionInformation,omitempty"`
	GroupingText       string              `xml:"GroupingText,omitempty"  json:"GroupingText,omitempty"`
	ReleaseNotesText   string              `xml:"ReleaseNotesText,omitempty"  json:"ReleaseNotesText,omitempty"`
	XMLName            xml.Name            `xml:"SoftwareInformation,omitempty"  json:"SoftwareInformation,omitempty"`
}

//SpdxElement ...
type SpdxElement struct {
	Annotation   *Annotation   `xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
	Name         string        `xml:"Name,omitempty"  json:"Name,omitempty"`
	CommentText  string        `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	Relationship *Relationship `xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
	XMLName      xml.Name      `xml:"SpdxElement,omitempty"  json:"SpdxElement,omitempty"`
}

//VersionInformation ...
type VersionInformation struct {
	SoftwareVersionText            string                      `xml:"SoftwareVersionText,omitempty"  json:"SoftwareVersionText,omitempty"`
	SemanticVersionIndicator       string                      `xml:"SemanticVersionIndicator,omitempty"  json:"SemanticVersionIndicator,omitempty"`
	SemanticVersionText            string                      `xml:"SemanticVersionText,omitempty"  json:"SemanticVersionText,omitempty"`
	SemanticVersionInformation     *SemanticVersionInformation `xml:"SemanticVersionInformation,omitempty"  json:"SemanticVersionInformation,omitempty"`
	SemanticVersionDate            string                      `xml:"SemanticVersionDate,omitempty"  json:"SemanticVersionDate,omitempty"`
	PythonSemanticVersionIndicator string                      `xml:"PythonSemanticVersionIndicator,omitempty"  json:"PythonSemanticVersionIndicator,omitempty"`
	PythonSemanticVersion          *PythonSemanticVersion      `xml:"PythonSemanticVersion,omitempty"  json:"PythonSemanticVersion,omitempty"`
	PythonVersionText              string                      `xml:"PythonVersionText,omitempty"  json:"PythonVersionText,omitempty"`
	XMLName                        xml.Name                    `xml:"VersionInformation,omitempty"  json:"VersionInformation,omitempty"`
}

//VirusInformation ...
type VirusInformation struct {
	KnownVirusQuantity       string         `xml:"KnownVirusQuantity,omitempty"  json:"KnownVirusQuantity,omitempty"`
	EngineVersionText        string         `xml:"EngineVersionText,omitempty"  json:"EngineVersionText,omitempty"`
	VirusEngineText          string         `xml:"VirusEngineText,omitempty"  json:"VirusEngineText,omitempty"`
	ScannedDirectoryQuantity string         `xml:"ScannedDirectoryQuantity,omitempty"  json:"ScannedDirectoryQuantity,omitempty"`
	ScannedFileQuantity      string         `xml:"ScannedFileQuantity,omitempty"  json:"ScannedFileQuantity,omitempty"`
	InfectedFileQuantity     string         `xml:"InfectedFileQuantity,omitempty"  json:"InfectedFileQuantity,omitempty"`
	DataScannedText          string         `xml:"DataScannedText,omitempty"  json:"DataScannedText,omitempty"`
	DataReadText             string         `xml:"DataReadText,omitempty"  json:"DataReadText,omitempty"`
	TimeText                 string         `xml:"TimeText,omitempty"  json:"TimeText,omitempty"`
	FileNotesText            string         `xml:"FileNotesText,omitempty"  json:"FileNotesText,omitempty"`
	EngineDetails            *EngineDetails `xml:"EngineDetails,omitempty"  json:"EngineDetails,omitempty"`
	XMLName                  xml.Name       `xml:"VirusInformation,omitempty"  json:"VirusInformation,omitempty"`
}

//Vulnerabilities ...
type Vulnerabilities struct {
	Vulnerability []*Vulnerability `xml:"Vulnerability,omitempty"  json:"Vulnerability,omitempty"`
	XMLName       xml.Name         `xml:"Vulnerabilities,omitempty"  json:"Vulnerabilities,omitempty"`
}

//Vulnerability ...
type Vulnerability struct {
	VulnerabilityIdentifierText     string                  `xml:"VulnerabilityIdentifierText,omitempty"  json:"VulnerabilityIdentifierText,omitempty"`
	SourceText                      string                  `xml:"SourceText,omitempty"  json:"SourceText,omitempty"`
	VulnerabilityTitleText          string                  `xml:"VulnerabilityTitleText,omitempty"  json:"VulnerabilityTitleText,omitempty"`
	VulnerabilitySummaryText        string                  `xml:"VulnerabilitySummaryText,omitempty"  json:"VulnerabilitySummaryText,omitempty"`
	VulnerabilityScoreValue         string                  `xml:"VulnerabilityScoreValue,omitempty"  json:"VulnerabilityScoreValue,omitempty"`
	VulnerabilityScoreVersionValue  string                  `xml:"VulnerabilityScoreVersionValue,omitempty"  json:"VulnerabilityScoreVersionValue,omitempty"`
	VulnerabilityScoreSystemText    string                  `xml:"VulnerabilityScoreSystemText,omitempty"  json:"VulnerabilityScoreSystemText,omitempty"`
	ScoreDetails                    *ScoreDetails           `xml:"ScoreDetails,omitempty"  json:"ScoreDetails,omitempty"`
	VulnerabilityVectorText         string                  `xml:"VulnerabilityVectorText,omitempty"  json:"VulnerabilityVectorText,omitempty"`
	AccessComplexityText            string                  `xml:"AccessComplexityText,omitempty"  json:"AccessComplexityText,omitempty"`
	VulnerabilityAuthenticationText string                  `xml:"VulnerabilityAuthenticationText,omitempty"  json:"VulnerabilityAuthenticationText,omitempty"`
	ConfidentialityImpactText       string                  `xml:"ConfidentialityImpactText,omitempty"  json:"ConfidentialityImpactText,omitempty"`
	IntegrityImpactText             string                  `xml:"IntegrityImpactText,omitempty"  json:"IntegrityImpactText,omitempty"`
	AvailabilityImpactText          string                  `xml:"AvailabilityImpactText,omitempty"  json:"AvailabilityImpactText,omitempty"`
	VulnerabilitySourceText         string                  `xml:"VulnerabilitySourceText,omitempty"  json:"VulnerabilitySourceText,omitempty"`
	AssessmentCheckText             string                  `xml:"AssessmentCheckText,omitempty"  json:"AssessmentCheckText,omitempty"`
	ScannerText                     string                  `xml:"ScannerText,omitempty"  json:"ScannerText,omitempty"`
	RecommendationText              string                  `xml:"RecommendationText,omitempty"  json:"RecommendationText,omitempty"`
	VulnerabilityReference          *VulnerabilityReference `xml:"VulnerabilityReference,omitempty"  json:"VulnerabilityReference,omitempty"`
	VulnerabilityModifiedDate       string                  `xml:"VulnerabilityModifiedDate,omitempty"  json:"VulnerabilityModifiedDate,omitempty"`
	VulnerabilityPublishedDate      string                  `xml:"VulnerabilityPublishedDate,omitempty"  json:"VulnerabilityPublishedDate,omitempty"`
	XMLName                         xml.Name                `xml:"Vulnerability,omitempty"  json:"Vulnerability,omitempty"`
}

//VulnerabilityInformation ...
type VulnerabilityInformation struct {
	Vulnerabilities  *Vulnerabilities  `xml:"Vulnerabilities,omitempty"  json:"Vulnerabilities,omitempty"`
	VirusInformation *VirusInformation `xml:"VirusInformation,omitempty"  json:"VirusInformation,omitempty"`
	XMLName          xml.Name          `xml:"VulnerabilityInformation,omitempty"  json:"VulnerabilityInformation,omitempty"`
}

//VulnerabilityReference ...
type VulnerabilityReference struct {
	TypeOfReferenceText string   `xml:"TypeOfReferenceText,omitempty"  json:"TypeOfReferenceText,omitempty"`
	ReferenceSourceText string   `xml:"ReferenceSourceText,omitempty"  json:"ReferenceSourceText,omitempty"`
	URIText             string   `xml:"URIText,omitempty"  json:"URIText,omitempty"`
	ReferenceText       string   `xml:"ReferenceText,omitempty"  json:"ReferenceText,omitempty"`
	XMLName             xml.Name `xml:"VulnerabilityReference,omitempty"  json:"VulnerabilityReference,omitempty"`
}
