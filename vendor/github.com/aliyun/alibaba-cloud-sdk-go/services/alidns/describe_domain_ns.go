package alidns

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

// DescribeDomainNs invokes the alidns.DescribeDomainNs API synchronously
// api document: https://help.aliyun.com/api/alidns/describedomainns.html
func (client *Client) DescribeDomainNs(request *DescribeDomainNsRequest) (response *DescribeDomainNsResponse, err error) {
	response = CreateDescribeDomainNsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainNsWithChan invokes the alidns.DescribeDomainNs API asynchronously
// api document: https://help.aliyun.com/api/alidns/describedomainns.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainNsWithChan(request *DescribeDomainNsRequest) (<-chan *DescribeDomainNsResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainNsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainNs(request)
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

// DescribeDomainNsWithCallback invokes the alidns.DescribeDomainNs API asynchronously
// api document: https://help.aliyun.com/api/alidns/describedomainns.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainNsWithCallback(request *DescribeDomainNsRequest, callback func(response *DescribeDomainNsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainNsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainNs(request)
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

// DescribeDomainNsRequest is the request struct for api DescribeDomainNs
type DescribeDomainNsRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeDomainNsResponse is the response struct for api DescribeDomainNs
type DescribeDomainNsResponse struct {
	*responses.BaseResponse
	RequestId        string                       `json:"RequestId" xml:"RequestId"`
	AllAliDns        bool                         `json:"AllAliDns" xml:"AllAliDns"`
	IncludeAliDns    bool                         `json:"IncludeAliDns" xml:"IncludeAliDns"`
	DnsServers       DnsServersInDescribeDomainNs `json:"DnsServers" xml:"DnsServers"`
	ExpectDnsServers ExpectDnsServers             `json:"ExpectDnsServers" xml:"ExpectDnsServers"`
}

// CreateDescribeDomainNsRequest creates a request to invoke DescribeDomainNs API
func CreateDescribeDomainNsRequest() (request *DescribeDomainNsRequest) {
	request = &DescribeDomainNsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomainNs", "alidns", "openAPI")
	return
}

// CreateDescribeDomainNsResponse creates a response to parse from DescribeDomainNs response
func CreateDescribeDomainNsResponse() (response *DescribeDomainNsResponse) {
	response = &DescribeDomainNsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
