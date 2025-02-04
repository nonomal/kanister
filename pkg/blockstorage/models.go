// Copyright 2019 The Kanister Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package blockstorage

import (
	strfmt "github.com/go-openapi/strfmt"
)

const (
	BytesInGi = 1024 * 1024 * 1024
	BytesInMi = 1024 * 1024
)

// SizeInGi converts bytes to GiBs
func SizeInGi(sizeInBytes int64) int64 {
	return (sizeInBytes + int64(BytesInGi) - 1) / int64(BytesInGi)
}

// Volume A storage provider volume
type Volume struct {

	// Availability zone
	Az string

	// Time stamp when volume creation was initiated
	CreationTime TimeStamp

	// Volume is encrypted
	Encrypted bool

	// A unique identifier generated by the storage provider
	ID string

	// Volume IOPS, if specified for this volume
	Iops int64

	// The size of the volume, in Bytes
	SizeInBytes int64

	// tags
	Tags VolumeTags

	// Storage type for this volume
	Type Type

	// Volume type
	VolumeType string

	// Attributes specific to the provider
	Attributes map[string]string
}

// Snapshot of Volume
type Snapshot struct {

	// Time stamp when snapshot creation was initiated
	CreationTime TimeStamp

	// Snapshot is encrypted
	Encrypted bool

	// A unique identifier generated by the storage provider
	ID string

	// Snapshot availability region
	Region string

	// The size of the snapshot, in Bytes
	SizeInBytes int64

	// tags
	Tags SnapshotTags

	// Storage type of the source volume for this snapshot
	Type Type

	// volume
	Volume *Volume

	// ProvisioningState is snapshot's provisioning state.
	ProvisioningState string
}

// TimeStamp Time stamp for an event related to an object, for example when the object was created.
type TimeStamp strfmt.DateTime

// VolumeTags volume tags
type VolumeTags []*KeyValue

// SnapshotTags snapshot tags
type SnapshotTags []*KeyValue

// KeyValue String key-value pairs
type KeyValue struct {

	// Key or index name
	Key string

	// Value string
	Value string
}
