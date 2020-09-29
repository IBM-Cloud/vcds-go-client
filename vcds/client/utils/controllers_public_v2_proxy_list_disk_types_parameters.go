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

// NewControllersPublicV2ProxyListDiskTypesParams creates a new ControllersPublicV2ProxyListDiskTypesParams object
// with the default values initialized.
func NewControllersPublicV2ProxyListDiskTypesParams() *ControllersPublicV2ProxyListDiskTypesParams {
	var ()
	return &ControllersPublicV2ProxyListDiskTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewControllersPublicV2ProxyListDiskTypesParamsWithTimeout creates a new ControllersPublicV2ProxyListDiskTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewControllersPublicV2ProxyListDiskTypesParamsWithTimeout(timeout time.Duration) *ControllersPublicV2ProxyListDiskTypesParams {
	var ()
	return &ControllersPublicV2ProxyListDiskTypesParams{

		timeout: timeout,
	}
}

// NewControllersPublicV2ProxyListDiskTypesParamsWithContext creates a new ControllersPublicV2ProxyListDiskTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewControllersPublicV2ProxyListDiskTypesParamsWithContext(ctx context.Context) *ControllersPublicV2ProxyListDiskTypesParams {
	var ()
	return &ControllersPublicV2ProxyListDiskTypesParams{

		Context: ctx,
	}
}

// NewControllersPublicV2ProxyListDiskTypesParamsWithHTTPClient creates a new ControllersPublicV2ProxyListDiskTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewControllersPublicV2ProxyListDiskTypesParamsWithHTTPClient(client *http.Client) *ControllersPublicV2ProxyListDiskTypesParams {
	var ()
	return &ControllersPublicV2ProxyListDiskTypesParams{
		HTTPClient: client,
	}
}

/*ControllersPublicV2ProxyListDiskTypesParams contains all the parameters to send to the API endpoint
for the controllers public v2 proxy list disk types operation typically these are written to a http.Request
*/
type ControllersPublicV2ProxyListDiskTypesParams struct {

	/*XGlobalTransactionID
	  Global transaction ID for request correlation.

	*/
	XGlobalTransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the controllers public v2 proxy list disk types params
func (o *ControllersPublicV2ProxyListDiskTypesParams) WithTimeout(timeout time.Duration) *ControllersPublicV2ProxyListDiskTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the controllers public v2 proxy list disk types params
func (o *ControllersPublicV2ProxyListDiskTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the controllers public v2 proxy list disk types params
func (o *ControllersPublicV2ProxyListDiskTypesParams) WithContext(ctx context.Context) *ControllersPublicV2ProxyListDiskTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the controllers public v2 proxy list disk types params
func (o *ControllersPublicV2ProxyListDiskTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the controllers public v2 proxy list disk types params
func (o *ControllersPublicV2ProxyListDiskTypesParams) WithHTTPClient(client *http.Client) *ControllersPublicV2ProxyListDiskTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the controllers public v2 proxy list disk types params
func (o *ControllersPublicV2ProxyListDiskTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXGlobalTransactionID adds the xGlobalTransactionID to the controllers public v2 proxy list disk types params
func (o *ControllersPublicV2ProxyListDiskTypesParams) WithXGlobalTransactionID(xGlobalTransactionID *string) *ControllersPublicV2ProxyListDiskTypesParams {
	o.SetXGlobalTransactionID(xGlobalTransactionID)
	return o
}

// SetXGlobalTransactionID adds the xGlobalTransactionId to the controllers public v2 proxy list disk types params
func (o *ControllersPublicV2ProxyListDiskTypesParams) SetXGlobalTransactionID(xGlobalTransactionID *string) {
	o.XGlobalTransactionID = xGlobalTransactionID
}

// WriteToRequest writes these params to a swagger request
func (o *ControllersPublicV2ProxyListDiskTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
