package main

import "encoding/xml"

//NewSoftwareEvidenceArchiveISM ...
func NewSoftwareEvidenceArchiveISM() *SoftwareEvidenceArchive {
	return &SoftwareEvidenceArchive{
		// Required for the proper namespacingSoftwareEvidenceArchive
		AttrXmlnsXsi: "http://www.w3.org/2001/XMLSchema-instance",
		AttrXmlns:    "urn:seva::1.0",
		AttrXmlnsIsm: "urn:us:gov:ic:ism",
		SoftwareInformation: &SoftwareInformation{
			VersionInformation: &VersionInformation{},
		},
		FileInformation:                &FileInformation{},
		AuthoritativeSourceInformation: &AuthoritativeSourceInformation{},
		EcosystemInformation: &EcosystemInformation{
			Ecosystem:          &Ecosystem{},
			CompanyInformation: &CompanyInformation{},
		},
		DependencyInformation: &DependencyInformation{
			Dependencies: &Dependencies{},
			Compiler:     &Compiler{},
		},
		LicenseInformation: &LicenseInformation{},
		VulnerabilityInformation: &VulnerabilityInformation{
			Vulnerabilities:  &Vulnerabilities{},
			VirusInformation: &VirusInformation{},
		},
		GovernanceRiskCompliance: &GovernanceRiskCompliance{},
		DeliveryInformation:      &DeliveryInformation{},
	}
}

//SoftwareEvidenceArchive ... A data item for automated software supply chain metadata.
type SoftwareEvidenceArchive struct {
	AttrXmlnsXsi                   string                          `xml:"xmlns:xsi,attr" json:"AttrXmlnsXsi,omitempty"`
	AttrXmlns                      string                          `xml:"xmlns,attr" json:"AttrXmlns,omitempty"`
	AttrXmlnsIsm                   string                          `xml:"xmlns:ism,attr" json:"AttrXmlnsIsm,omitempty"`
	SoftwareInformation            *SoftwareInformation            `xml:"SoftwareInformation,omitempty"  json:"SoftwareInformation,omitempty"`
	FileInformation                *FileInformation                `xml:"FileInformation,omitempty"  json:"FileInformation,omitempty"`
	AuthoritativeSourceInformation *AuthoritativeSourceInformation `xml:"AuthoritativeSourceInformation,omitempty"  json:"AuthoritativeSourceInformation,omitempty"`
	EcosystemInformation           *EcosystemInformation           `xml:"EcosystemInformation,omitempty"  json:"EcosystemInformation,omitempty"`
	DependencyInformation          *DependencyInformation          `xml:"DependencyInformation,omitempty"  json:"DependencyInformation,omitempty"`
	LicenseInformation             *LicenseInformation             `xml:"LicenseInformation,omitempty"  json:"LicenseInformation,omitempty"`
	VulnerabilityInformation       *VulnerabilityInformation       `xml:"VulnerabilityInformation,omitempty"  json:"VulnerabilityInformation,omitempty"`
	GovernanceRiskCompliance       *GovernanceRiskCompliance       `xml:"GovernanceRiskCompliance,omitempty"  json:"GovernanceRiskCompliance,omitempty"`
	DeliveryInformation            *DeliveryInformation            `xml:"DeliveryInformation,omitempty"  json:"DeliveryInformation,omitempty"`
	SecurityAttributesOptionGroup                                  //SecurityAttributesOptionGroup
	XMLName                        xml.Name                        `xml:"SoftwareEvidenceArchive,omitempty"  json:"SoftwareEvidenceArchive,omitempty"`
}

//APIName ... A data item for the name of the required public API declaration
type APIName struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"APIName,omitempty"  json:"APIName,omitempty"`
}

//AccessComplexityText ... A data item for a access complexity
type AccessComplexityText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"AccessComplexityText,omitempty"  json:"AccessComplexityText,omitempty"`
}

//AccessVectorText ... A data item for a vulnerability access vector
type AccessVectorText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"AccessVectorText,omitempty"  json:"AccessVectorText,omitempty"`
}

//AssessmentCheckText ... A data item for a vulnerabilty source
type AssessmentCheckText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"AssessmentCheckText,omitempty"  json:"AssessmentCheckText,omitempty"`
}

//AuthoritativeDigitalSignatureIndicator ... True if has an authoritative digital signatiure; false if not.
type AuthoritativeDigitalSignatureIndicator struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"AuthoritativeDigitalSignatureIndicator,omitempty"  json:"AuthoritativeDigitalSignatureIndicator,omitempty"`
}

//AuthoritativeSourceInformation ... A data item for information about what is to be considered the authoritative source for a given artifact or source repository. This includes an URL, a hash of the source, and whether or not the source has been signed
type AuthoritativeSourceInformation struct {
	SourceHashText                         *SourceHashText                         `xml:"SourceHashText,omitempty"  json:"SourceHashText,omitempty"`
	SourceURIText                          *SourceURIText                          `xml:"SourceURIText,omitempty"  json:"SourceURIText,omitempty"`
	AuthoritativeDigitalSignatureIndicator *AuthoritativeDigitalSignatureIndicator `xml:"AuthoritativeDigitalSignatureIndicator,omitempty"  json:"AuthoritativeDigitalSignatureIndicator,omitempty"`
	SecurityAttributesOptionGroup                                                  //SecurityAttributesOptionGroup
	XMLName                                xml.Name                                `xml:"AuthoritativeSourceInformation,omitempty"  json:"AuthoritativeSourceInformation,omitempty"`
}

//AvailabilityImpactText ... A data item for a availability impact
type AvailabilityImpactText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"AvailabilityImpactText,omitempty"  json:"AvailabilityImpactText,omitempty"`
}

//BaseSeverityText ... A data item for base severity
type BaseSeverityText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"BaseSeverityText,omitempty"  json:"BaseSeverityText,omitempty"`
}

//BuildMetaText ... A data item for Build metadata. Build metadata MAY be denoted by appending a plus sign and a series of dot separated identifiers immediately following the patch or pre-release version. Identifiers MUST comprise only ASCII alphanumerics and hyphen [0-9A-Za-z-]. Identifiers MUST NOT be empty. Build metadata SHOULD be ignored when determining version precedence. Thus two versions that differ only in the build metadata, have the same precedence. Examples: 1.0.0-alpha+001, 1.0.0+20130313144700, 1.0.0-beta+exp.sha.5114f85.(https://semver.org/spec/v2.0.0.html)
type BuildMetaText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"BuildMetaText,omitempty"  json:"BuildMetaText,omitempty"`
}

