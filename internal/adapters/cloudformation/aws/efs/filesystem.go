package efs

import (
	"github.com/aquasecurity/trivy/pkg/providers/aws/efs"
	"github.com/aquasecurity/trivy/pkg/scanners/cloudformation/parser"
)

func getFileSystems(ctx parser.FileContext) (filesystems []efs.FileSystem) {

	filesystemResources := ctx.GetResourcesByType("AWS::EFS::FileSystem")

	for _, r := range filesystemResources {

		filesystem := efs.FileSystem{
			Metadata:  r.Metadata(),
			Encrypted: r.GetBoolProperty("Encrypted"),
		}

		filesystems = append(filesystems, filesystem)
	}

	return filesystems
}
