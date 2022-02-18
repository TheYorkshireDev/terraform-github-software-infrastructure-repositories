package main

import (
	"log"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestIT_InputOutputExample(t *testing.T) {
	inputVariable := "input"

	tfOptions := &terraform.Options{
		TerraformDir: "../examples/input-output",
		Vars: map[string]interface{}{
			"input_var": inputVariable,
		},
	}

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

	// Assertions
	output := terraform.Output(t, tfOptions, "output_var")

	assert.Equal(t, inputVariable, output)
}

func TestIT_InputOutputExample_Default(t *testing.T) {
	expectedVariable := "xxx"

	tfOptions := &terraform.Options{
		TerraformDir: "../examples/input-output",
	}

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

	// Assertions
	output := terraform.Output(t, tfOptions, "output_var")

	assert.Equal(t, expectedVariable, output)
}