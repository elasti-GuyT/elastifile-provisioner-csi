// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gapic-generator. DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	iotpb "google.golang.org/genproto/googleapis/cloud/iot/v1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// DeviceManagerCallOptions contains the retry settings for each method of DeviceManagerClient.
type DeviceManagerCallOptions struct {
	CreateDeviceRegistry      []gax.CallOption
	GetDeviceRegistry         []gax.CallOption
	UpdateDeviceRegistry      []gax.CallOption
	DeleteDeviceRegistry      []gax.CallOption
	ListDeviceRegistries      []gax.CallOption
	CreateDevice              []gax.CallOption
	GetDevice                 []gax.CallOption
	UpdateDevice              []gax.CallOption
	DeleteDevice              []gax.CallOption
	ListDevices               []gax.CallOption
	ModifyCloudToDeviceConfig []gax.CallOption
	ListDeviceConfigVersions  []gax.CallOption
	ListDeviceStates          []gax.CallOption
	SetIamPolicy              []gax.CallOption
	GetIamPolicy              []gax.CallOption
	TestIamPermissions        []gax.CallOption
	SendCommandToDevice       []gax.CallOption
	BindDeviceToGateway       []gax.CallOption
	UnbindDeviceFromGateway   []gax.CallOption
}

func defaultDeviceManagerClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("cloudiot.googleapis.com:443"),
		option.WithScopes(DefaultAuthScopes()...),
	}
}

func defaultDeviceManagerCallOptions() *DeviceManagerCallOptions {
	retry := map[[2]string][]gax.CallOption{
		{"default", "idempotent"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
		{"rate_limited_aware", "rate_limited_aware"}: {
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.ResourceExhausted,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.3,
				})
			}),
		},
	}
	return &DeviceManagerCallOptions{
		CreateDeviceRegistry:      retry[[2]string{"default", "non_idempotent"}],
		GetDeviceRegistry:         retry[[2]string{"default", "idempotent"}],
		UpdateDeviceRegistry:      retry[[2]string{"default", "non_idempotent"}],
		DeleteDeviceRegistry:      retry[[2]string{"default", "idempotent"}],
		ListDeviceRegistries:      retry[[2]string{"default", "idempotent"}],
		CreateDevice:              retry[[2]string{"default", "non_idempotent"}],
		GetDevice:                 retry[[2]string{"default", "idempotent"}],
		UpdateDevice:              retry[[2]string{"default", "non_idempotent"}],
		DeleteDevice:              retry[[2]string{"default", "idempotent"}],
		ListDevices:               retry[[2]string{"default", "idempotent"}],
		ModifyCloudToDeviceConfig: retry[[2]string{"rate_limited_aware", "rate_limited_aware"}],
		ListDeviceConfigVersions:  retry[[2]string{"default", "idempotent"}],
		ListDeviceStates:          retry[[2]string{"default", "idempotent"}],
		SetIamPolicy:              retry[[2]string{"default", "non_idempotent"}],
		GetIamPolicy:              retry[[2]string{"default", "non_idempotent"}],
		TestIamPermissions:        retry[[2]string{"default", "non_idempotent"}],
		SendCommandToDevice:       retry[[2]string{"rate_limited_aware", "rate_limited_aware"}],
		BindDeviceToGateway:       retry[[2]string{"default", "non_idempotent"}],
		UnbindDeviceFromGateway:   retry[[2]string{"default", "non_idempotent"}],
	}
}

// DeviceManagerClient is a client for interacting with Cloud IoT API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type DeviceManagerClient struct {
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	deviceManagerClient iotpb.DeviceManagerClient

	// The call options for this service.
	CallOptions *DeviceManagerCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewDeviceManagerClient creates a new device manager client.
