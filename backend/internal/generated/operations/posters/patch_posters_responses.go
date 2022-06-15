// Code generated by go-swagger; DO NOT EDIT.

package posters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PatchPostersOKCode is the HTTP code returned for type PatchPostersOK
const PatchPostersOKCode int = 200

/*PatchPostersOK постер успешно обновлен

swagger:response patchPostersOK
*/
type PatchPostersOK struct {

	/*
	  In: Body
	*/
	Payload *models.Poster `json:"body,omitempty"`
}

// NewPatchPostersOK creates PatchPostersOK with default headers values
func NewPatchPostersOK() *PatchPostersOK {

	return &PatchPostersOK{}
}

// WithPayload adds the payload to the patch posters o k response
func (o *PatchPostersOK) WithPayload(payload *models.Poster) *PatchPostersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch posters o k response
func (o *PatchPostersOK) SetPayload(payload *models.Poster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchPostersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchPostersBadRequestCode is the HTTP code returned for type PatchPostersBadRequest
const PatchPostersBadRequestCode int = 400

/*PatchPostersBadRequest Bad request

swagger:response patchPostersBadRequest
*/
type PatchPostersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPatchPostersBadRequest creates PatchPostersBadRequest with default headers values
func NewPatchPostersBadRequest() *PatchPostersBadRequest {

	return &PatchPostersBadRequest{}
}

// WithPayload adds the payload to the patch posters bad request response
func (o *PatchPostersBadRequest) WithPayload(payload *models.ResultResponse) *PatchPostersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch posters bad request response
func (o *PatchPostersBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchPostersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchPostersInternalServerErrorCode is the HTTP code returned for type PatchPostersInternalServerError
const PatchPostersInternalServerErrorCode int = 500

/*PatchPostersInternalServerError Internal Server Error

swagger:response patchPostersInternalServerError
*/
type PatchPostersInternalServerError struct {
}

// NewPatchPostersInternalServerError creates PatchPostersInternalServerError with default headers values
func NewPatchPostersInternalServerError() *PatchPostersInternalServerError {

	return &PatchPostersInternalServerError{}
}

// WriteResponse to the client
func (o *PatchPostersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