//CommitLogText ... A data item for an entry associated with software repository commit event
type CommitLogText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"CommitLogText,omitempty"  json:"CommitLogText,omitempty"`
}

//CommitLogs ... A data item for entries associated with software repository commit events
type CommitLogs struct {
	CommitLogText                 *CommitLogText `xml:"CommitLogText,omitempty"  json:"CommitLogText,omitempty"`
	SecurityAttributesOptionGroup                //SecurityAttributesOptionGroup
	XMLName                       xml.Name       `xml:"CommitLogs,omitempty"  json:"CommitLogs,omitempty"`
}

//CommitterQuantity ... A data item for the number of committers to a software repository
type CommitterQuantity struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"CommitterQuantity,omitempty"  json:"CommitterQuantity,omitempty"`
}

//CommonVulnerabilityScoringSystemV2 ... A data type for Common Vulnerability Scoring System V2 (CVSSV2) Information
type CommonVulnerabilityScoringSystemV2 struct {
	VectorStringText                *VectorStringText                `xml:"VectorStringText,omitempty"  json:"VectorStringText,omitempty"`
	AccessVectorText                *AccessVectorText                `xml:"AccessVectorText,omitempty"  json:"AccessVectorText,omitempty"`
	AccessComplexityText            *AccessComplexityText            `xml:"AccessComplexityText,omitempty"  json:"AccessComplexityText,omitempty"`
	VulnerabilityAuthenticationText *VulnerabilityAuthenticationText `xml:"VulnerabilityAuthenticationText,omitempty"  json:"VulnerabilityAuthenticationText,omitempty"`
	ConfidentialityImpactText       *ConfidentialityImpactText       `xml:"ConfidentialityImpactText,omitempty"  json:"ConfidentialityImpactText,omitempty"`
	IntegrityImpactText             *IntegrityImpactText             `xml:"IntegrityImpactText,omitempty"  json:"IntegrityImpactText,omitempty"`
	AvailabilityImpactText          *AvailabilityImpactText          `xml:"AvailabilityImpactText,omitempty"  json:"AvailabilityImpactText,omitempty"`
	VulnerabilityBaseScoreValue     *VulnerabilityBaseScoreValue     `xml:"VulnerabilityBaseScoreValue,omitempty"  json:"VulnerabilityBaseScoreValue,omitempty"`
	SecurityAttributesOptionGroup                                    //SecurityAttributesOptionGroup
	XMLName                         xml.Name                         `xml:"CommonVulnerabilityScoringSystemV2,omitempty"  json:"CommonVulnerabilityScoringSystemV2,omitempty"`
}

//CommonVulnerabilityScoringSystemV3 ... A data type for Common Vulnerability Scoring System V3 (CVSSV3) Information
type CommonVulnerabilityScoringSystemV3 struct {
	VectorStringText              *VectorStringText            `xml:"VectorStringText,omitempty"  json:"VectorStringText,omitempty"`
	AccessVectorText              *AccessVectorText            `xml:"AccessVectorText,omitempty"  json:"AccessVectorText,omitempty"`
	AccessComplexityText          *AccessComplexityText        `xml:"AccessComplexityText,omitempty"  json:"AccessComplexityText,omitempty"`
	PrivilegesRequiredText        *PrivilegesRequiredText      `xml:"PrivilegesRequiredText,omitempty"  json:"PrivilegesRequiredText,omitempty"`
	UserInteractionText           *UserInteractionText         `xml:"UserInteractionText,omitempty"  json:"UserInteractionText,omitempty"`
	VulnerabilityScopeText        *VulnerabilityScopeText      `xml:"VulnerabilityScopeText,omitempty"  json:"VulnerabilityScopeText,omitempty"`
	ConfidentialityImpactText     *ConfidentialityImpactText   `xml:"ConfidentialityImpactText,omitempty"  json:"ConfidentialityImpactText,omitempty"`
	IntegrityImpactText           *IntegrityImpactText         `xml:"IntegrityImpactText,omitempty"  json:"IntegrityImpactText,omitempty"`
	AvailabilityImpactText        *AvailabilityImpactText      `xml:"AvailabilityImpactText,omitempty"  json:"AvailabilityImpactText,omitempty"`
	VulnerabilityBaseScoreValue   *VulnerabilityBaseScoreValue `xml:"VulnerabilityBaseScoreValue,omitempty"  json:"VulnerabilityBaseScoreValue,omitempty"`
	BaseSeverityText              *BaseSeverityText            `xml:"BaseSeverityText,omitempty"  json:"BaseSeverityText,omitempty"`
	SecurityAttributesOptionGroup                              //SecurityAttributesOptionGroup
	XMLName                       xml.Name                     `xml:"CommonVulnerabilityScoringSystemV3,omitempty"  json:"CommonVulnerabilityScoringSystemV3,omitempty"`
}

//CompanyInformation ... A data item for languages used in a software repository
type CompanyInformation struct {
	Name                          *Name    `xml:"Name,omitempty"  json:"Name,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"CompanyInformation,omitempty"  json:"CompanyInformation,omitempty"`
}

//Compiler ... A data item for languages used in a software repository
type Compiler struct {
	Name                          *Name                `xml:"Name,omitempty"  json:"Name,omitempty"`
	CompilerVersionText           *CompilerVersionText `xml:"CompilerVersionText,omitempty"  json:"CompilerVersionText,omitempty"`
	SecurityAttributesOptionGroup                      //SecurityAttributesOptionGroup
	XMLName                       xml.Name             `xml:"Compiler,omitempty"  json:"Compiler,omitempty"`
}

//CompilerVersionText ... A data item for a compiler version
type CompilerVersionText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"CompilerVersionText,omitempty"  json:"CompilerVersionText,omitempty"`
}

//ConfidentialityImpactText ... A data item for confidentiality impact
type ConfidentialityImpactText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"ConfidentialityImpactText,omitempty"  json:"ConfidentialityImpactText,omitempty"`
}

//DataQuantity ... A data item for the count of an artifact
type DataQuantity struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"DataQuantity,omitempty"  json:"DataQuantity,omitempty"`
}

//DataReadText ... A data item for an entry associated with software repository commit event
type DataReadText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"DataReadText,omitempty"  json:"DataReadText,omitempty"`
}

