package main

import "encoding/xml"

//NewLicense ...
func NewLicense() *License {
	return &License{
		// Required for the proper namespacingLicense
		AttrXmlnsXsi: "http://www.w3.org/2001/XMLSchema-instance",
		AttrXmlns:    "spdx:xsd::1.0",
	}
}

//License ...
type License struct {
	AttrXmlnsXsi            string   `xml:"xmlns:xsi,attr" json:"AttrXmlnsXsi,omitempty"`
	AttrXmlns               string   `xml:"xmlns,attr" json:"AttrXmlns,omitempty"`
	IsDeprecatedLicenseID   string   `xml:"IsDeprecatedLicenseId,omitempty"  json:"IsDeprecatedLicenseId,omitempty"`
	StandardLicenseHeader   string   `xml:"StandardLicenseHeader,omitempty"  json:"StandardLicenseHeader,omitempty"`
	StandardLicenseTemplate string   `xml:"StandardLicenseTemplate,omitempty"  json:"StandardLicenseTemplate,omitempty"`
	LicenseText             string   `xml:"LicenseText,omitempty"  json:"LicenseText,omitempty"`
	IsOsiApproved           string   `xml:"IsOsiApproved,omitempty"  json:"IsOsiApproved,omitempty"`
	IsFsfLibre              string   `xml:"IsFsfLibre,omitempty"  json:"IsFsfLibre,omitempty"`
	LicenseID               string   `xml:"LicenseId,omitempty"  json:"LicenseId,omitempty"`
	Name                    string   `xml:"Name,omitempty"  json:"Name,omitempty"`
	SeeAlso                 []string `xml:"SeeAlso,omitempty"  json:"SeeAlso[],omitempty"`
	Comment                 string   `xml:"Comment,omitempty"  json:"Comment,omitempty"`
	XMLName                 xml.Name `xml:"License,omitempty"  json:"License,omitempty"`
}

//Annotation ...
type Annotation struct {
	Date               string   `xml:"Date,omitempty"  json:"Date,omitempty"`
	AnnotationTypeCode string   `xml:"AnnotationTypeCode,omitempty"  json:"AnnotationTypeCode,omitempty"`
	Comment            string   `xml:"Comment,omitempty"  json:"Comment,omitempty"`
	Annotator          string   `xml:"Annotator,omitempty"  json:"Annotator,omitempty"`
	XMLName            xml.Name `xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
}

//Checksum ...
type Checksum struct {
	ChecksumValue string   `xml:"ChecksumValue,omitempty"  json:"ChecksumValue,omitempty"`
	AlgorithmCode string   `xml:"AlgorithmCode,omitempty"  json:"AlgorithmCode,omitempty"`
	XMLName       xml.Name `xml:"Checksum,omitempty"  json:"Checksum,omitempty"`
}

//ConjunctiveLicenseSet ...
type ConjunctiveLicenseSet struct {
	Container      string   `xml:"Container,omitempty"  json:"Container,omitempty"`
	AnyLicenseInfo string   `xml:"AnyLicenseInfo,omitempty"  json:"AnyLicenseInfo,omitempty"`
	XMLName        xml.Name `xml:"ConjunctiveLicenseSet,omitempty"  json:"ConjunctiveLicenseSet,omitempty"`
}

//CreationInfo ...
type CreationInfo struct {
	LicenseListVersion string   `xml:"LicenseListVersion,omitempty"  json:"LicenseListVersion,omitempty"`
	Created            string   `xml:"Created,omitempty"  json:"Created,omitempty"`
	Comment            string   `xml:"Comment,omitempty"  json:"Comment,omitempty"`
	Creator            string   `xml:"Creator,omitempty"  json:"Creator,omitempty"`
	XMLName            xml.Name `xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
}

