// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCluster openpitrix cluster
// swagger:model openpitrixCluster
type OpenpitrixCluster struct {

	// additional info
	AdditionalInfo string `json:"additional_info,omitempty"`

	// id of app run in cluster
	AppID string `json:"app_id,omitempty"`

	// cluster common set
	ClusterCommonSet OpenpitrixClusterClusterCommonSet `json:"cluster_common_set"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// cluster link set
	ClusterLinkSet OpenpitrixClusterClusterLinkSet `json:"cluster_link_set"`

	// cluster loadbalancer set
	ClusterLoadbalancerSet OpenpitrixClusterClusterLoadbalancerSet `json:"cluster_loadbalancer_set"`

	// cluster node set
	ClusterNodeSet OpenpitrixClusterClusterNodeSet `json:"cluster_node_set"`

	// cluster role set
	ClusterRoleSet OpenpitrixClusterClusterRoleSet `json:"cluster_role_set"`

	// cluster type, frontgate or normal cluster
	ClusterType int64 `json:"cluster_type,omitempty"`

	// the time when cluster create
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// cluster used to debug or not
	Debug bool `json:"debug,omitempty"`

	// cluster description
	Description string `json:"description,omitempty"`

	// endpoint of cluster
	Endpoints string `json:"endpoints,omitempty"`

	// cluster env
	Env string `json:"env,omitempty"`

	// frontgate id, a proxy for vpc to communicate
	FrontgateID string `json:"frontgate_id,omitempty"`

	// global uuid
	GlobalUUID string `json:"global_uuid,omitempty"`

	// metadata root access
	MetadataRootAccess bool `json:"metadata_root_access,omitempty"`

	// cluster name
	Name string `json:"name,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// owner path, concat string group_path:user_id
	OwnerPath string `json:"owner_path,omitempty"`

	// cluster runtime id
	RuntimeID string `json:"runtime_id,omitempty"`

	// cluster status eg.[active|used|enabled|disabled|deleted|stopped|ceased]
	Status string `json:"status,omitempty"`

	// record status changed time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// subnet id, cluster run in a subnet
	SubnetID string `json:"subnet_id,omitempty"`

	// cluster transition status eg.[creating|deleting|upgrading|updating|rollbacking|stopping|starting|recovering|ceasing|resizing|scaling]
	TransitionStatus string `json:"transition_status,omitempty"`

	// upgrade status, unused
	UpgradeStatus string `json:"upgrade_status,omitempty"`

	// cluster upgraded time
	UpgradeTime strfmt.DateTime `json:"upgrade_time,omitempty"`

	// id of version of app run in cluster
	VersionID string `json:"version_id,omitempty"`

	// vpc id, a vpc contain one more subnet
	VpcID string `json:"vpc_id,omitempty"`

	// zone of cluster eg.[pek3a|pek3b]
	Zone string `json:"zone,omitempty"`
}

// Validate validates this openpitrix cluster
func (m *OpenpitrixCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCluster) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
