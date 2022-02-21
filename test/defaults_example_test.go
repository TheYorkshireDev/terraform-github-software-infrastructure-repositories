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
		TerraformDir: "../examples/defaults",
		Vars: map[string]interface{}{
			"required_name": repositoryName,
		},
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
	plan, err := terraform.PlanE(t, tfOptions)

	if err != nil {
		log.Println(err)
	}

	t.Log(plan)

	// Run `terraform apply`
	apply, err := terraform.ApplyE(t, tfOptions)

	if err != nil {
		log.Println(err)
	}

	t.Log(apply)

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
