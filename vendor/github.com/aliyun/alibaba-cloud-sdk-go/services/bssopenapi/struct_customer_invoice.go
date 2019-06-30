package bssopenapi

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

// CustomerInvoice is a nested struct in bssopenapi response
type CustomerInvoice struct {
	Id                      int    `json:"Id" xml:"Id"`
	UserId                  int    `json:"UserId" xml:"UserId"`
	UserNick                string `json:"UserNick" xml:"UserNick"`
	InvoiceTitle            string `json:"InvoiceTitle" xml:"InvoiceTitle"`
	CustomerType            int    `json:"CustomerType" xml:"CustomerType"`
	TaxpayerType            int    `json:"TaxpayerType" xml:"TaxpayerType"`
	Bank                    string `json:"Bank" xml:"Bank"`
	BankNo                  string `json:"BankNo" xml:"BankNo"`
	OperatingLicenseAddress string `json:"OperatingLicenseAddress" xml:"OperatingLicenseAddress"`
	OperatingLicensePhone   string `json:"OperatingLicensePhone" xml:"OperatingLicensePhone"`
	RegisterNo              string `json:"RegisterNo" xml:"RegisterNo"`
	StartCycle              int    `json:"StartCycle" xml:"StartCycle"`
	Status                  int    `json:"Status" xml:"Status"`
	GmtCreate               string `json:"GmtCreate" xml:"GmtCreate"`
	TaxationLicense         string `json:"TaxationLicense" xml:"TaxationLicense"`
	AdjustType              int    `json:"AdjustType" xml:"AdjustType"`
	EndCycle                int    `json:"EndCycle" xml:"EndCycle"`
	TitleChangeInstructions string `json:"TitleChangeInstructions" xml:"TitleChangeInstructions"`
	IssueType               int    `json:"IssueType" xml:"IssueType"`
	Type                    int    `json:"Type" xml:"Type"`
	DefaultRemark           string `json:"DefaultRemark" xml:"DefaultRemark"`
}
