// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package infraattributes

import (
	"fybrik.io/fybrik/pkg/model/taxonomy"
)

// +kubebuilder:validation:Enum=numeric;string;bool
type AttributeType string

// List of attribute types
const (
	Numeric AttributeType = "numeric"
	String  AttributeType = "string"
	Bool    AttributeType = "bool"
)

type RangeType struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

// +kubebuilder:validation:Enum=fybrikmodule;fybrikstorageaccount;cluster
type InstanceType string

// List of instance types
const (
	Module         InstanceType = "fybrikmodule"
	Cluster        InstanceType = "cluster"
	StorageAccount InstanceType = "fybrikstorageaccount"
)

type InfrastructureElement struct {
	// Attribute name defined in the taxonomy
	// +required
	Attribute taxonomy.Attribute `json:"attribute"`
	// Description
	Description string `json:"description,omitempty"`
	// Attribute type, e.g. numeric or string
	// +required
	Type AttributeType `json:"type"`
	// Attribute value
	// +required
	Value string `json:"value"`
	// Measurement units defined in the tzxonomy
	Units taxonomy.Units `json:"units,omitempty"`
	// A resource defined by the attribute ("storageaccount","module","cluster")
	Object InstanceType `json:"object,omitempty"`
	// A reference to the resource instance, e.g. storage account name
	Instance string `json:"instance,omitempty"`
	// A scale of values (minimum and maximum) when applicable
	Scale *RangeType `json:"scale,omitempty"`
	// A list of arguments defining a specific metric, e.g. regions for a bandwidth
	Arguments []string `json:"arguments,omitempty"`
}

type Infrastructure struct {
	Items []InfrastructureElement `json:"infrastructure"`
}
