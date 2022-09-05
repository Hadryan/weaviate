//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewBackupsRestoreStatusParams creates a new BackupsRestoreStatusParams object
// with the default values initialized.
func NewBackupsRestoreStatusParams() *BackupsRestoreStatusParams {
	var ()
	return &BackupsRestoreStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBackupsRestoreStatusParamsWithTimeout creates a new BackupsRestoreStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBackupsRestoreStatusParamsWithTimeout(timeout time.Duration) *BackupsRestoreStatusParams {
	var ()
	return &BackupsRestoreStatusParams{

		timeout: timeout,
	}
}

// NewBackupsRestoreStatusParamsWithContext creates a new BackupsRestoreStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewBackupsRestoreStatusParamsWithContext(ctx context.Context) *BackupsRestoreStatusParams {
	var ()
	return &BackupsRestoreStatusParams{

		Context: ctx,
	}
}

// NewBackupsRestoreStatusParamsWithHTTPClient creates a new BackupsRestoreStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBackupsRestoreStatusParamsWithHTTPClient(client *http.Client) *BackupsRestoreStatusParams {
	var ()
	return &BackupsRestoreStatusParams{
		HTTPClient: client,
	}
}

/*
BackupsRestoreStatusParams contains all the parameters to send to the API endpoint
for the backups restore status operation typically these are written to a http.Request
*/
type BackupsRestoreStatusParams struct {

	/*ID
	  The ID of a backup. Must be URL-safe and work as a filesystem path, only lowercase, numbers, underscore, minus characters allowed.

	*/
	ID string
	/*StorageName
	  Storage name e.g. filesystem, gcs, s3.

	*/
	StorageName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the backups restore status params
func (o *BackupsRestoreStatusParams) WithTimeout(timeout time.Duration) *BackupsRestoreStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backups restore status params
func (o *BackupsRestoreStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backups restore status params
func (o *BackupsRestoreStatusParams) WithContext(ctx context.Context) *BackupsRestoreStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backups restore status params
func (o *BackupsRestoreStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backups restore status params
func (o *BackupsRestoreStatusParams) WithHTTPClient(client *http.Client) *BackupsRestoreStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backups restore status params
func (o *BackupsRestoreStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the backups restore status params
func (o *BackupsRestoreStatusParams) WithID(id string) *BackupsRestoreStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the backups restore status params
func (o *BackupsRestoreStatusParams) SetID(id string) {
	o.ID = id
}

// WithStorageName adds the storageName to the backups restore status params
func (o *BackupsRestoreStatusParams) WithStorageName(storageName string) *BackupsRestoreStatusParams {
	o.SetStorageName(storageName)
	return o
}

// SetStorageName adds the storageName to the backups restore status params
func (o *BackupsRestoreStatusParams) SetStorageName(storageName string) {
	o.StorageName = storageName
}

// WriteToRequest writes these params to a swagger request
func (o *BackupsRestoreStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param storageName
	if err := r.SetPathParam("storageName", o.StorageName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}