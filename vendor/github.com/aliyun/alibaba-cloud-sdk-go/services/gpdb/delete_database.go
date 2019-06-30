package gpdb

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

// DeleteDatabase invokes the gpdb.DeleteDatabase API synchronously
// api document: https://help.aliyun.com/api/gpdb/deletedatabase.html
func (client *Client) DeleteDatabase(request *DeleteDatabaseRequest) (response *DeleteDatabaseResponse, err error) {
	response = CreateDeleteDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDatabaseWithChan invokes the gpdb.DeleteDatabase API asynchronously
// api document: https://help.aliyun.com/api/gpdb/deletedatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDatabaseWithChan(request *DeleteDatabaseRequest) (<-chan *DeleteDatabaseResponse, <-chan error) {
	responseChan := make(chan *DeleteDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDatabase(request)
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

// DeleteDatabaseWithCallback invokes the gpdb.DeleteDatabase API asynchronously
// api document: https://help.aliyun.com/api/gpdb/deletedatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDatabaseWithCallback(request *DeleteDatabaseRequest, callback func(response *DeleteDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDatabaseResponse
		var err error
		defer close(result)
		response, err = client.DeleteDatabase(request)
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

// DeleteDatabaseRequest is the request struct for api DeleteDatabase
type DeleteDatabaseRequest struct {
	*requests.RpcRequest
	DBName       string `position:"Query" name:"DBName"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

// DeleteDatabaseResponse is the response struct for api DeleteDatabase
type DeleteDatabaseResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDatabaseRequest creates a request to invoke DeleteDatabase API
func CreateDeleteDatabaseRequest() (request *DeleteDatabaseRequest) {
	request = &DeleteDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DeleteDatabase", "gpdb", "openAPI")
	return
}

// CreateDeleteDatabaseResponse creates a response to parse from DeleteDatabase response
func CreateDeleteDatabaseResponse() (response *DeleteDatabaseResponse) {
	response = &DeleteDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
