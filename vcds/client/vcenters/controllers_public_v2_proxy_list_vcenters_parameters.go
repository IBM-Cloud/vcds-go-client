// Code generated by go-swagger; DO NOT EDIT.

package vcenters

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

// NewControllersPublicV2ProxyListVcentersParams creates a new ControllersPublicV2ProxyListVcentersParams object
// with the default values initialized.
func NewControllersPublicV2ProxyListVcentersParams() *ControllersPublicV2ProxyListVcentersParams {
	var ()
	return &ControllersPublicV2ProxyListVcentersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewControllersPublicV2ProxyListVcentersParamsWithTimeout creates a new ControllersPublicV2ProxyListVcentersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewControllersPublicV2ProxyListVcentersParamsWithTimeout(timeout time.Duration) *ControllersPublicV2ProxyListVcentersParams {
	var ()
	return &ControllersPublicV2ProxyListVcentersParams{

		timeout: timeout,
	}
}

// NewControllersPublicV2ProxyListVcentersParamsWithContext creates a new ControllersPublicV2ProxyListVcentersParams object
// with the default values initialized, and the ability to set a context for a request
func NewControllersPublicV2ProxyListVcentersParamsWithContext(ctx context.Context) *ControllersPublicV2ProxyListVcentersParams {
	var ()
	return &ControllersPublicV2ProxyListVcentersParams{

		Context: ctx,
	}
}

// NewControllersPublicV2ProxyListVcentersParamsWithHTTPClient creates a new ControllersPublicV2ProxyListVcentersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewControllersPublicV2ProxyListVcentersParamsWithHTTPClient(client *http.Client) *ControllersPublicV2ProxyListVcentersParams {
	var ()
	return &ControllersPublicV2ProxyListVcentersParams{
		HTTPClient: client,
	}
}

/*ControllersPublicV2ProxyListVcentersParams contains all the parameters to send to the API endpoint
for the controllers public v2 proxy list vcenters operation typically these are written to a http.Request
*/
type ControllersPublicV2ProxyListVcentersParams struct {

	/*Authorization
	  Your IBM Cloud Identity and Access Management (IAM) token. To retrieve your IAM token, run `ibmcloud iam oauth-tokens`.

	*/
	Authorization string
	/*XGlobalTransactionID
	  Global transaction ID for request correlation.

	*/
	XGlobalTransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the controllers public v2 proxy list vcenters params
func (o *ControllersPublicV2ProxyListVcentersParams) WithTimeout(timeout time.Duration) *ControllersPublicV2ProxyListVcentersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the controllers public v2 proxy list vcenters params
func (o *ControllersPublicV2ProxyListVcentersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the controllers public v2 proxy list vcenters params
func (o *ControllersPublicV2ProxyListVcentersParams) WithContext(ctx context.Context) *ControllersPublicV2ProxyListVcentersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the controllers public v2 proxy list vcenters params
func (o *ControllersPublicV2ProxyListVcentersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the controllers public v2 proxy list vcenters params
func (o *ControllersPublicV2ProxyListVcentersParams) WithHTTPClient(client *http.Client) *ControllersPublicV2ProxyListVcentersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the controllers public v2 proxy list vcenters params
func (o *ControllersPublicV2ProxyListVcentersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the controllers public v2 proxy list vcenters params
func (o *ControllersPublicV2ProxyListVcentersParams) WithAuthorization(authorization string) *ControllersPublicV2ProxyListVcentersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the controllers public v2 proxy list vcenters params
func (o *ControllersPublicV2ProxyListVcentersParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXGlobalTransactionID adds the xGlobalTransactionID to the controllers public v2 proxy list vcenters params
func (o *ControllersPublicV2ProxyListVcentersParams) WithXGlobalTransactionID(xGlobalTransactionID *string) *ControllersPublicV2ProxyListVcentersParams {
	o.SetXGlobalTransactionID(xGlobalTransactionID)
	return o
}

// SetXGlobalTransactionID adds the xGlobalTransactionId to the controllers public v2 proxy list vcenters params
func (o *ControllersPublicV2ProxyListVcentersParams) SetXGlobalTransactionID(xGlobalTransactionID *string) {
	o.XGlobalTransactionID = xGlobalTransactionID
}

// WriteToRequest writes these params to a swagger request
func (o *ControllersPublicV2ProxyListVcentersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

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