//DataScannedText ... A data item for an entry associated with software repository commit event
type DataScannedText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"DataScannedText,omitempty"  json:"DataScannedText,omitempty"`
}

//DatabaseVersionText ... A data item for a virus database version
type DatabaseVersionText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"DatabaseVersionText,omitempty"  json:"DatabaseVersionText,omitempty"`
}

//DateTime ... A data type for date and time information
type DateTime struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"DateTime,omitempty"  json:"DateTime,omitempty"`
}

//DeliveringOrganizationText ... A data item for the Organization of Delivery for the SEvA
type DeliveringOrganizationText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"DeliveringOrganizationText,omitempty"  json:"DeliveringOrganizationText,omitempty"`
}

//DeliveryInformation ... A data item for delivery information including the target url and last delivered date and time for a previously signed Software Evidence Archive.
type DeliveryInformation struct {
	DeliveringOrganizationText    *DeliveringOrganizationText `xml:"DeliveringOrganizationText,omitempty"  json:"DeliveringOrganizationText,omitempty"`
	DestinationText               *DestinationText            `xml:"DestinationText,omitempty"  json:"DestinationText,omitempty"`
	HashText                      *HashText                   `xml:"HashText,omitempty"  json:"HashText,omitempty"`
	LastRegistryUpdateDate        *LastRegistryUpdateDate     `xml:"LastRegistryUpdateDate,omitempty"  json:"LastRegistryUpdateDate,omitempty"`
	SecurityAttributesOptionGroup                             //SecurityAttributesOptionGroup
	XMLName                       xml.Name                    `xml:"DeliveryInformation,omitempty"  json:"DeliveryInformation,omitempty"`
}

//Dependencies ... A data item for software dependencies
type Dependencies struct {
	Dependency                    *Dependency `xml:"Dependency,omitempty"  json:"Dependency,omitempty"`
	SecurityAttributesOptionGroup             //SecurityAttributesOptionGroup
	XMLName                       xml.Name    `xml:"Dependencies,omitempty"  json:"Dependencies,omitempty"`
}

//Dependency ... A data item for software dependency information
type Dependency struct {
	Name                          *Name                  `xml:"Name,omitempty"  json:"Name,omitempty"`
	DependencyVersionText         *DependencyVersionText `xml:"DependencyVersionText,omitempty"  json:"DependencyVersionText,omitempty"`
	SecurityAttributesOptionGroup                        //SecurityAttributesOptionGroup
	XMLName                       xml.Name               `xml:"Dependency,omitempty"  json:"Dependency,omitempty"`
}

//DependencyInformation ... A data item for a list of dependencies for a given software project derived directly from the artifact or source dependency definition file. Entries include names, versions, and vulnerabilities
type DependencyInformation struct {
	Dependencies                  *Dependencies `xml:"Dependencies,omitempty"  json:"Dependencies,omitempty"`
	Compiler                      *Compiler     `xml:"Compiler,omitempty"  json:"Compiler,omitempty"`
	SecurityAttributesOptionGroup               //SecurityAttributesOptionGroup
	XMLName                       xml.Name      `xml:"DependencyInformation,omitempty"  json:"DependencyInformation,omitempty"`
}

//DependencyVersionText ... A data item for a dependency version
type DependencyVersionText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"DependencyVersionText,omitempty"  json:"DependencyVersionText,omitempty"`
}

//DestinationText ... A data item for a destination
type DestinationText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"DestinationText,omitempty"  json:"DestinationText,omitempty"`
}

//Ecosystem ... A data item for software ecosystem information
type Ecosystem struct {
	CommitterQuantity             *CommitterQuantity `xml:"CommitterQuantity,omitempty"  json:"CommitterQuantity,omitempty"`
	Languages                     *Languages         `xml:"Languages,omitempty"  json:"Languages,omitempty"`
	CommitLogs                    *CommitLogs        `xml:"CommitLogs,omitempty"  json:"CommitLogs,omitempty"`
	SecurityAttributesOptionGroup                    //SecurityAttributesOptionGroup
	XMLName                       xml.Name           `xml:"Ecosystem,omitempty"  json:"Ecosystem,omitempty"`
}

//EcosystemInformation ... A data item for information pertaining to a software project's ecosystem which includes programming languages, references, number of committers, mailing list activity, overall sentiment, and other information
type EcosystemInformation struct {
	Ecosystem                     *Ecosystem          `xml:"Ecosystem,omitempty"  json:"Ecosystem,omitempty"`
	CompanyInformation            *CompanyInformation `xml:"CompanyInformation,omitempty"  json:"CompanyInformation,omitempty"`
	SecurityAttributesOptionGroup                     //SecurityAttributesOptionGroup
	XMLName                       xml.Name            `xml:"EcosystemInformation,omitempty"  json:"EcosystemInformation,omitempty"`
}

//EndOfLifeIndicator ... True if software has end of life status; false if not.
type EndOfLifeIndicator struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"EndOfLifeIndicator,omitempty"  json:"EndOfLifeIndicator,omitempty"`
}

//EngineDetails ... A data item for for details of a virus engine
type EngineDetails struct {
	EngineVersionText             *EngineVersionText   `xml:"EngineVersionText,omitempty"  json:"EngineVersionText,omitempty"`
	DatabaseVersionText           *DatabaseVersionText `xml:"DatabaseVersionText,omitempty"  json:"DatabaseVersionText,omitempty"`
	SecurityAttributesOptionGroup                      //SecurityAttributesOptionGroup
	XMLName                       xml.Name             `xml:"EngineDetails,omitempty"  json:"EngineDetails,omitempty"`
}

//EngineVersionText ... A data item for a virus engine version
type EngineVersionText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"EngineVersionText,omitempty"  json:"EngineVersionText,omitempty"`
}

//FileExtensionText ... A data item for a computer file extension
type FileExtensionText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"FileExtensionText,omitempty"  json:"FileExtensionText,omitempty"`
}

//FileInformation ... A data item for file types and related Multipurpose Internet Mail Extensions (MIME) types
type FileInformation struct {
	FileNameText                  *FileNameText      `xml:"FileNameText,omitempty"  json:"FileNameText,omitempty"`
	FileExtensionText             *FileExtensionText `xml:"FileExtensionText,omitempty"  json:"FileExtensionText,omitempty"`
	SecurityAttributesOptionGroup                    //SecurityAttributesOptionGroup
	XMLName                       xml.Name           `xml:"FileInformation,omitempty"  json:"FileInformation,omitempty"`
}

//FileNameText ... A data item for a computer file name without extension
type FileNameText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"FileNameText,omitempty"  json:"FileNameText,omitempty"`
}

