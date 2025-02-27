// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DatabaseInitParameters struct {

	// (String) ID of the Postgres Flex instance.
	// ID of the Postgres Flex instance.
	// +crossplane:generate:reference:type=github.com/stackitcloud/crossplane-provider-stackit/apis/postgresflex/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/stackitcloud/crossplane-provider-stackit/config/postgresflex.AtProviderExtractor("instanceId")
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in postgresflex to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in postgresflex to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// (String) Database name.
	// Database name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Username of the database owner.
	// Username of the database owner.
	// +crossplane:generate:reference:type=github.com/stackitcloud/crossplane-provider-stackit/apis/postgresflex/v1alpha1.User
	// +crossplane:generate:reference:extractor=github.com/stackitcloud/crossplane-provider-stackit/config/postgresflex.AtProviderExtractor("username")
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Reference to a User in postgresflex to populate owner.
	// +kubebuilder:validation:Optional
	OwnerRef *v1.Reference `json:"ownerRef,omitempty" tf:"-"`

	// Selector for a User in postgresflex to populate owner.
	// +kubebuilder:validation:Optional
	OwnerSelector *v1.Selector `json:"ownerSelector,omitempty" tf:"-"`

	// (String) STACKIT project ID to which the instance is associated.
	// STACKIT project ID to which the instance is associated.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type DatabaseObservation struct {

	// (String) Database ID.
	// Database ID.
	DatabaseID *string `json:"databaseId,omitempty" tf:"database_id,omitempty"`

	// It is structured as "project_id,instance_id,database_id".
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) ID of the Postgres Flex instance.
	// ID of the Postgres Flex instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// (String) Database name.
	// Database name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Username of the database owner.
	// Username of the database owner.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// (String) STACKIT project ID to which the instance is associated.
	// STACKIT project ID to which the instance is associated.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type DatabaseParameters struct {

	// (String) ID of the Postgres Flex instance.
	// ID of the Postgres Flex instance.
	// +crossplane:generate:reference:type=github.com/stackitcloud/crossplane-provider-stackit/apis/postgresflex/v1alpha1.Instance
	// +crossplane:generate:reference:extractor=github.com/stackitcloud/crossplane-provider-stackit/config/postgresflex.AtProviderExtractor("instanceId")
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in postgresflex to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in postgresflex to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// (String) Database name.
	// Database name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Username of the database owner.
	// Username of the database owner.
	// +crossplane:generate:reference:type=github.com/stackitcloud/crossplane-provider-stackit/apis/postgresflex/v1alpha1.User
	// +crossplane:generate:reference:extractor=github.com/stackitcloud/crossplane-provider-stackit/config/postgresflex.AtProviderExtractor("username")
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Reference to a User in postgresflex to populate owner.
	// +kubebuilder:validation:Optional
	OwnerRef *v1.Reference `json:"ownerRef,omitempty" tf:"-"`

	// Selector for a User in postgresflex to populate owner.
	// +kubebuilder:validation:Optional
	OwnerSelector *v1.Selector `json:"ownerSelector,omitempty" tf:"-"`

	// (String) STACKIT project ID to which the instance is associated.
	// STACKIT project ID to which the instance is associated.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DatabaseInitParameters `json:"initProvider,omitempty"`
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Database is the Schema for the Databases API. Postgres Flex database resource schema. Must have a region specified in the provider configuration.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,stackit}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectId) || (has(self.initProvider) && has(self.initProvider.projectId))",message="spec.forProvider.projectId is a required parameter"
	Spec   DatabaseSpec   `json:"spec"`
	Status DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Repository type metadata.
var (
	Database_Kind             = "Database"
	Database_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Database_Kind}.String()
	Database_KindAPIVersion   = Database_Kind + "." + CRDGroupVersion.String()
	Database_GroupVersionKind = CRDGroupVersion.WithKind(Database_Kind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
