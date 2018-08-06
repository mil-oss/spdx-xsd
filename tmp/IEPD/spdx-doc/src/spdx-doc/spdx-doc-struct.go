package main

import "encoding/xml"

//NewSpdxDocument ...
func NewSpdxDocument() *SpdxDocumentType{
    return &SpdxDocumentType{
        // Required for the proper namespacing
        AttrXmlnsXsi:"http://www.w3.org/2001/XMLSchema-instance",
        AttrXmlns:"spdx:xsd::1.0",
    }
}
//SpdxDocumentType ... 
type SpdxDocumentType struct {
        AttrXmlnsXsi                             string                                   `xml:"xmlns:xsi,attr" json:"AttrXmlnsXsi,omitempty"`
        AttrXmlns                                string                                   `xml:"xmlns,attr" json:"AttrXmlns,omitempty"`
        Annotation                               *AnnotationType                          `xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
        Name                                     string                                   `xml:"Name,omitempty"  json:"Name,omitempty"`
        Comment                                  string                                   `xml:"Comment,omitempty"  json:"Comment,omitempty"`
        Relationship                             *RelationshipType                        `xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
        CreationInfo                             *CreationInfoType                        `xml:"CreationInfo,omitempty"  json:"CreationInfo,omitempty"`
        SpecVersion                              string                                   `xml:"SpecVersion,omitempty"  json:"SpecVersion,omitempty"`
        ExternalDocumentRef                      *ExternalDocumentRefType                 `xml:"ExternalDocumentRef,omitempty"  json:"ExternalDocumentRef,omitempty"`
        DataLicense                              string                                   `xml:"DataLicense,omitempty"  json:"DataLicense,omitempty"`
        XMLName                                  xml.Name                                 `xml:"SpdxDocument,omitempty"  json:"SpdxDocument,omitempty"`
}
//AnnotationType ... 
type AnnotationType struct {
        Date                                     string                                   `xml:"Date,omitempty"  json:"Date,omitempty"`
        AnnotationTypeCode                       string                                   `xml:"AnnotationTypeCode,omitempty"  json:"AnnotationTypeCode,omitempty"`
        Comment                                  string                                   `xml:"Comment,omitempty"  json:"Comment,omitempty"`
        Annotator                                string                                   `xml:"Annotator,omitempty"  json:"Annotator,omitempty"`
}
//ChecksumType ... 
type ChecksumType struct {
        ChecksumValue                            string                                   `xml:"ChecksumValue,omitempty"  json:"ChecksumValue,omitempty"`
        AlgorithmCode                            string                                   `xml:"AlgorithmCode,omitempty"  json:"AlgorithmCode,omitempty"`
}
//CreationInfoType ... 
type CreationInfoType struct {
        LicenseListVersion                       string                                   `xml:"LicenseListVersion,omitempty"  json:"LicenseListVersion,omitempty"`
        Created                                  string                                   `xml:"Created,omitempty"  json:"Created,omitempty"`
        Comment                                  string                                   `xml:"Comment,omitempty"  json:"Comment,omitempty"`
        Creator                                  string                                   `xml:"Creator,omitempty"  json:"Creator,omitempty"`
}
//ExternalDocumentRefType ... 
type ExternalDocumentRefType struct {
        Checksum                                 *ChecksumType                            `xml:"Checksum,omitempty"  json:"Checksum,omitempty"`
        ExternalDocumentID                       string                                   `xml:"ExternalDocumentID,omitempty"  json:"ExternalDocumentID,omitempty"`
        SpdxDocument                             *SpdxDocumentType                        `xml:"SpdxDocument,omitempty"  json:"SpdxDocument,omitempty"`
}
//RelationshipType ... 
type RelationshipType struct {
        RelationshipTypeCode                     string                                   `xml:"RelationshipTypeCode,omitempty"  json:"RelationshipTypeCode,omitempty"`
        Comment                                  string                                   `xml:"Comment,omitempty"  json:"Comment,omitempty"`
        RelatedSpdxElement                       *RelationshipType                        `xml:"RelatedSpdxElement,omitempty"  json:"RelatedSpdxElement,omitempty"`
}
//SpdxElementType ... 
type SpdxElementType struct {
        Annotation                               *AnnotationType                          `xml:"Annotation,omitempty"  json:"Annotation,omitempty"`
        Name                                     string                                   `xml:"Name,omitempty"  json:"Name,omitempty"`
        Comment                                  string                                   `xml:"Comment,omitempty"  json:"Comment,omitempty"`
        Relationship                             *RelationshipType                        `xml:"Relationship,omitempty"  json:"Relationship,omitempty"`
}
