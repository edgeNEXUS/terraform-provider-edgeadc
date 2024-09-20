# Example Terraform configuration file for the edgeNEXUS EdgeADC provider

terraform {
  required_providers {
    edgeadc = {
      source = "edgeNEXUS/edgeadc"
    }
  }
}

# EdgeNexus EdgeADC Configuration
provider "edgeadc" {
  username = "admin"
  password = "jetnexus"
  endpoint = "http://192.168.1.1:8080"
}

# Create a Virtual Service
resource "edgeadc_ip_services" "test_ip_service_1" {
    ip_addr = "192.168.1.101"
    subnet_mask = "255.255.255.255"
    service_name = "test service 1"
    local_port_enabled_checked = "true"
    primary_checked = "Active"
    service_type = "HTTP"
    port = "101"
}

# Create a Server and associate it with the Virtual Service
resource "edgeadc_server" "example_server_1" {
    ip_service = edgeadc_ip_services.test_ip_service_1.id
	cs_activity = "1"
	cs_ip_addr = "192.168.2.101"
	cs_port = "111"
	cs_notes = "test server 1"
	weight_factor = "100"
}

# Create a Certificate
resource "edgeadc_ssl_certificates" "example_certificate" {
    id = "TestTerraformCertificate"
    file_path = "c:\\temp\\sslcert_TerraformTest.pfx"
    password = "123456"
}

# Create a second Virtual Service
# With SSL and advanced settings
resource "edgeadc_ip_services" "test_ip_service_2" {
    ip_addr = "192.168.1.102"
    subnet_mask = "255.255.255.255"
    service_name = "test service 2"
    local_port_enabled_checked = "true"
    primary_checked = "Active"
    service_type = "HTTP"
    port = "102"
    # Advanced settings
    ssl_certificate = edgeadc_ssl_certificates.example_certificate.id
}

# Create a Server and associate it with the second Virtual Service
resource "edgeadc_server" "example_server_2" {
    ip_service = edgeadc_ip_services.test_ip_service_2.id
	cs_activity = "1"
	cs_ip_addr = "192.168.2.102"
	cs_port = "222"
	cs_notes = "test server 2"
	weight_factor = "100"
}

# Create an API user
resource "edgeadc_users" "test_api_user" {
    user_name =  "test_api_user"
	is_add_on = "0"
	is_admin = "0"
	is_api = "1"
	is_guir = "0"
	is_guiw = "0"
	is_ssh = "0"
	new_password = "123456"
}
