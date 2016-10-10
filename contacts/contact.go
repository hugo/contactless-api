package contacts

import "encoding/xml"

// Category is the atom category of the entry
type Category struct {
	XMLName xml.Name `xml:"category"`
	Scheme  string   `xml:"scheme,attr"`
	Term    string   `xml:"term,attr"`
}

// ContactName is the name of a contact comprised of several parts
type ContactName struct {
	XMLName    xml.Name `xml:"name"`
	GivenName  string   `xml:"givenName"`
	FamilyName string   `xml:"familyName"`
	FullName   string   `xml:"fullName"`
}

// Content is an arbitraty bit of content
type Content struct {
	Data string `xml:",innerxml"`
	Type string `xml:"type,attr"`
}

// Contents is a list of content items
type Contents []Content

// Email is an email address
type Email struct {
	Address     string `xml:"address,attr"`
	Primary     bool   `xml:"primary,attr,omitempty"`
	Rel         string `xml:"rel,attr"`
	DisplayName string `xml:"displayName,attr"`
}

// Emails is a list of email addresses
type Emails []Email

// PhoneNumber is a telephone number
type PhoneNumber struct {
	Number  string `xml:",innerxml"`
	Primary bool   `xml:"primary,attr,omitempty"`
	Rel     string `xml:"rel,attr"`
}

// PhoneNumbers is a list of telephone numbers
type PhoneNumbers []PhoneNumber

// Address is an address in Google's structured format
type Address struct {
	Primary  bool   `xml:"primary,attr"`
	Rel      string `xml:"rel,attr"`
	Street   string `xml:"street"`
	City     string `xml:"city"`
	Region   string `xml:"region"`
	Postcode string `xml:"postcode"`
	Country  string `xml:"country"`
}

// Addresses is a list of addresses
type Addresses []Address

// Contact is a Contact as parsed from XML
type Contact struct {
	XMLName      xml.Name    `xml:"entry"`
	Category     Category    `xml:"category"`
	ID           string      `xml:"id"`
	Etag         string      `xml:"etag,attr"`
	Name         ContactName `xml:"name"`
	Contents     `xml:"content"`
	Emails       `xml:"email"`
	PhoneNumbers `xml:"phoneNumber"`
	Addresses    `xml:"structuredPostalAddress"`
}
