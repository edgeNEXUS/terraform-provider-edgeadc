# Troubleshooting

See - https://developer.hashicorp.com/terraform/internals/debugging

## Debugging Terraform
This provider exposes an environment variable that can be used to enable debug logging for the provider.

This provider logs http request URI, request body and response body at the debug level.

To enable debug logging, set the `TF_LOG_PROVIDER` environment variable to `DEBUG` before running Terraform.

```bash
set TF_LOG_PROVIDER=DEBUG
```

You can also configure the provider to write the logs to disk
```bash
set TF_LOG_PATH=terraform-provider-edgeadc_output.log
```
