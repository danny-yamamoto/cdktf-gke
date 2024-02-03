package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/computeinstance"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here
	provider.NewGoogleProvider(stack, jsii.String("google"), &provider.GoogleProviderConfig{
		Region:  jsii.String("asia-northeast1"),
		Project: jsii.String("sanbox-334000"),
		Zone:    jsii.String("asia-northeast1-b"),
	})
	computeinstance.NewComputeInstance(stack, jsii.String("test-vm"), &computeinstance.ComputeInstanceConfig{
		Name:                   jsii.String("test-vm"),
		MachineType:            jsii.String("e2-micro"),
		AllowStoppingForUpdate: jsii.Bool(true),
		Zone:                   jsii.String("asia-northeast1-b"),
		BootDisk: &computeinstance.ComputeInstanceBootDisk{
			InitializeParams: &computeinstance.ComputeInstanceBootDiskInitializeParams{
				Image: jsii.String("projects/debian-cloud/global/images/family/debian-11"),
				Size:  jsii.Number(50),
				Type:  jsii.String("pd-standard"),
			},
		},
		NetworkInterface: &[]computeinstance.ComputeInstanceNetworkInterface{
			{
				Network: jsii.String("default"),
			},
		},
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktf-gke")

	app.Synth()
}
