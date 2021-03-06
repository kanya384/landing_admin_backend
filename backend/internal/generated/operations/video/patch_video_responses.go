// Code generated by go-swagger; DO NOT EDIT.

package video

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// PatchVideoOKCode is the HTTP code returned for type PatchVideoOK
const PatchVideoOKCode int = 200

/*PatchVideoOK Request success

swagger:response patchVideoOK
*/
type PatchVideoOK struct {

	/*
	  In: Body
	*/
	Payload *models.Video `json:"body,omitempty"`
}

// NewPatchVideoOK creates PatchVideoOK with default headers values
func NewPatchVideoOK() *PatchVideoOK {

	return &PatchVideoOK{}
}

// WithPayload adds the payload to the patch video o k response
func (o *PatchVideoOK) WithPayload(payload *models.Video) *PatchVideoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch video o k response
func (o *PatchVideoOK) SetPayload(payload *models.Video) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchVideoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchVideoBadRequestCode is the HTTP code returned for type PatchVideoBadRequest
const PatchVideoBadRequestCode int = 400

/*PatchVideoBadRequest Bad request

swagger:response patchVideoBadRequest
*/
type PatchVideoBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewPatchVideoBadRequest creates PatchVideoBadRequest with default headers values
func NewPatchVideoBadRequest() *PatchVideoBadRequest {

	return &PatchVideoBadRequest{}
}

// WithPayload adds the payload to the patch video bad request response
func (o *PatchVideoBadRequest) WithPayload(payload *models.ResultResponse) *PatchVideoBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch video bad request response
func (o *PatchVideoBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchVideoBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchVideoInternalServerErrorCode is the HTTP code returned for type PatchVideoInternalServerError
const PatchVideoInternalServerErrorCode int = 500

/*PatchVideoInternalServerError Internal Server Error

swagger:response patchVideoInternalServerError
*/
type PatchVideoInternalServerError struct {
}

// NewPatchVideoInternalServerError creates PatchVideoInternalServerError with default headers values
func NewPatchVideoInternalServerError() *PatchVideoInternalServerError {

	return &PatchVideoInternalServerError{}
}

// WriteResponse to the client
func (o *PatchVideoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