//FileNotesText ... A data item a for a multi-line entry for file notes
type FileNotesText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"FileNotesText,omitempty"  json:"FileNotesText,omitempty"`
}

//FileText ... A data item for a computer file name with extension
type FileText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"FileText,omitempty"  json:"FileText,omitempty"`
}

//GovernanceRiskCompliance ... A data item for calculated risk and compliance to governance for a software artifact or source repository
type GovernanceRiskCompliance struct {
	RiskCode                      *RiskCode                 `xml:"RiskCode,omitempty"  json:"RiskCode,omitempty"`
	StatementOfAssuranceText      *StatementOfAssuranceText `xml:"StatementOfAssuranceText,omitempty"  json:"StatementOfAssuranceText,omitempty"`
	SecurityAttributesOptionGroup                           //SecurityAttributesOptionGroup
	XMLName                       xml.Name                  `xml:"GovernanceRiskCompliance,omitempty"  json:"GovernanceRiskCompliance,omitempty"`
}

//GroupingText ... A data item for a software family or relationship
type GroupingText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"GroupingText,omitempty"  json:"GroupingText,omitempty"`
}

//HashText ... A data item for a hash of the artifact delivered
type HashText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"HashText,omitempty"  json:"HashText,omitempty"`
}

//InfectedFileQuantity ... A data item for the number of files infected by a virus
type InfectedFileQuantity struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"InfectedFileQuantity,omitempty"  json:"InfectedFileQuantity,omitempty"`
}

//IntegrityImpactText ... A data item for a integrity impact
type IntegrityImpactText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"IntegrityImpactText,omitempty"  json:"IntegrityImpactText,omitempty"`
}

//KnownVirusQuantity ... A data item for the number of known viruses
type KnownVirusQuantity struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"KnownVirusQuantity,omitempty"  json:"KnownVirusQuantity,omitempty"`
}

//LanguageText ... A data item for the name of a programming language
type LanguageText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"LanguageText,omitempty"  json:"LanguageText,omitempty"`
}

//Languages ... A data item for languages used in a software repository
type Languages struct {
	LanguageText                  *LanguageText `xml:"LanguageText,omitempty"  json:"LanguageText,omitempty"`
	SecurityAttributesOptionGroup               //SecurityAttributesOptionGroup
	XMLName                       xml.Name      `xml:"Languages,omitempty"  json:"Languages,omitempty"`
}

//LastRegistryUpdateDate ... A data type for date and time of last registry update
type LastRegistryUpdateDate struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"LastRegistryUpdateDate,omitempty"  json:"LastRegistryUpdateDate,omitempty"`
}

//LicenseCategoryCode ... A data item for a software license category
type LicenseCategoryCode struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"LicenseCategoryCode,omitempty"  json:"LicenseCategoryCode,omitempty"`
}

//LicenseCode ... A data item for a type of software license
type LicenseCode struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"LicenseCode,omitempty"  json:"LicenseCode,omitempty"`
}

//LicenseInformation ... A data item for software license information for a software artifact or source repository
type LicenseInformation struct {
	LicenseCategoryCode           *LicenseCategoryCode `xml:"LicenseCategoryCode,omitempty"  json:"LicenseCategoryCode,omitempty"`
	LicenseCode                   *LicenseCode         `xml:"LicenseCode,omitempty"  json:"LicenseCode,omitempty"`
	EndOfLifeIndicator            *EndOfLifeIndicator  `xml:"EndOfLifeIndicator,omitempty"  json:"EndOfLifeIndicator,omitempty"`
	SecurityAttributesOptionGroup                      //SecurityAttributesOptionGroup
	XMLName                       xml.Name             `xml:"LicenseInformation,omitempty"  json:"LicenseInformation,omitempty"`
}

//LineText ... A data item for a short text entry suitable for a name or a title
type LineText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"LineText,omitempty"  json:"LineText,omitempty"`
}

//LogEntryText ... A data item for a log entry
type LogEntryText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"LogEntryText,omitempty"  json:"LogEntryText,omitempty"`
}

//MajorVersionNumeric ... A data item for the Major version. The value zero (0.y.z) is for initial development. Anything may change at any time. The public API should not be considered stable. Version 1.0.0 defines the public API. The way in which the version number is incremented after this release is dependent on this public API and how it changes. Major version X (X.y.z | X > 0) MUST be incremented if any backwards incompatible changes are introduced to the public API. It MAY include minor and patch level changes. Patch and minor version MUST be reset to 0 when major version is incremented.
type MajorVersionNumeric struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"MajorVersionNumeric,omitempty"  json:"MajorVersionNumeric,omitempty"`
}

//MinorVersionNumeric ... A data item for the Minor version. Minor version Y (x.Y.z | x > 0) MUST be incremented if new, backwards compatible functionality is introduced to the public API. It MUST be incremented if any public API functionality is marked as deprecated. It MAY be incremented if substantial new functionality or improvements are introduced within the private code. It MAY include patch level changes. Patch version MUST be reset to 0 when minor version is incremented.
type MinorVersionNumeric struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"MinorVersionNumeric,omitempty"  json:"MinorVersionNumeric,omitempty"`
}

//Name ... A data item for a name
type Name struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"Name,omitempty"  json:"Name,omitempty"`
}

//ParagraphText ... A data item a for a multi-line text entry
type ParagraphText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"ParagraphText,omitempty"  json:"ParagraphText,omitempty"`
}

//PatchVersionText ... A data item for the Patch version. Patch version Z (x.y.Z | x > 0) MUST be incremented if only backwards compatible bug fixes are introduced. A bug fix is defined as an internal change that fixes incorrect behavior.
type PatchVersionText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"PatchVersionText,omitempty"  json:"PatchVersionText,omitempty"`
}

