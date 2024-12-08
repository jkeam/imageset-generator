package lib

import (
	"gorm.io/gorm"
)

// name: 4.17
type Version struct {
	gorm.Model
	Name     string
	Channels []Channel `json:"-"`
	Catalogs []Catalog `json:"-"`
}

// name: stable-4.17
type Channel struct {
	gorm.Model
	Name      string
	VersionID uint
	Releases []Release `json:"-"`
}

// name: 4.17.1
// pull_spec: quay.io/openshift-release-dev/ocp-release@sha256:e16ac60ac6971e5b6f89c1d818f5ae711c0d63ad6a6a26ffe795c738e8cc4dde
type Release struct {
	gorm.Model
	Name      string
	PullSpec  string
	ChannelID uint
	Images    []Image `json:"-"`
}

// name: agent-installer-api-server,
// pull_spec: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:f6f5ebb3320f12bc582d3c3884500a890c81358af0c47fa3f9a2cb40fb4d5746
type Image struct {
	gorm.Model
	Name      string
	PullSpec  string
	ReleaseID uint
}

// name: registry.redhat.io/redhat/redhat-operator-index:v4.17
type Catalog struct {
	gorm.Model
	Name      string
	VersionID uint
	Operators []Operator `json:"-"`
}

// name: cluster-logging
// display: Red Hat OpenShift Logging
// default_channel: stable-6.0
type Operator struct {
	gorm.Model
	Name           string
	Display        string
	DefaultChannel string
	CatalogID      uint
}
