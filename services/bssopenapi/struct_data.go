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

// Data is a nested struct in bssopenapi response
type Data struct {
	BusinessType           string                                 `json:"BusinessType" xml:"BusinessType"`
	TradePrice             float64                                `json:"TradePrice" xml:"TradePrice"`
	HostId                 string                                 `json:"HostId" xml:"HostId"`
	PageSize               int                                    `json:"PageSize" xml:"PageSize"`
	Numerator              int                                    `json:"Numerator" xml:"Numerator"`
	OriginalPrice          float64                                `json:"OriginalPrice" xml:"OriginalPrice"`
	AvailableAmount        string                                 `json:"AvailableAmount" xml:"AvailableAmount"`
	Amount                 string                                 `json:"Amount" xml:"Amount"`
	MybankCreditAmount     string                                 `json:"MybankCreditAmount" xml:"MybankCreditAmount"`
	OrderId                string                                 `json:"OrderId" xml:"OrderId"`
	CreditAmount           string                                 `json:"CreditAmount" xml:"CreditAmount"`
	ThresholdType          int                                    `json:"ThresholdType" xml:"ThresholdType"`
	InstanceId             string                                 `json:"InstanceId" xml:"InstanceId"`
	ItemCode               string                                 `json:"ItemCode" xml:"ItemCode"`
	ThresholdAmount        string                                 `json:"ThresholdAmount" xml:"ThresholdAmount"`
	BillingCycle           string                                 `json:"BillingCycle" xml:"BillingCycle"`
	Uid                    int                                    `json:"Uid" xml:"Uid"`
	TotalCount             int                                    `json:"TotalCount" xml:"TotalCount"`
	OutstandingAmount      float64                                `json:"OutstandingAmount" xml:"OutstandingAmount"`
	InvalidTimeStamp       int                                    `json:"InvalidTimeStamp" xml:"InvalidTimeStamp"`
	Bid                    string                                 `json:"Bid" xml:"Bid"`
	Currency               string                                 `json:"Currency" xml:"Currency"`
	Quantity               int                                    `json:"Quantity" xml:"Quantity"`
	AvailableCashAmount    string                                 `json:"AvailableCashAmount" xml:"AvailableCashAmount"`
	PageNum                int                                    `json:"PageNum" xml:"PageNum"`
	DiscountPrice          float64                                `json:"DiscountPrice" xml:"DiscountPrice"`
	EffectTimeStamp        int                                    `json:"EffectTimeStamp" xml:"EffectTimeStamp"`
	PrimaryAccount         string                                 `json:"PrimaryAccount" xml:"PrimaryAccount"`
	TotalOutstandingAmount float64                                `json:"TotalOutstandingAmount" xml:"TotalOutstandingAmount"`
	Status                 string                                 `json:"Status" xml:"Status"`
	UserId                 int                                    `json:"UserId" xml:"UserId"`
	Denominator            int                                    `json:"Denominator" xml:"Denominator"`
	NewInvoiceAmount       float64                                `json:"NewInvoiceAmount" xml:"NewInvoiceAmount"`
	ModuleList             ModuleList                             `json:"ModuleList" xml:"ModuleList"`
	Redeem                 Redeem                                 `json:"Redeem" xml:"Redeem"`
	Promotions             Promotions                             `json:"Promotions" xml:"Promotions"`
	ResourcePackages       ResourcePackages                       `json:"ResourcePackages" xml:"ResourcePackages"`
	ModuleDetails          ModuleDetailsInGetSubscriptionPrice    `json:"ModuleDetails" xml:"ModuleDetails"`
	Modules                ModulesInQueryPriceList                `json:"Modules" xml:"Modules"`
	AttributeList          AttributeList                          `json:"AttributeList" xml:"AttributeList"`
	PromotionDetails       PromotionDetailsInGetSubscriptionPrice `json:"PromotionDetails" xml:"PromotionDetails"`
	Items                  ItemsInQueryMonthlyBill                `json:"Items" xml:"Items"`
}