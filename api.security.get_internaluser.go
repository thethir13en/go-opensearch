// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package opensearch

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"net/http"
	"strings"
)

func newGetInternalUserFunc(t esapi.Transport) GetInternalUser {
	return func(o ...func(*GetInternalUserRequest)) (*esapi.Response, error) {
		var r = GetInternalUserRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// GetInternalUser - Retrieves internal users in the native realm.
//
// See full documentation at https://docs-beta.opensearch.org/security-plugin/access-control/api/#users.
//
type GetInternalUser func(o ...func(*GetInternalUserRequest)) (*esapi.Response, error)

// GetInternalUserRequest configures the Security Get InternalUser API request.
//
type GetInternalUserRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r GetInternalUserRequest) Do(ctx context.Context, transport esapi.Transport) (*esapi.Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_opendistro") + 1 + len("_security") + 1 + len("api") + 1 + len("internalusers") + 1 + len(r.Name))
	path.WriteString("/_opendistro")
	path.WriteString("/")
	path.WriteString("_security")
	path.WriteString("/")
	path.WriteString("api")
	path.WriteString("/")
	path.WriteString("internalusers")
	path.WriteString("/")
	path.WriteString(r.Name)

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := http.NewRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := esapi.Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f GetInternalUser) WithContext(v context.Context) func(*GetInternalUserRequest) {
	return func(r *GetInternalUserRequest) {
		r.ctx = v
	}
}

// WithName - sets user name.
//
func (f GetInternalUser) WithName(v string) func(*GetInternalUserRequest) {
	return func(r *GetInternalUserRequest) {
		r.Name = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f GetInternalUser) WithPretty() func(*GetInternalUserRequest) {
	return func(r *GetInternalUserRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f GetInternalUser) WithHuman() func(*GetInternalUserRequest) {
	return func(r *GetInternalUserRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f GetInternalUser) WithErrorTrace() func(*GetInternalUserRequest) {
	return func(r *GetInternalUserRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f GetInternalUser) WithFilterPath(v ...string) func(*GetInternalUserRequest) {
	return func(r *GetInternalUserRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f GetInternalUser) WithHeader(h map[string]string) func(*GetInternalUserRequest) {
	return func(r *GetInternalUserRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f GetInternalUser) WithOpaqueID(s string) func(*GetInternalUserRequest) {
	return func(r *GetInternalUserRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
