package contacts

// JSONContact is a contact in the format we desire for JSON
type JSONContact struct {
	ID   string `json:"id"`
	Etag string `json:"etag"`
	Name struct {
		GivenName  string `json:"givenName"`
		FamilyName string `json:"familyName"`
		FullName   string `json:"fullName"`
	} `json:"name"`
	Emails []struct {
		Email       string `json:"email"`
		Primary     bool   `json:"primary"`
		Type        string `json:"type"`
		DisplayName string `json:"displayName"`
	} `json:"emailAddresses"`
	PhoneNumbers []struct {
		Number  string `json:"number"`
		Primary bool   `json:"primary"`
		Type    string `json:"type"`
	} `json:"emailAddresses"`
	Addresses []struct {
		Primary  bool   `json:"primary"`
		Type     string `json:"type"`
		Street   string `json:"street"`
		City     string `json:"city"`
		Region   string `json:"region"`
		Postcode string `json:"postcode"`
		Country  string `json:"country"`
	} `json:"addresses"`
}

// ToJSON takes a contact and sets it up for rendering to JSON
func ToJSON(contact Contact) JSONContact {
	var jsonContact JSONContact

	jsonContact.ID = contact.ID
	jsonContact.Etag = contact.Etag

	jsonContact.Name.GivenName = contact.Name.GivenName
	jsonContact.Name.FamilyName = contact.Name.FamilyName
	jsonContact.Name.FullName = contact.Name.FullName

	for _, email := range contact.Emails {
		jsonContact.Emails = append(
			jsonContact.Emails,
			struct {
				Email       string `json:"email"`
				Primary     bool   `json:"primary"`
				Type        string `json:"type"`
				DisplayName string `json:"displayName"`
			}{
				email.Data,
				email.Primary,
				email.Rel,
				email.DisplayName,
			},
		)
	}

	for _, phoneNumber := range contact.PhoneNumbers {
		jsonContact.PhoneNumbers = append(
			jsonContact.PhoneNumbers,
			struct {
				Number  string `json:"number"`
				Primary bool   `json:"primary"`
				Type    string `json:"type"`
			}{
				phoneNumber.Number,
				phoneNumber.Primary,
				phoneNumber.Rel,
			},
		)
	}

	for _, address := range contact.Addresses {
		jsonContact.Addresses = append(
			jsonContact.Addresses,
			struct {
				Primary  bool   `json:"primary"`
				Type     string `json:"type"`
				Street   string `json:"street"`
				City     string `json:"city"`
				Region   string `json:"region"`
				Postcode string `json:"postcode"`
				Country  string `json:"country"`
			}{
				address.Primary,
				address.Rel,
				address.Street,
				address.City,
				address.Region,
				address.Postcode,
				address.Country,
			},
		)
	}

	return jsonContact
}

// FromJSON prepares a JSON for conversion to XML
func FromJSON(jsonContact JSONContact) Contact {
	var contact Contact

	contact.ID = jsonContact.ID
	contact.Etag = jsonContact.Etag

	contact.Name.GivenName = jsonContact.Name.GivenName
	contact.Name.FamilyName = jsonContact.Name.FamilyName
	contact.Name.FullName = jsonContact.Name.FullName

	for _, email := range jsonContact.Emails {
		contact.Emails = append(contact.Emails,
			Email{
				email.Email,
				email.Primary,
				email.Type,
				email.DisplayName,
			},
		)
	}

	for _, phoneNumber := range jsonContact.PhoneNumbers {
		contact.PhoneNumbers = append(contact.PhoneNumbers,
			PhoneNumber{
				phoneNumber.Number,
				phoneNumber.Primary,
				phoneNumber.Type,
			},
		)
	}

	for _, address := range contact.Addresses {
		contact.Addresses = append(contact.Addresses,
			Address{
				address.Primary,
				address.Rel,
				address.Street,
				address.City,
				address.Region,
				address.Postcode,
				address.Country,
			},
		)
	}

	return contact
}