//DescribesFile ... A data item for The describesFile property relates an SpdxDocument to the file which it describes
type DescribesFile struct {
	SpdxElement          *SpdxElement          `xml:"SpdxElement,omitempty"  json:"SpdxElement,omitempty"`
	Package              *Package              `xml:"Package,omitempty"  json:"Package,omitempty"`
	File                 *File                 `xml:"File,omitempty"  json:"File,omitempty"`
	ExtractedLicenseInfo *ExtractedLicenseInfo `xml:"ExtractedLicenseInfo,omitempty"  json:"ExtractedLicenseInfo,omitempty"`
	CreationInfo         *CreationInfo         `xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
	SpecVersion          string                `xml:"SpecVersion,omitempty"  json:"SpecVersion,omitempty"`
	ExternalDocumentRef  *ExternalDocumentRef  `xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
	DataLicense          string                `xml:"DataLicense,omitempty"  json:"DataLicense,omitempty"`
	XMLName              xml.Name              `xml:"DescribesFile,omitempty"  json:"DescribesFile,omitempty"`
}

//DescribesPackage ... A data item for The describesPackage property relates an SpdxDocument to the package which it describes
type DescribesPackage struct {
	SpdxElement          *SpdxElement          `xml:"SpdxElement,omitempty"  json:"SpdxElement,omitempty"`
	Package              *Package              `xml:"Package,omitempty"  json:"Package,omitempty"`
	File                 *File                 `xml:"File,omitempty"  json:"File,omitempty"`
	ExtractedLicenseInfo *ExtractedLicenseInfo `xml:"ExtractedLicenseInfo,omitempty"  json:"ExtractedLicenseInfo,omitempty"`
	CreationInfo         *CreationInfo         `xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
	SpecVersion          string                `xml:"SpecVersion,omitempty"  json:"SpecVersion,omitempty"`
	ExternalDocumentRef  *ExternalDocumentRef  `xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
	DataLicense          string                `xml:"DataLicense,omitempty"  json:"DataLicense,omitempty"`
	XMLName              xml.Name              `xml:"DescribesPackage,omitempty"  json:"DescribesPackage,omitempty"`
}

//DisjunctiveLicenseSet ...
type DisjunctiveLicenseSet struct {
	Container      string   `xml:"Container,omitempty"  json:"Container,omitempty"`
	AnyLicenseInfo string   `xml:"AnyLicenseInfo,omitempty"  json:"AnyLicenseInfo,omitempty"`
	XMLName        xml.Name `xml:"DisjunctiveLicenseSet,omitempty"  json:"DisjunctiveLicenseSet,omitempty"`
}

