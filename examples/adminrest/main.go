/*
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package main

// Code Setup
import (
	"fmt"
	"net/http"
	"os"

	"github.com/IBM/eventstreams-go-sdk/pkg/adminrestv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

// End Code Setup

func main() {
	URL := os.Getenv("KAFKA_ADMIN_URL")
	apiKey := os.Getenv("API_KEY")
	bearerToken := os.Getenv("BEARER_TOKEN")

	if URL == "" {
		fmt.Println("Please set env KAFKA_ADMIN_URL")
		os.Exit(1)
	}

	if apiKey == "" && bearerToken == "" {
		fmt.Println("Please set either an API_KEY or a BEARER_TOKEN")
		os.Exit(1)
	}

	if apiKey != "" && bearerToken != "" {
		fmt.Println("Please set either an API_KEY or a BEARER_TOKEN not both")
		os.Exit(1)
	}
	// Create Authenticator
	var authenticator core.Authenticator

	if apiKey != "" {
		var err error
		// Create an Basic IAM authenticator.
		authenticator, err = core.NewBasicAuthenticator("token", apiKey)
		if err != nil {
			fmt.Printf("failed to create new basic authenticator: %s\n", err.Error())
			os.Exit(1)
		}
	} else {
		var err error
		// Create an IAM Bearer Token authenticator.
		authenticator, err = core.NewBearerTokenAuthenticator(bearerToken)
		if err != nil {
			fmt.Printf("failed to create new bearer token authenticator: %s\n", err.Error())
			os.Exit(1)
		}
	}
	// End Authenticator

	// Create Service
	serviceAPI, serviceErr := adminrestv1.NewAdminrestV1(&adminrestv1.AdminrestV1Options{
		URL:           URL,
		Authenticator: authenticator,
	})
	// End Create Service

	if serviceErr != nil {
		fmt.Printf("Error Creating Service")
		os.Exit(1)
	}

	// Always try to delete test-topic
	fmt.Printf("Delete Topic\n")
	_ = deleteTopic(serviceAPI)

	fmt.Printf("List Topics\n")
	err := listTopics(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Create Topic\n")
	err = createTopic(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Print Topic Details\n")
	err = topicDetails(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("List Topics\n")
	err = listTopics(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Update Topic Details\n")
	err = updateTopicDetails(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Print Topic Details\n")
	err = topicDetails(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	// Uncomment these examples if you are running against a Event Streams Mirrored Target Cluster.
	// fmt.Printf("List Active Mirroring Topics\n")
	// err = getMirroringActiveTopics(serviceAPI)
	// if err != nil {
	// 	fmt.Printf("%s\n", err.Error())
	// 	os.Exit(1)
	// }

	// fmt.Printf("Replace Mirroring Topics\n")
	// err = replaceMirroringTopicSelection(serviceAPI)
	// if err != nil {
	// 	fmt.Printf("%s\n", err.Error())
	// 	os.Exit(1)
	// }

	// fmt.Printf("List Mirroring Topic Selection\n")
	// err = listMirroringTopicSelection(serviceAPI)
	// if err != nil {
	// 	fmt.Printf("%s\n", err.Error())
	// 	os.Exit(1)
	// }

	fmt.Printf("Delete Topic\n")
	err = deleteTopic(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("List Topics\n")
	err = listTopics(serviceAPI)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}

func listTopics(serviceAPI *adminrestv1.AdminrestV1) error {
	// Construct an instance of the ListTopicsOptions model
	listTopicsOptionsModel := new(adminrestv1.ListTopicsOptions)

	// Call ListTopics.
	result, response, operationErr := serviceAPI.ListTopics(listTopicsOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Listing Topics" + operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Error Listing Topics: status %d\n", response.StatusCode)
	}

	// Loop and print topics.
	for _, topicDetail := range result {
		fmt.Printf("\tname: %s\n", *topicDetail.Name)
	}
	return nil
} // func.end

func topicDetails(serviceAPI *adminrestv1.AdminrestV1) error {
	// Construct an instance of the GetTopicOptions model
	getTopicOptionsModel := new(adminrestv1.GetTopicOptions)
	getTopicOptionsModel.TopicName = core.StringPtr("test-topic")

	// Call List Topic Details.
	result, response, operationErr := serviceAPI.GetTopic(getTopicOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Listing Topic Details" + operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Error Listing Topic Details: status %d\n", response.StatusCode)
	}

	// Print topics details.
	fmt.Printf("\tname: \t\t\t%s\n", *result.Name)

	// The number of partitions.
	fmt.Printf("\tno of partitions: \t%d\n", *result.Partitions)

	// The number of replication factor.
	fmt.Printf("\treplication factor: \t%d\n", *result.ReplicationFactor)

	// // The value of config property 'retention.ms'.
	fmt.Printf("\tretention (ms): \t%d\n", *result.RetentionMs)

	// // The value of config property 'cleanup.policy'.
	fmt.Printf("\tcleanup policy: \t%s\n", *result.CleanupPolicy)

	// Configs *TopicConfigs
	fmt.Printf("\ttopic configs: \t\t%+v\n", *result.Configs)

	// The replia assignment of the topic.
	// ReplicaAssignments []ReplicaAssignment
	for _, assignment := range result.ReplicaAssignments {
		fmt.Printf("\tassignment:  \t\tid:%d,  \tbrokers: %+v\n", assignment.ID, assignment.Brokers)
	}

	return nil

} // func.end

func createTopic(serviceAPI *adminrestv1.AdminrestV1) error {
	// Construct an instance of the createTopicOptionsModel.
	createTopicOptionsModel := new(adminrestv1.CreateTopicOptions)
	createTopicOptionsModel.Name = core.StringPtr("test-topic")
	createTopicOptionsModel.PartitionCount = core.Int64Ptr(int64(1))

	// Create the Topic.
	response, operationErr := serviceAPI.CreateTopic(createTopicOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Creating Topics: %s\n", operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error Creating Topic: status %d\n", response.StatusCode)
	}

	fmt.Printf("\tname: %s created\n", *createTopicOptionsModel.Name)

	return nil
} // func.end

func deleteTopic(serviceAPI *adminrestv1.AdminrestV1) error {
	// Construct an instance of the DeleteTopicOptions model
	deleteTopicOptionsModel := new(adminrestv1.DeleteTopicOptions)
	deleteTopicOptionsModel.TopicName = core.StringPtr("test-topic")

	// Delete Topic
	response, operationErr := serviceAPI.DeleteTopic(deleteTopicOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Deleting Topic: %s\n", operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error Deleting Topic: status %d\n", response.StatusCode)
	}

	fmt.Printf("\tname: %s deleted\n", *deleteTopicOptionsModel.TopicName)
	return nil
} // func.end

func updateTopicDetails(serviceAPI *adminrestv1.AdminrestV1) error {
	// Construct an instance of the UpdateTopicOptions model
	updateTopicOptionsModel := new(adminrestv1.UpdateTopicOptions)
	updateTopicOptionsModel.TopicName = core.StringPtr("test-topic")
	updateTopicOptionsModel.NewTotalPartitionCount = core.Int64Ptr(int64(6))

	// Invoke operation with valid options model.
	response, operationErr := serviceAPI.UpdateTopic(updateTopicOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Updating Topic: %s\n", operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error Updating Topics: status %d\n", response.StatusCode)
	}

	fmt.Printf("\tname: %s updated\n", *updateTopicOptionsModel.TopicName)

	return nil
} // func.end

// nolint
func replaceMirroringTopicSelection(serviceAPI *adminrestv1.AdminrestV1) error {
	// Construct an instance of the ReplaceMirroringTopicSelectionOptions model
	replaceMirroringTopicSelectionOptionsModel := new(adminrestv1.ReplaceMirroringTopicSelectionOptions)
	replaceMirroringTopicSelectionOptionsModel.Includes = []string{"test-topic"}

	// Invoke operation with valid options model.
	result, response, operationErr := serviceAPI.ReplaceMirroringTopicSelection(replaceMirroringTopicSelectionOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Replacing Mirroring Topics: %s\n", operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error Replacing Mirroring Topics: status %d\n", response.StatusCode)
	}

	// Loop and print mirroring topics.
	for _, topicName := range result.Includes {
		fmt.Printf("\ttopic added: %s\n", topicName)
	}

	return nil
} // func.end

// nolint
func listMirroringTopicSelection(serviceAPI *adminrestv1.AdminrestV1) error {
	// Construct an instance of the GetMirroringTopicSelectionOptions model
	getMirroringTopicSelectionOptionsModel := new(adminrestv1.GetMirroringTopicSelectionOptions)

	// Call GetMirroringTopicSelection.
	result, response, operationErr := serviceAPI.GetMirroringTopicSelection(getMirroringTopicSelectionOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Listing Mirroring Topics: %s\n", operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error Listing Mirroring Topics: status %d\n", response.StatusCode)
	}

	// Loop and print mirroring topics.
	for _, topicName := range result.Includes {
		fmt.Printf("\tname: %s\n", topicName)
	}

	return nil
} // func.end

// nolint
func getMirroringActiveTopics(serviceAPI *adminrestv1.AdminrestV1) error {
	// Construct an instance of the GetMirroringActiveTopicsOptions model
	getMirroringActiveTopicsOptionsModel := new(adminrestv1.GetMirroringActiveTopicsOptions)

	// Call GetMirroringActiveTopics.
	result, response, operationErr := serviceAPI.GetMirroringActiveTopics(getMirroringActiveTopicsOptionsModel)
	if operationErr != nil {
		return fmt.Errorf("Error Listing Active Mirroring Topics: %s\n", operationErr.Error())
	}

	// Check the result.
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error Listing Active Mirroring Topics: status %d\n", response.StatusCode)
	}

	// Loop and print mirroring topics.
	for _, topicName := range result.ActiveTopics {
		fmt.Printf("\tname: %s\n", topicName)
	}

	return nil
} // func.end
