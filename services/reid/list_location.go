package reid

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListLocation invokes the reid.ListLocation API synchronously
// api document: https://help.aliyun.com/api/reid/listlocation.html
func (client *Client) ListLocation(request *ListLocationRequest) (response *ListLocationResponse, err error) {
	response = CreateListLocationResponse()
	err = client.DoAction(request, response)
	return
}

// ListLocationWithChan invokes the reid.ListLocation API asynchronously
// api document: https://help.aliyun.com/api/reid/listlocation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListLocationWithChan(request *ListLocationRequest) (<-chan *ListLocationResponse, <-chan error) {
	responseChan := make(chan *ListLocationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLocation(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListLocationWithCallback invokes the reid.ListLocation API asynchronously
// api document: https://help.aliyun.com/api/reid/listlocation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListLocationWithCallback(request *ListLocationRequest, callback func(response *ListLocationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLocationResponse
		var err error
		defer close(result)
		response, err = client.ListLocation(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListLocationRequest is the request struct for api ListLocation
type ListLocationRequest struct {
	*requests.RpcRequest
	StoreId requests.Integer `position:"Body" name:"StoreId"`
}

// ListLocationResponse is the response struct for api ListLocation
type ListLocationResponse struct {
	*responses.BaseResponse
	ErrorCode         string            `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage      string            `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	Success           bool              `json:"Success" xml:"Success"`
	LocationInfoItems LocationInfoItems `json:"LocationInfoItems" xml:"LocationInfoItems"`
}

// CreateListLocationRequest creates a request to invoke ListLocation API
func CreateListLocationRequest() (request *ListLocationRequest) {
	request = &ListLocationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "ListLocation", "1.1.7", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListLocationResponse creates a response to parse from ListLocation response
func CreateListLocationResponse() (response *ListLocationResponse) {
	response = &ListLocationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