//ExternalDocumentRef ...
type ExternalDocumentRef struct {
	Checksum           *Checksum     `xml:"Checksum,omitempty"  json:"Checksum,omitempty"`
	ExternalDocumentID string        `xml:"ExternalDocumentId,omitempty"  json:"ExternalDocumentId,omitempty"`
	SpdxDocument       *SpdxDocument `xml:"SpdxDocument,omitempty"  json:"SpdxDocument,omitempty"`
	XMLName            xml.Name      `xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
}

//ExternalRef ...
type ExternalRef struct {
	Comment               string         `xml:"Comment,omitempty"  json:"Comment,omitempty"`
	ReferenceLocator      string         `xml:"ReferenceLocator,omitempty"  json:"ReferenceLocator,omitempty"`
	ReferenceType         *ReferenceType `xml:"ReferenceType,omitempty"  json:"ReferenceType,omitempty"`
	ReferenceCategoryCode string         `xml:"ReferenceCategoryCode,omitempty"  json:"ReferenceCategoryCode,omitempty"`
	XMLName               xml.Name       `xml:"ExternalRef,omitempty"  json:"ExternalRef,omitempty"`
}

//ExtractedLicenseInfo ...
type ExtractedLicenseInfo struct {
	SimpleLicensingInfo *SimpleLicensingInfo `xml:"SimpleLicensingInfo,omitempty"  json:"SimpleLicensingInfo,omitempty"`
	ExtractedText       string               `xml:"ExtractedText,omitempty"  json:"ExtractedText,omitempty"`
	XMLName             xml.Name             `xml:"ExtractedLicenseInfo,omitempty"  json:"ExtractedLicenseInfo,omitempty"`
}

//File ...
type File struct {
	SpdxItem        *SpdxItem `xml:"SpdxItem,omitempty"  json:"SpdxItem,omitempty"`
	NoticeText      string    `xml:"NoticeText,omitempty"  json:"NoticeText,omitempty"`
	FileContributor string    `xml:"FileContributor,omitempty"  json:"FileContributor,omitempty"`
	Project         string    `xml:"Project,omitempty"  json:"Project,omitempty"`
	FileName        string    `xml:"FileName,omitempty"  json:"FileName,omitempty"`
	FileTypeCode    string    `xml:"FileTypeCode,omitempty"  json:"FileTypeCode,omitempty"`
	Checksum        *Checksum `xml:"Checksum,omitempty"  json:"Checksum,omitempty"`
	XMLName         xml.Name  `xml:"File,omitempty"  json:"File,omitempty"`
}

//HasExtractedLicensingInfo ... A data item for Indicates that a particular ExtractedLicensingInfo was defined in the subject SpdxDocument
type HasExtractedLicensingInfo struct {
	SpdxElement          *SpdxElement          `xml:"SpdxElement,omitempty"  json:"SpdxElement,omitempty"`
	Package              *Package              `xml:"Package,omitempty"  json:"Package,omitempty"`
	File                 *File                 `xml:"File,omitempty"  json:"File,omitempty"`
	ExtractedLicenseInfo *ExtractedLicenseInfo `xml:"ExtractedLicenseInfo,omitempty"  json:"ExtractedLicenseInfo,omitempty"`
	CreationInfo         *CreationInfo         `xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
	SpecVersion          string                `xml:"SpecVersion,omitempty"  json:"SpecVersion,omitempty"`
	ExternalDocumentRef  *ExternalDocumentRef  `xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
	DataLicense          string                `xml:"DataLicense,omitempty"  json:"DataLicense,omitempty"`
	XMLName              xml.Name              `xml:"HasExtractedLicensingInfo,omitempty"  json:"HasExtractedLicensingInfo,omitempty"`
}

//HasFile ... A data item for Indicates that a particular file belongs to a package
type HasFile struct {
	SpdxItem                *SpdxItem                `xml:"SpdxItem,omitempty"  json:"SpdxItem,omitempty"`
	ExternalRef             *ExternalRef             `xml:"ExternalRef,omitempty"  json:"ExternalRef,omitempty"`
	PackageFileName         string                   `xml:"PackageFileName,omitempty"  json:"PackageFileName,omitempty"`
	FilesAnalyzed           string                   `xml:"FilesAnalyzed,omitempty"  json:"FilesAnalyzed,omitempty"`
	Description             string                   `xml:"Description,omitempty"  json:"Description,omitempty"`
	Homepage                string                   `xml:"Homepage,omitempty"  json:"Homepage,omitempty"`
	Originator              string                   `xml:"Originator,omitempty"  json:"Originator,omitempty"`
	DownloadLocation        string                   `xml:"DownloadLocation,omitempty"  json:"DownloadLocation,omitempty"`
	SourceInfo              string                   `xml:"SourceInfo,omitempty"  json:"SourceInfo,omitempty"`
	Supplier                string                   `xml:"Supplier,omitempty"  json:"Supplier,omitempty"`
	PackageVerificationCode *PackageVerificationCode `xml:"PackageVerificationCode,omitempty"  json:"PackageVerificationCode,omitempty"`
	LicenseDeclaredCode     string                   `xml:"LicenseDeclaredCode,omitempty"  json:"LicenseDeclaredCode,omitempty"`
	VersionInfo             string                   `xml:"VersionInfo,omitempty"  json:"VersionInfo,omitempty"`
	Summary                 string                   `xml:"Summary,omitempty"  json:"Summary,omitempty"`
	Checksum                *Checksum                `xml:"Checksum,omitempty"  json:"Checksum,omitempty"`
	XMLName                 xml.Name                 `xml:"HasFile,omitempty"  json:"HasFile,omitempty"`
}

//LicenseDeclared ... A data item for The licensing that the creators of the software in the package, or the packager, have declared
type LicenseDeclared struct {
	SpdxElement          *SpdxElement         `xml:"SpdxElement,omitempty"  json:"SpdxElement,omitempty"`
	LicenseConcludedCode string               `xml:"LicenseConcludedCode,omitempty"  json:"LicenseConcludedCode,omitempty"`
	CopyrightText        string               `xml:"CopyrightText,omitempty"  json:"CopyrightText,omitempty"`
	SimpleLicensingInfo  *SimpleLicensingInfo `xml:"SimpleLicensingInfo,omitempty"  json:"SimpleLicensingInfo,omitempty"`
	LicenseComments      string               `xml:"LicenseComments,omitempty"  json:"LicenseComments,omitempty"`
	XMLName              xml.Name             `xml:"LicenseDeclared,omitempty"  json:"LicenseDeclared,omitempty"`
}

//LicenseException ...
type LicenseException struct {
	LicenseExceptionText string   `xml:"LicenseExceptionText,omitempty"  json:"LicenseExceptionText,omitempty"`
	Comment              string   `xml:"Comment,omitempty"  json:"Comment,omitempty"`
	Example              string   `xml:"Example,omitempty"  json:"Example,omitempty"`
	Name                 string   `xml:"Name,omitempty"  json:"Name,omitempty"`
	SeeAlso              string   `xml:"SeeAlso,omitempty"  json:"SeeAlso,omitempty"`
	LicenseExceptionID   string   `xml:"LicenseExceptionId,omitempty"  json:"LicenseExceptionId,omitempty"`
	XMLName              xml.Name `xml:"LicenseException,omitempty"  json:"LicenseException,omitempty"`
}

//ListedLicense ...
type ListedLicense struct {
	IsDeprecatedLicenseID   string   `xml:"IsDeprecatedLicenseId,omitempty"  json:"IsDeprecatedLicenseId,omitempty"`
	StandardLicenseHeader   string   `xml:"StandardLicenseHeader,omitempty"  json:"StandardLicenseHeader,omitempty"`
	StandardLicenseTemplate string   `xml:"StandardLicenseTemplate,omitempty"  json:"StandardLicenseTemplate,omitempty"`
	LicenseText             string   `xml:"LicenseText,omitempty"  json:"LicenseText,omitempty"`
	IsOsiApproved           string   `xml:"IsOsiApproved,omitempty"  json:"IsOsiApproved,omitempty"`
	IsFsfLibre              string   `xml:"IsFsfLibre,omitempty"  json:"IsFsfLibre,omitempty"`
	LicenseID               string   `xml:"LicenseId,omitempty"  json:"LicenseId,omitempty"`
	Name                    string   `xml:"Name,omitempty"  json:"Name,omitempty"`
	SeeAlso                 []string `xml:"SeeAlso,omitempty"  json:"SeeAlso[],omitempty"`
	Comment                 string   `xml:"Comment,omitempty"  json:"Comment,omitempty"`
	XMLName                 xml.Name `xml:"ListedLicense,omitempty"  json:"ListedLicense,omitempty"`
}

//OrLaterOperator ...
type OrLaterOperator struct {
	SimpleLicensingInfo *SimpleLicensingInfo `xml:"SimpleLicensingInfo,omitempty"  json:"SimpleLicensingInfo,omitempty"`
	XMLName             xml.Name             `xml:"OrLaterOperator,omitempty"  json:"OrLaterOperator,omitempty"`
}

//Package ...
type Package struct {
	SpdxItem                *SpdxItem                `xml:"SpdxItem,omitempty"  json:"SpdxItem,omitempty"`
	ExternalRef             *ExternalRef             `xml:"ExternalRef,omitempty"  json:"ExternalRef,omitempty"`
	PackageFileName         string                   `xml:"PackageFileName,omitempty"  json:"PackageFileName,omitempty"`
	FilesAnalyzed           string                   `xml:"FilesAnalyzed,omitempty"  json:"FilesAnalyzed,omitempty"`
	Description             string                   `xml:"Description,omitempty"  json:"Description,omitempty"`
	Homepage                string                   `xml:"Homepage,omitempty"  json:"Homepage,omitempty"`
	Originator              string                   `xml:"Originator,omitempty"  json:"Originator,omitempty"`
	DownloadLocation        string                   `xml:"DownloadLocation,omitempty"  json:"DownloadLocation,omitempty"`
	SourceInfo              string                   `xml:"SourceInfo,omitempty"  json:"SourceInfo,omitempty"`
	Supplier                string                   `xml:"Supplier,omitempty"  json:"Supplier,omitempty"`
	PackageVerificationCode *PackageVerificationCode `xml:"PackageVerificationCode,omitempty"  json:"PackageVerificationCode,omitempty"`
	LicenseDeclaredCode     string                   `xml:"LicenseDeclaredCode,omitempty"  json:"LicenseDeclaredCode,omitempty"`
	VersionInfo             string                   `xml:"VersionInfo,omitempty"  json:"VersionInfo,omitempty"`
	Summary                 string                   `xml:"Summary,omitempty"  json:"Summary,omitempty"`
	Checksum                *Checksum                `xml:"Checksum,omitempty"  json:"Checksum,omitempty"`
	XMLName                 xml.Name                 `xml:"Package,omitempty"  json:"Package,omitempty"`
}

//PackageVerificationCode ...
type PackageVerificationCode struct {
	PackageVerificationCodeValue string   `xml:"PackageVerificationCodeValue,omitempty"  json:"PackageVerificationCodeValue,omitempty"`
	VerificationCodeExcludedFile string   `xml:"VerificationCodeExcludedFile,omitempty"  json:"VerificationCodeExcludedFile,omitempty"`
	XMLName                      xml.Name `xml:"PackageVerificationCode,omitempty"  json:"PackageVerificationCode,omitempty"`
}

//Range ... A data item for This field defines the byte range in the original host file (in X
type Range struct {
	SpdxItem        *SpdxItem `xml:"SpdxItem,omitempty"  json:"SpdxItem,omitempty"`
	File            *File     `xml:"File,omitempty"  json:"File,omitempty"`
	CompoundPointer string    `xml:"CompoundPointer,omitempty"  json:"CompoundPointer,omitempty"`
	XMLName         xml.Name  `xml:"Range,omitempty"  json:"Range,omitempty"`
}

//ReferenceType ...
type ReferenceType struct {
	ExternalReferenceSite string   `xml:"ExternalReferenceSite,omitempty"  json:"ExternalReferenceSite,omitempty"`
	ContextualExample     string   `xml:"ContextualExample,omitempty"  json:"ContextualExample,omitempty"`
	Documentation         string   `xml:"Documentation,omitempty"  json:"Documentation,omitempty"`
	XMLName               xml.Name `xml:"ReferenceType,omitempty"  json:"ReferenceType,omitempty"`
}

//RelatedSpdxElement ... A data item for A related SpdxElement
type RelatedSpdxElement struct {
	RelationshipTypeCode string       `xml:"RelationshipTypeCode,omitempty"  json:"RelationshipTypeCode,omitempty"`
	Comment              string       `xml:"Comment,omitempty"  json:"Comment,omitempty"`
	SpdxElement          *SpdxElement `xml:"SpdxElement,omitempty"  json:"SpdxElement,omitempty"`
	XMLName              xml.Name     `xml:"RelatedSpdxElement,omitempty"  json:"RelatedSpdxElement,omitempty"`
}

//Relationship ...
type Relationship struct {
	RelationshipTypeCode string       `xml:"RelationshipTypeCode,omitempty"  json:"RelationshipTypeCode,omitempty"`
	Comment              string       `xml:"Comment,omitempty"  json:"Comment,omitempty"`
	SpdxElement          *SpdxElement `xml:"SpdxElement,omitempty"  json:"SpdxElement,omitempty"`
	XMLName              xml.Name     `xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
}

