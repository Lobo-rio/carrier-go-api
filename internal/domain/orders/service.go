package orders

import (
	contracts "carrierCheck/internal/contracts/order"
	internalerrors "carrierCheck/internal/internal-errors"
	"errors"
)

const (
	OrderPlaced        = "Order Placed"
	OrderCanceled      = "Order Canceled"
	PaymentApproved    = "Payment Approved"
	SeparatedInStock   = "Separated In Stock"
	InvoiceIssued      = "Invoice Issued"
	InTransit          = "In Transit"
	OutForDelivery     = "Out For Delivery"
	DeliveryCompleted  = "Delivery Completed"
)

type OrdersService interface {
	Create(createOrder contracts.CreateOrder) (string, error)
	GetById(id string) (*contracts.ResponseOrder, error)
	GetAll() ([]contracts.ResponseOrder, error)
	Update(id string, request contracts.UpdateOrder) error
	UpdateCarrier(id string, request contracts.UpdateCarrierOrder) error
	UpdateStatusOrderCanceled(id string) error
	UpdateStatusPaymentApproved(id string) error
	UpdateStatusSeparatedInStock(id string) error
	UpdateStatusInvoiceIssued(id string) error
	UpdateStatusInTransit(id string) error
	UpdateStatusOutForDelivery(id string) error
	UpdateStatusDeliveryCompleted(id string) error
	Delete(id string) error
}

type OrdersServiceImp struct {
	Repository OrdersRepository
}

func (s *OrdersServiceImp) Create(createOrder contracts.CreateOrder) (string, error) {
	product := make([]OrdersProducts, len(createOrder.Products))
	for i, prd := range createOrder.Products {
		product[i] = OrdersProducts{
			ProductId:      prd.ProductId,
			Quantity:       prd.Quantity,
			Price:   prd.Price,
		}
	}
	order, err := NewOrder(createOrder.ClientId, createOrder.AddressId, product)
	if err != nil {
		return "", err
	}

	order.Status = OrderPlaced

	err = s.Repository.Save(order)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return order.ID, nil
}

func (s *OrdersServiceImp) GetById(id string) (*contracts.ResponseOrder, error) {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return nil, internalerrors.ProcessErrorToReturn(err)
	}
	
	product := make([]OrdersProducts, len(order.OrderProduct))
	for i, prd := range order.OrderProduct {
		product[i] = OrdersProducts{
			ProductId: prd.ProductId,
			Quantity:  prd.Quantity,
			Price:     prd.Price,
		}
	}
	return &contracts.ResponseOrder{
		ID:        order.ID,
		ClientId: order.ClientId,
		AddressId: order.AddressId,
		Products:   product,
		CreatedAt: order.CreatedAt.String(),
	}, nil
}

func (s *OrdersServiceImp) GetAll() ([]contracts.ResponseOrder, error) {
	orders, err := s.Repository.GetAll()
	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	responseOrders := make([]contracts.ResponseOrder, len(orders))
	for i, order := range orders {
		product := make([]OrdersProducts, len(order.OrderProduct))
		for j, prd := range order.OrderProduct {
			product[j] = OrdersProducts{
				ProductId: prd.ProductId,
				Quantity:  prd.Quantity,
				Price:     prd.Price,
			}
		}
		responseOrders[i] = contracts.ResponseOrder{
			ID:        order.ID,
			ClientId:  order.ClientId,
			AddressId: order.AddressId,
			Products:  product,
			CreatedAt: order.CreatedAt.String(),
		}
	}

	return responseOrders, nil
}

func (s *OrdersServiceImp) Update(id string, request contracts.UpdateOrder) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	order.ClientId = request.ClientId
	order.AddressId = request.AddressId

	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}

func (s *OrdersServiceImp) UpdateCarrier(id string, request contracts.UpdateCarrierOrder) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}

func (s *OrdersServiceImp) UpdateStatusOrderCanceled(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	order.Status = OrderCanceled

	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}
	return nil
}

func (s *OrdersServiceImp) UpdateStatusPaymentApproved(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	if order.Status == OrderCanceled {
		return errors.New("order is already canceled")
	}

	order.Status = PaymentApproved

	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}
	return nil
}

func (s *OrdersServiceImp) UpdateStatusSeparatedInStock(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	if order.Status == OrderCanceled {
		return errors.New("order is already canceled")
	}

	order.Status = SeparatedInStock
	
	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}
	return nil
}

func (s *OrdersServiceImp) UpdateStatusInvoiceIssued(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	if order.Status == OrderCanceled {
		return errors.New("order is already canceled")
	}

	order.Status = InvoiceIssued
	
	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}
	return nil
}

func (s *OrdersServiceImp) UpdateStatusInTransit(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	if order.Status == OrderCanceled {
		return errors.New("order is already canceled")
	}

	order.Status = InTransit
	
	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}
	return nil
}

func (s *OrdersServiceImp) UpdateStatusOutForDelivery(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	if order.Status == OrderCanceled {
		return errors.New("order is already canceled")
	}

	order.Status = OutForDelivery
	
	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}
	return nil
}

func (s *OrdersServiceImp) UpdateStatusDeliveryCompleted(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	if order.Status == OrderCanceled {
		return errors.New("order is already canceled")
	}

	order.Status = DeliveryCompleted

	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}
	return nil
}

func (s *OrdersServiceImp) Delete(id string) error {
	carrier, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	err = s.Repository.Delete(carrier)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}
			