package carrier

type CarrierRepository interface {
	Save(carrier *Carrier) error
	GetAll() ([]Carrier, error)
	GetById(id string) (*Carrier, error)
	Update(carrier *Carrier) error
	Delete(carrier *Carrier) error
}