//SimpleLicensingInfo ...
type SimpleLicensingInfo struct {
	LicenseID string   `xml:"LicenseId,omitempty"  json:"LicenseId,omitempty"`
	Comment   string   `xml:"Comment,omitempty"  json:"Comment,omitempty"`
	SeeAlso   string   `xml:"SeeAlso,omitempty"  json:"SeeAlso,omitempty"`
	Name      string   `xml:"Name,omitempty"  json:"Name,omitempty"`
	XMLName   xml.Name `xml:"SimpleLicensingInfo,omitempty"  json:"SimpleLicensingInfo,omitempty"`
}

//Snippet ...
type Snippet struct {
	SpdxItem        *SpdxItem `xml:"SpdxItem,omitempty"  json:"SpdxItem,omitempty"`
	File            *File     `xml:"File,omitempty"  json:"File,omitempty"`
	CompoundPointer string    `xml:"CompoundPointer,omitempty"  json:"CompoundPointer,omitempty"`
	XMLName         xml.Name  `xml:"Snippet,omitempty"  json:"Snippet,omitempty"`
}

//SnippetFromFile ... A data item for File containing the SPDX element (e
type SnippetFromFile struct {
	SpdxItem        *SpdxItem `xml:"SpdxItem,omitempty"  json:"SpdxItem,omitempty"`
	File            *File     `xml:"File,omitempty"  json:"File,omitempty"`
	CompoundPointer string    `xml:"CompoundPointer,omitempty"  json:"CompoundPointer,omitempty"`
	XMLName         xml.Name  `xml:"SnippetFromFile,omitempty"  json:"SnippetFromFile,omitempty"`
}

