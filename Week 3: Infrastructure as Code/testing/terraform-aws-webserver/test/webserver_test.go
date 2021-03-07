package test

import (
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

//The values to pass into the terrafrom CLI
func TestTerraformWebserverExample(t *testing.T) {

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		//The path where terrafrom code is located
		TerraformDir: "../examples/webserver",

		//variables to pass to terrafrom code
		Vars: map[string]interface{}{
			"region":     "eu-central-1",
			"servername": "testwebserver",
		},
	})

	//run terraform init and apply options
	terraform.InitAndApply(t, terraformOptions)

	//run teraform destroy at the end
	defer terraform.Destroy(t, terraformOptions)

	publicIP := terraform.Output(t, terraformOptions, "public_ip")

	url := fmt.Sprintf("http://%s:8080", publicIP)

	http_helper.HttpGetWithRetry(t, url, nil, 200, "I made a Terrafrom Module!", 30, 5*time.Second)

}
