/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.25.0-2b3f843a-20210115-164628
 */

// Package adminrestv1 : Operations and models for the AdminrestV1 service
package adminrestv1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/eventstreams-go-sdk/pkg/common"
	"github.com/IBM/go-sdk-core/v4/core"
	"net/http"
	"reflect"
	"time"
)

// AdminrestV1 : The administration REST API for IBM Event Streams on Cloud.
//
// Version: 1.1.0
type AdminrestV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://adminrest.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "adminrest"

// AdminrestV1Options : Service options
type AdminrestV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewAdminrestV1UsingExternalConfig : constructs an instance of AdminrestV1 with passed in options and external configuration.
func NewAdminrestV1UsingExternalConfig(options *AdminrestV1Options) (adminrest *AdminrestV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	adminrest, err = NewAdminrestV1(options)
	if err != nil {
		return
	}

	err = adminrest.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = adminrest.Service.SetServiceURL(options.URL)
	}
	return
}

// NewAdminrestV1 : constructs an instance of AdminrestV1 with passed in options.
func NewAdminrestV1(options *AdminrestV1Options) (service *AdminrestV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &AdminrestV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "adminrest" suitable for processing requests.
func (adminrest *AdminrestV1) Clone() *AdminrestV1 {
	if core.IsNil(adminrest) {
		return nil
	}
	clone := *adminrest
	clone.Service = adminrest.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (adminrest *AdminrestV1) SetServiceURL(url string) error {
	return adminrest.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (adminrest *AdminrestV1) GetServiceURL() string {
	return adminrest.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (adminrest *AdminrestV1) SetDefaultHeaders(headers http.Header) {
	adminrest.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (adminrest *AdminrestV1) SetEnableGzipCompression(enableGzip bool) {
	adminrest.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (adminrest *AdminrestV1) GetEnableGzipCompression() bool {
	return adminrest.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (adminrest *AdminrestV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	adminrest.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (adminrest *AdminrestV1) DisableRetries() {
	adminrest.Service.DisableRetries()
}

// CreateTopic : Create a new topic
// Create a new topic.
func (adminrest *AdminrestV1) CreateTopic(createTopicOptions *CreateTopicOptions) (response *core.DetailedResponse, err error) {
	return adminrest.CreateTopicWithContext(context.Background(), createTopicOptions)
}

// CreateTopicWithContext is an alternate form of the CreateTopic method which supports a Context parameter
func (adminrest *AdminrestV1) CreateTopicWithContext(ctx context.Context, createTopicOptions *CreateTopicOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createTopicOptions, "createTopicOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createTopicOptions, "createTopicOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/topics`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createTopicOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "CreateTopic")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createTopicOptions.Name != nil {
		body["name"] = createTopicOptions.Name
	}
	if createTopicOptions.Partitions != nil {
		body["partitions"] = createTopicOptions.Partitions
	}
	if createTopicOptions.PartitionCount != nil {
		body["partition_count"] = createTopicOptions.PartitionCount
	}
	if createTopicOptions.Configs != nil {
		body["configs"] = createTopicOptions.Configs
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = adminrest.Service.Request(request, nil)

	return
}

// ListTopics : Get a list of topics
// Returns a list containing information about all of the Kafka topics that are defined for an instance of the Event
// Streams service. If there are currently no topics defined then an empty list is returned.
func (adminrest *AdminrestV1) ListTopics(listTopicsOptions *ListTopicsOptions) (result []TopicDetail, response *core.DetailedResponse, err error) {
	return adminrest.ListTopicsWithContext(context.Background(), listTopicsOptions)
}

// ListTopicsWithContext is an alternate form of the ListTopics method which supports a Context parameter
func (adminrest *AdminrestV1) ListTopicsWithContext(ctx context.Context, listTopicsOptions *ListTopicsOptions) (result []TopicDetail, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listTopicsOptions, "listTopicsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/topics`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTopicsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "ListTopics")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listTopicsOptions.TopicFilter != nil {
		builder.AddQuery("topic_filter", fmt.Sprint(*listTopicsOptions.TopicFilter))
	}
	if listTopicsOptions.PerPage != nil {
		builder.AddQuery("per_page", fmt.Sprint(*listTopicsOptions.PerPage))
	}
	if listTopicsOptions.Page != nil {
		builder.AddQuery("page", fmt.Sprint(*listTopicsOptions.Page))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse []json.RawMessage
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTopicDetail)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetTopic : Get detailed information on a topic
// Get detailed information on a topic.
func (adminrest *AdminrestV1) GetTopic(getTopicOptions *GetTopicOptions) (result *TopicDetail, response *core.DetailedResponse, err error) {
	return adminrest.GetTopicWithContext(context.Background(), getTopicOptions)
}

// GetTopicWithContext is an alternate form of the GetTopic method which supports a Context parameter
func (adminrest *AdminrestV1) GetTopicWithContext(ctx context.Context, getTopicOptions *GetTopicOptions) (result *TopicDetail, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTopicOptions, "getTopicOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTopicOptions, "getTopicOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"topic_name": *getTopicOptions.TopicName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/topics/{topic_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTopicOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "GetTopic")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTopicDetail)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteTopic : Delete a topic
// Delete a topic.
func (adminrest *AdminrestV1) DeleteTopic(deleteTopicOptions *DeleteTopicOptions) (response *core.DetailedResponse, err error) {
	return adminrest.DeleteTopicWithContext(context.Background(), deleteTopicOptions)
}

// DeleteTopicWithContext is an alternate form of the DeleteTopic method which supports a Context parameter
func (adminrest *AdminrestV1) DeleteTopicWithContext(ctx context.Context, deleteTopicOptions *DeleteTopicOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTopicOptions, "deleteTopicOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTopicOptions, "deleteTopicOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"topic_name": *deleteTopicOptions.TopicName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/topics/{topic_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTopicOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "DeleteTopic")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = adminrest.Service.Request(request, nil)

	return
}

// UpdateTopic : Increase the number of partitions and/or update one or more topic configuration parameters
// Increase the number of partitions and/or update one or more topic configuration parameters.
func (adminrest *AdminrestV1) UpdateTopic(updateTopicOptions *UpdateTopicOptions) (response *core.DetailedResponse, err error) {
	return adminrest.UpdateTopicWithContext(context.Background(), updateTopicOptions)
}

// UpdateTopicWithContext is an alternate form of the UpdateTopic method which supports a Context parameter
func (adminrest *AdminrestV1) UpdateTopicWithContext(ctx context.Context, updateTopicOptions *UpdateTopicOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateTopicOptions, "updateTopicOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateTopicOptions, "updateTopicOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"topic_name": *updateTopicOptions.TopicName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/topics/{topic_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateTopicOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "UpdateTopic")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateTopicOptions.NewTotalPartitionCount != nil {
		body["new_total_partition_count"] = updateTopicOptions.NewTotalPartitionCount
	}
	if updateTopicOptions.Configs != nil {
		body["configs"] = updateTopicOptions.Configs
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = adminrest.Service.Request(request, nil)

	return
}

// GetMirroringTopicSelection : Get current topic selection for mirroring
// Get current topic selection for mirroring.
func (adminrest *AdminrestV1) GetMirroringTopicSelection(getMirroringTopicSelectionOptions *GetMirroringTopicSelectionOptions) (result *MirroringTopicSelection, response *core.DetailedResponse, err error) {
	return adminrest.GetMirroringTopicSelectionWithContext(context.Background(), getMirroringTopicSelectionOptions)
}

// GetMirroringTopicSelectionWithContext is an alternate form of the GetMirroringTopicSelection method which supports a Context parameter
func (adminrest *AdminrestV1) GetMirroringTopicSelectionWithContext(ctx context.Context, getMirroringTopicSelectionOptions *GetMirroringTopicSelectionOptions) (result *MirroringTopicSelection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMirroringTopicSelectionOptions, "getMirroringTopicSelectionOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/mirroring/topic-selection`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMirroringTopicSelectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "GetMirroringTopicSelection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMirroringTopicSelection)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReplaceMirroringTopicSelection : Replace topic selection for mirroring
// Replace topic selection for mirroring. This operation replaces the complete set of mirroring topic selections.
func (adminrest *AdminrestV1) ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptions *ReplaceMirroringTopicSelectionOptions) (result *MirroringTopicSelection, response *core.DetailedResponse, err error) {
	return adminrest.ReplaceMirroringTopicSelectionWithContext(context.Background(), replaceMirroringTopicSelectionOptions)
}

// ReplaceMirroringTopicSelectionWithContext is an alternate form of the ReplaceMirroringTopicSelection method which supports a Context parameter
func (adminrest *AdminrestV1) ReplaceMirroringTopicSelectionWithContext(ctx context.Context, replaceMirroringTopicSelectionOptions *ReplaceMirroringTopicSelectionOptions) (result *MirroringTopicSelection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceMirroringTopicSelectionOptions, "replaceMirroringTopicSelectionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceMirroringTopicSelectionOptions, "replaceMirroringTopicSelectionOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/mirroring/topic-selection`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceMirroringTopicSelectionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "ReplaceMirroringTopicSelection")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceMirroringTopicSelectionOptions.Includes != nil {
		body["includes"] = replaceMirroringTopicSelectionOptions.Includes
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMirroringTopicSelection)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetMirroringActiveTopics : Get topics that are being actively mirrored
// Get topics that are being actively mirrored.
func (adminrest *AdminrestV1) GetMirroringActiveTopics(getMirroringActiveTopicsOptions *GetMirroringActiveTopicsOptions) (result *MirroringActiveTopics, response *core.DetailedResponse, err error) {
	return adminrest.GetMirroringActiveTopicsWithContext(context.Background(), getMirroringActiveTopicsOptions)
}

// GetMirroringActiveTopicsWithContext is an alternate form of the GetMirroringActiveTopics method which supports a Context parameter
func (adminrest *AdminrestV1) GetMirroringActiveTopicsWithContext(ctx context.Context, getMirroringActiveTopicsOptions *GetMirroringActiveTopicsOptions) (result *MirroringActiveTopics, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMirroringActiveTopicsOptions, "getMirroringActiveTopicsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = adminrest.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(adminrest.Service.Options.URL, `/admin/mirroring/active-topics`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMirroringActiveTopicsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("adminrest", "V1", "GetMirroringActiveTopics")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = adminrest.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMirroringActiveTopics)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateTopicOptions : The CreateTopic options.
type CreateTopicOptions struct {
	// The name of topic to be created.
	Name *string

	// The number of partitions.
	Partitions *int64

	// The number of partitions, this field takes precedence over 'partitions'. Default value is 1 if not specified.
	PartitionCount *int64

	// The config properties to be set for the new topic.
	Configs []ConfigCreate

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateTopicOptions : Instantiate CreateTopicOptions
func (*AdminrestV1) NewCreateTopicOptions() *CreateTopicOptions {
	return &CreateTopicOptions{}
}

// SetName : Allow user to set Name
func (options *CreateTopicOptions) SetName(name string) *CreateTopicOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetPartitions : Allow user to set Partitions
func (options *CreateTopicOptions) SetPartitions(partitions int64) *CreateTopicOptions {
	options.Partitions = core.Int64Ptr(partitions)
	return options
}

// SetPartitionCount : Allow user to set PartitionCount
func (options *CreateTopicOptions) SetPartitionCount(partitionCount int64) *CreateTopicOptions {
	options.PartitionCount = core.Int64Ptr(partitionCount)
	return options
}

// SetConfigs : Allow user to set Configs
func (options *CreateTopicOptions) SetConfigs(configs []ConfigCreate) *CreateTopicOptions {
	options.Configs = configs
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTopicOptions) SetHeaders(param map[string]string) *CreateTopicOptions {
	options.Headers = param
	return options
}

// DeleteTopicOptions : The DeleteTopic options.
type DeleteTopicOptions struct {
	// The topic name for the topic to be listed.
	TopicName *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteTopicOptions : Instantiate DeleteTopicOptions
func (*AdminrestV1) NewDeleteTopicOptions(topicName string) *DeleteTopicOptions {
	return &DeleteTopicOptions{
		TopicName: core.StringPtr(topicName),
	}
}

// SetTopicName : Allow user to set TopicName
func (options *DeleteTopicOptions) SetTopicName(topicName string) *DeleteTopicOptions {
	options.TopicName = core.StringPtr(topicName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTopicOptions) SetHeaders(param map[string]string) *DeleteTopicOptions {
	options.Headers = param
	return options
}

// GetMirroringActiveTopicsOptions : The GetMirroringActiveTopics options.
type GetMirroringActiveTopicsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetMirroringActiveTopicsOptions : Instantiate GetMirroringActiveTopicsOptions
func (*AdminrestV1) NewGetMirroringActiveTopicsOptions() *GetMirroringActiveTopicsOptions {
	return &GetMirroringActiveTopicsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetMirroringActiveTopicsOptions) SetHeaders(param map[string]string) *GetMirroringActiveTopicsOptions {
	options.Headers = param
	return options
}

// GetMirroringTopicSelectionOptions : The GetMirroringTopicSelection options.
type GetMirroringTopicSelectionOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetMirroringTopicSelectionOptions : Instantiate GetMirroringTopicSelectionOptions
func (*AdminrestV1) NewGetMirroringTopicSelectionOptions() *GetMirroringTopicSelectionOptions {
	return &GetMirroringTopicSelectionOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetMirroringTopicSelectionOptions) SetHeaders(param map[string]string) *GetMirroringTopicSelectionOptions {
	options.Headers = param
	return options
}

// GetTopicOptions : The GetTopic options.
type GetTopicOptions struct {
	// The topic name for the topic to be listed.
	TopicName *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTopicOptions : Instantiate GetTopicOptions
func (*AdminrestV1) NewGetTopicOptions(topicName string) *GetTopicOptions {
	return &GetTopicOptions{
		TopicName: core.StringPtr(topicName),
	}
}

// SetTopicName : Allow user to set TopicName
func (options *GetTopicOptions) SetTopicName(topicName string) *GetTopicOptions {
	options.TopicName = core.StringPtr(topicName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetTopicOptions) SetHeaders(param map[string]string) *GetTopicOptions {
	options.Headers = param
	return options
}

// ListTopicsOptions : The ListTopics options.
type ListTopicsOptions struct {
	// A filter to be applied to the topic names. A simple filter can be specified as a string with asterisk (`*`)
	// wildcards representing 0 or more characters, e.g. `topic-name*` will filter all topic names that begin with the
	// string `topic-name` followed by any character sequence. A more complex filter pattern can be used by surrounding a
	// regular expression in forward slash (`/`) delimiters, e.g. `/topic-name.* /`.
	TopicFilter *string

	// The number of topic names to be returns.
	PerPage *int64

	// The page number to be returned. The number 1 represents the first page. The default value is 1.
	Page *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListTopicsOptions : Instantiate ListTopicsOptions
func (*AdminrestV1) NewListTopicsOptions() *ListTopicsOptions {
	return &ListTopicsOptions{}
}

// SetTopicFilter : Allow user to set TopicFilter
func (options *ListTopicsOptions) SetTopicFilter(topicFilter string) *ListTopicsOptions {
	options.TopicFilter = core.StringPtr(topicFilter)
	return options
}

// SetPerPage : Allow user to set PerPage
func (options *ListTopicsOptions) SetPerPage(perPage int64) *ListTopicsOptions {
	options.PerPage = core.Int64Ptr(perPage)
	return options
}

// SetPage : Allow user to set Page
func (options *ListTopicsOptions) SetPage(page int64) *ListTopicsOptions {
	options.Page = core.Int64Ptr(page)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListTopicsOptions) SetHeaders(param map[string]string) *ListTopicsOptions {
	options.Headers = param
	return options
}

// ReplaceMirroringTopicSelectionOptions : The ReplaceMirroringTopicSelection options.
type ReplaceMirroringTopicSelectionOptions struct {
	Includes []string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceMirroringTopicSelectionOptions : Instantiate ReplaceMirroringTopicSelectionOptions
func (*AdminrestV1) NewReplaceMirroringTopicSelectionOptions() *ReplaceMirroringTopicSelectionOptions {
	return &ReplaceMirroringTopicSelectionOptions{}
}

// SetIncludes : Allow user to set Includes
func (options *ReplaceMirroringTopicSelectionOptions) SetIncludes(includes []string) *ReplaceMirroringTopicSelectionOptions {
	options.Includes = includes
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceMirroringTopicSelectionOptions) SetHeaders(param map[string]string) *ReplaceMirroringTopicSelectionOptions {
	options.Headers = param
	return options
}

// ReplicaAssignmentBrokers : ReplicaAssignmentBrokers struct
type ReplicaAssignmentBrokers struct {
	Replicas []int64 `json:"replicas,omitempty"`
}

// UnmarshalReplicaAssignmentBrokers unmarshals an instance of ReplicaAssignmentBrokers from the specified map of raw messages.
func UnmarshalReplicaAssignmentBrokers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ReplicaAssignmentBrokers)
	err = core.UnmarshalPrimitive(m, "replicas", &obj.Replicas)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateTopicOptions : The UpdateTopic options.
type UpdateTopicOptions struct {
	// The topic name for the topic to be listed.
	TopicName *string `validate:"required,ne="`

	// The new partition number to be increased.
	NewTotalPartitionCount *int64

	// The config properties to be updated for the topic. Valid config keys are 'cleanup.policy', 'retention.ms',
	// 'retention.bytes', 'segment.bytes', 'segment.ms', 'segment.index.bytes'.
	Configs []ConfigUpdate

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateTopicOptions : Instantiate UpdateTopicOptions
func (*AdminrestV1) NewUpdateTopicOptions(topicName string) *UpdateTopicOptions {
	return &UpdateTopicOptions{
		TopicName: core.StringPtr(topicName),
	}
}

// SetTopicName : Allow user to set TopicName
func (options *UpdateTopicOptions) SetTopicName(topicName string) *UpdateTopicOptions {
	options.TopicName = core.StringPtr(topicName)
	return options
}

// SetNewTotalPartitionCount : Allow user to set NewTotalPartitionCount
func (options *UpdateTopicOptions) SetNewTotalPartitionCount(newTotalPartitionCount int64) *UpdateTopicOptions {
	options.NewTotalPartitionCount = core.Int64Ptr(newTotalPartitionCount)
	return options
}

// SetConfigs : Allow user to set Configs
func (options *UpdateTopicOptions) SetConfigs(configs []ConfigUpdate) *UpdateTopicOptions {
	options.Configs = configs
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTopicOptions) SetHeaders(param map[string]string) *UpdateTopicOptions {
	options.Headers = param
	return options
}

// ConfigCreate : ConfigCreate struct
type ConfigCreate struct {
	// The name of the config property.
	Name *string `json:"name,omitempty"`

	// The value for a config property.
	Value *string `json:"value,omitempty"`
}

// UnmarshalConfigCreate unmarshals an instance of ConfigCreate from the specified map of raw messages.
func UnmarshalConfigCreate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigCreate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigUpdate : ConfigUpdate struct
type ConfigUpdate struct {
	// The name of the config property.
	Name *string `json:"name,omitempty"`

	// The value for a config property.
	Value *string `json:"value,omitempty"`

	// When true, the value of the config property is reset to its default value.
	ResetToDefault *bool `json:"reset_to_default,omitempty"`
}

// UnmarshalConfigUpdate unmarshals an instance of ConfigUpdate from the specified map of raw messages.
func UnmarshalConfigUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigUpdate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reset_to_default", &obj.ResetToDefault)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MirroringActiveTopics : Topics that are being actively mirrored.
type MirroringActiveTopics struct {
	ActiveTopics []string `json:"active_topics,omitempty"`
}

// UnmarshalMirroringActiveTopics unmarshals an instance of MirroringActiveTopics from the specified map of raw messages.
func UnmarshalMirroringActiveTopics(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MirroringActiveTopics)
	err = core.UnmarshalPrimitive(m, "active_topics", &obj.ActiveTopics)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MirroringTopicSelection : Mirroring topic selection payload.
type MirroringTopicSelection struct {
	Includes []string `json:"includes,omitempty"`
}

// UnmarshalMirroringTopicSelection unmarshals an instance of MirroringTopicSelection from the specified map of raw messages.
func UnmarshalMirroringTopicSelection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MirroringTopicSelection)
	err = core.UnmarshalPrimitive(m, "includes", &obj.Includes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReplicaAssignment : ReplicaAssignment struct
type ReplicaAssignment struct {
	// The ID of the partition.
	ID *int64 `json:"id,omitempty"`

	Brokers *ReplicaAssignmentBrokers `json:"brokers,omitempty"`
}

// UnmarshalReplicaAssignment unmarshals an instance of ReplicaAssignment from the specified map of raw messages.
func UnmarshalReplicaAssignment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ReplicaAssignment)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "brokers", &obj.Brokers, UnmarshalReplicaAssignmentBrokers)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TopicConfigs : TopicConfigs struct
type TopicConfigs struct {
	// The value of config property 'cleanup.policy'.
	CleanupPolicy *string `json:"cleanup.policy,omitempty"`

	// The value of config property 'min.insync.replicas'.
	MinInsyncReplicas *string `json:"min.insync.replicas,omitempty"`

	// The value of config property 'retention.bytes'.
	RetentionBytes *string `json:"retention.bytes,omitempty"`

	// The value of config property 'retention.ms'.
	RetentionMs *string `json:"retention.ms,omitempty"`

	// The value of config property 'segment.bytes'.
	SegmentBytes *string `json:"segment.bytes,omitempty"`

	// The value of config property 'segment.index.bytes'.
	SegmentIndexBytes *string `json:"segment.index.bytes,omitempty"`

	// The value of config property 'segment.ms'.
	SegmentMs *string `json:"segment.ms,omitempty"`
}

// UnmarshalTopicConfigs unmarshals an instance of TopicConfigs from the specified map of raw messages.
func UnmarshalTopicConfigs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TopicConfigs)
	err = core.UnmarshalPrimitive(m, "cleanup.policy", &obj.CleanupPolicy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "min.insync.replicas", &obj.MinInsyncReplicas)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "retention.bytes", &obj.RetentionBytes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "retention.ms", &obj.RetentionMs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "segment.bytes", &obj.SegmentBytes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "segment.index.bytes", &obj.SegmentIndexBytes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "segment.ms", &obj.SegmentMs)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TopicDetail : TopicDetail struct
type TopicDetail struct {
	// The name of the topic.
	Name *string `json:"name,omitempty"`

	// The number of partitions.
	Partitions *int64 `json:"partitions,omitempty"`

	// The number of replication factor.
	ReplicationFactor *int64 `json:"replicationFactor,omitempty"`

	// The value of config property 'retention.ms'.
	RetentionMs *int64 `json:"retentionMs,omitempty"`

	// The value of config property 'cleanup.policy'.
	CleanupPolicy *string `json:"cleanupPolicy,omitempty"`

	Configs *TopicConfigs `json:"configs,omitempty"`

	// The replia assignment of the topic.
	ReplicaAssignments []ReplicaAssignment `json:"replicaAssignments,omitempty"`
}

// UnmarshalTopicDetail unmarshals an instance of TopicDetail from the specified map of raw messages.
func UnmarshalTopicDetail(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TopicDetail)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "partitions", &obj.Partitions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "replicationFactor", &obj.ReplicationFactor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "retentionMs", &obj.RetentionMs)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cleanupPolicy", &obj.CleanupPolicy)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalTopicConfigs)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "replicaAssignments", &obj.ReplicaAssignments, UnmarshalReplicaAssignment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
