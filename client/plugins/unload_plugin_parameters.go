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

// NewUnloadPluginParams creates a new UnloadPluginParams object
// with the default values initialized.
func NewUnloadPluginParams() *UnloadPluginParams {
	var ()
	return &UnloadPluginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnloadPluginParamsWithTimeout creates a new UnloadPluginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnloadPluginParamsWithTimeout(timeout time.Duration) *UnloadPluginParams {
	var ()
	return &UnloadPluginParams{

		timeout: timeout,
	}
}

// NewUnloadPluginParamsWithContext creates a new UnloadPluginParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnloadPluginParamsWithContext(ctx context.Context) *UnloadPluginParams {
	var ()
	return &UnloadPluginParams{

		Context: ctx,
	}
}

// NewUnloadPluginParamsWithHTTPClient creates a new UnloadPluginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnloadPluginParamsWithHTTPClient(client *http.Client) *UnloadPluginParams {
	var ()
	return &UnloadPluginParams{
		HTTPClient: client,
	}
}

/*UnloadPluginParams contains all the parameters to send to the API endpoint
for the unload plugin operation typically these are written to a http.Request
*/
type UnloadPluginParams struct {

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

// WithTimeout adds the timeout to the unload plugin params
func (o *UnloadPluginParams) WithTimeout(timeout time.Duration) *UnloadPluginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unload plugin params
func (o *UnloadPluginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unload plugin params
func (o *UnloadPluginParams) WithContext(ctx context.Context) *UnloadPluginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unload plugin params
func (o *UnloadPluginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unload plugin params
func (o *UnloadPluginParams) WithHTTPClient(client *http.Client) *UnloadPluginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unload plugin params
func (o *UnloadPluginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPname adds the pname to the unload plugin params
func (o *UnloadPluginParams) WithPname(pname string) *UnloadPluginParams {
	o.SetPname(pname)
	return o
}

// SetPname adds the pname to the unload plugin params
func (o *UnloadPluginParams) SetPname(pname string) {
	o.Pname = pname
}

// WithPtype adds the ptype to the unload plugin params
func (o *UnloadPluginParams) WithPtype(ptype string) *UnloadPluginParams {
	o.SetPtype(ptype)
	return o
}

// SetPtype adds the ptype to the unload plugin params
func (o *UnloadPluginParams) SetPtype(ptype string) {
	o.Ptype = ptype
}

// WithPversion adds the pversion to the unload plugin params
func (o *UnloadPluginParams) WithPversion(pversion int64) *UnloadPluginParams {
	o.SetPversion(pversion)
	return o
}

// SetPversion adds the pversion to the unload plugin params
func (o *UnloadPluginParams) SetPversion(pversion int64) {
	o.Pversion = pversion
}

// WriteToRequest writes these params to a swagger request
func (o *UnloadPluginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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