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

func (client *Client) ListSkillGroups(request *ListSkillGroupsRequest) (response *ListSkillGroupsResponse, err error) {
	response = CreateListSkillGroupsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListSkillGroupsWithChan(request *ListSkillGroupsRequest) (<-chan *ListSkillGroupsResponse, <-chan error) {
	responseChan := make(chan *ListSkillGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSkillGroups(request)
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

func (client *Client) ListSkillGroupsWithCallback(request *ListSkillGroupsRequest, callback func(response *ListSkillGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSkillGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListSkillGroups(request)
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

type ListSkillGroupsRequest struct {
	*requests.RpcRequest
	InstanceId  string `position:"Query" name:"InstanceId"`
	AccessKeyId string `position:"Query" name:"AccessKeyId"`
}

type ListSkillGroupsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId"`
	Success        bool   `json:"Success"`
	Code           string `json:"Code"`
	Message        string `json:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode"`
	SkillGroups    []struct {
		SkillGroupId          string `json:"SkillGroupId"`
		InstanceId            string `json:"InstanceId"`
		SkillGroupName        string `json:"SkillGroupName"`
		AccSkillGroupName     string `json:"AccSkillGroupName"`
		AccQueueName          string `json:"AccQueueName"`
		SkillGroupDescription string `json:"SkillGroupDescription"`
		UserCount             int    `json:"UserCount"`
		OutboundPhoneNumbers  []struct {
			PhoneNumberId          string `json:"PhoneNumberId"`
			InstanceId             string `json:"InstanceId"`
			Number                 string `json:"Number"`
			PhoneNumberDescription string `json:"PhoneNumberDescription"`
			TestOnly               bool   `json:"TestOnly"`
			RemainingTime          int    `json:"RemainingTime"`
			AllowOutbound          bool   `json:"AllowOutbound"`
			Usage                  string `json:"Usage"`
			Trunks                 int    `json:"Trunks"`
		} `json:"OutboundPhoneNumbers"`
	} `json:"SkillGroups"`
}

func CreateListSkillGroupsRequest() (request *ListSkillGroupsRequest) {
	request = &ListSkillGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListSkillGroups", "", "")
	return
}

func CreateListSkillGroupsResponse() (response *ListSkillGroupsResponse) {
	response = &ListSkillGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
