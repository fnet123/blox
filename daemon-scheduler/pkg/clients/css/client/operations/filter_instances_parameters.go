// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFilterInstancesParams creates a new FilterInstancesParams object
// with the default values initialized.
func NewFilterInstancesParams() *FilterInstancesParams {
	var ()
	return &FilterInstancesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFilterInstancesParamsWithTimeout creates a new FilterInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFilterInstancesParamsWithTimeout(timeout time.Duration) *FilterInstancesParams {
	var ()
	return &FilterInstancesParams{

		timeout: timeout,
	}
}

// NewFilterInstancesParamsWithContext creates a new FilterInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewFilterInstancesParamsWithContext(ctx context.Context) *FilterInstancesParams {
	var ()
	return &FilterInstancesParams{

		Context: ctx,
	}
}

/*FilterInstancesParams contains all the parameters to send to the API endpoint
for the filter instances operation typically these are written to a http.Request
*/
type FilterInstancesParams struct {

	/*Cluster
	  Cluster name or ARN to filter instances by

	*/
	Cluster string
	/*Status
	  Status to filter instances by

	*/
	Status string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the filter instances params
func (o *FilterInstancesParams) WithTimeout(timeout time.Duration) *FilterInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the filter instances params
func (o *FilterInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the filter instances params
func (o *FilterInstancesParams) WithContext(ctx context.Context) *FilterInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the filter instances params
func (o *FilterInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCluster adds the cluster to the filter instances params
func (o *FilterInstancesParams) WithCluster(cluster string) *FilterInstancesParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the filter instances params
func (o *FilterInstancesParams) SetCluster(cluster string) {
	o.Cluster = cluster
}

// WithStatus adds the status to the filter instances params
func (o *FilterInstancesParams) WithStatus(status string) *FilterInstancesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the filter instances params
func (o *FilterInstancesParams) SetStatus(status string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *FilterInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// query param cluster
	qrCluster := o.Cluster
	qCluster := qrCluster
	if qCluster != "" {
		if err := r.SetQueryParam("cluster", qCluster); err != nil {
			return err
		}
	}

	// query param status
	qrStatus := o.Status
	qStatus := qrStatus
	if qStatus != "" {
		if err := r.SetQueryParam("status", qStatus); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
