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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/vcds-go-client/vcds/models"
)

// NewControllersPublicV1ProxyOrderVcenterClusterParams creates a new ControllersPublicV1ProxyOrderVcenterClusterParams object
// with the default values initialized.
func NewControllersPublicV1ProxyOrderVcenterClusterParams() *ControllersPublicV1ProxyOrderVcenterClusterParams {
	var (
		checkPriceDefault = bool(false)
		verifyOnlyDefault = bool(false)
	)
	return &ControllersPublicV1ProxyOrderVcenterClusterParams{
		CheckPrice: &checkPriceDefault,
		VerifyOnly: &verifyOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewControllersPublicV1ProxyOrderVcenterClusterParamsWithTimeout creates a new ControllersPublicV1ProxyOrderVcenterClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewControllersPublicV1ProxyOrderVcenterClusterParamsWithTimeout(timeout time.Duration) *ControllersPublicV1ProxyOrderVcenterClusterParams {
	var (
		checkPriceDefault = bool(false)
		verifyOnlyDefault = bool(false)
	)
	return &ControllersPublicV1ProxyOrderVcenterClusterParams{
		CheckPrice: &checkPriceDefault,
		VerifyOnly: &verifyOnlyDefault,

		timeout: timeout,
	}
}

// NewControllersPublicV1ProxyOrderVcenterClusterParamsWithContext creates a new ControllersPublicV1ProxyOrderVcenterClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewControllersPublicV1ProxyOrderVcenterClusterParamsWithContext(ctx context.Context) *ControllersPublicV1ProxyOrderVcenterClusterParams {
	var (
		checkPriceDefault = bool(false)
		verifyOnlyDefault = bool(false)
	)
	return &ControllersPublicV1ProxyOrderVcenterClusterParams{
		CheckPrice: &checkPriceDefault,
		VerifyOnly: &verifyOnlyDefault,

		Context: ctx,
	}
}

// NewControllersPublicV1ProxyOrderVcenterClusterParamsWithHTTPClient creates a new ControllersPublicV1ProxyOrderVcenterClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewControllersPublicV1ProxyOrderVcenterClusterParamsWithHTTPClient(client *http.Client) *ControllersPublicV1ProxyOrderVcenterClusterParams {
	var (
		checkPriceDefault = bool(false)
		verifyOnlyDefault = bool(false)
	)
	return &ControllersPublicV1ProxyOrderVcenterClusterParams{
		CheckPrice: &checkPriceDefault,
		VerifyOnly: &verifyOnlyDefault,
		HTTPClient: client,
	}
}

/*ControllersPublicV1ProxyOrderVcenterClusterParams contains all the parameters to send to the API endpoint
for the controllers public v1 proxy order vcenter cluster operation typically these are written to a http.Request
*/
type ControllersPublicV1ProxyOrderVcenterClusterParams struct {

	/*Authorization
	  Your IBM Cloud Identity and Access Management (IAM) token. To retrieve your IAM token, run `ibmcloud iam oauth-tokens`.

	*/
	Authorization string
	/*CheckPrice
	  Whether to display the high-level price associated with the order.

	*/
	CheckPrice *bool
	/*ClusterOrderData
	  Cluster order specification with configurations.

	*/
	ClusterOrderData *models.ClusterOrderData
	/*InstanceID
	  Instance ID.

	*/
	InstanceID string
	/*VerifyOnly
	  Whether to verify the order only instead of performing the order for real.

	*/
	VerifyOnly *bool
	/*XGlobalTransactionID
	  Global transaction ID for request correlation.

	*/
	XGlobalTransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) WithTimeout(timeout time.Duration) *ControllersPublicV1ProxyOrderVcenterClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) WithContext(ctx context.Context) *ControllersPublicV1ProxyOrderVcenterClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) WithHTTPClient(client *http.Client) *ControllersPublicV1ProxyOrderVcenterClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) WithAuthorization(authorization string) *ControllersPublicV1ProxyOrderVcenterClusterParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithCheckPrice adds the checkPrice to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) WithCheckPrice(checkPrice *bool) *ControllersPublicV1ProxyOrderVcenterClusterParams {
	o.SetCheckPrice(checkPrice)
	return o
}

// SetCheckPrice adds the checkPrice to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) SetCheckPrice(checkPrice *bool) {
	o.CheckPrice = checkPrice
}

// WithClusterOrderData adds the clusterOrderData to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) WithClusterOrderData(clusterOrderData *models.ClusterOrderData) *ControllersPublicV1ProxyOrderVcenterClusterParams {
	o.SetClusterOrderData(clusterOrderData)
	return o
}

// SetClusterOrderData adds the clusterOrderData to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) SetClusterOrderData(clusterOrderData *models.ClusterOrderData) {
	o.ClusterOrderData = clusterOrderData
}

// WithInstanceID adds the instanceID to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) WithInstanceID(instanceID string) *ControllersPublicV1ProxyOrderVcenterClusterParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithVerifyOnly adds the verifyOnly to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) WithVerifyOnly(verifyOnly *bool) *ControllersPublicV1ProxyOrderVcenterClusterParams {
	o.SetVerifyOnly(verifyOnly)
	return o
}

// SetVerifyOnly adds the verifyOnly to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) SetVerifyOnly(verifyOnly *bool) {
	o.VerifyOnly = verifyOnly
}

// WithXGlobalTransactionID adds the xGlobalTransactionID to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) WithXGlobalTransactionID(xGlobalTransactionID *string) *ControllersPublicV1ProxyOrderVcenterClusterParams {
	o.SetXGlobalTransactionID(xGlobalTransactionID)
	return o
}

// SetXGlobalTransactionID adds the xGlobalTransactionId to the controllers public v1 proxy order vcenter cluster params
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) SetXGlobalTransactionID(xGlobalTransactionID *string) {
	o.XGlobalTransactionID = xGlobalTransactionID
}

// WriteToRequest writes these params to a swagger request
func (o *ControllersPublicV1ProxyOrderVcenterClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.CheckPrice != nil {

		// query param check_price
		var qrCheckPrice bool
		if o.CheckPrice != nil {
			qrCheckPrice = *o.CheckPrice
		}
		qCheckPrice := swag.FormatBool(qrCheckPrice)
		if qCheckPrice != "" {
			if err := r.SetQueryParam("check_price", qCheckPrice); err != nil {
				return err
			}
		}

	}

	if o.ClusterOrderData != nil {
		if err := r.SetBodyParam(o.ClusterOrderData); err != nil {
			return err
		}
	}

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
		return err
	}

	if o.VerifyOnly != nil {

		// query param verify_only
		var qrVerifyOnly bool
		if o.VerifyOnly != nil {
			qrVerifyOnly = *o.VerifyOnly
		}
		qVerifyOnly := swag.FormatBool(qrVerifyOnly)
		if qVerifyOnly != "" {
			if err := r.SetQueryParam("verify_only", qVerifyOnly); err != nil {
				return err
			}
		}

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
