package main

import (
	"log"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformModule(t *testing.T) {
	options := &terraform.Options{
		TerraformDir: "../",
	}

	defer terraform.Destroy(t, options)

	// Run `terraform init`
	init, err := terraform.InitE(t, options)

	if err != nil {
		log.Println(err)
	}

	t.Log(init)

	// Run `terraform plan`
	plan, err := terraform.PlanE(t, options)

	if err != nil {
		log.Println(err)
	}

	t.Log(plan)

	// Run `terraform apply`
	apply, err := terraform.ApplyE(t, options)

	if err != nil {
		log.Println(err)
	}

	t.Log(apply)
}
