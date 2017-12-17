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

func (client *Client) ListUsersOfSkillGroup(request *ListUsersOfSkillGroupRequest) (response *ListUsersOfSkillGroupResponse, err error) {
	response = CreateListUsersOfSkillGroupResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListUsersOfSkillGroupWithChan(request *ListUsersOfSkillGroupRequest) (<-chan *ListUsersOfSkillGroupResponse, <-chan error) {
	responseChan := make(chan *ListUsersOfSkillGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUsersOfSkillGroup(request)
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

func (client *Client) ListUsersOfSkillGroupWithCallback(request *ListUsersOfSkillGroupRequest, callback func(response *ListUsersOfSkillGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUsersOfSkillGroupResponse
		var err error
		defer close(result)
		response, err = client.ListUsersOfSkillGroup(request)
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

type ListUsersOfSkillGroupRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	SkillGroupId string `position:"Query" name:"SkillGroupId"`
	PageSize     string `position:"Query" name:"PageSize"`
	PageNumber   string `position:"Query" name:"PageNumber"`
	AccessKeyId  string `position:"Query" name:"AccessKeyId"`
}

type ListUsersOfSkillGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId"`
	Success        bool   `json:"Success"`
	Code           string `json:"Code"`
	Message        string `json:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode"`
	Users          struct {
		TotalCount int `json:"TotalCount"`
		PageNumber int `json:"PageNumber"`
		PageSize   int `json:"PageSize"`
		List       []struct {
			UserId     string `json:"UserId"`
			RamId      string `json:"RamId"`
			InstanceId string `json:"InstanceId"`
			Detail     struct {
				LoginName   string `json:"LoginName"`
				DisplayName string `json:"DisplayName"`
				Phone       string `json:"Phone"`
				Email       string `json:"Email"`
				Department  string `json:"Department"`
			} `json:"Detail"`
			Roles []struct {
				RoleId          string `json:"RoleId"`
				InstanceId      string `json:"InstanceId"`
				RoleName        string `json:"RoleName"`
				RoleDescription string `json:"RoleDescription"`
				UserCount       int    `json:"UserCount"`
				Privileges      []struct {
					PrivilegeId          string `json:"PrivilegeId"`
					PrivilegeName        string `json:"PrivilegeName"`
					PrivilegeDescription string `json:"PrivilegeDescription"`
				} `json:"Privileges"`
			} `json:"Roles"`
			SkillLevels []struct {
				SkillLevelId string `json:"SkillLevelId"`
				Level        int    `json:"Level"`
				Skill        struct {
					SkillGroupId          string `json:"SkillGroupId"`
					InstanceId            string `json:"InstanceId"`
					SkillGroupName        string `json:"SkillGroupName"`
					SkillGroupDescription string `json:"SkillGroupDescription"`
				} `json:"Skill"`
			} `json:"SkillLevels"`
		} `json:"List"`
	} `json:"Users"`
}

func CreateListUsersOfSkillGroupRequest() (request *ListUsersOfSkillGroupRequest) {
	request = &ListUsersOfSkillGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListUsersOfSkillGroup", "", "")
	return
}

func CreateListUsersOfSkillGroupResponse() (response *ListUsersOfSkillGroupResponse) {
	response = &ListUsersOfSkillGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
