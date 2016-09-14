package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/mrfusion/models"
)

/*PatchSourcesIDUsersUserIDNoContent Users's configuration was changed

swagger:response patchSourcesIdUsersUserIdNoContent
*/
type PatchSourcesIDUsersUserIDNoContent struct {
}

// NewPatchSourcesIDUsersUserIDNoContent creates PatchSourcesIDUsersUserIDNoContent with default headers values
func NewPatchSourcesIDUsersUserIDNoContent() *PatchSourcesIDUsersUserIDNoContent {
	return &PatchSourcesIDUsersUserIDNoContent{}
}

// WriteResponse to the client
func (o *PatchSourcesIDUsersUserIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*PatchSourcesIDUsersUserIDNotFound Happens when trying to access a non-existent user.

swagger:response patchSourcesIdUsersUserIdNotFound
*/
type PatchSourcesIDUsersUserIDNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchSourcesIDUsersUserIDNotFound creates PatchSourcesIDUsersUserIDNotFound with default headers values
func NewPatchSourcesIDUsersUserIDNotFound() *PatchSourcesIDUsersUserIDNotFound {
	return &PatchSourcesIDUsersUserIDNotFound{}
}

// WithPayload adds the payload to the patch sources Id users user Id not found response
func (o *PatchSourcesIDUsersUserIDNotFound) WithPayload(payload *models.Error) *PatchSourcesIDUsersUserIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch sources Id users user Id not found response
func (o *PatchSourcesIDUsersUserIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchSourcesIDUsersUserIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PatchSourcesIDUsersUserIDDefault A processing or an unexpected error.

swagger:response patchSourcesIdUsersUserIdDefault
*/
type PatchSourcesIDUsersUserIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchSourcesIDUsersUserIDDefault creates PatchSourcesIDUsersUserIDDefault with default headers values
func NewPatchSourcesIDUsersUserIDDefault(code int) *PatchSourcesIDUsersUserIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PatchSourcesIDUsersUserIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the patch sources ID users user ID default response
func (o *PatchSourcesIDUsersUserIDDefault) WithStatusCode(code int) *PatchSourcesIDUsersUserIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the patch sources ID users user ID default response
func (o *PatchSourcesIDUsersUserIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the patch sources ID users user ID default response
func (o *PatchSourcesIDUsersUserIDDefault) WithPayload(payload *models.Error) *PatchSourcesIDUsersUserIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch sources ID users user ID default response
func (o *PatchSourcesIDUsersUserIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchSourcesIDUsersUserIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
