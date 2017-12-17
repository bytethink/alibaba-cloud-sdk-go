package ccc

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

func (client *Client) ListRoles(request *ListRolesRequest) (response *ListRolesResponse, err error) {
	response = CreateListRolesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListRolesWithChan(request *ListRolesRequest) (<-chan *ListRolesResponse, <-chan error) {
	responseChan := make(chan *ListRolesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRoles(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) ListRolesWithCallback(request *ListRolesRequest, callback func(response *ListRolesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRolesResponse
		var err error
		defer close(result)
		response, err = client.ListRoles(request)
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

type ListRolesRequest struct {
	*requests.RpcRequest
	InstanceId  string `position:"Query" name:"InstanceId"`
	AccessKeyId string `position:"Query" name:"AccessKeyId"`
}

type ListRolesResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId"`
	Success        bool   `json:"Success"`
	Code           string `json:"Code"`
	Message        string `json:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode"`
	Roles          []struct {
		RoleId          string `json:"RoleId"`
		InstanceId      string `json:"InstanceId"`
		RoleName        string `json:"RoleName"`
		RoleDescription string `json:"RoleDescription"`
	} `json:"Roles"`
}

func CreateListRolesRequest() (request *ListRolesRequest) {
	request = &ListRolesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListRoles", "", "")
	return
}

func CreateListRolesResponse() (response *ListRolesResponse) {
	response = &ListRolesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
