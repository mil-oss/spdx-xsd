package main

import "encoding/xml"

//NewLicense ...
func NewLicense()*License{
    return &License{
        // Required for the proper namespacingLicense
        AttrXmlnsXsi:"http://www.w3.org/2001/XMLSchema-instance",
        AttrXmlns:"spdx:xsd::1.0",
    }
}
//License ... 
type License struct {
        AttrXmlnsXsi                             string                                   `xml:"xmlns:xsi,attr" json:"AttrXmlnsXsi,omitempty"`
        AttrXmlns                                string                                   `xml:"xmlns,attr" json:"AttrXmlns,omitempty"`
        IsDeprecatedLicenseID                    string                                   `xml:"IsDeprecatedLicenseId,omitempty"  json:"IsDeprecatedLicenseId,omitempty"`
        StandardLicenseHeader                    string                                   `xml:"StandardLicenseHeader,omitempty"  json:"StandardLicenseHeader,omitempty"`
        StandardLicenseTemplate                  string                                   `xml:"StandardLicenseTemplate,omitempty"  json:"StandardLicenseTemplate,omitempty"`
        LicenseText                              string                                   `xml:"LicenseText,omitempty"  json:"LicenseText,omitempty"`
        IsOsiApproved                            string                                   `xml:"IsOsiApproved,omitempty"  json:"IsOsiApproved,omitempty"`
        IsFsfLibre                               string                                   `xml:"IsFsfLibre,omitempty"  json:"IsFsfLibre,omitempty"`
        LicenseID                                string                                   `xml:"LicenseId,omitempty"  json:"LicenseId,omitempty"`
        Name                                     string                                   `xml:"Name,omitempty"  json:"Name,omitempty"`
        SeeAlso                                  []string                                 `xml:"SeeAlso,omitempty"  json:"SeeAlso[],omitempty"`
        Comment                                  string                                   `xml:"Comment,omitempty"  json:"Comment,omitempty"`
        XMLName                                  xml.Name                                 `xml:"License,omitempty"  json:"License,omitempty"`
}
