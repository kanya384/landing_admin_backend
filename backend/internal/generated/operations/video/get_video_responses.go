// Code generated by go-swagger; DO NOT EDIT.

package video

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"landing_admin_backend/models"
)

// GetVideoOKCode is the HTTP code returned for type GetVideoOK
const GetVideoOKCode int = 200

/*GetVideoOK returns video list

swagger:response getVideoOK
*/
type GetVideoOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Video `json:"body,omitempty"`
}

// NewGetVideoOK creates GetVideoOK with default headers values
func NewGetVideoOK() *GetVideoOK {

	return &GetVideoOK{}
}

// WithPayload adds the payload to the get video o k response
func (o *GetVideoOK) WithPayload(payload []*models.Video) *GetVideoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get video o k response
func (o *GetVideoOK) SetPayload(payload []*models.Video) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVideoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Video, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetVideoBadRequestCode is the HTTP code returned for type GetVideoBadRequest
const GetVideoBadRequestCode int = 400

/*GetVideoBadRequest Bad request

swagger:response getVideoBadRequest
*/
type GetVideoBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResultResponse `json:"body,omitempty"`
}

// NewGetVideoBadRequest creates GetVideoBadRequest with default headers values
func NewGetVideoBadRequest() *GetVideoBadRequest {

	return &GetVideoBadRequest{}
}

// WithPayload adds the payload to the get video bad request response
func (o *GetVideoBadRequest) WithPayload(payload *models.ResultResponse) *GetVideoBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get video bad request response
func (o *GetVideoBadRequest) SetPayload(payload *models.ResultResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVideoBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVideoInternalServerErrorCode is the HTTP code returned for type GetVideoInternalServerError
const GetVideoInternalServerErrorCode int = 500

/*GetVideoInternalServerError Internal Server Error

swagger:response getVideoInternalServerError
*/
type GetVideoInternalServerError struct {
}

// NewGetVideoInternalServerError creates GetVideoInternalServerError with default headers values
func NewGetVideoInternalServerError() *GetVideoInternalServerError {

	return &GetVideoInternalServerError{}
}

// WriteResponse to the client
func (o *GetVideoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}