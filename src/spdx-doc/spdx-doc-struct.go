package main

import "encoding/xml"

//Annotation ...
type Annotation struct {
	Date               []string `xml:"Date,omitempty"  json:"Date,omitempty"`
	AnnotationTypeCode string   `xml:"AnnotationTypeCode,omitempty"  json:"AnnotationTypeCode,omitempty"`
	CommentText        []string `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	AnnotatorText      []string `xml:"AnnotatorText,omitempty"  json:"AnnotatorText,omitempty"`
	XMLName            xml.Name `xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
}

//Checksum ...
type Checksum struct {
	ChecksumValue []string `xml:"ChecksumValue,omitempty"  json:"ChecksumValue,omitempty"`
	AlgorithmCode string   `xml:"AlgorithmCode,omitempty"  json:"AlgorithmCode,omitempty"`
	XMLName       xml.Name `xml:"Checksum,omitempty"  json:"Checksum,omitempty"`
}

//CreationInfo ...
type CreationInfo struct {
	LicenseListVersionText string   `xml:"LicenseListVersionText,omitempty"  json:"LicenseListVersionText,omitempty"`
	CreatedDateTime        []string `xml:"CreatedDateTime,omitempty"  json:"CreatedDateTime,omitempty"`
	CommentText            string   `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	CreatorText            string   `xml:"CreatorText,omitempty"  json:"CreatorText,omitempty"`
	XMLName                xml.Name `xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
}

//DescribesFile ... A data item for The describesFile property relates an SpdxDocument to the file which it describes
type DescribesFile struct {
	Annotation                string   `xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
	Name                      []string `xml:"Name,omitempty"  json:"Name,omitempty"`
	CommentText               []string `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	Relationship              string   `xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
	DescribesPackage          string   `xml:"DescribesPackage,omitempty"  json:"DescribesPackage,omitempty"`
	DescribesFile             string   `xml:"DescribesFile,omitempty"  json:"DescribesFile,omitempty"`
	HasExtractedLicensingInfo string   `xml:"HasExtractedLicensingInfo,omitempty"  json:"HasExtractedLicensingInfo,omitempty"`
	CreationInfo              []string `xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
	SpecVersionText           []string `xml:"SpecVersionText,omitempty"  json:"SpecVersionText,omitempty"`
	ExternalDocumentRef       string   `xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
	DataLicense               string   `xml:"DataLicense,omitempty"  json:"DataLicense,omitempty"`
	XMLName                   xml.Name `xml:"DescribesFile,omitempty"  json:"DescribesFile,omitempty"`
}

//DescribesPackage ... A data item for The describesPackage property relates an SpdxDocument to the package which it describes
type DescribesPackage struct {
	Annotation                string   `xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
	Name                      []string `xml:"Name,omitempty"  json:"Name,omitempty"`
	CommentText               []string `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	Relationship              string   `xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
	DescribesPackage          string   `xml:"DescribesPackage,omitempty"  json:"DescribesPackage,omitempty"`
	DescribesFile             string   `xml:"DescribesFile,omitempty"  json:"DescribesFile,omitempty"`
	HasExtractedLicensingInfo string   `xml:"HasExtractedLicensingInfo,omitempty"  json:"HasExtractedLicensingInfo,omitempty"`
	CreationInfo              []string `xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
	SpecVersionText           []string `xml:"SpecVersionText,omitempty"  json:"SpecVersionText,omitempty"`
	ExternalDocumentRef       string   `xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
	DataLicense               string   `xml:"DataLicense,omitempty"  json:"DataLicense,omitempty"`
	XMLName                   xml.Name `xml:"DescribesPackage,omitempty"  json:"DescribesPackage,omitempty"`
}

//ExternalDocumentRef ...
type ExternalDocumentRef struct {
	Checksum           []string `xml:"Checksum,omitempty"  json:"Checksum,omitempty"`
	ExternalDocumentID []string `xml:"ExternalDocumentID,omitempty"  json:"ExternalDocumentID,omitempty"`
	SpdxDocument       []string `xml:"SpdxDocument,omitempty"  json:"SpdxDocument,omitempty"`
	XMLName            xml.Name `xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
}

//HasExtractedLicensingInfo ... A data item for Indicates that a particular ExtractedLicensingInfo was defined in the subject SpdxDocument
type HasExtractedLicensingInfo struct {
	Annotation                string   `xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
	Name                      []string `xml:"Name,omitempty"  json:"Name,omitempty"`
	CommentText               []string `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	Relationship              string   `xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
	DescribesPackage          string   `xml:"DescribesPackage,omitempty"  json:"DescribesPackage,omitempty"`
	DescribesFile             string   `xml:"DescribesFile,omitempty"  json:"DescribesFile,omitempty"`
	HasExtractedLicensingInfo string   `xml:"HasExtractedLicensingInfo,omitempty"  json:"HasExtractedLicensingInfo,omitempty"`
	CreationInfo              []string `xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
	SpecVersionText           []string `xml:"SpecVersionText,omitempty"  json:"SpecVersionText,omitempty"`
	ExternalDocumentRef       string   `xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
	DataLicense               string   `xml:"DataLicense,omitempty"  json:"DataLicense,omitempty"`
	XMLName                   xml.Name `xml:"HasExtractedLicensingInfo,omitempty"  json:"HasExtractedLicensingInfo,omitempty"`
}

//RelatedSpdxElement ... A data item for A related SpdxElement
type RelatedSpdxElement struct {
	RelationshipTypeCode string   `xml:"RelationshipTypeCode,omitempty"  json:"RelationshipTypeCode,omitempty"`
	CommentText          []string `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	RelatedSpdxElement   []string `xml:"RelatedSpdxElement,omitempty"  json:"RelatedSpdxElement,omitempty"`
	XMLName              xml.Name `xml:"RelatedSpdxElement,omitempty"  json:"RelatedSpdxElement,omitempty"`
}

//Relationship ...
type Relationship struct {
	RelationshipTypeCode string   `xml:"RelationshipTypeCode,omitempty"  json:"RelationshipTypeCode,omitempty"`
	CommentText          []string `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	RelatedSpdxElement   []string `xml:"RelatedSpdxElement,omitempty"  json:"RelatedSpdxElement,omitempty"`
	XMLName              xml.Name `xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
}

//SpdxDocument ...
type SpdxDocument struct {
	Annotation                string   `xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
	Name                      []string `xml:"Name,omitempty"  json:"Name,omitempty"`
	CommentText               []string `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	Relationship              string   `xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
	DescribesPackage          string   `xml:"DescribesPackage,omitempty"  json:"DescribesPackage,omitempty"`
	DescribesFile             string   `xml:"DescribesFile,omitempty"  json:"DescribesFile,omitempty"`
	HasExtractedLicensingInfo string   `xml:"HasExtractedLicensingInfo,omitempty"  json:"HasExtractedLicensingInfo,omitempty"`
	CreationInfo              []string `xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
	SpecVersionText           []string `xml:"SpecVersionText,omitempty"  json:"SpecVersionText,omitempty"`
	ExternalDocumentRef       string   `xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
	DataLicense               string   `xml:"DataLicense,omitempty"  json:"DataLicense,omitempty"`
	XMLName                   xml.Name `xml:"SpdxDocument,omitempty"  json:"SpdxDocument,omitempty"`
}
