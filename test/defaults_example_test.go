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
	repositoryName := "RepoName"

	tfOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/defaults",

		EnvVars: map[string]string{
			"TF_IN_AUTOMATION": "true",
		},

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"required_name": repositoryName,
		},

		NoColor: true,
	}

	//
	// Act
	//
	defer terraform.Destroy(t, tfOptions)

	// Run `terraform init`
	init, err := terraform.InitE(t, tfOptions)

	if err != nil {
		log.Println(err)
	}

	t.Log(init)

	// Run `terraform plan`
	// terraform.Plan(t, tfOptions)

	// Run `terraform apply`
	// terraform.Apply(t, tfOptions)

	//
	// Assert
	//
	//output := terraform.Output(t, tfOptions, "output_var")

	//assert.Equal(t, repositoryName, output)
}

// func TestDefaultsExampleWithInputs(t *testing.T) {
// 	//
// 	// Arrange
// 	//
// 	repositoryName := "acustomreponame"

// 	tfOptions := &terraform.Options{
// 		TerraformDir: "../examples/defaults",
// 		Vars: map[string]interface{}{
// 			"input_var": inputVariable,
// 		},
// 	}

// 	//
// 	// Act
// 	//
// 	defer terraform.Destroy(t, tfOptions)

// 	// Run `terraform init`
// 	init, err := terraform.InitE(t, tfOptions)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	t.Log(init)

// 	// Run `terraform plan`
// 	plan, err := terraform.PlanE(t, tfOptions)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	t.Log(plan)

// 	// Run `terraform apply`
// 	apply, err := terraform.ApplyE(t, tfOptions)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	t.Log(apply)

// 	//
// 	// Assert
// 	//
// 	output := terraform.Output(t, tfOptions, "output_var")

// 	assert.Equal(t, inputVariable, output)
// }
