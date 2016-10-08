package contacts

// Googlify takes a contact and turns it into a GoogleContact
func Googlify(contact Contact) *GoogleContact {
	var googleContact = GoogleContact{
		AtomNamespace: atomNamespace,
		GDNamespace:   googleNamespace,
		ID:            contact.ID,
		Etag:          contact.Etag,
		Category: GoogleCategory{
			Scheme: contact.Category.Scheme,
			Term:   contact.Category.Term,
		},
		Name: GoogleContactName{
			GivenName:  contact.Name.GivenName,
			FamilyName: contact.Name.FamilyName,
			FullName:   contact.Name.FullName,
		},
	}

	for _, email := range contact.Emails {
		googleContact.GoogleEmails = append(
			googleContact.GoogleEmails,
			GoogleEmail{
				email.Data,
				email.Primary,
				email.Rel,
				email.DisplayName,
			},
		)
	}

	for _, phoneNumber := range contact.PhoneNumbers {
		googleContact.GooglePhoneNumbers = append(
			googleContact.GooglePhoneNumbers,
			GooglePhoneNumber{
				phoneNumber.Number,
				phoneNumber.Primary,
				phoneNumber.Rel,
			},
		)
	}

	for _, address := range contact.Addresses {
		googleContact.GoogleAddresses = append(
			googleContact.GoogleAddresses,
			GoogleAddress{
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

	return &googleContact
}
