package main

import "encoding/xml"

//NewLicense ...
func NewLicense() *License {
	return &License{
		// Required for the proper namespacing
		AttrXmlnsXsi: "http://www.w3.org/2001/XMLSchema-instance",
		AttrXmlns:    "spdx:xsd::1.0",
	}
}

//License ...
type License struct {
	AttrXmlnsXsi            string   `xml:"xmlns:xsi,attr" json:"AttrXmlnsXsi,omitempty"`
	AttrXmlns               string   `xml:"xmlns,attr" json:"AttrXmlns,omitempty"`
	LicenseID               string   `xml:"LicenseID,omitempty"  json:"LicenseID,omitempty"`
	CommentText             string   `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	SeeAlsoURI              string   `xml:"SeeAlsoURI,omitempty"  json:"SeeAlsoURI,omitempty"`
	Name                    string   `xml:"Name,omitempty"  json:"Name,omitempty"`
	IsDeprecatedLicenseID   bool     `xml:"IsDeprecatedLicenseID,omitempty"  json:"IsDeprecatedLicenseID,omitempty"`
	IsOsiApprovedIndicator  bool     `xml:"IsOsiApprovedIndicator,omitempty"  json:"IsOsiApprovedIndicator,omitempty"`
	IsFsfLibreIndicator     bool     `xml:"IsFsfLibreIndicator,omitempty"  json:"IsFsfLibreIndicator,omitempty"`
	StandardLicenseHeader   string   `xml:"StandardLicenseHeader,omitempty"  json:"StandardLicenseHeader,omitempty"`
	LicenseText             string   `xml:"LicenseText,omitempty"  json:"LicenseText,omitempty"`
	StandardLicenseTemplate string   `xml:"StandardLicenseTemplate,omitempty"  json:"StandardLicenseTemplate,omitempty"`
	XMLName                 xml.Name `xml:"License,omitempty"  json:"License,omitempty"`
}
