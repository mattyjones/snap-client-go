package snap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSwapPluginsParams creates a new SwapPluginsParams object
// with the default values initialized.
func NewSwapPluginsParams() *SwapPluginsParams {
	var ()
	return &SwapPluginsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSwapPluginsParamsWithTimeout creates a new SwapPluginsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSwapPluginsParamsWithTimeout(timeout time.Duration) *SwapPluginsParams {
	var ()
	return &SwapPluginsParams{

		timeout: timeout,
	}
}

// NewSwapPluginsParamsWithContext creates a new SwapPluginsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSwapPluginsParamsWithContext(ctx context.Context) *SwapPluginsParams {
	var ()
	return &SwapPluginsParams{

		Context: ctx,
	}
}

// NewSwapPluginsParamsWithHTTPClient creates a new SwapPluginsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSwapPluginsParamsWithHTTPClient(client *http.Client) *SwapPluginsParams {
	var ()
	return &SwapPluginsParams{
		HTTPClient: client,
	}
}

/*SwapPluginsParams contains all the parameters to send to the API endpoint
for the swap plugins operation typically these are written to a http.Request
*/
type SwapPluginsParams struct {

	/*PluginData
	  loads a plugin.

	*/
	PluginData *os.File
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

// WithTimeout adds the timeout to the swap plugins params
func (o *SwapPluginsParams) WithTimeout(timeout time.Duration) *SwapPluginsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the swap plugins params
func (o *SwapPluginsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the swap plugins params
func (o *SwapPluginsParams) WithContext(ctx context.Context) *SwapPluginsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the swap plugins params
func (o *SwapPluginsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the swap plugins params
func (o *SwapPluginsParams) WithHTTPClient(client *http.Client) *SwapPluginsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the swap plugins params
func (o *SwapPluginsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPluginData adds the pluginData to the swap plugins params
func (o *SwapPluginsParams) WithPluginData(pluginData *os.File) *SwapPluginsParams {
	o.SetPluginData(pluginData)
	return o
}

// SetPluginData adds the pluginData to the swap plugins params
func (o *SwapPluginsParams) SetPluginData(pluginData *os.File) {
	o.PluginData = pluginData
}

// WithPname adds the pname to the swap plugins params
func (o *SwapPluginsParams) WithPname(pname string) *SwapPluginsParams {
	o.SetPname(pname)
	return o
}

// SetPname adds the pname to the swap plugins params
func (o *SwapPluginsParams) SetPname(pname string) {
	o.Pname = pname
}

// WithPtype adds the ptype to the swap plugins params
func (o *SwapPluginsParams) WithPtype(ptype string) *SwapPluginsParams {
	o.SetPtype(ptype)
	return o
}

// SetPtype adds the ptype to the swap plugins params
func (o *SwapPluginsParams) SetPtype(ptype string) {
	o.Ptype = ptype
}

// WithPversion adds the pversion to the swap plugins params
func (o *SwapPluginsParams) WithPversion(pversion int64) *SwapPluginsParams {
	o.SetPversion(pversion)
	return o
}

// SetPversion adds the pversion to the swap plugins params
func (o *SwapPluginsParams) SetPversion(pversion int64) {
	o.Pversion = pversion
}

// WriteToRequest writes these params to a swagger request
func (o *SwapPluginsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.PluginData != nil {

		if o.PluginData != nil {

			// form file param pluginData
			if err := r.SetFileParam("pluginData", o.PluginData); err != nil {
				return err
			}

		}

	}

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