package main

import (
	"encoding/xml"
)

//SecurityAttributesGroup ... The group of Information Security Marking attributes in which the use of attributes 'classification' and 'ownerProducer' is required. This group is to be contrasted with group 'SecurityAttributesOptionGroup' in which use of those attributes is optional.
type SecurityAttributesGroup struct {
	Classification           Attrib   `xml:"urn:us:gov:ic:ism classification,attr"  json:"classification,omitempty"`
	OwnerProducer            []Attrib `xml:"urn:us:gov:ic:ism ownerProducer,attr,omitempty"  json:"ownerProducer,omitempty"`
	SCIcontrols              []Attrib `xml:"urn:us:gov:ic:ism SCIcontrols,attr,omitempty"  json:"SCIcontrols,omitempty"`
	SARIdentifier            []Attrib `xml:"urn:us:gov:ic:ism SARIdentifier,attr,omitempty"  json:"SARIdentifier,omitempty"`
	DisseminationControls    []Attrib `xml:"urn:us:gov:ic:ism disseminationControls,attr,omitempty"  json:"disseminationControls,omitempty"`
	FGIsourceOpen            []Attrib `xml:"urn:us:gov:ic:ism FGIsourceOpen,attr,omitempty"  json:"FGIsourceOpen,omitempty"`
	FGIsourceProtected       []Attrib `xml:"urn:us:gov:ic:ism FGIsourceProtected,attr,omitempty"  json:"FGIsourceProtected,omitempty"`
	ReleasableTo             []Attrib `xml:"urn:us:gov:ic:ism releasableTo,attr,omitempty"  json:"releasableTo,omitempty"`
	NonICmarkings            []Attrib `xml:"urn:us:gov:ic:ism nonICmarkings,attr,omitempty"  json:"nonICmarkings,omitempty"`
	ClassifiedBy             Attrib   `xml:"urn:us:gov:ic:ism classifiedBy,attr,omitempty"  json:"classifiedBy,omitempty"`
	DerivativelyClassifiedBy Attrib   `xml:"urn:us:gov:ic:ism derivativelyClassifiedBy,attr,omitempty"  json:"derivativelyClassifiedBy,omitempty"`
	ClassificationReason     Attrib   `xml:"urn:us:gov:ic:ism classificationReason,attr,omitempty"  json:"classificationReason,omitempty"`
	DerivedFrom              Attrib   `xml:"urn:us:gov:ic:ism derivedFrom,attr,omitempty"  json:"derivedFrom,omitempty"`
	DeclassDate              Attrib   `xml:"urn:us:gov:ic:ism declassDate,attr,omitempty"  json:"declassDate,omitempty"`
	DeclassEvent             Attrib   `xml:"urn:us:gov:ic:ism declassEvent,attr,omitempty"  json:"declassEvent,omitempty"`
	DeclassException         []Attrib `xml:"urn:us:gov:ic:ism declassException,attr,omitempty"  json:"declassException,omitempty"`
	TypeOfExemptedSource     []Attrib `xml:"urn:us:gov:ic:ism typeOfExemptedSource,attr,omitempty"  json:"typeOfExemptedSource,omitempty"`
	DateOfExemptedSource     Attrib   `xml:"urn:us:gov:ic:ism dateOfExemptedSource,attr,omitempty"  json:"dateOfExemptedSource,omitempty"`
	DeclassManualReview      Attrib   `xml:"urn:us:gov:ic:ism declassManualReview,attr,omitempty"  json:"declassManualReview,omitempty"`
}

//SecurityAttributesOptionGroup ... The group of Information Security Marking attributes in which the use of attributes 'classification' and 'ownerProducer' is optional. This group is to be contrasted with group 'SecurityAttributesGroup' in which use of these attributes is required.
type SecurityAttributesOptionGroup struct {
	Classification           Attrib   `xml:"urn:us:gov:ic:ism classification,attr,omitempty"  json:"classification,omitempty"`
	OwnerProducer            []Attrib `xml:"urn:us:gov:ic:ism ownerProducer,attr,omitempty"  json:"ownerProducer,omitempty"`
	SCIcontrols              []Attrib `xml:"urn:us:gov:ic:ism SCIcontrols,attr,omitempty"  json:"SCIcontrols,omitempty"`
	SARIdentifier            []Attrib `xml:"urn:us:gov:ic:ism SARIdentifier,attr,omitempty"  json:"SARIdentifier,omitempty"`
	DisseminationControls    []Attrib `xml:"urn:us:gov:ic:ism disseminationControls,attr,omitempty"  json:"disseminationControls,omitempty"`
	FGIsourceOpen            []Attrib `xml:"urn:us:gov:ic:ism FGIsourceOpen,attr,omitempty"  json:"FGIsourceOpen,omitempty"`
	FGIsourceProtected       []Attrib `xml:"urn:us:gov:ic:ism FGIsourceProtected,attr,omitempty"  json:"FGIsourceProtected,omitempty"`
	ReleasableTo             []Attrib `xml:"urn:us:gov:ic:ism releasableTo,attr,omitempty"  json:"releasableTo,omitempty"`
	NonICmarkings            []Attrib `xml:"urn:us:gov:ic:ism nonICmarkings,attr,omitempty"  json:"nonICmarkings,omitempty"`
	ClassifiedBy             Attrib   `xml:"urn:us:gov:ic:ism classifiedBy,attr,omitempty"  json:"classifiedBy,omitempty"`
	DerivativelyClassifiedBy Attrib   `xml:"urn:us:gov:ic:ism derivativelyClassifiedBy,attr,omitempty"  json:"derivativelyClassifiedBy,omitempty"`
	ClassificationReason     Attrib   `xml:"urn:us:gov:ic:ism classificationReason,attr,omitempty"  json:"classificationReason,omitempty"`
	DerivedFrom              Attrib   `xml:"urn:us:gov:ic:ism derivedFrom,attr,omitempty"  json:"derivedFrom,omitempty"`
	DeclassDate              Attrib   `xml:"urn:us:gov:ic:ism declassDate,attr,omitempty"  json:"declassDate,omitempty"`
	DeclassEvent             Attrib   `xml:"urn:us:gov:ic:ism declassEvent,attr,omitempty"  json:"declassEvent,omitempty"`
	DeclassException         []Attrib `xml:"urn:us:gov:ic:ism declassException,attr,omitempty"  json:"declassException,omitempty"`
	TypeOfExemptedSource     []Attrib `xml:"urn:us:gov:ic:ism typeOfExemptedSource,attr,omitempty"  json:"typeOfExemptedSource,omitempty"`
	DateOfExemptedSource     Attrib   `xml:"urn:us:gov:ic:ism dateOfExemptedSource,attr,omitempty"  json:"dateOfExemptedSource,omitempty"`
	DeclassManualReview      Attrib   `xml:"urn:us:gov:ic:ism declassManualReview,attr,omitempty"  json:"declassManualReview,omitempty"`
}

//Attrib ...
type Attrib xml.Attr

//UnmarshalXMLAttr ...
func (a *Attrib) UnmarshalXMLAttr(attr xml.Attr) error {
	//log.Println("UnmarshalXMLAttr")
	//log.Println(attr)
	a.Name.Local = "ism:" + a.Name.Local
	*a = Attrib(attr)
	return nil
}

//MarshalXMLAttr ...
func (a Attrib) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if len(a.Name.Local) > 0 {
		a.Name.Local = "ism:" + a.Name.Local
	}
	a.Name.Space = ""
	//log.Println(a.Name)
	return xml.Attr(a), nil
}
