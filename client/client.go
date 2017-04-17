package client

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the eventory service client.
type Client struct {
	*goaclient.Client
	UserTokenSigner goaclient.Signer
	CronTokenSigner goaclient.Signer
	Encoder         *goa.HTTPEncoder
	Decoder         *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	client.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")
	client.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	client.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}

// SetUserTokenSigner sets the request signer for the userToken security scheme.
func (c *Client) SetUserTokenSigner(signer goaclient.Signer) {
	c.UserTokenSigner = signer
}

// SetCronTokenSigner sets the request signer for the cronToken security scheme.
func (c *Client) SetCronTokenSigner(signer goaclient.Signer) {
	c.CronTokenSigner = signer
}
