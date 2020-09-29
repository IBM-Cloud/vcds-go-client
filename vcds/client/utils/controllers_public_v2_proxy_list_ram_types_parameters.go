// Code generated by go-swagger; DO NOT EDIT.

package utils

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewControllersPublicV2ProxyListRAMTypesParams creates a new ControllersPublicV2ProxyListRAMTypesParams object
// with the default values initialized.
func NewControllersPublicV2ProxyListRAMTypesParams() *ControllersPublicV2ProxyListRAMTypesParams {
	var ()
	return &ControllersPublicV2ProxyListRAMTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewControllersPublicV2ProxyListRAMTypesParamsWithTimeout creates a new ControllersPublicV2ProxyListRAMTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewControllersPublicV2ProxyListRAMTypesParamsWithTimeout(timeout time.Duration) *ControllersPublicV2ProxyListRAMTypesParams {
	var ()
	return &ControllersPublicV2ProxyListRAMTypesParams{

		timeout: timeout,
	}
}

// NewControllersPublicV2ProxyListRAMTypesParamsWithContext creates a new ControllersPublicV2ProxyListRAMTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewControllersPublicV2ProxyListRAMTypesParamsWithContext(ctx context.Context) *ControllersPublicV2ProxyListRAMTypesParams {
	var ()
	return &ControllersPublicV2ProxyListRAMTypesParams{

		Context: ctx,
	}
}

// NewControllersPublicV2ProxyListRAMTypesParamsWithHTTPClient creates a new ControllersPublicV2ProxyListRAMTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewControllersPublicV2ProxyListRAMTypesParamsWithHTTPClient(client *http.Client) *ControllersPublicV2ProxyListRAMTypesParams {
	var ()
	return &ControllersPublicV2ProxyListRAMTypesParams{
		HTTPClient: client,
	}
}

/*ControllersPublicV2ProxyListRAMTypesParams contains all the parameters to send to the API endpoint
for the controllers public v2 proxy list ram types operation typically these are written to a http.Request
*/
type ControllersPublicV2ProxyListRAMTypesParams struct {

	/*XGlobalTransactionID
	  Global transaction ID for request correlation.

	*/
	XGlobalTransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the controllers public v2 proxy list ram types params
func (o *ControllersPublicV2ProxyListRAMTypesParams) WithTimeout(timeout time.Duration) *ControllersPublicV2ProxyListRAMTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the controllers public v2 proxy list ram types params
func (o *ControllersPublicV2ProxyListRAMTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the controllers public v2 proxy list ram types params
func (o *ControllersPublicV2ProxyListRAMTypesParams) WithContext(ctx context.Context) *ControllersPublicV2ProxyListRAMTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the controllers public v2 proxy list ram types params
func (o *ControllersPublicV2ProxyListRAMTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the controllers public v2 proxy list ram types params
func (o *ControllersPublicV2ProxyListRAMTypesParams) WithHTTPClient(client *http.Client) *ControllersPublicV2ProxyListRAMTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the controllers public v2 proxy list ram types params
func (o *ControllersPublicV2ProxyListRAMTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXGlobalTransactionID adds the xGlobalTransactionID to the controllers public v2 proxy list ram types params
func (o *ControllersPublicV2ProxyListRAMTypesParams) WithXGlobalTransactionID(xGlobalTransactionID *string) *ControllersPublicV2ProxyListRAMTypesParams {
	o.SetXGlobalTransactionID(xGlobalTransactionID)
	return o
}

// SetXGlobalTransactionID adds the xGlobalTransactionId to the controllers public v2 proxy list ram types params
func (o *ControllersPublicV2ProxyListRAMTypesParams) SetXGlobalTransactionID(xGlobalTransactionID *string) {
	o.XGlobalTransactionID = xGlobalTransactionID
}

// WriteToRequest writes these params to a swagger request
func (o *ControllersPublicV2ProxyListRAMTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XGlobalTransactionID != nil {

		// header param x-global-transaction-id
		if err := r.SetHeaderParam("x-global-transaction-id", *o.XGlobalTransactionID); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
