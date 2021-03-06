// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PatchActionsOKCode is the HTTP code returned for type PatchActionsOK
const PatchActionsOKCode int = 200

/*PatchActionsOK постер успешно обновлен

swagger:response patchActionsOK
*/
type PatchActionsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Action `json:"body,omitempty"`
}

// NewPatchActionsOK creates PatchActionsOK with default headers values
func NewPatchActionsOK() *PatchActionsOK {

	return &PatchActionsOK{}
}

// WithPayload adds the payload to the patch actions o k response
func (o *PatchActionsOK) WithPayload(payload *models.Action) *PatchActionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch actions o k response
func (o *PatchActionsOK) SetPayload(payload *models.Action) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchActionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchActionsBadRequestCode is the HTTP code returned for type PatchActionsBadRequest
const PatchActionsBadRequestCode int = 400

/*PatchActionsBadRequest Bad request

swagger:response patchActionsBadRequest
*/
type PatchActionsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPatchActionsBadRequest creates PatchActionsBadRequest with default headers values
func NewPatchActionsBadRequest() *PatchActionsBadRequest {

	return &PatchActionsBadRequest{}
}

// WithPayload adds the payload to the patch actions bad request response
func (o *PatchActionsBadRequest) WithPayload(payload *models.ResultResponse) *PatchActionsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch actions bad request response
func (o *PatchActionsBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchActionsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchActionsInternalServerErrorCode is the HTTP code returned for type PatchActionsInternalServerError
const PatchActionsInternalServerErrorCode int = 500

/*PatchActionsInternalServerError Internal Server Error

swagger:response patchActionsInternalServerError
*/
type PatchActionsInternalServerError struct {
}

// NewPatchActionsInternalServerError creates PatchActionsInternalServerError with default headers values
func NewPatchActionsInternalServerError() *PatchActionsInternalServerError {

	return &PatchActionsInternalServerError{}
}

// WriteResponse to the client
func (o *PatchActionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
