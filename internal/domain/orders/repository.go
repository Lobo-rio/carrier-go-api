package orders

type OrdersRepository interface {
	Save(order *Order) error
	GetAll() ([]Order, error)
	GetById(id string) (*Order, error)
	Update(order *Order) error
	Delete(order *Order) error
}