//SpdxDocument ...
type SpdxDocument struct {
	SpdxElement          *SpdxElement          `xml:"SpdxElement,omitempty"  json:"SpdxElement,omitempty"`
	Package              *Package              `xml:"Package,omitempty"  json:"Package,omitempty"`
	File                 *File                 `xml:"File,omitempty"  json:"File,omitempty"`
	ExtractedLicenseInfo *ExtractedLicenseInfo `xml:"ExtractedLicenseInfo,omitempty"  json:"ExtractedLicenseInfo,omitempty"`
	CreationInfo         *CreationInfo         `xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
	SpecVersion          string                `xml:"SpecVersion,omitempty"  json:"SpecVersion,omitempty"`
	ExternalDocumentRef  *ExternalDocumentRef  `xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
	DataLicense          string                `xml:"DataLicense,omitempty"  json:"DataLicense,omitempty"`
	XMLName              xml.Name              `xml:"SpdxDocument,omitempty"  json:"SpdxDocument,omitempty"`
}

//SpdxElement ...
type SpdxElement struct {
	Annotation   *Annotation   `xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
	Name         string        `xml:"Name,omitempty"  json:"Name,omitempty"`
	Comment      string        `xml:"Comment,omitempty"  json:"Comment,omitempty"`
	Relationship *Relationship `xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
	XMLName      xml.Name      `xml:"SpdxElement,omitempty"  json:"SpdxElement,omitempty"`
}

