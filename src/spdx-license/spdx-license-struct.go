package main

import "encoding/xml"

//License ...
type License struct {
	LicenseID               []string `xml:"LicenseID,omitempty"  json:"LicenseID,omitempty"`
	CommentText             []string `xml:"CommentText,omitempty"  json:"CommentText,omitempty"`
	SeeAlsoURI              string   `xml:"SeeAlsoURI,omitempty"  json:"SeeAlsoURI,omitempty"`
	Name                    []string `xml:"Name,omitempty"  json:"Name,omitempty"`
	IsDeprecatedLicenseID   bool     `xml:"IsDeprecatedLicenseID,omitempty"  json:"IsDeprecatedLicenseID,omitempty"`
	IsOsiApprovedIndicator  bool     `xml:"IsOsiApprovedIndicator,omitempty"  json:"IsOsiApprovedIndicator,omitempty"`
	IsFsfLibreIndicator     bool     `xml:"IsFsfLibreIndicator,omitempty"  json:"IsFsfLibreIndicator,omitempty"`
	StandardLicenseHeader   string   `xml:"StandardLicenseHeader,omitempty"  json:"StandardLicenseHeader,omitempty"`
	LicenseText             string   `xml:"LicenseText,omitempty"  json:"LicenseText,omitempty"`
	StandardLicenseTemplate string   `xml:"StandardLicenseTemplate,omitempty"  json:"StandardLicenseTemplate,omitempty"`
	XMLName                 xml.Name `xml:"License,omitempty"  json:"License,omitempty"`
}
