package orders

import (
	contracts "carrierCheck/internal/contracts/order"
	internalerrors "carrierCheck/internal/internal-errors"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	OrderPlaced        = "Pedido Criado"
	OrderCanceled      = "Pedido Cancelado"
	OrderWithTheCarrier = "Pedido com Transportadora"
	PaymentApproved    = "Pagamento Aprovado"
	SeparatedInStock   = "Separado em Estoque"
	InvoiceIssued      = "Nota Fiscal Emitida"
	InTransit          = "Em Trânsito"
	OutForDelivery     = "Saiu para Entrega"
	DeliveryCompleted  = "Entrega Concluída"
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
	SendMail   func(to string, subject string, body string) error
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

	go s.SendMailCreatedAt(&createOrder, order.ID)

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
		CarrierId: order.CarrierId,
		Status: order.Status,
		Products:   product,
		CreatedAt: order.CreatedAt.String(),
		UpdatedAt: order.UpdatedAt.String(),
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
	
	if order.Status == DeliveryCompleted || order.Status == OrderCanceled{
		return errors.New("order is not in " + order.Status + " status")
	}

	order.CarrierId = request.CarrierId
	order.Status = OrderWithTheCarrier
	order.UpdatedAt = time.Now()

	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}

	go s.SendMailStatusClient(order.ClientId, order.ID, order.Status)
	go s.SendMailStatusCarrier(order.CarrierId, order.ID, order.Status)

	return nil
}

func (s *OrdersServiceImp) UpdateStatusOrderCanceled(id string) error {
	order, err := s.Repository.GetById(id)
	fmt.Println("Updating order status to canceled", order.ID, order.Status)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	if order.Status == DeliveryCompleted || order.Status == OrderCanceled{
		return errors.New("order is not in " + order.Status + " status")
	}

	order.CarrierId = id
	order.Status = OrderCanceled
	order.UpdatedAt = time.Now()
	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}

	go s.SendMailStatusClient(order.ClientId, order.ID, order.Status)
	
	return nil
}

func (s *OrdersServiceImp) UpdateStatusPaymentApproved(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	if order.Status == DeliveryCompleted || order.Status == OrderCanceled{
		return errors.New("order is not in " + order.Status + " status")
	}

	order.Status = PaymentApproved
    order.UpdatedAt = time.Now()

	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}

	go s.SendMailStatusClient(order.ClientId, order.ID, order.Status)

	return nil
}

func (s *OrdersServiceImp) UpdateStatusSeparatedInStock(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	if order.Status == DeliveryCompleted || order.Status == OrderCanceled{
		return errors.New("order is not in " + order.Status + " status")
	}

	order.Status = SeparatedInStock
	order.UpdatedAt = time.Now()

	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}

	go s.SendMailStatusClient(order.ClientId, order.ID, order.Status)

	return nil
}

func (s *OrdersServiceImp) UpdateStatusInvoiceIssued(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	if order.Status == DeliveryCompleted || order.Status == OrderCanceled{
		return errors.New("order is not in " + order.Status + " status")
	}

	order.Status = InvoiceIssued
	order.UpdatedAt = time.Now()

	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}

	go s.SendMailStatusClient(order.ClientId, order.ID, order.Status)

	return nil
}

func (s *OrdersServiceImp) UpdateStatusInTransit(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	if order.Status == DeliveryCompleted || order.Status == OrderCanceled{
		return errors.New("order is not in " + order.Status + " status")
	}

	order.Status = InTransit
	order.UpdatedAt = time.Now()

	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}

	go s.SendMailStatusClient(order.ClientId, order.ID, order.Status)

	return nil
}

func (s *OrdersServiceImp) UpdateStatusOutForDelivery(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}
	
	if order.Status == DeliveryCompleted || order.Status == OrderCanceled{
		return errors.New("order is not in " + order.Status + " status")
	}

	order.Status = OutForDelivery
	order.UpdatedAt = time.Now()

	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}

	go s.SendMailStatusClient(order.ClientId, order.ID, order.Status)

	return nil
}

func (s *OrdersServiceImp) UpdateStatusDeliveryCompleted(id string) error {
	order, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	if order.Status == DeliveryCompleted || order.Status == OrderCanceled{
		return errors.New("order is not in " + order.Status + " status")
	}

	order.Status = DeliveryCompleted
    order.UpdatedAt = time.Now()

	err = s.Repository.Update(order)
	if err != nil {
		return internalerrors.ErrInternal
	}

	go s.SendMailStatusClient(order.ClientId, order.ID, order.Status)

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

func (s *OrdersServiceImp) SendMailCreatedAt(order *contracts.CreateOrder, numberOrder string) error {
	client, _ := s.Repository.GetByIdClients(order.ClientId)

	to := client.Name + "<" + client.Email + ">"
	subject := "Order Created Successfully"
    body := "<h1>Tente e Invente Novamente</h1><br>" +
	        "<h3>Tudo bem " + client.Name + " :-) !!, obrigado pela sua compra, o pedido foi criado com o número: " + numberOrder + 
			", acesse nosso site para acompanhar os status do seu pedido, clicando no link abaixo</h3><br><br>" +
			"<a href='https://lobo.rio.br'>Acompanhar Pedido</a><br><br>" +
	        "<h5>Seu pedido foi criado com sucesso!</h5><br>"

	err := s.SendMail(to, subject, body)
	if err != nil {
		return err
	}
	fmt.Println("Email enviado com sucesso!")
	return nil
}

func (s *OrdersServiceImp) SendMailStatusClient(id string, numberOrder string, status string) error {
	client, _ := s.Repository.GetByIdClients(id)

	to := client.Name + "<" + client.Email + ">"
	subject := "Order Status Update"
    body := "<h1>Tente e Invente Novamente</h1><br>" +
	        "<h3>Tudo bem " + client.Name + " :-) !!, O seu pedido entrou no status de " + status + 
			", para maiores informações, favor clicar no link abaixo e acessar o site. </h3><br><br>" + 
			"<a href='https://lobo.rio.br'>Acompanhar Pedido</a><br><br>" +
	        "<h5>Status do pedido " + numberOrder +  " foi alterado com sucesso!</h5><br>"

	err := s.SendMail(to, subject, body)
	if err != nil {
		return err
	}
	fmt.Println("Email enviado com sucesso!")
	return nil
}

func (s *OrdersServiceImp) SendMailStatusCarrier(id string, numberOrder string, status string) error {
	carrier, _ := s.Repository.GetByIdCarrier(id)
     emails := []string{}
	for _, e := range carrier.Email {
		emails = append(emails, e.Email)
	}	

	result := strings.Join(emails, ",")
	
	to := carrier.Contact + "<" + result + ">"
	subject := "Order Status Update Carrier"
    body := "<h1>Tente e Invente Novamente</h1><br>" +
	        "<h3>Tudo bem " + carrier.Contact + " :-) !!, O pedido entrou no status de " + status + 
			", então a sua empresa, está liberada para efetuar a entrega do(s) produto(s), para maiores informações, favor clicar no link abaixo e acessar o site. </h3><br><br>" + 
			"<a href='https://lobo.rio.br'>Acompanhar Pedido</a><br><br>" +
	        "<h5>Status do pedido " + numberOrder +  " foi alterado com sucesso!</h5><br>"

	err := s.SendMail(to, subject, body)
	if err != nil {
		return err
	}
	fmt.Println("Email enviado com sucesso carrier!")
	return nil
}