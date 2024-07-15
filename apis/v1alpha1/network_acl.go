// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkAclSpec defines the desired state of NetworkAcl.
//
// Describes a network ACL.
type NetworkACLSpec struct {
	Associations []*NetworkACLAssociation `json:"associations,omitempty"`
	Entries      []*NetworkACLEntry       `json:"entries,omitempty"`
	// The tags. The value parameter is required, but if you don't want the tag
	// to have a value, specify the parameter with no value, and we set the value
	// to an empty string.
	Tags []*Tag `json:"tags,omitempty"`
	// The ID of the VPC.
	VPCID  *string                                  `json:"vpcID,omitempty"`
	VPCRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"vpcRef,omitempty"`
}

// NetworkACLStatus defines the observed state of NetworkACL
type NetworkACLStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The ID of the network ACL.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty"`
	// Indicates whether this is the default network ACL for the VPC.
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty"`
	// The ID of the Amazon Web Services account that owns the network ACL.
	// +kubebuilder:validation:Optional
	OwnerID *string `json:"ownerID,omitempty"`
}

// NetworkACL is the Schema for the NetworkACLS API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ID",type=string,priority=0,JSONPath=`.status.id`
type NetworkACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkACLSpec   `json:"spec,omitempty"`
	Status            NetworkACLStatus `json:"status,omitempty"`
}

// NetworkACLList contains a list of NetworkACL
// +kubebuilder:object:root=true
type NetworkACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkACL `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkACL{}, &NetworkACLList{})
}