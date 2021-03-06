// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPluginParams creates a new GetPluginParams object
// with the default values initialized.
func NewGetPluginParams() *GetPluginParams {
	var ()
	return &GetPluginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPluginParamsWithTimeout creates a new GetPluginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPluginParamsWithTimeout(timeout time.Duration) *GetPluginParams {
	var ()
	return &GetPluginParams{

		timeout: timeout,
	}
}

// NewGetPluginParamsWithContext creates a new GetPluginParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPluginParamsWithContext(ctx context.Context) *GetPluginParams {
	var ()
	return &GetPluginParams{

		Context: ctx,
	}
}

// NewGetPluginParamsWithHTTPClient creates a new GetPluginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPluginParamsWithHTTPClient(client *http.Client) *GetPluginParams {
	var ()
	return &GetPluginParams{
		HTTPClient: client,
	}
}

/*GetPluginParams contains all the parameters to send to the API endpoint
for the get plugin operation typically these are written to a http.Request
*/
type GetPluginParams struct {

	/*Pname*/
	Pname string
	/*Ptype*/
	Ptype string
	/*Pversion*/
	Pversion int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get plugin params
func (o *GetPluginParams) WithTimeout(timeout time.Duration) *GetPluginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plugin params
func (o *GetPluginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plugin params
func (o *GetPluginParams) WithContext(ctx context.Context) *GetPluginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plugin params
func (o *GetPluginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plugin params
func (o *GetPluginParams) WithHTTPClient(client *http.Client) *GetPluginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plugin params
func (o *GetPluginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPname adds the pname to the get plugin params
func (o *GetPluginParams) WithPname(pname string) *GetPluginParams {
	o.SetPname(pname)
	return o
}

// SetPname adds the pname to the get plugin params
func (o *GetPluginParams) SetPname(pname string) {
	o.Pname = pname
}

// WithPtype adds the ptype to the get plugin params
func (o *GetPluginParams) WithPtype(ptype string) *GetPluginParams {
	o.SetPtype(ptype)
	return o
}

// SetPtype adds the ptype to the get plugin params
func (o *GetPluginParams) SetPtype(ptype string) {
	o.Ptype = ptype
}

// WithPversion adds the pversion to the get plugin params
func (o *GetPluginParams) WithPversion(pversion int64) *GetPluginParams {
	o.SetPversion(pversion)
	return o
}

// SetPversion adds the pversion to the get plugin params
func (o *GetPluginParams) SetPversion(pversion int64) {
	o.Pversion = pversion
}

// WriteToRequest writes these params to a swagger request
func (o *GetPluginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pname
	if err := r.SetPathParam("pname", o.Pname); err != nil {
		return err
	}

	// path param ptype
	if err := r.SetPathParam("ptype", o.Ptype); err != nil {
		return err
	}

	// path param pversion
	if err := r.SetPathParam("pversion", swag.FormatInt64(o.Pversion)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
