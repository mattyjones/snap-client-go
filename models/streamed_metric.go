package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// StreamedMetric StreamedMetric defines a streamed metric.
// swagger:model StreamedMetric
type StreamedMetric struct {

	// data
	Data interface{} `json:"data,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`

	// timestamp
	Timestamp interface{} `json:"timestamp,omitempty"`
}

// Validate validates this streamed metric
func (m *StreamedMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}