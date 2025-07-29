package carrier

type CarrierRepository interface {
	Save(carrier *Carrier) error
}