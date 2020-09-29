// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/vcds-go-client/vcds/client/services"
	"github.com/IBM-Cloud/vcds-go-client/vcds/client/users_and_accounts"
	"github.com/IBM-Cloud/vcds-go-client/vcds/client/utils"
	"github.com/IBM-Cloud/vcds-go-client/vcds/client/vcenters"
)

// Default i b m cloud for vmware solutions HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new i b m cloud for vmware solutions HTTP client.
func NewHTTPClient(formats strfmt.Registry) *IBMCloudForVmwareSolutions {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new i b m cloud for vmware solutions HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *IBMCloudForVmwareSolutions {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new i b m cloud for vmware solutions client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *IBMCloudForVmwareSolutions {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(IBMCloudForVmwareSolutions)
	cli.Transport = transport

	cli.Services = services.New(transport, formats)

	cli.UsersAndAccounts = users_and_accounts.New(transport, formats)

	cli.Utils = utils.New(transport, formats)

	cli.Vcenters = vcenters.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// IBMCloudForVmwareSolutions is a client for i b m cloud for vmware solutions
type IBMCloudForVmwareSolutions struct {
	Services *services.Client

	UsersAndAccounts *users_and_accounts.Client

	Utils *utils.Client

	Vcenters *vcenters.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *IBMCloudForVmwareSolutions) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Services.SetTransport(transport)

	c.UsersAndAccounts.SetTransport(transport)

	c.Utils.SetTransport(transport)

	c.Vcenters.SetTransport(transport)

}
