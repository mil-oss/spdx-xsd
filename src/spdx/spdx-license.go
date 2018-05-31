package main

import "encoding/xml"

//NewLicense ...
func NewLicense()*License{
    return &License{
        // Required for the proper namespacingLicense
        AttrXmlnsXsi:"http://www.w3.org/2001/XMLSchema-instance",
        AttrXmlns:"spdx:xsd::1.0",
        LicenseID                                []string                                 `xml:"LicenseID,omitempty"  json:"LicenseID,omitempty"`
        Comment                                  []string                                 `xml:"Comment,omitempty"  json:"Comment,omitempty"`
        SeeAlso                                  string                                   `xml:"SeeAlso,omitempty"  json:"SeeAlso,omitempty"`
        Name                                     []string                                 `xml:"Name,omitempty"  json:"Name,omitempty"`
        IsDeprecatedLicenseID                    boolean                                  `xml:"IsDeprecatedLicenseId,omitempty"  json:"IsDeprecatedLicenseId,omitempty"`
        IsOsiApproved                            boolean                                  `xml:"IsOsiApproved,omitempty"  json:"IsOsiApproved,omitempty"`
        IsFsfLibre                               boolean                                  `xml:"IsFsfLibre,omitempty"  json:"IsFsfLibre,omitempty"`
        StandardLicenseHeader                    string                                   `xml:"StandardLicenseHeader,omitempty"  json:"StandardLicenseHeader,omitempty"`
        LicenseText                              string                                   `xml:"LicenseText,omitempty"  json:"LicenseText,omitempty"`
        StandardLicenseTemplate                  string                                   `xml:"StandardLicenseTemplate,omitempty"  json:"StandardLicenseTemplate,omitempty"`
    }
}
//License ... 
type License struct {
        AttrXmlnsXsi                             string                                   `xml:"xmlns:xsi,attr" json:"AttrXmlnsXsi,omitempty"`
        AttrXmlns                                string                                   `xml:"xmlns,attr" json:"AttrXmlns,omitempty"`
        LicenseID                                []string                                 `xml:"LicenseID,omitempty"  json:"LicenseID,omitempty"`
        Comment                                  []string                                 `xml:"Comment,omitempty"  json:"Comment,omitempty"`
        SeeAlso                                  string                                   `xml:"SeeAlso,omitempty"  json:"SeeAlso,omitempty"`
        Name                                     []string                                 `xml:"Name,omitempty"  json:"Name,omitempty"`
        IsDeprecatedLicenseID                    boolean                                  `xml:"IsDeprecatedLicenseId,omitempty"  json:"IsDeprecatedLicenseId,omitempty"`
        IsOsiApproved                            boolean                                  `xml:"IsOsiApproved,omitempty"  json:"IsOsiApproved,omitempty"`
        IsFsfLibre                               boolean                                  `xml:"IsFsfLibre,omitempty"  json:"IsFsfLibre,omitempty"`
        StandardLicenseHeader                    string                                   `xml:"StandardLicenseHeader,omitempty"  json:"StandardLicenseHeader,omitempty"`
        LicenseText                              string                                   `xml:"LicenseText,omitempty"  json:"LicenseText,omitempty"`
        StandardLicenseTemplate                  string                                   `xml:"StandardLicenseTemplate,omitempty"  json:"StandardLicenseTemplate,omitempty"`
        XMLName                                  xml.Name                                 `xml:"License,omitempty"  json:"License,omitempty"`
}
