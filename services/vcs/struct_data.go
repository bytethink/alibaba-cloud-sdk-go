package vcs

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

// Data is a nested struct in vcs response
type Data struct {
	PhoneNo      string             `json:"PhoneNo" xml:"PhoneNo"`
	Name         string             `json:"Name" xml:"Name"`
	TotalPage    int                `json:"TotalPage" xml:"TotalPage"`
	Attachment   string             `json:"Attachment" xml:"Attachment"`
	CatalogId    int                `json:"CatalogId" xml:"CatalogId"`
	UserId       int                `json:"UserId" xml:"UserId"`
	FaceUrl      string             `json:"FaceUrl" xml:"FaceUrl"`
	IsvSubId     string             `json:"IsvSubId" xml:"IsvSubId"`
	Age          string             `json:"Age" xml:"Age"`
	PageSize     int                `json:"PageSize" xml:"PageSize"`
	OssPath      string             `json:"OssPath" xml:"OssPath"`
	BizId        string             `json:"BizId" xml:"BizId"`
	TaskId       string             `json:"TaskId" xml:"TaskId"`
	PageNo       int                `json:"PageNo" xml:"PageNo"`
	TotalCount   int                `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int                `json:"PageNumber" xml:"PageNumber"`
	Description  string             `json:"Description" xml:"Description"`
	LiveAddress  string             `json:"LiveAddress" xml:"LiveAddress"`
	DataSourceId string             `json:"DataSourceId" xml:"DataSourceId"`
	SceneType    string             `json:"SceneType" xml:"SceneType"`
	UserGroupId  int                `json:"UserGroupId" xml:"UserGroupId"`
	PersonId     string             `json:"PersonId" xml:"PersonId"`
	PageNum      int                `json:"PageNum" xml:"PageNum"`
	KafkaTopic   string             `json:"KafkaTopic" xml:"KafkaTopic"`
	StructList   string             `json:"StructList" xml:"StructList"`
	Address      string             `json:"Address" xml:"Address"`
	FaceImageUrl string             `json:"FaceImageUrl" xml:"FaceImageUrl"`
	IdNumber     string             `json:"IdNumber" xml:"IdNumber"`
	Gender       string             `json:"Gender" xml:"Gender"`
	QualityScore string             `json:"QualityScore" xml:"QualityScore"`
	UserName     string             `json:"UserName" xml:"UserName"`
	PicUrl       string             `json:"PicUrl" xml:"PicUrl"`
	PlateNo      string             `json:"PlateNo" xml:"PlateNo"`
	ProfileId    int                `json:"ProfileId" xml:"ProfileId"`
	Attributes   Attributes         `json:"Attributes" xml:"Attributes"`
	ResultObject []ResultObjectItem `json:"ResultObject" xml:"ResultObject"`
	Records      []Record           `json:"Records" xml:"Records"`
	TagList      []TagListItem      `json:"TagList" xml:"TagList"`
	FaceList     []Face             `json:"FaceList" xml:"FaceList"`
	BodyList     []Body             `json:"BodyList" xml:"BodyList"`
}
