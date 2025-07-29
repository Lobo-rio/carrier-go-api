package clients

type ClientsRepository interface {
	Save(client *Client) error
}