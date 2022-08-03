//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package commercialmarketplace

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SubscriptionOperationsClient contains the methods for the SubscriptionOperations group.
// Don't use this type directly, use NewSubscriptionOperationsClient() instead.
type SubscriptionOperationsClient struct {
	host string
	pl runtime.Pipeline
}

// NewSubscriptionOperationsClient creates a new instance of SubscriptionOperationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSubscriptionOperationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SubscriptionOperationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SubscriptionOperationsClient{
		host: ep,
pl: pl,
	}
	return client, nil
}

// GetOperationStatus - Enables the publisher to track the status of the specified triggered async operation (such as Subscribe,
// Unsubscribe, ChangePlan, or ChangeQuantity).
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-08-31
// options - SubscriptionOperationsClientGetOperationStatusOptions contains the optional parameters for the SubscriptionOperationsClient.GetOperationStatus
// method.
func (client *SubscriptionOperationsClient) GetOperationStatus(ctx context.Context, subscriptionID string, operationID string, options *SubscriptionOperationsClientGetOperationStatusOptions) (SubscriptionOperationsClientGetOperationStatusResponse, error) {
	req, err := client.getOperationStatusCreateRequest(ctx, subscriptionID, operationID, options)
	if err != nil {
		return SubscriptionOperationsClientGetOperationStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionOperationsClientGetOperationStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionOperationsClientGetOperationStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOperationStatusHandleResponse(resp)
}

// getOperationStatusCreateRequest creates the GetOperationStatus request.
func (client *SubscriptionOperationsClient) getOperationStatusCreateRequest(ctx context.Context, subscriptionID string, operationID string, options *SubscriptionOperationsClientGetOperationStatusOptions) (*policy.Request, error) {
	urlPath := "/saas/subscriptions/{subscriptionId}/operations/{operationId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-requestid"] = []string{*options.RequestID}
	}
	if options != nil && options.CorrelationID != nil {
		req.Raw().Header["x-ms-correlationid"] = []string{*options.CorrelationID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOperationStatusHandleResponse handles the GetOperationStatus response.
func (client *SubscriptionOperationsClient) getOperationStatusHandleResponse(resp *http.Response) (SubscriptionOperationsClientGetOperationStatusResponse, error) {
	result := SubscriptionOperationsClientGetOperationStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Operation); err != nil {
		return SubscriptionOperationsClientGetOperationStatusResponse{}, err
	}
	return result, nil
}

// ListOperations - Lists the outstanding operations for the current publisher.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-08-31
// options - SubscriptionOperationsClientListOperationsOptions contains the optional parameters for the SubscriptionOperationsClient.ListOperations
// method.
func (client *SubscriptionOperationsClient) ListOperations(ctx context.Context, subscriptionID string, options *SubscriptionOperationsClientListOperationsOptions) (SubscriptionOperationsClientListOperationsResponse, error) {
	req, err := client.listOperationsCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return SubscriptionOperationsClientListOperationsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionOperationsClientListOperationsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionOperationsClientListOperationsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listOperationsHandleResponse(resp)
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *SubscriptionOperationsClient) listOperationsCreateRequest(ctx context.Context, subscriptionID string, options *SubscriptionOperationsClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/saas/subscriptions/{subscriptionId}/operations"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-requestid"] = []string{*options.RequestID}
	}
	if options != nil && options.CorrelationID != nil {
		req.Raw().Header["x-ms-correlationid"] = []string{*options.CorrelationID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *SubscriptionOperationsClient) listOperationsHandleResponse(resp *http.Response) (SubscriptionOperationsClientListOperationsResponse, error) {
	result := SubscriptionOperationsClientListOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationList); err != nil {
		return SubscriptionOperationsClientListOperationsResponse{}, err
	}
	return result, nil
}

// UpdateOperationStatus - Update the status of an operation to indicate success or failure with the provided values.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-08-31
// options - SubscriptionOperationsClientUpdateOperationStatusOptions contains the optional parameters for the SubscriptionOperationsClient.UpdateOperationStatus
// method.
func (client *SubscriptionOperationsClient) UpdateOperationStatus(ctx context.Context, subscriptionID string, operationID string, body UpdateOperation, options *SubscriptionOperationsClientUpdateOperationStatusOptions) (SubscriptionOperationsClientUpdateOperationStatusResponse, error) {
	req, err := client.updateOperationStatusCreateRequest(ctx, subscriptionID, operationID, body, options)
	if err != nil {
		return SubscriptionOperationsClientUpdateOperationStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionOperationsClientUpdateOperationStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionOperationsClientUpdateOperationStatusResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionOperationsClientUpdateOperationStatusResponse{}, nil
}

// updateOperationStatusCreateRequest creates the UpdateOperationStatus request.
func (client *SubscriptionOperationsClient) updateOperationStatusCreateRequest(ctx context.Context, subscriptionID string, operationID string, body UpdateOperation, options *SubscriptionOperationsClientUpdateOperationStatusOptions) (*policy.Request, error) {
	urlPath := "/saas/subscriptions/{subscriptionId}/operations/{operationId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.RequestID != nil {
		req.Raw().Header["x-ms-requestid"] = []string{*options.RequestID}
	}
	if options != nil && options.CorrelationID != nil {
		req.Raw().Header["x-ms-correlationid"] = []string{*options.CorrelationID}
	}
	return req, runtime.MarshalAsJSON(req, body)
}