//PreReleaseVersionText ... A pre-release version MAY be denoted by appending a hyphen and a series of dot separated identifiers immediately following the patch version. Identifiers MUST comprise only ASCII alphanumerics and hyphen [0-9A-Za-z-]. Identifiers MUST NOT be empty. Numeric identifiers MUST NOT include leading zeroes. Pre-release versions have a lower precedence than the associated normal version. A pre-release version indicates that the version is unstable and might not satisfy the intended compatibility requirements as denoted by its associated normal version. Examples: 1.0.0-alpha, 1.0.0-alpha.1, 1.0.0-0.3.7, 1.0.0-x.7.z.92.
type PreReleaseVersionText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"PreReleaseVersionText,omitempty"  json:"PreReleaseVersionText,omitempty"`
}

//PrivilegesRequiredText ... A data item for privileges required
type PrivilegesRequiredText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"PrivilegesRequiredText,omitempty"  json:"PrivilegesRequiredText,omitempty"`
}

//ProductTitleText ... A data item for the product title
type ProductTitleText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"ProductTitleText,omitempty"  json:"ProductTitleText,omitempty"`
}

//PythonDevelopmentReleaseSegmentText ... A data type for the Python development release segment with the scheme: .devN: [N!]N(.N)*[{a|b|rc}N][.postN][.devN] (https://www.python.org/dev/peps/pep-0440/#public-version-identifiers).
type PythonDevelopmentReleaseSegmentText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"PythonDevelopmentReleaseSegmentText,omitempty"  json:"PythonDevelopmentReleaseSegmentText,omitempty"`
}

//PythonEpochSegmentText ... A data type for the Python Epoch Segment with the scheme: [N!]. (https://www.python.org/dev/peps/pep-0440/#public-version-identifiers).
type PythonEpochSegmentText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"PythonEpochSegmentText,omitempty"  json:"PythonEpochSegmentText,omitempty"`
}

//PythonPostReleaseSegmentText ... A data type for the Python Post-Release Segment with the scheme: .postN (https://www.python.org/dev/peps/pep-0440/#public-version-identifiers).
type PythonPostReleaseSegmentText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"PythonPostReleaseSegmentText,omitempty"  json:"PythonPostReleaseSegmentText,omitempty"`
}

//PythonPreReleaseSegmentText ... A data type for the Python Pre-Release Segment with the scheme: {a|b|rc}N. (https://www.python.org/dev/peps/pep-0440/#public-version-identifiers).
type PythonPreReleaseSegmentText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"PythonPreReleaseSegmentText,omitempty"  json:"PythonPreReleaseSegmentText,omitempty"`
}

//PythonReleaseSegmentText ... A data type for the Python Release Segment with the scheme: N(.N)*. (https://www.python.org/dev/peps/pep-0440/#public-version-identifiers).
type PythonReleaseSegmentText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"PythonReleaseSegmentText,omitempty"  json:"PythonReleaseSegmentText,omitempty"`
}

//PythonSemanticVersion ... A data item for a Python Distribution identified by a public version identifier which supports all defined version comparison operations. The version scheme is used both to describe the distribution version provided by a particular distribution archive, as well as to place constraints on the version of dependencies needed in order to build or run the software.The canonical public version identifiers MUST comply with the following scheme: [N!]N(.N)*[{a|b|rc}N][.postN][.devN] (https://www.python.org/dev/peps/pep-0440/#public-version-identifiers)
type PythonSemanticVersion struct {
	PythonEpochSegmentText              *PythonEpochSegmentText              `xml:"PythonEpochSegmentText,omitempty"  json:"PythonEpochSegmentText,omitempty"`
	PythonReleaseSegmentText            *PythonReleaseSegmentText            `xml:"PythonReleaseSegmentText,omitempty"  json:"PythonReleaseSegmentText,omitempty"`
	PythonPreReleaseSegmentText         *PythonPreReleaseSegmentText         `xml:"PythonPreReleaseSegmentText,omitempty"  json:"PythonPreReleaseSegmentText,omitempty"`
	PythonPostReleaseSegmentText        *PythonPostReleaseSegmentText        `xml:"PythonPostReleaseSegmentText,omitempty"  json:"PythonPostReleaseSegmentText,omitempty"`
	PythonDevelopmentReleaseSegmentText *PythonDevelopmentReleaseSegmentText `xml:"PythonDevelopmentReleaseSegmentText,omitempty"  json:"PythonDevelopmentReleaseSegmentText,omitempty"`
	SecurityAttributesOptionGroup                                            //SecurityAttributesOptionGroup
	XMLName                             xml.Name                             `xml:"PythonSemanticVersion,omitempty"  json:"PythonSemanticVersion,omitempty"`
}

//PythonSemanticVersionIndicator ... True if python semantic version; false if not or not known.
type PythonSemanticVersionIndicator struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"PythonSemanticVersionIndicator,omitempty"  json:"PythonSemanticVersionIndicator,omitempty"`
}

//PythonVersionText ... An augmentation point for SemanticVersionType
type PythonVersionText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"PythonVersionText,omitempty"  json:"PythonVersionText,omitempty"`
}

//RecommendationText ... A data item for a recommendation paragraph
type RecommendationText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"RecommendationText,omitempty"  json:"RecommendationText,omitempty"`
}

//ReferenceSourceText ... A data item for the reference source
type ReferenceSourceText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"ReferenceSourceText,omitempty"  json:"ReferenceSourceText,omitempty"`
}

//ReferenceText ... A data item for a reference description
type ReferenceText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"ReferenceText,omitempty"  json:"ReferenceText,omitempty"`
}

//ReleaseNotesText ... A data item for a release notes
type ReleaseNotesText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"ReleaseNotesText,omitempty"  json:"ReleaseNotesText,omitempty"`
}

//RiskCode ... A data item for the level of risk associated with a software artifact
type RiskCode struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"RiskCode,omitempty"  json:"RiskCode,omitempty"`
}

//ScannedDirectoryQuantity ... A data item for the number of directories scanned
type ScannedDirectoryQuantity struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"ScannedDirectoryQuantity,omitempty"  json:"ScannedDirectoryQuantity,omitempty"`
}

//ScannedFileQuantity ... A data item for the number of files scanned
type ScannedFileQuantity struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"ScannedFileQuantity,omitempty"  json:"ScannedFileQuantity,omitempty"`
}

//ScannerText ... A data item for a vulnerabilty source
type ScannerText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"ScannerText,omitempty"  json:"ScannerText,omitempty"`
}

