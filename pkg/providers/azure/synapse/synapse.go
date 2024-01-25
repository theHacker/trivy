package synapse

import (
	defsecTypes "github.com/aquasecurity/trivy/pkg/types"
)

type Synapse struct {
	Workspaces []Workspace
}

type Workspace struct {
	Metadata                    defsecTypes.MisconfigMetadata
	EnableManagedVirtualNetwork defsecTypes.BoolValue
}
