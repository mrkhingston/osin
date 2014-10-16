package osin

// Client information
type Client interface {
	// Client id
	GetId() string

	// Client secret
	GetSecret() string

	// Base client uri
	GetRedirectUri() string

	// Allow client secret to be ignored for mobile and desktop clients
	// Set to true means no client secret required and only password
	// token requests allowed.
	GetIsPasswordOnlyClient() bool

	// Data to be passed to storage. Not used by the library.
	GetUserData() interface{}
}

// DefaultClient stores all data in struct variables
type DefaultClient struct {
	Id                   string
	Secret               string
	RedirectUri          string
	IsPasswordOnlyClient bool
	UserData             interface{}
}

func (d *DefaultClient) GetId() string {
	return d.Id
}

func (d *DefaultClient) GetSecret() string {
	return d.Secret
}

func (d *DefaultClient) GetRedirectUri() string {
	return d.RedirectUri
}

func (d *DefaultClient) GetIsPasswordOnlyClient() bool {
	return d.IsPasswordOnlyClient
}

func (d *DefaultClient) GetUserData() interface{} {
	return d.UserData
}

func (d *DefaultClient) CopyFrom(client Client) {
	d.Id = client.GetId()
	d.Secret = client.GetSecret()
	d.RedirectUri = client.GetRedirectUri()
	d.IsPasswordOnlyClient = client.GetIsPasswordOnlyClient()
	d.UserData = client.GetUserData()
}
