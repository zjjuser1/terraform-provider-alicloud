package actiontrail

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

// StartLogging invokes the actiontrail.StartLogging API synchronously
// api document: https://help.aliyun.com/api/actiontrail/startlogging.html
func (client *Client) StartLogging(request *StartLoggingRequest) (response *StartLoggingResponse, err error) {
	response = CreateStartLoggingResponse()
	err = client.DoAction(request, response)
	return
}

// StartLoggingWithChan invokes the actiontrail.StartLogging API asynchronously
// api document: https://help.aliyun.com/api/actiontrail/startlogging.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartLoggingWithChan(request *StartLoggingRequest) (<-chan *StartLoggingResponse, <-chan error) {
	responseChan := make(chan *StartLoggingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartLogging(request)
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

// StartLoggingWithCallback invokes the actiontrail.StartLogging API asynchronously
// api document: https://help.aliyun.com/api/actiontrail/startlogging.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartLoggingWithCallback(request *StartLoggingRequest, callback func(response *StartLoggingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartLoggingResponse
		var err error
		defer close(result)
		response, err = client.StartLogging(request)
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

// StartLoggingRequest is the request struct for api StartLogging
type StartLoggingRequest struct {
	*requests.RpcRequest
	Name string `position:"Query" name:"Name"`
}

// StartLoggingResponse is the response struct for api StartLogging
type StartLoggingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Param     string `json:"Param" xml:"Param"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateStartLoggingRequest creates a request to invoke StartLogging API
func CreateStartLoggingRequest() (request *StartLoggingRequest) {
	request = &StartLoggingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Actiontrail", "2017-12-04", "StartLogging", "actiontrail", "openAPI")
	return
}

// CreateStartLoggingResponse creates a response to parse from StartLogging response
func CreateStartLoggingResponse() (response *StartLoggingResponse) {
	response = &StartLoggingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