//
// Internet of Things (IoT) service. Securely connect and manage IoT devices.
func NewDeviceManagerClient(ctx context.Context, opts ...option.ClientOption) (*DeviceManagerClient, error) {
	conn, err := transport.DialGRPC(ctx, append(defaultDeviceManagerClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &DeviceManagerClient{
		conn:        conn,
		CallOptions: defaultDeviceManagerCallOptions(),

		deviceManagerClient: iotpb.NewDeviceManagerClient(conn),
	}
	c.setGoogleClientInfo()
	return c, nil
}

// Connection returns the client's connection to the API service.
func (c *DeviceManagerClient) Connection() *grpc.ClientConn {
	return c.conn
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *DeviceManagerClient) Close() error {
	return c.conn.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *DeviceManagerClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateDeviceRegistry creates a device registry that contains devices.
func (c *DeviceManagerClient) CreateDeviceRegistry(ctx context.Context, req *iotpb.CreateDeviceRegistryRequest, opts ...gax.CallOption) (*iotpb.DeviceRegistry, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateDeviceRegistry[0:len(c.CallOptions.CreateDeviceRegistry):len(c.CallOptions.CreateDeviceRegistry)], opts...)
	var resp *iotpb.DeviceRegistry
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.CreateDeviceRegistry(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDeviceRegistry gets a device registry configuration.
func (c *DeviceManagerClient) GetDeviceRegistry(ctx context.Context, req *iotpb.GetDeviceRegistryRequest, opts ...gax.CallOption) (*iotpb.DeviceRegistry, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetDeviceRegistry[0:len(c.CallOptions.GetDeviceRegistry):len(c.CallOptions.GetDeviceRegistry)], opts...)
	var resp *iotpb.DeviceRegistry
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.GetDeviceRegistry(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateDeviceRegistry updates a device registry configuration.
func (c *DeviceManagerClient) UpdateDeviceRegistry(ctx context.Context, req *iotpb.UpdateDeviceRegistryRequest, opts ...gax.CallOption) (*iotpb.DeviceRegistry, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "device_registry.name", req.GetDeviceRegistry().GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateDeviceRegistry[0:len(c.CallOptions.UpdateDeviceRegistry):len(c.CallOptions.UpdateDeviceRegistry)], opts...)
	var resp *iotpb.DeviceRegistry
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.UpdateDeviceRegistry(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteDeviceRegistry deletes a device registry configuration.
func (c *DeviceManagerClient) DeleteDeviceRegistry(ctx context.Context, req *iotpb.DeleteDeviceRegistryRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteDeviceRegistry[0:len(c.CallOptions.DeleteDeviceRegistry):len(c.CallOptions.DeleteDeviceRegistry)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.deviceManagerClient.DeleteDeviceRegistry(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ListDeviceRegistries lists device registries.
func (c *DeviceManagerClient) ListDeviceRegistries(ctx context.Context, req *iotpb.ListDeviceRegistriesRequest, opts ...gax.CallOption) *DeviceRegistryIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListDeviceRegistries[0:len(c.CallOptions.ListDeviceRegistries):len(c.CallOptions.ListDeviceRegistries)], opts...)
	it := &DeviceRegistryIterator{}
	req = proto.Clone(req).(*iotpb.ListDeviceRegistriesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*iotpb.DeviceRegistry, string, error) {
		var resp *iotpb.ListDeviceRegistriesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.deviceManagerClient.ListDeviceRegistries(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.DeviceRegistries, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// CreateDevice creates a device in a device registry.
func (c *DeviceManagerClient) CreateDevice(ctx context.Context, req *iotpb.CreateDeviceRequest, opts ...gax.CallOption) (*iotpb.Device, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateDevice[0:len(c.CallOptions.CreateDevice):len(c.CallOptions.CreateDevice)], opts...)
	var resp *iotpb.Device
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.CreateDevice(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDevice gets details about a device.
func (c *DeviceManagerClient) GetDevice(ctx context.Context, req *iotpb.GetDeviceRequest, opts ...gax.CallOption) (*iotpb.Device, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetDevice[0:len(c.CallOptions.GetDevice):len(c.CallOptions.GetDevice)], opts...)
	var resp *iotpb.Device
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.GetDevice(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateDevice updates a device.
func (c *DeviceManagerClient) UpdateDevice(ctx context.Context, req *iotpb.UpdateDeviceRequest, opts ...gax.CallOption) (*iotpb.Device, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "device.name", req.GetDevice().GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateDevice[0:len(c.CallOptions.UpdateDevice):len(c.CallOptions.UpdateDevice)], opts...)
	var resp *iotpb.Device
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.UpdateDevice(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteDevice deletes a device.
func (c *DeviceManagerClient) DeleteDevice(ctx context.Context, req *iotpb.DeleteDeviceRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteDevice[0:len(c.CallOptions.DeleteDevice):len(c.CallOptions.DeleteDevice)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.deviceManagerClient.DeleteDevice(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ListDevices list devices in a device registry.
func (c *DeviceManagerClient) ListDevices(ctx context.Context, req *iotpb.ListDevicesRequest, opts ...gax.CallOption) *DeviceIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListDevices[0:len(c.CallOptions.ListDevices):len(c.CallOptions.ListDevices)], opts...)
	it := &DeviceIterator{}
	req = proto.Clone(req).(*iotpb.ListDevicesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*iotpb.Device, string, error) {
		var resp *iotpb.ListDevicesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.deviceManagerClient.ListDevices(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}
		return resp.Devices, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// ModifyCloudToDeviceConfig modifies the configuration for the device, which is eventually sent from
// the Cloud IoT Core servers. Returns the modified configuration version and
// its metadata.
func (c *DeviceManagerClient) ModifyCloudToDeviceConfig(ctx context.Context, req *iotpb.ModifyCloudToDeviceConfigRequest, opts ...gax.CallOption) (*iotpb.DeviceConfig, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ModifyCloudToDeviceConfig[0:len(c.CallOptions.ModifyCloudToDeviceConfig):len(c.CallOptions.ModifyCloudToDeviceConfig)], opts...)
	var resp *iotpb.DeviceConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.ModifyCloudToDeviceConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListDeviceConfigVersions lists the last few versions of the device configuration in descending
// order (i.e.: newest first).
func (c *DeviceManagerClient) ListDeviceConfigVersions(ctx context.Context, req *iotpb.ListDeviceConfigVersionsRequest, opts ...gax.CallOption) (*iotpb.ListDeviceConfigVersionsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListDeviceConfigVersions[0:len(c.CallOptions.ListDeviceConfigVersions):len(c.CallOptions.ListDeviceConfigVersions)], opts...)
	var resp *iotpb.ListDeviceConfigVersionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.ListDeviceConfigVersions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListDeviceStates lists the last few versions of the device state in descending order (i.e.:
// newest first).
func (c *DeviceManagerClient) ListDeviceStates(ctx context.Context, req *iotpb.ListDeviceStatesRequest, opts ...gax.CallOption) (*iotpb.ListDeviceStatesResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListDeviceStates[0:len(c.CallOptions.ListDeviceStates):len(c.CallOptions.ListDeviceStates)], opts...)
	var resp *iotpb.ListDeviceStatesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.ListDeviceStates(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any
// existing policy.
func (c *DeviceManagerClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", req.GetResource()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SetIamPolicy[0:len(c.CallOptions.SetIamPolicy):len(c.CallOptions.SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetIamPolicy gets the access control policy for a resource.
// Returns an empty policy if the resource exists and does not have a policy
// set.
func (c *DeviceManagerClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", req.GetResource()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetIamPolicy[0:len(c.CallOptions.GetIamPolicy):len(c.CallOptions.GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
// If the resource does not exist, this will return an empty set of
// permissions, not a NOT_FOUND error.
func (c *DeviceManagerClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", req.GetResource()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.TestIamPermissions[0:len(c.CallOptions.TestIamPermissions):len(c.CallOptions.TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendCommandToDevice sends a command to the specified device. In order for a device to be able
// to receive commands, it must:
// 1) be connected to Cloud IoT Core using the MQTT protocol, and
// 2) be subscribed to the group of MQTT topics specified by
//    /devices/{device-id}/commands/#. This subscription will receive commands
//    at the top-level topic /devices/{device-id}/commands as well as commands
//    for subfolders, like /devices/{device-id}/commands/subfolder.
//    Note that subscribing to specific subfolders is not supported.
// If the command could not be delivered to the device, this method will
// return an error; in particular, if the device is not subscribed, this
// method will return FAILED_PRECONDITION. Otherwise, this method will
// return OK. If the subscription is QoS 1, at least once delivery will be
// guaranteed; for QoS 0, no acknowledgment will be expected from the device.
func (c *DeviceManagerClient) SendCommandToDevice(ctx context.Context, req *iotpb.SendCommandToDeviceRequest, opts ...gax.CallOption) (*iotpb.SendCommandToDeviceResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", req.GetName()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SendCommandToDevice[0:len(c.CallOptions.SendCommandToDevice):len(c.CallOptions.SendCommandToDevice)], opts...)
	var resp *iotpb.SendCommandToDeviceResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.SendCommandToDevice(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BindDeviceToGateway associates the device with the gateway.
func (c *DeviceManagerClient) BindDeviceToGateway(ctx context.Context, req *iotpb.BindDeviceToGatewayRequest, opts ...gax.CallOption) (*iotpb.BindDeviceToGatewayResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.BindDeviceToGateway[0:len(c.CallOptions.BindDeviceToGateway):len(c.CallOptions.BindDeviceToGateway)], opts...)
	var resp *iotpb.BindDeviceToGatewayResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.BindDeviceToGateway(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UnbindDeviceFromGateway deletes the association between the device and the gateway.
func (c *DeviceManagerClient) UnbindDeviceFromGateway(ctx context.Context, req *iotpb.UnbindDeviceFromGatewayRequest, opts ...gax.CallOption) (*iotpb.UnbindDeviceFromGatewayResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", req.GetParent()))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UnbindDeviceFromGateway[0:len(c.CallOptions.UnbindDeviceFromGateway):len(c.CallOptions.UnbindDeviceFromGateway)], opts...)
	var resp *iotpb.UnbindDeviceFromGatewayResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.deviceManagerClient.UnbindDeviceFromGateway(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeviceIterator manages a stream of *iotpb.Device.
type DeviceIterator struct {
	items    []*iotpb.Device
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*iotpb.Device, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *DeviceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *DeviceIterator) Next() (*iotpb.Device, error) {
	var item *iotpb.Device
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *DeviceIterator) bufLen() int {
	return len(it.items)
}

func (it *DeviceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// DeviceRegistryIterator manages a stream of *iotpb.DeviceRegistry.
type DeviceRegistryIterator struct {
	items    []*iotpb.DeviceRegistry
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*iotpb.DeviceRegistry, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *DeviceRegistryIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *DeviceRegistryIterator) Next() (*iotpb.DeviceRegistry, error) {
	var item *iotpb.DeviceRegistry
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *DeviceRegistryIterator) bufLen() int {
	return len(it.items)
}

func (it *DeviceRegistryIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
