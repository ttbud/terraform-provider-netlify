// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// ConfigureDNSForSiteReader is a Reader for the ConfigureDNSForSite structure.
type ConfigureDNSForSiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigureDNSForSiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigureDNSForSiteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConfigureDNSForSiteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigureDNSForSiteOK creates a ConfigureDNSForSiteOK with default headers values
func NewConfigureDNSForSiteOK() *ConfigureDNSForSiteOK {
	return &ConfigureDNSForSiteOK{}
}

/*ConfigureDNSForSiteOK handles this case with default header values.

OK
*/
type ConfigureDNSForSiteOK struct {
	Payload []*models.DNSZone
}

func (o *ConfigureDNSForSiteOK) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/dns][%d] configureDnsForSiteOK  %+v", 200, o.Payload)
}

func (o *ConfigureDNSForSiteOK) GetPayload() []*models.DNSZone {
	return o.Payload
}

func (o *ConfigureDNSForSiteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigureDNSForSiteDefault creates a ConfigureDNSForSiteDefault with default headers values
func NewConfigureDNSForSiteDefault(code int) *ConfigureDNSForSiteDefault {
	return &ConfigureDNSForSiteDefault{
		_statusCode: code,
	}
}

/*ConfigureDNSForSiteDefault handles this case with default header values.

error
*/
type ConfigureDNSForSiteDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the configure DNS for site default response
func (o *ConfigureDNSForSiteDefault) Code() int {
	return o._statusCode
}

func (o *ConfigureDNSForSiteDefault) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/dns][%d] configureDNSForSite default  %+v", o._statusCode, o.Payload)
}

func (o *ConfigureDNSForSiteDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ConfigureDNSForSiteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
