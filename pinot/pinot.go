package pinot

import (
	"fmt"
	"strings"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

// Client is a client for working with the Pinot API.
type Client struct {
	sling *sling.Sling
	*Options
}

// Options holds certain values useful to run the requests
type Options struct {
	controllerURL []string
	zookeeperURL  []string
	brokerURL     []string
}

// NewClient creates a new client to send HTTP requests to an Apache Pinot CLuster
func NewClient(opts ...func(*Options)) *Client {
	o := &Options{}
	// call option functions on instance to set options on it
	for _, opt := range opts {
		opt(o)
	}

	// add all custom validations here
	validate.RegisterValidation("role", ValidateRole)

	return &Client{
		sling:   sling.New(),
		Options: o,
	}
}

// WitchControllerURL functional option. The parameter in input can be a list of
// urls separate by comma. eg: controller-1:9000,controller-2:9000,controller-3:9000,
func WitchControllerURL(c string) func(*Options) {
	return func(r *Options) {
		r.controllerURL = strings.Split(c, ",")
	}
}

// WitchBrokerURL functional option. The parameter in input can be a list of
// urls separate by comma. eg: broker-1:9000,broker-2:9000,broker-3:9000,
func WitchBrokerURL(b string) func(*Options) {
	return func(r *Options) {
		r.brokerURL = strings.Split(b, ",")
	}
}

// WitchZookeeperURL functional option. The parameter in input can be a list of
// urls separate by comma. eg: zookeeper-1:9000,zookeeper-2:9000,zookeeper-3:9000,
func WitchZookeeperURL(zk string) func(*Options) {
	return func(r *Options) {
		r.zookeeperURL = strings.Split(zk, ",")
	}
}

// ErrorPinot represents an error returned by the Pinot API.
type ErrorPinot struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

// Error string reformat
func (e ErrorPinot) Error() string {
	return fmt.Sprintf("pinot: %d %s", e.Status, e.Detail)
}

// post makes a POST request to the specified Pinot endpoint
func (c *Client) post(url string, body interface{}, result interface{}) error {
	ep := new(ErrorPinot)
	_, err := c.sling.Post(url).BodyJSON(body).Receive(result, ep)
	if err != nil {
		return err
	}
	if ep.Status != 0 {
		return ep
	}
	return nil
}

// get makes a GET request to the specified Pinot endpoint
func (c *Client) get(url string, result interface{}) error {
	ep := new(ErrorPinot)
	_, err := c.sling.Get(url).Receive(result, ep)
	if err != nil {
		return err
	}
	if ep.Status != 0 {
		return ep
	}
	return nil
}

// put makes a PUT request to the specified Pinot endpoint
func (c *Client) put(url string, body interface{}, result interface{}) error {
	ep := new(ErrorPinot)
	_, err := c.sling.Put(url).BodyJSON(body).Receive(result, ep)
	if err != nil {
		return err
	}
	if ep.Status != 0 {
		return ep
	}
	return nil
}

// patch makes a PATCH request to the specified Pinot endpoint
func (c *Client) patch(url string, result interface{}) error {
	ep := new(ErrorPinot)
	_, err := c.sling.Patch(url).Receive(result, ep)
	if err != nil {
		return err
	}
	if ep.Status != 0 {
		return ep
	}
	return nil
}

// delete makes a DELETE request to the specified Pinot endpoint
func (c *Client) delete(url string, result interface{}) error {
	ep := new(ErrorPinot)
	_, err := c.sling.Delete(url).Receive(result, ep)
	if err != nil {
		return err
	}
	if ep.Status != 0 {
		return ep
	}
	return nil
}
