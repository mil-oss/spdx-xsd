package main

import "encoding/xml"

//NewSoftwareEvidenceArchive ...
func NewSoftwareEvidenceArchive()*SoftwareEvidenceArchive{
    return &SoftwareEvidenceArchive{
        // Required for the proper namespacingSoftwareEvidenceArchive
        AttrXmlnsXsi:"http://www.w3.org/2001/XMLSchema-instance",
        AttrXmlns:"urn:seva::1.0",
        SoftwareInformation: &SoftwareInformation{
            VersionInformation:&VersionInformation{},
        },
        FileInformation: &FileInformation{},
        AuthoritativeSourceInformation: &AuthoritativeSourceInformation{},
        EcosystemInformation: &EcosystemInformation{
            Ecosystem:&Ecosystem{},
            CompanyInformation:&CompanyInformation{},
        },
        DependencyInformation: &DependencyInformation{
            Dependencies:&Dependencies{},
            Compiler:&Compiler{},
        },
        LicenseInformation: &LicenseInformation{},
        VulnerabilityInformation: &VulnerabilityInformation{
            Vulnerabilities:&Vulnerabilities{},
            VirusInformation:&VirusInformation{},
        },
        GovernanceRiskCompliance: &GovernanceRiskCompliance{},
        DeliveryInformation: &DeliveryInformation{},
    }
}
//SoftwareEvidenceArchive ... A data item for automated software supply chain metadata.
type SoftwareEvidenceArchive struct {
        AttrXmlnsXsi                             string                                   `xml:"xmlns:xsi,attr" json:"AttrXmlnsXsi,omitempty"`
        AttrXmlns                                string                                   `xml:"xmlns,attr" json:"AttrXmlns,omitempty"`
        SoftwareInformation                      *SoftwareInformation                     `xml:"SoftwareInformation,omitempty"  json:"SoftwareInformation,omitempty"`
        FileInformation                          *FileInformation                         `xml:"FileInformation,omitempty"  json:"FileInformation,omitempty"`
        AuthoritativeSourceInformation           *AuthoritativeSourceInformation          `xml:"AuthoritativeSourceInformation,omitempty"  json:"AuthoritativeSourceInformation,omitempty"`
        EcosystemInformation                     *EcosystemInformation                    `xml:"EcosystemInformation,omitempty"  json:"EcosystemInformation,omitempty"`
        DependencyInformation                    *DependencyInformation                   `xml:"DependencyInformation,omitempty"  json:"DependencyInformation,omitempty"`
        LicenseInformation                       *LicenseInformation                      `xml:"LicenseInformation,omitempty"  json:"LicenseInformation,omitempty"`
        VulnerabilityInformation                 *VulnerabilityInformation                `xml:"VulnerabilityInformation,omitempty"  json:"VulnerabilityInformation,omitempty"`
        GovernanceRiskCompliance                 *GovernanceRiskCompliance                `xml:"GovernanceRiskCompliance,omitempty"  json:"GovernanceRiskCompliance,omitempty"`
        DeliveryInformation                      *DeliveryInformation                     `xml:"DeliveryInformation,omitempty"  json:"DeliveryInformation,omitempty"`
        XMLName                                  xml.Name                                 `xml:"SoftwareEvidenceArchive,omitempty"  json:"SoftwareEvidenceArchive,omitempty"`
}
//AuthoritativeSourceInformation ... A data item for information about what is to be considered the authoritative source for a given artifact or source repository. This includes an URL, a hash of the source, and whether or not the source has been signed
type AuthoritativeSourceInformation struct {
        SourceHashText                           string                                   `xml:"SourceHashText,omitempty"  json:"SourceHashText,omitempty"`
        SourceURIText                            string                                   `xml:"SourceURIText,omitempty"  json:"SourceURIText,omitempty"`
        AuthoritativeDigitalSignatureIndicator   string                                   `xml:"AuthoritativeDigitalSignatureIndicator,omitempty"  json:"AuthoritativeDigitalSignatureIndicator,omitempty"`
        XMLName                                  xml.Name                                 `xml:"AuthoritativeSourceInformation,omitempty"  json:"AuthoritativeSourceInformation,omitempty"`
}
//CommitLogs ... A data item for entries associated with software repository commit events
type CommitLogs struct {
        CommitLogText                            []string                                 `xml:"CommitLogText,omitempty"  json:"CommitLogText[],omitempty"`
        XMLName                                  xml.Name                                 `xml:"CommitLogs,omitempty"  json:"CommitLogs,omitempty"`
}
//CommonVulnerabilityScoringSystemV2 ... A data type for Common Vulnerability Scoring System V2 (CVSSV2) Information
type CommonVulnerabilityScoringSystemV2 struct {
        VectorStringText                         string                                   `xml:"VectorStringText,omitempty"  json:"VectorStringText,omitempty"`
        AccessVectorText                         string                                   `xml:"AccessVectorText,omitempty"  json:"AccessVectorText,omitempty"`
        AccessComplexityText                     string                                   `xml:"AccessComplexityText,omitempty"  json:"AccessComplexityText,omitempty"`
        VulnerabilityAuthenticationText          string                                   `xml:"VulnerabilityAuthenticationText,omitempty"  json:"VulnerabilityAuthenticationText,omitempty"`
        ConfidentialityImpactText                string                                   `xml:"ConfidentialityImpactText,omitempty"  json:"ConfidentialityImpactText,omitempty"`
        IntegrityImpactText                      string                                   `xml:"IntegrityImpactText,omitempty"  json:"IntegrityImpactText,omitempty"`
        AvailabilityImpactText                   string                                   `xml:"AvailabilityImpactText,omitempty"  json:"AvailabilityImpactText,omitempty"`
        VulnerabilityBaseScoreValue              string                                   `xml:"VulnerabilityBaseScoreValue,omitempty"  json:"VulnerabilityBaseScoreValue,omitempty"`
        XMLName                                  xml.Name                                 `xml:"CommonVulnerabilityScoringSystemV2,omitempty"  json:"CommonVulnerabilityScoringSystemV2,omitempty"`
}
//CommonVulnerabilityScoringSystemV3 ... A data type for Common Vulnerability Scoring System V3 (CVSSV3) Information
type CommonVulnerabilityScoringSystemV3 struct {
        VectorStringText                         string                                   `xml:"VectorStringText,omitempty"  json:"VectorStringText,omitempty"`
        AccessVectorText                         string                                   `xml:"AccessVectorText,omitempty"  json:"AccessVectorText,omitempty"`
        AccessComplexityText                     string                                   `xml:"AccessComplexityText,omitempty"  json:"AccessComplexityText,omitempty"`
        PrivilegesRequiredText                   string                                   `xml:"PrivilegesRequiredText,omitempty"  json:"PrivilegesRequiredText,omitempty"`
        UserInteractionText                      string                                   `xml:"UserInteractionText,omitempty"  json:"UserInteractionText,omitempty"`
        VulnerabilityScopeText                   string                                   `xml:"VulnerabilityScopeText,omitempty"  json:"VulnerabilityScopeText,omitempty"`
        ConfidentialityImpactText                string                                   `xml:"ConfidentialityImpactText,omitempty"  json:"ConfidentialityImpactText,omitempty"`
        IntegrityImpactText                      string                                   `xml:"IntegrityImpactText,omitempty"  json:"IntegrityImpactText,omitempty"`
        AvailabilityImpactText                   string                                   `xml:"AvailabilityImpactText,omitempty"  json:"AvailabilityImpactText,omitempty"`
        VulnerabilityBaseScoreValue              string                                   `xml:"VulnerabilityBaseScoreValue,omitempty"  json:"VulnerabilityBaseScoreValue,omitempty"`
        BaseSeverityText                         string                                   `xml:"BaseSeverityText,omitempty"  json:"BaseSeverityText,omitempty"`
        XMLName                                  xml.Name                                 `xml:"CommonVulnerabilityScoringSystemV3,omitempty"  json:"CommonVulnerabilityScoringSystemV3,omitempty"`
}
//CompanyInformation ... A data item for languages used in a software repository
type CompanyInformation struct {
        Name                                     string                                   `xml:"Name,omitempty"  json:"Name,omitempty"`
        XMLName                                  xml.Name                                 `xml:"CompanyInformation,omitempty"  json:"CompanyInformation,omitempty"`
}
//Compiler ... A data item for languages used in a software repository
type Compiler struct {
        Name                                     string                                   `xml:"Name,omitempty"  json:"Name,omitempty"`
        CompilerVersionText                      string                                   `xml:"CompilerVersionText,omitempty"  json:"CompilerVersionText,omitempty"`
        XMLName                                  xml.Name                                 `xml:"Compiler,omitempty"  json:"Compiler,omitempty"`
}
//DeliveryInformation ... A data item for delivery information including the target url and last delivered date and time for a previously signed Software Evidence Archive.
type DeliveryInformation struct {
        DeliveringOrganizationText               string                                   `xml:"DeliveringOrganizationText,omitempty"  json:"DeliveringOrganizationText,omitempty"`
        DestinationText                          string                                   `xml:"DestinationText,omitempty"  json:"DestinationText,omitempty"`
        HashText                                 string                                   `xml:"HashText,omitempty"  json:"HashText,omitempty"`
        LastRegistryUpdateDate                   string                                   `xml:"LastRegistryUpdateDate,omitempty"  json:"LastRegistryUpdateDate,omitempty"`
        XMLName                                  xml.Name                                 `xml:"DeliveryInformation,omitempty"  json:"DeliveryInformation,omitempty"`
}
//Dependencies ... A data item for software dependencies
type Dependencies struct {
        Dependency                               []Dependency                             `xml:"Dependency,omitempty"  json:"Dependency[],omitempty"`
        XMLName                                  xml.Name                                 `xml:"Dependencies,omitempty"  json:"Dependencies,omitempty"`
}
//Dependency ... A data item for software dependency information
type Dependency struct {
        Name                                     string                                   `xml:"Name,omitempty"  json:"Name,omitempty"`
        DependencyVersionText                    string                                   `xml:"DependencyVersionText,omitempty"  json:"DependencyVersionText,omitempty"`
        XMLName                                  xml.Name                                 `xml:"Dependency,omitempty"  json:"Dependency,omitempty"`
}
//DependencyInformation ... A data item for a list of dependencies for a given software project derived directly from the artifact or source dependency definition file. Entries include names, versions, and vulnerabilities
type DependencyInformation struct {
        Dependencies                             *Dependencies                            `xml:"Dependencies,omitempty"  json:"Dependencies,omitempty"`
        Compiler                                 *Compiler                                `xml:"Compiler,omitempty"  json:"Compiler,omitempty"`
        XMLName                                  xml.Name                                 `xml:"DependencyInformation,omitempty"  json:"DependencyInformation,omitempty"`
}
//Ecosystem ... A data item for software ecosystem information
type Ecosystem struct {
        CommitterQuantity                        string                                   `xml:"CommitterQuantity,omitempty"  json:"CommitterQuantity,omitempty"`
        Languages                                *Languages                               `xml:"Languages,omitempty"  json:"Languages,omitempty"`
        CommitLogs                               *CommitLogs                              `xml:"CommitLogs,omitempty"  json:"CommitLogs,omitempty"`
        XMLName                                  xml.Name                                 `xml:"Ecosystem,omitempty"  json:"Ecosystem,omitempty"`
}
//EcosystemInformation ... A data item for information pertaining to a software project's ecosystem which includes programming languages, references, number of committers, mailing list activity, overall sentiment, and other information
type EcosystemInformation struct {
        Ecosystem                                *Ecosystem                               `xml:"Ecosystem,omitempty"  json:"Ecosystem,omitempty"`
        CompanyInformation                       *CompanyInformation                      `xml:"CompanyInformation,omitempty"  json:"CompanyInformation,omitempty"`
        XMLName                                  xml.Name                                 `xml:"EcosystemInformation,omitempty"  json:"EcosystemInformation,omitempty"`
}
//EngineDetails ... A data item for for details of a virus engine
type EngineDetails struct {
        EngineVersionText                        string                                   `xml:"EngineVersionText,omitempty"  json:"EngineVersionText,omitempty"`
        DatabaseVersionText                      string                                   `xml:"DatabaseVersionText,omitempty"  json:"DatabaseVersionText,omitempty"`
        XMLName                                  xml.Name                                 `xml:"EngineDetails,omitempty"  json:"EngineDetails,omitempty"`
}
//FileInformation ... A data item for file types and related Multipurpose Internet Mail Extensions (MIME) types
type FileInformation struct {
        FileNameText                             string                                   `xml:"FileNameText,omitempty"  json:"FileNameText,omitempty"`
        FileExtensionText                        string                                   `xml:"FileExtensionText,omitempty"  json:"FileExtensionText,omitempty"`
        XMLName                                  xml.Name                                 `xml:"FileInformation,omitempty"  json:"FileInformation,omitempty"`
}
//GovernanceRiskCompliance ... A data item for calculated risk and compliance to governance for a software artifact or source repository
type GovernanceRiskCompliance struct {
        RiskCode                                 string                                   `xml:"RiskCode,omitempty"  json:"RiskCode,omitempty"`
        StatementOfAssuranceText                 string                                   `xml:"StatementOfAssuranceText,omitempty"  json:"StatementOfAssuranceText,omitempty"`
        XMLName                                  xml.Name                                 `xml:"GovernanceRiskCompliance,omitempty"  json:"GovernanceRiskCompliance,omitempty"`
}
//Languages ... A data item for languages used in a software repository
type Languages struct {
        LanguageText                             []string                                 `xml:"LanguageText,omitempty"  json:"LanguageText[],omitempty"`
        XMLName                                  xml.Name                                 `xml:"Languages,omitempty"  json:"Languages,omitempty"`
}
//LicenseInformation ... A data item for software license information for a software artifact or source repository
type LicenseInformation struct {
        LicenseCategoryCode                      string                                   `xml:"LicenseCategoryCode,omitempty"  json:"LicenseCategoryCode,omitempty"`
        LicenseCode                              []string                                 `xml:"LicenseCode,omitempty"  json:"LicenseCode[],omitempty"`
        EndOfLifeIndicator                       string                                   `xml:"EndOfLifeIndicator,omitempty"  json:"EndOfLifeIndicator,omitempty"`
        XMLName                                  xml.Name                                 `xml:"LicenseInformation,omitempty"  json:"LicenseInformation,omitempty"`
}
//PythonSemanticVersion ... A data item for a Python Distribution identified by a public version identifier which supports all defined version comparison operations. The version scheme is used both to describe the distribution version provided by a particular distribution archive, as well as to place constraints on the version of dependencies needed in order to build or run the software.The canonical public version identifiers MUST comply with the following scheme: [N!]N(.N)*[{a|b|rc}N][.postN][.devN] (https://www.python.org/dev/peps/pep-0440/#public-version-identifiers)
type PythonSemanticVersion struct {
        PythonEpochSegmentText                   string                                   `xml:"PythonEpochSegmentText,omitempty"  json:"PythonEpochSegmentText,omitempty"`
        PythonReleaseSegmentText                 string                                   `xml:"PythonReleaseSegmentText,omitempty"  json:"PythonReleaseSegmentText,omitempty"`
        PythonPreReleaseSegmentText              string                                   `xml:"PythonPreReleaseSegmentText,omitempty"  json:"PythonPreReleaseSegmentText,omitempty"`
        PythonPostReleaseSegmentText             string                                   `xml:"PythonPostReleaseSegmentText,omitempty"  json:"PythonPostReleaseSegmentText,omitempty"`
        PythonDevelopmentReleaseSegmentText      string                                   `xml:"PythonDevelopmentReleaseSegmentText,omitempty"  json:"PythonDevelopmentReleaseSegmentText,omitempty"`
        XMLName                                  xml.Name                                 `xml:"PythonSemanticVersion,omitempty"  json:"PythonSemanticVersion,omitempty"`
}
//ScoreDetails ... A data item for common vulnerability score details
type ScoreDetails struct {
        CommonVulnerabilityScoringSystemV2       *CommonVulnerabilityScoringSystemV2      `xml:"CommonVulnerabilityScoringSystemV2,omitempty"  json:"CommonVulnerabilityScoringSystemV2,omitempty"`
        CommonVulnerabilityScoringSystemV3       *CommonVulnerabilityScoringSystemV3      `xml:"CommonVulnerabilityScoringSystemV3,omitempty"  json:"CommonVulnerabilityScoringSystemV3,omitempty"`
        XMLName                                  xml.Name                                 `xml:"ScoreDetails,omitempty"  json:"ScoreDetails,omitempty"`
}
//SemanticVersionInformation ... A data type for semantic versioning information. A normal version number MUST take the form X.Y.Z where X, Y, and Z are non-negative integers, and MUST NOT contain leading zeroes. X is the major version, Y is the minor version, and Z is the patch version. Each element MUST increase numerically. For instance: 1.9.0 -> 1.10.0 -> 1.11.0.(https://semver.org/spec/v2.0.0.html)
type SemanticVersionInformation struct {
        APIName                                  string                                   `xml:"APIName,omitempty"  json:"APIName,omitempty"`
        MajorVersionNumeric                      string                                   `xml:"MajorVersionNumeric,omitempty"  json:"MajorVersionNumeric,omitempty"`
        MinorVersionNumeric                      string                                   `xml:"MinorVersionNumeric,omitempty"  json:"MinorVersionNumeric,omitempty"`
        PatchVersionText                         string                                   `xml:"PatchVersionText,omitempty"  json:"PatchVersionText,omitempty"`
        PreReleaseVersionText                    string                                   `xml:"PreReleaseVersionText,omitempty"  json:"PreReleaseVersionText,omitempty"`
        BuildMetaText                            string                                   `xml:"BuildMetaText,omitempty"  json:"BuildMetaText,omitempty"`
        XMLName                                  xml.Name                                 `xml:"SemanticVersionInformation,omitempty"  json:"SemanticVersionInformation,omitempty"`
}
//SoftwareInformation ... A data item for software product naming or version related information
type SoftwareInformation struct {
        ProductTitleText                         string                                   `xml:"ProductTitleText,omitempty"  json:"ProductTitleText,omitempty"`
        SoftwareNameText                         string                                   `xml:"SoftwareNameText,omitempty"  json:"SoftwareNameText,omitempty"`
        SoftwareOrgText                          string                                   `xml:"SoftwareOrgText,omitempty"  json:"SoftwareOrgText,omitempty"`
        VersionInformation                       *VersionInformation                      `xml:"VersionInformation,omitempty"  json:"VersionInformation,omitempty"`
        GroupingText                             string                                   `xml:"GroupingText,omitempty"  json:"GroupingText,omitempty"`
        ReleaseNotesText                         string                                   `xml:"ReleaseNotesText,omitempty"  json:"ReleaseNotesText,omitempty"`
        XMLName                                  xml.Name                                 `xml:"SoftwareInformation,omitempty"  json:"SoftwareInformation,omitempty"`
}
//VersionInformation ... A data item for version information
type VersionInformation struct {
        SoftwareVersionText                      string                                   `xml:"SoftwareVersionText,omitempty"  json:"SoftwareVersionText,omitempty"`
        SemanticVersionIndicator                 string                                   `xml:"SemanticVersionIndicator,omitempty"  json:"SemanticVersionIndicator,omitempty"`
        SemanticVersionText                      string                                   `xml:"SemanticVersionText,omitempty"  json:"SemanticVersionText,omitempty"`
        SemanticVersionInformation               *SemanticVersionInformation              `xml:"SemanticVersionInformation,omitempty"  json:"SemanticVersionInformation,omitempty"`
        SemanticVersionDate                      string                                   `xml:"SemanticVersionDate,omitempty"  json:"SemanticVersionDate,omitempty"`
        PythonSemanticVersionIndicator           string                                   `xml:"PythonSemanticVersionIndicator,omitempty"  json:"PythonSemanticVersionIndicator,omitempty"`
        PythonSemanticVersion                    *PythonSemanticVersion                   `xml:"PythonSemanticVersion,omitempty"  json:"PythonSemanticVersion,omitempty"`
        PythonVersionText                        string                                   `xml:"PythonVersionText,omitempty"  json:"PythonVersionText,omitempty"`
        XMLName                                  xml.Name                                 `xml:"VersionInformation,omitempty"  json:"VersionInformation,omitempty"`
}
//VirusInformation ... A data item for information collected from virus scans
type VirusInformation struct {
        KnownVirusQuantity                       string                                   `xml:"KnownVirusQuantity,omitempty"  json:"KnownVirusQuantity,omitempty"`
        EngineVersionText                        string                                   `xml:"EngineVersionText,omitempty"  json:"EngineVersionText,omitempty"`
        VirusEngineText                          string                                   `xml:"VirusEngineText,omitempty"  json:"VirusEngineText,omitempty"`
        ScannedDirectoryQuantity                 string                                   `xml:"ScannedDirectoryQuantity,omitempty"  json:"ScannedDirectoryQuantity,omitempty"`
        ScannedFileQuantity                      string                                   `xml:"ScannedFileQuantity,omitempty"  json:"ScannedFileQuantity,omitempty"`
        InfectedFileQuantity                     string                                   `xml:"InfectedFileQuantity,omitempty"  json:"InfectedFileQuantity,omitempty"`
        DataScannedText                          string                                   `xml:"DataScannedText,omitempty"  json:"DataScannedText,omitempty"`
        DataReadText                             string                                   `xml:"DataReadText,omitempty"  json:"DataReadText,omitempty"`
        TimeText                                 string                                   `xml:"TimeText,omitempty"  json:"TimeText,omitempty"`
        FileNotesText                            string                                   `xml:"FileNotesText,omitempty"  json:"FileNotesText,omitempty"`
        EngineDetails                            *EngineDetails                           `xml:"EngineDetails,omitempty"  json:"EngineDetails,omitempty"`
        XMLName                                  xml.Name                                 `xml:"VirusInformation,omitempty"  json:"VirusInformation,omitempty"`
}
//Vulnerabilities ... A data item for vulnerabilty information
type Vulnerabilities struct {
        Vulnerability                            []Vulnerability                          `xml:"Vulnerability,omitempty"  json:"Vulnerability[],omitempty"`
        XMLName                                  xml.Name                                 `xml:"Vulnerabilities,omitempty"  json:"Vulnerabilities,omitempty"`
}
//Vulnerability ... A data item for vulnerabilty information
type Vulnerability struct {
        VulnerabilityIdentifierText              string                                   `xml:"VulnerabilityIdentifierText,omitempty"  json:"VulnerabilityIdentifierText,omitempty"`
        SourceText                               string                                   `xml:"SourceText,omitempty"  json:"SourceText,omitempty"`
        VulnerabilityTitleText                   string                                   `xml:"VulnerabilityTitleText,omitempty"  json:"VulnerabilityTitleText,omitempty"`
        VulnerabilitySummaryText                 string                                   `xml:"VulnerabilitySummaryText,omitempty"  json:"VulnerabilitySummaryText,omitempty"`
        VulnerabilityScoreValue                  string                                   `xml:"VulnerabilityScoreValue,omitempty"  json:"VulnerabilityScoreValue,omitempty"`
        VulnerabilityScoreVersionValue           string                                   `xml:"VulnerabilityScoreVersionValue,omitempty"  json:"VulnerabilityScoreVersionValue,omitempty"`
        VulnerabilityScoreSystemText             string                                   `xml:"VulnerabilityScoreSystemText,omitempty"  json:"VulnerabilityScoreSystemText,omitempty"`
        ScoreDetails                             *ScoreDetails                            `xml:"ScoreDetails,omitempty"  json:"ScoreDetails,omitempty"`
        VulnerabilityVectorText                  string                                   `xml:"VulnerabilityVectorText,omitempty"  json:"VulnerabilityVectorText,omitempty"`
        AccessComplexityText                     string                                   `xml:"AccessComplexityText,omitempty"  json:"AccessComplexityText,omitempty"`
        VulnerabilityAuthenticationText          string                                   `xml:"VulnerabilityAuthenticationText,omitempty"  json:"VulnerabilityAuthenticationText,omitempty"`
        ConfidentialityImpactText                string                                   `xml:"ConfidentialityImpactText,omitempty"  json:"ConfidentialityImpactText,omitempty"`
        IntegrityImpactText                      string                                   `xml:"IntegrityImpactText,omitempty"  json:"IntegrityImpactText,omitempty"`
        AvailabilityImpactText                   string                                   `xml:"AvailabilityImpactText,omitempty"  json:"AvailabilityImpactText,omitempty"`
        VulnerabilitySourceText                  string                                   `xml:"VulnerabilitySourceText,omitempty"  json:"VulnerabilitySourceText,omitempty"`
        AssessmentCheckText                      string                                   `xml:"AssessmentCheckText,omitempty"  json:"AssessmentCheckText,omitempty"`
        ScannerText                              string                                   `xml:"ScannerText,omitempty"  json:"ScannerText,omitempty"`
        RecommendationText                       string                                   `xml:"RecommendationText,omitempty"  json:"RecommendationText,omitempty"`
        VulnerabilityReference                   *VulnerabilityReference                  `xml:"VulnerabilityReference,omitempty"  json:"VulnerabilityReference,omitempty"`
        VulnerabilityModifiedDate                string                                   `xml:"VulnerabilityModifiedDate,omitempty"  json:"VulnerabilityModifiedDate,omitempty"`
        VulnerabilityPublishedDate               string                                   `xml:"VulnerabilityPublishedDate,omitempty"  json:"VulnerabilityPublishedDate,omitempty"`
        XMLName                                  xml.Name                                 `xml:"Vulnerability,omitempty"  json:"Vulnerability,omitempty"`
}
//VulnerabilityInformation ... A data item for vulnerability or virus information for a software artifact or source repository
type VulnerabilityInformation struct {
        Vulnerabilities                          *Vulnerabilities                         `xml:"Vulnerabilities,omitempty"  json:"Vulnerabilities,omitempty"`
        VirusInformation                         *VirusInformation                        `xml:"VirusInformation,omitempty"  json:"VirusInformation,omitempty"`
        XMLName                                  xml.Name                                 `xml:"VulnerabilityInformation,omitempty"  json:"VulnerabilityInformation,omitempty"`
}
//VulnerabilityReference ... A data item for reference details for a vulnerability
type VulnerabilityReference struct {
        TypeOfReferenceText                      string                                   `xml:"TypeOfReferenceText,omitempty"  json:"TypeOfReferenceText,omitempty"`
        ReferenceSourceText                      string                                   `xml:"ReferenceSourceText,omitempty"  json:"ReferenceSourceText,omitempty"`
        URIText                                  string                                   `xml:"URIText,omitempty"  json:"URIText,omitempty"`
        ReferenceText                            string                                   `xml:"ReferenceText,omitempty"  json:"ReferenceText,omitempty"`
        XMLName                                  xml.Name                                 `xml:"VulnerabilityReference,omitempty"  json:"VulnerabilityReference,omitempty"`
}
