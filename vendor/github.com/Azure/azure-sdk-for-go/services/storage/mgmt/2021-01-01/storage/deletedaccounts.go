package storage

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DeletedAccountsClient is the the Azure Storage Management API.
type DeletedAccountsClient struct {
	BaseClient
}

// NewDeletedAccountsClient creates an instance of the DeletedAccountsClient client.
func NewDeletedAccountsClient(subscriptionID string) DeletedAccountsClient {
	return NewDeletedAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDeletedAccountsClientWithBaseURI creates an instance of the DeletedAccountsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDeletedAccountsClientWithBaseURI(baseURI string, subscriptionID string) DeletedAccountsClient {
	return DeletedAccountsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get properties of specified deleted account resource.
// Parameters:
// deletedAccountName - name of the deleted storage account.
// location - the location of the deleted storage account.
func (client DeletedAccountsClient) Get(ctx context.Context, deletedAccountName string, location string) (result DeletedAccount, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedAccountsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: deletedAccountName,
			Constraints: []validation.Constraint{{Target: "deletedAccountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "deletedAccountName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storage.DeletedAccountsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, deletedAccountName, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.DeletedAccountsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.DeletedAccountsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.DeletedAccountsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DeletedAccountsClient) GetPreparer(ctx context.Context, deletedAccountName string, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deletedAccountName": autorest.Encode("path", deletedAccountName),
		"location":           autorest.Encode("path", location),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/locations/{location}/deletedAccounts/{deletedAccountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeletedAccountsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeletedAccountsClient) GetResponder(resp *http.Response) (result DeletedAccount, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists deleted accounts under the subscription.
func (client DeletedAccountsClient) List(ctx context.Context) (result DeletedAccountListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedAccountsClient.List")
		defer func() {
			sc := -1
			if result.dalr.Response.Response != nil {
				sc = result.dalr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storage.DeletedAccountsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.DeletedAccountsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dalr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storage.DeletedAccountsClient", "List", resp, "Failure sending request")
		return
	}

	result.dalr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.DeletedAccountsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.dalr.hasNextLink() && result.dalr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DeletedAccountsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/deletedAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeletedAccountsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeletedAccountsClient) ListResponder(resp *http.Response) (result DeletedAccountListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DeletedAccountsClient) listNextResults(ctx context.Context, lastResults DeletedAccountListResult) (result DeletedAccountListResult, err error) {
	req, err := lastResults.deletedAccountListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storage.DeletedAccountsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storage.DeletedAccountsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.DeletedAccountsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeletedAccountsClient) ListComplete(ctx context.Context) (result DeletedAccountListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedAccountsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
