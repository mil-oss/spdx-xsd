package main

import "encoding/xml"

//NewSpdxDocument ...
func NewSpdxDocument() *SpdxDocument{
    return &SpdxDocument{
        // Required for the proper namespacing
        AttrXmlnsXsi:"http://www.w3.org/2001/XMLSchema-instance",
        AttrXmlns:"spdx:xsd::1.0",
    }
}
//SpdxDocument ... 
type SpdxDocument struct {
        AttrXmlnsXsi	string		`xml:"xmlns:xsi,attr" json:"AttrXmlnsXsi,omitempty"`
	AttrXmlns	string		`xml:"xmlns,attr" json:"AttrXmlns,omitempty"`
	Annotation	*Annotation		`xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
	Name	string		`xml:"Name,omitempty"  json:"Name,omitempty"`
	CommentText	string		`xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	Relationship	*Relationship		`xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
	DescribesPackage	*DescribesPackage		`xml:"DescribesPackage,omitempty"  json:"DescribesPackage,omitempty"`
	DescribesFile	*DescribesFile		`xml:"DescribesFile,omitempty"  json:"DescribesFile,omitempty"`
	HasExtractedLicensingInfo	*HasExtractedLicensingInfo		`xml:"HasExtractedLicensingInfo,omitempty"  json:"HasExtractedLicensingInfo,omitempty"`
	CreationInfo	*CreationInfo		`xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
	SpecVersionText	string		`xml:"SpecVersionText,omitempty"  json:"SpecVersionText,omitempty"`
	ExternalDocumentRef	*ExternalDocumentRef		`xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
	DataLicense	string		`xml:"DataLicense,omitempty"  json:"DataLicense,omitempty"`
	XMLName	xml.Name		`xml:"SpdxDocument,omitempty"  json:"SpdxDocument,omitempty"`
}
//Annotation ... 
type Annotation struct {
	Date	string		`xml:"Date,omitempty"  json:"Date,omitempty"`
	AnnotationTypeCode	string		`xml:"AnnotationTypeCode,omitempty"  json:"AnnotationTypeCode,omitempty"`
	CommentText	string		`xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	AnnotatorText	string		`xml:"AnnotatorText,omitempty"  json:"AnnotatorText,omitempty"`
	XMLName	xml.Name		`xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
}
//Checksum ... 
type Checksum struct {
	ChecksumValue	string		`xml:"ChecksumValue,omitempty"  json:"ChecksumValue,omitempty"`
	AlgorithmCode	string		`xml:"AlgorithmCode,omitempty"  json:"AlgorithmCode,omitempty"`
	XMLName	xml.Name		`xml:"Checksum,omitempty"  json:"Checksum,omitempty"`
}
//CreationInfo ... 
type CreationInfo struct {
	LicenseListVersionText	string		`xml:"LicenseListVersionText,omitempty"  json:"LicenseListVersionText,omitempty"`
	CreatedDateTime	string		`xml:"CreatedDateTime,omitempty"  json:"CreatedDateTime,omitempty"`
	CommentText	string		`xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	CreatorText	string		`xml:"CreatorText,omitempty"  json:"CreatorText,omitempty"`
	XMLName	xml.Name		`xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
}
//ExternalDocumentRef ... 
type ExternalDocumentRef struct {
	Checksum	*Checksum		`xml:"Checksum,omitempty"  json:"Checksum,omitempty"`
	ExternalDocumentID	string		`xml:"ExternalDocumentID,omitempty"  json:"ExternalDocumentID,omitempty"`
	SpdxDocument	*SpdxDocument		`xml:"SpdxDocument,omitempty"  json:"SpdxDocument,omitempty"`
	XMLName	xml.Name		`xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
}
//RelatedSpdxElement ... A data item for A related SpdxElement
type RelatedSpdxElement struct {
	Annotation	*Annotation		`xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
	Name	string		`xml:"Name,omitempty"  json:"Name,omitempty"`
	CommentText	string		`xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	Relationship	*Relationship		`xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
	XMLName	xml.Name		`xml:"RelatedSpdxElement,omitempty"  json:"RelatedSpdxElement,omitempty"`
}
//Relationship ... 
type Relationship struct {
	RelationshipTypeCode	string		`xml:"RelationshipTypeCode,omitempty"  json:"RelationshipTypeCode,omitempty"`
	CommentText	string		`xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	RelatedSpdxElement	*RelatedSpdxElement		`xml:"RelatedSpdxElement,omitempty"  json:"RelatedSpdxElement,omitempty"`
	XMLName	xml.Name		`xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
}
// RelatedSpdxElement ... 
type RelatedSpdxElement struct {
	Annotation	*Annotation		`xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
	Name	string		`xml:"Name,omitempty"  json:"Name,omitempty"`
	CommentText	string		`xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	XMLName	xml.Name		`xml:"RelatedSpdxElement,omitempty"  json:"RelatedSpdxElement,omitempty"`
}
