package contact

type ContactHandler struct {
	Svc *ContactService
}

func (hdl *ContactHandler) CreateContactByID(contact *Contact) error {
	return hdl.Svc.CreateContact(contact)
}

func (hdl *ContactHandler) GetContactsByName(name string) ([]Contact, error) {
	return hdl.Svc.GetContactsByName(name)
}

func (hdl *ContactHandler) GetContactsByPhone(phone string) ([]Contact, error) {
	return hdl.Svc.GetContactsByPhone(phone)
}

func (hdl *ContactHandler) GetContactsByMessage(message string) ([]Contact, error) {
	return hdl.Svc.GetContactsByMessage(message)
}

func (hdl *ContactHandler) UpdateContact(contact *Contact) error {
	return hdl.Svc.UpdateContact(contact)
}

func (hdl *ContactHandler) DeleteContact(id uint) error {
	return hdl.Svc.DeleteContact(id)
}