//ScoreDetails ... A data item for common vulnerability score details
type ScoreDetails struct {
	CommonVulnerabilityScoringSystemV2 *CommonVulnerabilityScoringSystemV2 `xml:"CommonVulnerabilityScoringSystemV2,omitempty"  json:"CommonVulnerabilityScoringSystemV2,omitempty"`
	CommonVulnerabilityScoringSystemV3 *CommonVulnerabilityScoringSystemV3 `xml:"CommonVulnerabilityScoringSystemV3,omitempty"  json:"CommonVulnerabilityScoringSystemV3,omitempty"`
	SecurityAttributesOptionGroup                                          //SecurityAttributesOptionGroup
	XMLName                            xml.Name                            `xml:"ScoreDetails,omitempty"  json:"ScoreDetails,omitempty"`
}

//ScoreSystemText ... A data item for the name of a score system
type ScoreSystemText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"ScoreSystemText,omitempty"  json:"ScoreSystemText,omitempty"`
}

//SemanticVersionDate ... A data type for date and time information
type SemanticVersionDate struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"SemanticVersionDate,omitempty"  json:"SemanticVersionDate,omitempty"`
}

//SemanticVersionIndicator ... True if semantic version; false if not or not known.
type SemanticVersionIndicator struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"SemanticVersionIndicator,omitempty"  json:"SemanticVersionIndicator,omitempty"`
}

//SemanticVersionInformation ... A data type for semantic versioning information. A normal version number MUST take the form X.Y.Z where X, Y, and Z are non-negative integers, and MUST NOT contain leading zeroes. X is the major version, Y is the minor version, and Z is the patch version. Each element MUST increase numerically. For instance: 1.9.0 -> 1.10.0 -> 1.11.0.(https://semver.org/spec/v2.0.0.html)
type SemanticVersionInformation struct {
	APIName                       *APIName               `xml:"APIName,omitempty"  json:"APIName,omitempty"`
	MajorVersionNumeric           *MajorVersionNumeric   `xml:"MajorVersionNumeric,omitempty"  json:"MajorVersionNumeric,omitempty"`
	MinorVersionNumeric           *MinorVersionNumeric   `xml:"MinorVersionNumeric,omitempty"  json:"MinorVersionNumeric,omitempty"`
	PatchVersionText              *PatchVersionText      `xml:"PatchVersionText,omitempty"  json:"PatchVersionText,omitempty"`
	PreReleaseVersionText         *PreReleaseVersionText `xml:"PreReleaseVersionText,omitempty"  json:"PreReleaseVersionText,omitempty"`
	BuildMetaText                 *BuildMetaText         `xml:"BuildMetaText,omitempty"  json:"BuildMetaText,omitempty"`
	SecurityAttributesOptionGroup                        //SecurityAttributesOptionGroup
	XMLName                       xml.Name               `xml:"SemanticVersionInformation,omitempty"  json:"SemanticVersionInformation,omitempty"`
}

//SemanticVersionText ... A data item a software product version
type SemanticVersionText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"SemanticVersionText,omitempty"  json:"SemanticVersionText,omitempty"`
}

//Sha256Text ... A data item a for a sha256 encoded string
type Sha256Text struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"Sha256Text,omitempty"  json:"Sha256Text,omitempty"`
}

//SoftwareInformation ... A data item for software product naming or version related information
type SoftwareInformation struct {
	ProductTitleText              *ProductTitleText   `xml:"ProductTitleText,omitempty"  json:"ProductTitleText,omitempty"`
	SoftwareNameText              *SoftwareNameText   `xml:"SoftwareNameText,omitempty"  json:"SoftwareNameText,omitempty"`
	SoftwareOrgText               *SoftwareOrgText    `xml:"SoftwareOrgText,omitempty"  json:"SoftwareOrgText,omitempty"`
	VersionInformation            *VersionInformation `xml:"VersionInformation,omitempty"  json:"VersionInformation,omitempty"`
	GroupingText                  *GroupingText       `xml:"GroupingText,omitempty"  json:"GroupingText,omitempty"`
	ReleaseNotesText              *ReleaseNotesText   `xml:"ReleaseNotesText,omitempty"  json:"ReleaseNotesText,omitempty"`
	SecurityAttributesOptionGroup                     //SecurityAttributesOptionGroup
	XMLName                       xml.Name            `xml:"SoftwareInformation,omitempty"  json:"SoftwareInformation,omitempty"`
}

//SoftwareNameText ... A data item for the Name of the software
type SoftwareNameText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"SoftwareNameText,omitempty"  json:"SoftwareNameText,omitempty"`
}

//SoftwareOrgText ... A data item for the Organization of the software
type SoftwareOrgText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"SoftwareOrgText,omitempty"  json:"SoftwareOrgText,omitempty"`
}

//SoftwareVersionText ... A data item for the raw Version of the software
type SoftwareVersionText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"SoftwareVersionText,omitempty"  json:"SoftwareVersionText,omitempty"`
}

//SourceHashText ... A data item a for a sha1 encoded string
type SourceHashText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"SourceHashText,omitempty"  json:"SourceHashText,omitempty"`
}

//SourceText ... A data item describing the vulnerabilty source
type SourceText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"SourceText,omitempty"  json:"SourceText,omitempty"`
}

//SourceURIText ... A data item for a World Wide Web Consortium Uniform Reference Indicator
type SourceURIText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"SourceURIText,omitempty"  json:"SourceURIText,omitempty"`
}

//StatementOfAssuranceText ... A data item for a statement of assurance
type StatementOfAssuranceText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"StatementOfAssuranceText,omitempty"  json:"StatementOfAssuranceText,omitempty"`
}

//SummaryText ... A data item for a summary
type SummaryText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"SummaryText,omitempty"  json:"SummaryText,omitempty"`
}

//TimeText ... A data item for a text time entry
type TimeText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"TimeText,omitempty"  json:"TimeText,omitempty"`
}

//TypeOfReferenceText ... A data item for the type of reference
type TypeOfReferenceText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"TypeOfReferenceText,omitempty"  json:"TypeOfReferenceText,omitempty"`
}

//URIText ... A data item for a World Wide Web Consortium Uniform Reference Indicator
type URIText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"URIText,omitempty"  json:"URIText,omitempty"`
}

//UserInteractionText ... A data item for user interaction
type UserInteractionText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"UserInteractionText,omitempty"  json:"UserInteractionText,omitempty"`
}

//VectorStringText ... A data item for a vulnerability vector string
type VectorStringText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VectorStringText,omitempty"  json:"VectorStringText,omitempty"`
}

