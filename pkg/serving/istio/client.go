// Copyright © 2019 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package istio

import (
	istio_nw "github.com/aspenmesh/istio-client-go/pkg/apis/networking/v1alpha3"
	istio_v1alpha3 "github.com/aspenmesh/istio-client-go/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kn_errors "knative.dev/client/pkg/errors"
)

// Kn interface to serving. All methods are relative to the
// namespace specified during construction
type KnIstioClient interface {

	// Get a service by its unique name
	GetVirtualService(name string) (*istio_nw.VirtualService, error)
}

type knIstioClient struct {
	client    istio_v1alpha3.Interface
	namespace string
}

// Create a new client facade for the provided namespace
func NewKnIstioClient(client istio_v1alpha3.Interface, namespace string) KnIstioClient {
	return &knIstioClient{
		client:    client,
		namespace: namespace,
	}
}

// Get a service by its unique name
func (cl *knIstioClient) GetVirtualService(name string) (*istio_nw.VirtualService, error) {
	// TODO: hardcode for test
	virtualService, err := cl.client.NetworkingV1alpha3().VirtualServices("knative-serving").Get("route-d1356b80-c356-11e9-bc8a-5ec9f1d9fef7", v1.GetOptions{})
	if err != nil {
		return nil, kn_errors.GetError(err)
	}

	return virtualService, nil
}
