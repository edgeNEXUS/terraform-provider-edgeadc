
# EdgeADC Terraform Provider

Terraform provider for EdgeADC Device - [EdgeNexus](https://www.edgenexus.io/)

## Documentation
- [Terraform Resource Documentation](/docs/index.md)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 1.0+

## Using the provider

To use the latest version of the provider in your Terraform environment, run `terraform init` and Terraform will automatically install the provider.

If you wish to pin your environment to a specific release of the provider, you can do so with a `required_providers` statement in your Terraform manifest. The `terraform` [configuration block](https://www.terraform.io/docs/configuration/provider-requirements.html) varies slightly depending on which Terraform version you're using. See below for more examples of configuring the provider version for the different versions of Terraform.

For Terraform version 1.x and above

```hcl
terraform {
  required_version = "~> 1.0"
  required_providers {
    newrelic = {
      source  = "edgeNEXUS/edgeadc"
    }
  }
}
```

For more information on using the provider and the associated resources, please see the [provider documentation][provider_docs] page.

# Getting Started

Authentication
```hcl
provider "edgeadc" {
  username = "admin"
  password = "jetnexus"
  endpoint = "http://192.168.1.1:8080"
}
```

## Examples
- [See examples directory](/examples/main.tf)

# Resource Support
The following resources are currently supported:
- `ip_services` - [Documentation](/docs/resources/ip_services.md)
- `servers` - [Documentation](/docs/resources/server.md)
- `ssl_certificates` - [Documentation](/docs/resources/ssl_certificates.md)
- `users` - [Documentation](/docs/resources/users.md)

# Importing Existing Resources
See the documentation for each resource for information on importing existing resources.
`ip_services` can be imported by using `ip:port` as the ID.
```bash
terraform import edgeadc_ip_services.test_ip_service_1 192.168.1.101:101
```

## Contributing
- [See contributing guide](/contributing)