//VersionInformation ... A data item for version information
type VersionInformation struct {
	SoftwareVersionText            *SoftwareVersionText            `xml:"SoftwareVersionText,omitempty"  json:"SoftwareVersionText,omitempty"`
	SemanticVersionIndicator       *SemanticVersionIndicator       `xml:"SemanticVersionIndicator,omitempty"  json:"SemanticVersionIndicator,omitempty"`
	SemanticVersionText            *SemanticVersionText            `xml:"SemanticVersionText,omitempty"  json:"SemanticVersionText,omitempty"`
	SemanticVersionInformation     *SemanticVersionInformation     `xml:"SemanticVersionInformation,omitempty"  json:"SemanticVersionInformation,omitempty"`
	SemanticVersionDate            *SemanticVersionDate            `xml:"SemanticVersionDate,omitempty"  json:"SemanticVersionDate,omitempty"`
	PythonSemanticVersionIndicator *PythonSemanticVersionIndicator `xml:"PythonSemanticVersionIndicator,omitempty"  json:"PythonSemanticVersionIndicator,omitempty"`
	PythonSemanticVersion          *PythonSemanticVersion          `xml:"PythonSemanticVersion,omitempty"  json:"PythonSemanticVersion,omitempty"`
	PythonVersionText              *PythonVersionText              `xml:"PythonVersionText,omitempty"  json:"PythonVersionText,omitempty"`
	SecurityAttributesOptionGroup                                  //SecurityAttributesOptionGroup
	XMLName                        xml.Name                        `xml:"VersionInformation,omitempty"  json:"VersionInformation,omitempty"`
}

//VersionText ... A data item for a version
type VersionText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VersionText,omitempty"  json:"VersionText,omitempty"`
}

//VirusEngineText ... A data item for an entry associated with software repository commit event
type VirusEngineText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VirusEngineText,omitempty"  json:"VirusEngineText,omitempty"`
}

//VirusInformation ... A data item for information collected from virus scans
type VirusInformation struct {
	KnownVirusQuantity            *KnownVirusQuantity       `xml:"KnownVirusQuantity,omitempty"  json:"KnownVirusQuantity,omitempty"`
	EngineVersionText             *EngineVersionText        `xml:"EngineVersionText,omitempty"  json:"EngineVersionText,omitempty"`
	VirusEngineText               *VirusEngineText          `xml:"VirusEngineText,omitempty"  json:"VirusEngineText,omitempty"`
	ScannedDirectoryQuantity      *ScannedDirectoryQuantity `xml:"ScannedDirectoryQuantity,omitempty"  json:"ScannedDirectoryQuantity,omitempty"`
	ScannedFileQuantity           *ScannedFileQuantity      `xml:"ScannedFileQuantity,omitempty"  json:"ScannedFileQuantity,omitempty"`
	InfectedFileQuantity          *InfectedFileQuantity     `xml:"InfectedFileQuantity,omitempty"  json:"InfectedFileQuantity,omitempty"`
	DataScannedText               *DataScannedText          `xml:"DataScannedText,omitempty"  json:"DataScannedText,omitempty"`
	DataReadText                  *DataReadText             `xml:"DataReadText,omitempty"  json:"DataReadText,omitempty"`
	TimeText                      *TimeText                 `xml:"TimeText,omitempty"  json:"TimeText,omitempty"`
	FileNotesText                 *FileNotesText            `xml:"FileNotesText,omitempty"  json:"FileNotesText,omitempty"`
	EngineDetails                 *EngineDetails            `xml:"EngineDetails,omitempty"  json:"EngineDetails,omitempty"`
	SecurityAttributesOptionGroup                           //SecurityAttributesOptionGroup
	XMLName                       xml.Name                  `xml:"VirusInformation,omitempty"  json:"VirusInformation,omitempty"`
}

//Vulnerabilities ... A data item for vulnerabilty information
type Vulnerabilities struct {
	Vulnerability                 *Vulnerability `xml:"Vulnerability,omitempty"  json:"Vulnerability,omitempty"`
	SecurityAttributesOptionGroup                //SecurityAttributesOptionGroup
	XMLName                       xml.Name       `xml:"Vulnerabilities,omitempty"  json:"Vulnerabilities,omitempty"`
}

//Vulnerability ... A data item for vulnerabilty information
type Vulnerability struct {
	VulnerabilityIdentifierText     *VulnerabilityIdentifierText     `xml:"VulnerabilityIdentifierText,omitempty"  json:"VulnerabilityIdentifierText,omitempty"`
	SourceText                      *SourceText                      `xml:"SourceText,omitempty"  json:"SourceText,omitempty"`
	VulnerabilityTitleText          *VulnerabilityTitleText          `xml:"VulnerabilityTitleText,omitempty"  json:"VulnerabilityTitleText,omitempty"`
	VulnerabilitySummaryText        *VulnerabilitySummaryText        `xml:"VulnerabilitySummaryText,omitempty"  json:"VulnerabilitySummaryText,omitempty"`
	VulnerabilityScoreValue         *VulnerabilityScoreValue         `xml:"VulnerabilityScoreValue,omitempty"  json:"VulnerabilityScoreValue,omitempty"`
	VulnerabilityScoreVersionValue  *VulnerabilityScoreVersionValue  `xml:"VulnerabilityScoreVersionValue,omitempty"  json:"VulnerabilityScoreVersionValue,omitempty"`
	VulnerabilityScoreSystemText    *VulnerabilityScoreSystemText    `xml:"VulnerabilityScoreSystemText,omitempty"  json:"VulnerabilityScoreSystemText,omitempty"`
	ScoreDetails                    *ScoreDetails                    `xml:"ScoreDetails,omitempty"  json:"ScoreDetails,omitempty"`
	VulnerabilityVectorText         *VulnerabilityVectorText         `xml:"VulnerabilityVectorText,omitempty"  json:"VulnerabilityVectorText,omitempty"`
	AccessComplexityText            *AccessComplexityText            `xml:"AccessComplexityText,omitempty"  json:"AccessComplexityText,omitempty"`
	VulnerabilityAuthenticationText *VulnerabilityAuthenticationText `xml:"VulnerabilityAuthenticationText,omitempty"  json:"VulnerabilityAuthenticationText,omitempty"`
	ConfidentialityImpactText       *ConfidentialityImpactText       `xml:"ConfidentialityImpactText,omitempty"  json:"ConfidentialityImpactText,omitempty"`
	IntegrityImpactText             *IntegrityImpactText             `xml:"IntegrityImpactText,omitempty"  json:"IntegrityImpactText,omitempty"`
	AvailabilityImpactText          *AvailabilityImpactText          `xml:"AvailabilityImpactText,omitempty"  json:"AvailabilityImpactText,omitempty"`
	VulnerabilitySourceText         *VulnerabilitySourceText         `xml:"VulnerabilitySourceText,omitempty"  json:"VulnerabilitySourceText,omitempty"`
	AssessmentCheckText             *AssessmentCheckText             `xml:"AssessmentCheckText,omitempty"  json:"AssessmentCheckText,omitempty"`
	ScannerText                     *ScannerText                     `xml:"ScannerText,omitempty"  json:"ScannerText,omitempty"`
	RecommendationText              *RecommendationText              `xml:"RecommendationText,omitempty"  json:"RecommendationText,omitempty"`
	VulnerabilityReference          *VulnerabilityReference          `xml:"VulnerabilityReference,omitempty"  json:"VulnerabilityReference,omitempty"`
	VulnerabilityModifiedDate       *VulnerabilityModifiedDate       `xml:"VulnerabilityModifiedDate,omitempty"  json:"VulnerabilityModifiedDate,omitempty"`
	VulnerabilityPublishedDate      *VulnerabilityPublishedDate      `xml:"VulnerabilityPublishedDate,omitempty"  json:"VulnerabilityPublishedDate,omitempty"`
	SecurityAttributesOptionGroup                                    //SecurityAttributesOptionGroup
	XMLName                         xml.Name                         `xml:"Vulnerability,omitempty"  json:"Vulnerability,omitempty"`
}

