package main

import (
	"log"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestDefaultsExampleWithOnlyRequiredInputs(t *testing.T) {
	//
	// Arrange
	//

	tfOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/defaults",

		EnvVars: map[string]string{
			"TF_IN_AUTOMATION": "true",
		},

		// NoColor: true,
		// NoStderr: true,

		// Logger: logger.Discard,
	}

	//
	// Act
	//
	defer terraform.Destroy(t, tfOptions)

	init, err := terraform.InitE(t, tfOptions)

	if err != nil {
		log.Println(err)
	}

	t.Log(init)
}