//SpdxItem ...
type SpdxItem struct {
	SpdxElement          *SpdxElement         `xml:"SpdxElement,omitempty"  json:"SpdxElement,omitempty"`
	LicenseConcludedCode string               `xml:"LicenseConcludedCode,omitempty"  json:"LicenseConcludedCode,omitempty"`
	CopyrightText        string               `xml:"CopyrightText,omitempty"  json:"CopyrightText,omitempty"`
	SimpleLicensingInfo  *SimpleLicensingInfo `xml:"SimpleLicensingInfo,omitempty"  json:"SimpleLicensingInfo,omitempty"`
	LicenseComments      string               `xml:"LicenseComments,omitempty"  json:"LicenseComments,omitempty"`
	XMLName              xml.Name             `xml:"SpdxItem,omitempty"  json:"SpdxItem,omitempty"`
}

//WithExceptionOperator ...
type WithExceptionOperator struct {
	SimpleLicensingInfo *SimpleLicensingInfo `xml:"SimpleLicensingInfo,omitempty"  json:"SimpleLicensingInfo,omitempty"`
	LicenseException    *LicenseException    `xml:"LicenseException,omitempty"  json:"LicenseException,omitempty"`
	XMLName             xml.Name             `xml:"WithExceptionOperator,omitempty"  json:"WithExceptionOperator,omitempty"`
}
