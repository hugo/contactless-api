package contacts

import "encoding/xml"

// GoogleCategory is the atom category of the entry
type GoogleCategory struct {
	XMLName xml.Name `xml:"atom:category"`
	Scheme  string   `xml:"scheme,attr"`
	Term    string   `xml:"term,attr"`
}

// GoogleContactName is the name of a contact comprised of several parts
type GoogleContactName struct {
	XMLName    xml.Name `xml:"gd:name,omitempty"`
	GivenName  string   `xml:"gd:givenName"`
	FamilyName string   `xml:"gd:familyName"`
	FullName   string   `xml:"gd:fullName"`
}

// GoogleContent is an arbitraty bit of content
type GoogleContent struct {
	Data string `xml:",innerxml"`
	Type string `xml:"type,attr"`
}

// GoogleContents is a list of content items
type GoogleContents []GoogleContent

// GoogleEmail is an email address
type GoogleEmail struct {
	Address     string `xml:"address,attr"`
	Primary     bool   `xml:"primary,attr,omitempty"`
	Rel         string `xml:"rel,attr"`
	DisplayName string `xml:"gd:displayName,attr"`
}

// GoogleEmails is a list of email addresses
type GoogleEmails []GoogleEmail

// GooglePhoneNumber is a telephone number
type GooglePhoneNumber struct {
	Number  string `xml:",innerxml"`
	Primary bool   `xml:"primary,attr,omitempty"`
	Rel     string `xml:"rel,attr"`
}

// GooglePhoneNumbers is a list of telephone numbers
type GooglePhoneNumbers []GooglePhoneNumber

// GoogleAddress is an address in Google's structured format
type GoogleAddress struct {
	Primary  bool   `xml:"primary,attr"`
	Rel      string `xml:"rel,attr"`
	Street   string `xml:"gd:street"`
	City     string `xml:"gd:city"`
	Region   string `xml:"gd:region"`
	Postcode string `xml:"gd:postcode"`
	Country  string `xml:"gd:country"`
}

// GoogleAddresses is a list of addresses
type GoogleAddresses []GoogleAddress

// GoogleContact is a Contact as parsed from XML
type GoogleContact struct {
	XMLName            xml.Name          `xml:"atom:entry"`
	AtomNamespace      string            `xml:"xmlns:atom,attr"`
	GDNamespace        string            `xml:"xmlns:gd,attr"`
	Category           GoogleCategory    `xml:"atom:category"`
	ID                 string            `xml:"id,omitempty"`
	Etag               string            `xml:"etag,attr,omitempty"`
	Name               GoogleContactName `xml:"gd:name"`
	GoogleContents     `xml:"atom:content"`
	GoogleEmails       `xml:"gd:email"`
	GooglePhoneNumbers `xml:"gd:phoneNumber"`
	GoogleAddresses    `xml:"gd:structuredPostalAddress"`
}
