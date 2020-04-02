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

package dts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ConfigureSynchronizationJobAlert invokes the dts.ConfigureSynchronizationJobAlert API synchronously
// api document: https://help.aliyun.com/api/dts/configuresynchronizationjobalert.html
func (client *Client) ConfigureSynchronizationJobAlert(request *ConfigureSynchronizationJobAlertRequest) (response *ConfigureSynchronizationJobAlertResponse, err error) {
	response = CreateConfigureSynchronizationJobAlertResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigureSynchronizationJobAlertWithChan invokes the dts.ConfigureSynchronizationJobAlert API asynchronously
// api document: https://help.aliyun.com/api/dts/configuresynchronizationjobalert.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigureSynchronizationJobAlertWithChan(request *ConfigureSynchronizationJobAlertRequest) (<-chan *ConfigureSynchronizationJobAlertResponse, <-chan error) {
	responseChan := make(chan *ConfigureSynchronizationJobAlertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigureSynchronizationJobAlert(request)
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

// ConfigureSynchronizationJobAlertWithCallback invokes the dts.ConfigureSynchronizationJobAlert API asynchronously
// api document: https://help.aliyun.com/api/dts/configuresynchronizationjobalert.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigureSynchronizationJobAlertWithCallback(request *ConfigureSynchronizationJobAlertRequest, callback func(response *ConfigureSynchronizationJobAlertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigureSynchronizationJobAlertResponse
		var err error
		defer close(result)
		response, err = client.ConfigureSynchronizationJobAlert(request)
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

// ConfigureSynchronizationJobAlertRequest is the request struct for api ConfigureSynchronizationJobAlert
type ConfigureSynchronizationJobAlertRequest struct {
	*requests.RpcRequest
	SynchronizationJobId     string `position:"Query" name:"SynchronizationJobId"`
	SynchronizationDirection string `position:"Query" name:"SynchronizationDirection"`
	DelayAlertStatus         string `position:"Query" name:"DelayAlertStatus"`
	DelayAlertPhone          string `position:"Query" name:"DelayAlertPhone"`
	ErrorAlertStatus         string `position:"Query" name:"ErrorAlertStatus"`
	ErrorAlertPhone          string `position:"Query" name:"ErrorAlertPhone"`
	DelayOverSeconds         string `position:"Query" name:"DelayOverSeconds"`
	OwnerId                  string `position:"Query" name:"OwnerId"`
}

// ConfigureSynchronizationJobAlertResponse is the response struct for api ConfigureSynchronizationJobAlert
type ConfigureSynchronizationJobAlertResponse struct {
	*responses.BaseResponse
	Success    string `json:"Success" xml:"Success"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigureSynchronizationJobAlertRequest creates a request to invoke ConfigureSynchronizationJobAlert API
func CreateConfigureSynchronizationJobAlertRequest() (request *ConfigureSynchronizationJobAlertRequest) {
	request = &ConfigureSynchronizationJobAlertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "ConfigureSynchronizationJobAlert", "dts", "openAPI")
	return
}

// CreateConfigureSynchronizationJobAlertResponse creates a response to parse from ConfigureSynchronizationJobAlert response
func CreateConfigureSynchronizationJobAlertResponse() (response *ConfigureSynchronizationJobAlertResponse) {
	response = &ConfigureSynchronizationJobAlertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}