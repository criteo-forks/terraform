//
// This file is automatically generated by scripts/generate-plugins.go -- Do not edit!
//
package command

import (
	chefprovisioner "github.com/hashicorp/terraform/builtin/provisioners/chef"
	chefzeroprovisioner "github.com/hashicorp/terraform/builtin/provisioners/chef-zero"
	fileprovisioner "github.com/hashicorp/terraform/builtin/provisioners/file"
	habitatprovisioner "github.com/hashicorp/terraform/builtin/provisioners/habitat"
	localexecprovisioner "github.com/hashicorp/terraform/builtin/provisioners/local-exec"
	remoteexecprovisioner "github.com/hashicorp/terraform/builtin/provisioners/remote-exec"
	saltmasterlessprovisioner "github.com/hashicorp/terraform/builtin/provisioners/salt-masterless"

	"github.com/hashicorp/terraform/plugin"
)

var InternalProviders = map[string]plugin.ProviderFunc{}

var InternalProvisioners = map[string]plugin.ProvisionerFunc{
	"chef":            chefprovisioner.Provisioner,
	"chef-zero":       chefzeroprovisioner.Provisioner,
	"file":            fileprovisioner.Provisioner,
	"habitat":         habitatprovisioner.Provisioner,
	"local-exec":      localexecprovisioner.Provisioner,
	"remote-exec":     remoteexecprovisioner.Provisioner,
	"salt-masterless": saltmasterlessprovisioner.Provisioner,
}
