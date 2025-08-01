package clients

type ClientsRepository interface {
	Save(client *Client) error
	GetAll() ([]Client, error)
	GetById(id string) (*Client, error)
	Update(client *Client) error
	Delete(client *Client) error
}