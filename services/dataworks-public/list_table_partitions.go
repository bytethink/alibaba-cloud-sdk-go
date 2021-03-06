package dataworks_public

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

// ListTablePartitions invokes the dataworks_public.ListTablePartitions API synchronously
func (client *Client) ListTablePartitions(request *ListTablePartitionsRequest) (response *ListTablePartitionsResponse, err error) {
	response = CreateListTablePartitionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTablePartitionsWithChan invokes the dataworks_public.ListTablePartitions API asynchronously
func (client *Client) ListTablePartitionsWithChan(request *ListTablePartitionsRequest) (<-chan *ListTablePartitionsResponse, <-chan error) {
	responseChan := make(chan *ListTablePartitionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTablePartitions(request)
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

// ListTablePartitionsWithCallback invokes the dataworks_public.ListTablePartitions API asynchronously
func (client *Client) ListTablePartitionsWithCallback(request *ListTablePartitionsRequest, callback func(response *ListTablePartitionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTablePartitionsResponse
		var err error
		defer close(result)
		response, err = client.ListTablePartitions(request)
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

// ListTablePartitionsRequest is the request struct for api ListTablePartitions
type ListTablePartitionsRequest struct {
	*requests.RpcRequest
	DatabaseName string           `position:"Query" name:"DatabaseName"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	ClusterId    string           `position:"Query" name:"ClusterId"`
	Sort         string           `position:"Query" name:"Sort"`
	TableName    string           `position:"Query" name:"TableName"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	Order        string           `position:"Query" name:"Order"`
}

// ListTablePartitionsResponse is the response struct for api ListTablePartitions
type ListTablePartitionsResponse struct {
	*responses.BaseResponse
	ErrorCode    string                    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string                    `json:"RequestId" xml:"RequestId"`
	Data         DataInListTablePartitions `json:"Data" xml:"Data"`
}

// CreateListTablePartitionsRequest creates a request to invoke ListTablePartitions API
func CreateListTablePartitionsRequest() (request *ListTablePartitionsRequest) {
	request = &ListTablePartitionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2018-06-01", "ListTablePartitions", "", "")
	request.Method = requests.POST
	return
}

// CreateListTablePartitionsResponse creates a response to parse from ListTablePartitions response
func CreateListTablePartitionsResponse() (response *ListTablePartitionsResponse) {
	response = &ListTablePartitionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
