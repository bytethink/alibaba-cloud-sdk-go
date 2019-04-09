package ivision

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

// DescribeProjects invokes the ivision.DescribeProjects API synchronously
// api document: https://help.aliyun.com/api/ivision/describeprojects.html
func (client *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
	response = CreateDescribeProjectsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProjectsWithChan invokes the ivision.DescribeProjects API asynchronously
// api document: https://help.aliyun.com/api/ivision/describeprojects.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeProjectsWithChan(request *DescribeProjectsRequest) (<-chan *DescribeProjectsResponse, <-chan error) {
	responseChan := make(chan *DescribeProjectsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProjects(request)
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

// DescribeProjectsWithCallback invokes the ivision.DescribeProjects API asynchronously
// api document: https://help.aliyun.com/api/ivision/describeprojects.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeProjectsWithCallback(request *DescribeProjectsRequest, callback func(response *DescribeProjectsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProjectsResponse
		var err error
		defer close(result)
		response, err = client.DescribeProjects(request)
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

// DescribeProjectsRequest is the request struct for api DescribeProjects
type DescribeProjectsRequest struct {
	*requests.RpcRequest
	NextPageToken string           `position:"Query" name:"NextPageToken"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	ShowLog       string           `position:"Query" name:"ShowLog"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	ProjectIds    string           `position:"Query" name:"ProjectIds"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeProjectsResponse is the response struct for api DescribeProjects
type DescribeProjectsResponse struct {
	*responses.BaseResponse
	RequestId     string   `json:"RequestId" xml:"RequestId"`
	TotalNum      int      `json:"TotalNum" xml:"TotalNum"`
	CurrentPage   int      `json:"CurrentPage" xml:"CurrentPage"`
	PageSize      int      `json:"PageSize" xml:"PageSize"`
	NextPageToken string   `json:"NextPageToken" xml:"NextPageToken"`
	Projects      Projects `json:"Projects" xml:"Projects"`
}

// CreateDescribeProjectsRequest creates a request to invoke DescribeProjects API
func CreateDescribeProjectsRequest() (request *DescribeProjectsRequest) {
	request = &DescribeProjectsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivision", "2019-03-08", "DescribeProjects", "ivision", "openAPI")
	return
}

// CreateDescribeProjectsResponse creates a response to parse from DescribeProjects response
func CreateDescribeProjectsResponse() (response *DescribeProjectsResponse) {
	response = &DescribeProjectsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}