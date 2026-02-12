package contact

type ContactService struct {
	Repo *ContactRepository
}

func (svc *ContactService) CreateContact(contact *Contact) error {
	return svc.Repo.Create(contact)
}

func (svc *ContactService) GetContectsByID(id uint) (*Contact, error) {
	return svc.Repo.FindByID(id)
}

func (svc *ContactService) GetContactsByName(name string) ([]Contact, error) {
	return svc.Repo.FindByName(name)
}

func (svc *ContactService) GetContactsByPhone(phone string) ([]Contact, error) {
	return svc.Repo.FindByPhone(phone)
}

func (svc *ContactService) GetContactsByMessage(message string) ([]Contact, error) {
	return svc.Repo.FindByMessage(message)
}

func (svc *ContactService) UpdateContact(contact *Contact) error {
	return svc.Repo.Update(contact)
}

func (svc *ContactService) DeleteContact(id uint) error {
	return svc.Repo.Delete(id)
}
