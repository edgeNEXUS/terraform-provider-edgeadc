# Local Development
See - https://developer.hashicorp.com/terraform/cli/config/config-file

# Prerequisites
* Go 1.22 or later
* Terraform binary

# Steps
* Create a `tf.tfrc` file or `tf-windows.tfrc` file with the name of the provider and your go path.
```hcl
provider_installation {

  dev_overrides {
    # Example GOBIN path, will need to be replaced with your own GOBIN path. Default is $GOPATH/bin
    "registry.terraform.io/edgeNEXUS/edgeadc" = "/home/USER/gopath/bin"
    "edgeNEXUS/edgeadc" = "/home/USER/gopath/bin"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

* Set the `TF_CLI_CONFIG_FILE` environment variable to the path of the `tf.tfrc` file.
```bash
export TF_CLI_CONFIG_FILE=/tmp/terraform-provider-edgeadc/tf.tfrc
```

* Build and install the provider into your go bin directory
```bash
go install .
```
* Run the provider (there is no need `terraform init` as the provider is installed in the go bin directory)
```bash
terraform plan
terraform apply
```