//VulnerabilityAuthenticationText ... A data item for authentication
type VulnerabilityAuthenticationText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilityAuthenticationText,omitempty"  json:"VulnerabilityAuthenticationText,omitempty"`
}

//VulnerabilityBaseScoreValue ... A data item for a vulnerability base score
type VulnerabilityBaseScoreValue struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilityBaseScoreValue,omitempty"  json:"VulnerabilityBaseScoreValue,omitempty"`
}

//VulnerabilityIdentifierText ... A data item for a vulnerabilty identifier
type VulnerabilityIdentifierText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilityIdentifierText,omitempty"  json:"VulnerabilityIdentifierText,omitempty"`
}

//VulnerabilityInformation ... A data item for vulnerability or virus information for a software artifact or source repository
type VulnerabilityInformation struct {
	Vulnerabilities               *Vulnerabilities  `xml:"Vulnerabilities,omitempty"  json:"Vulnerabilities,omitempty"`
	VirusInformation              *VirusInformation `xml:"VirusInformation,omitempty"  json:"VirusInformation,omitempty"`
	SecurityAttributesOptionGroup                   //SecurityAttributesOptionGroup
	XMLName                       xml.Name          `xml:"VulnerabilityInformation,omitempty"  json:"VulnerabilityInformation,omitempty"`
}

//VulnerabilityModifiedDate ... A data type for date and time vulnerability modified
type VulnerabilityModifiedDate struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilityModifiedDate,omitempty"  json:"VulnerabilityModifiedDate,omitempty"`
}

//VulnerabilityPublishedDate ... A data type for date and time vulnerability published
type VulnerabilityPublishedDate struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilityPublishedDate,omitempty"  json:"VulnerabilityPublishedDate,omitempty"`
}

//VulnerabilityReference ... A data item for reference details for a vulnerability
type VulnerabilityReference struct {
	TypeOfReferenceText           *TypeOfReferenceText `xml:"TypeOfReferenceText,omitempty"  json:"TypeOfReferenceText,omitempty"`
	ReferenceSourceText           *ReferenceSourceText `xml:"ReferenceSourceText,omitempty"  json:"ReferenceSourceText,omitempty"`
	URIText                       *URIText             `xml:"URIText,omitempty"  json:"URIText,omitempty"`
	ReferenceText                 *ReferenceText       `xml:"ReferenceText,omitempty"  json:"ReferenceText,omitempty"`
	SecurityAttributesOptionGroup                      //SecurityAttributesOptionGroup
	XMLName                       xml.Name             `xml:"VulnerabilityReference,omitempty"  json:"VulnerabilityReference,omitempty"`
}

//VulnerabilityScopeText ... A data item for vulnerability scope
type VulnerabilityScopeText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilityScopeText,omitempty"  json:"VulnerabilityScopeText,omitempty"`
}

//VulnerabilityScoreSystemText ... A data item for a vulnerability score system
type VulnerabilityScoreSystemText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilityScoreSystemText,omitempty"  json:"VulnerabilityScoreSystemText,omitempty"`
}

//VulnerabilityScoreValue ... A data item for a vulnerability score
type VulnerabilityScoreValue struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilityScoreValue,omitempty"  json:"VulnerabilityScoreValue,omitempty"`
}

//VulnerabilityScoreVersionValue ... A data item for a vulnerability score version
type VulnerabilityScoreVersionValue struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilityScoreVersionValue,omitempty"  json:"VulnerabilityScoreVersionValue,omitempty"`
}

//VulnerabilitySourceText ... A data item for a vulnerabilty source
type VulnerabilitySourceText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilitySourceText,omitempty"  json:"VulnerabilitySourceText,omitempty"`
}

//VulnerabilitySummaryText ... A data item for a vulnerability summary paragraph
type VulnerabilitySummaryText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilitySummaryText,omitempty"  json:"VulnerabilitySummaryText,omitempty"`
}

//VulnerabilityTitleText ... A data item for a vulnerability title
type VulnerabilityTitleText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilityTitleText,omitempty"  json:"VulnerabilityTitleText,omitempty"`
}

//VulnerabilityVectorText ... A data item for a vector
type VulnerabilityVectorText struct {
	Data                          string   `xml:",chardata"  json:"Data,omitempty"`
	SecurityAttributesOptionGroup          //SecurityAttributesOptionGroup
	XMLName                       xml.Name `xml:"VulnerabilityVectorText,omitempty"  json:"VulnerabilityVectorText,omitempty"`
}
