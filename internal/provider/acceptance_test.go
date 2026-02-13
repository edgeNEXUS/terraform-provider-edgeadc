package provider

// Acceptance tests that run against a real EdgeADC appliance.
//
// These tests are SKIPPED by default (no ADC required for unit tests).
// To run them, set the following environment variables:
//
//   export EDGEADC_ENDPOINT=https://192.168.1.1
//   export EDGEADC_USERNAME=admin
//   export EDGEADC_PASSWORD=jetnexus
//   export TF_ACC=1
//
// Optional:
//   export EDGEADC_SKIP_CERT_VERIFY=true   # skip TLS certificate verification
//
// Then run:
//   go test ./internal/provider/... -v -run "TestAcc" -timeout 300s
//
// WARNING: These tests create and destroy real resources on the ADC.
//          Use a test/staging appliance, not production.

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// testAccProtoV6ProviderFactories configures the provider for acceptance tests.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"edgeadc": providerserver.NewProtocol6WithError(New()()),
}

// providerConfig returns the provider HCL block using env vars.
func providerConfig() string {
	skipCert := os.Getenv("EDGEADC_SKIP_CERT_VERIFY")
	if skipCert == "" {
		skipCert = "true"
	}
	return fmt.Sprintf(`
provider "edgeadc" {
  endpoint              = %q
  username              = %q
  password              = %q
  skip_cert_verification = %s
}
`, os.Getenv("EDGEADC_ENDPOINT"),
		os.Getenv("EDGEADC_USERNAME"),
		os.Getenv("EDGEADC_PASSWORD"),
		skipCert,
	)
}

// testAccPreCheck validates required env vars are set before running acceptance tests.
func testAccPreCheck(t *testing.T) {
	t.Helper()
	if os.Getenv("EDGEADC_ENDPOINT") == "" {
		t.Skip("EDGEADC_ENDPOINT not set — skipping acceptance test")
	}
	if os.Getenv("EDGEADC_USERNAME") == "" {
		t.Skip("EDGEADC_USERNAME not set — skipping acceptance test")
	}
	if os.Getenv("EDGEADC_PASSWORD") == "" {
		t.Skip("EDGEADC_PASSWORD not set — skipping acceptance test")
	}
}

// ============================================================================
// Realserver Monitor acceptance tests
// ============================================================================

// TestAccRealserverMonitor_CreateAndRead creates a monitor, verifies all
// attributes, then destroys it.
func TestAccRealserverMonitor_CreateAndRead(t *testing.T) {
	testAccPreCheck(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "edgeadc_realserver_monitor" "test" {
  name        = "tf-acc-test-monitor"
  description = "Acceptance test monitor"
  type        = "Check200"
  url         = "/health"
  content     = "200 OK"
  username    = ""
  password    = ""
  threshold   = "3"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.test", "name", "tf-acc-test-monitor"),
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.test", "type", "Check200"),
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.test", "threshold", "3"),
					resource.TestCheckResourceAttrSet("edgeadc_realserver_monitor.test", "id"),
				),
			},
		},
	})
}

// TestAccRealserverMonitor_FriendlyTypeName verifies that using a friendly
// UI name like "HTTP Head" works correctly — the monitor is created with
// the backend name "CheckHead" and functions properly.
func TestAccRealserverMonitor_FriendlyTypeName(t *testing.T) {
	testAccPreCheck(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "edgeadc_realserver_monitor" "friendly" {
  name        = "tf-acc-test-friendly"
  description = "Test friendly type name"
  type        = "HTTP Head"
  url         = "/"
  content     = ""
  username    = ""
  password    = ""
  threshold   = "3"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.friendly", "name", "tf-acc-test-friendly"),
					// The type should be normalized to the backend name
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.friendly", "type", "CheckHead"),
				),
			},
		},
	})
}

// TestAccRealserverMonitor_Update verifies updating a monitor's attributes
// doesn't produce inconsistent state.
func TestAccRealserverMonitor_Update(t *testing.T) {
	testAccPreCheck(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create
			{
				Config: providerConfig() + `
resource "edgeadc_realserver_monitor" "update_test" {
  name        = "tf-acc-test-update"
  description = "Before update"
  type        = "Check200"
  url         = "/health"
  content     = "200 OK"
  username    = ""
  password    = ""
  threshold   = "3"
}
`,
				Check: resource.TestCheckResourceAttr("edgeadc_realserver_monitor.update_test", "description", "Before update"),
			},
			// Step 2: Update description and threshold
			{
				Config: providerConfig() + `
resource "edgeadc_realserver_monitor" "update_test" {
  name        = "tf-acc-test-update"
  description = "After update"
  type        = "Check200"
  url         = "/health"
  content     = "200 OK"
  username    = ""
  password    = ""
  threshold   = "5"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.update_test", "description", "After update"),
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.update_test", "threshold", "5"),
				),
			},
		},
	})
}

// ============================================================================
// IP Services acceptance tests
// ============================================================================

// TestAccIpServices_CreateAndRead creates a virtual service and verifies
// no state drift on immediate re-plan.
func TestAccIpServices_CreateAndRead(t *testing.T) {
	testAccPreCheck(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "edgeadc_ip_services" "test" {
  ip_addr                   = "10.254.254.1"
  subnet_mask               = "255.255.255.255"
  service_name              = "tf-acc-test-vs"
  local_port_enabled_checked = "true"
  primary_checked           = "Active"
  service_type              = "HTTP"
  port                      = "19080"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_ip_services.test", "ip_addr", "10.254.254.1"),
					resource.TestCheckResourceAttr("edgeadc_ip_services.test", "port", "19080"),
					resource.TestCheckResourceAttr("edgeadc_ip_services.test", "primary_checked", "Active"),
					resource.TestCheckResourceAttr("edgeadc_ip_services.test", "service_type", "HTTP"),
					resource.TestCheckResourceAttrSet("edgeadc_ip_services.test", "id"),
				),
			},
		},
	})
}

// TestAccIpServices_PrimaryCheckedEmpty verifies that setting primary_checked
// to empty string does NOT cause an inconsistent state error. The API will
// return either "Active" or "Passive", and the provider must accept it.
func TestAccIpServices_PrimaryCheckedEmpty(t *testing.T) {
	testAccPreCheck(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "edgeadc_ip_services" "pc_test" {
  ip_addr                   = "10.254.254.2"
  subnet_mask               = "255.255.255.255"
  service_name              = "tf-acc-test-pc"
  local_port_enabled_checked = "true"
  primary_checked           = ""
  service_type              = "HTTP"
  port                      = "19081"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_ip_services.pc_test", "ip_addr", "10.254.254.2"),
					// primary_checked will be whatever the API returns — we just verify it's set
					resource.TestCheckResourceAttrSet("edgeadc_ip_services.pc_test", "primary_checked"),
				),
			},
		},
	})
}

// TestAccIpServices_Update verifies updating a VS doesn't produce
// inconsistent state errors.
func TestAccIpServices_Update(t *testing.T) {
	testAccPreCheck(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create
			{
				Config: providerConfig() + `
resource "edgeadc_ip_services" "update_test" {
  ip_addr                   = "10.254.254.3"
  subnet_mask               = "255.255.255.255"
  service_name              = "tf-acc-before-update"
  local_port_enabled_checked = "true"
  primary_checked           = "Active"
  service_type              = "HTTP"
  port                      = "19082"
}
`,
				Check: resource.TestCheckResourceAttr("edgeadc_ip_services.update_test", "service_name", "tf-acc-before-update"),
			},
			// Step 2: Update service_name
			{
				Config: providerConfig() + `
resource "edgeadc_ip_services" "update_test" {
  ip_addr                   = "10.254.254.3"
  subnet_mask               = "255.255.255.255"
  service_name              = "tf-acc-after-update"
  local_port_enabled_checked = "true"
  primary_checked           = "Active"
  service_type              = "HTTP"
  port                      = "19082"
}
`,
				Check: resource.TestCheckResourceAttr("edgeadc_ip_services.update_test", "service_name", "tf-acc-after-update"),
			},
		},
	})
}

// ============================================================================
// Server acceptance tests
// ============================================================================

// TestAccServer_CreateUpdateDestroy creates a VS, adds a server, updates
// the server, then destroys everything. This exercises the full lifecycle.
func TestAccServer_CreateUpdateDestroy(t *testing.T) {
	testAccPreCheck(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create VS + Server
			{
				Config: providerConfig() + `
resource "edgeadc_ip_services" "server_test_vs" {
  ip_addr                   = "10.254.254.4"
  subnet_mask               = "255.255.255.255"
  service_name              = "tf-acc-server-test"
  local_port_enabled_checked = "true"
  primary_checked           = "Active"
  service_type              = "HTTP"
  port                      = "19083"
}

resource "edgeadc_server" "test" {
  ip_service   = edgeadc_ip_services.server_test_vs.id
  cs_activity  = "1"
  cs_ip_addr   = "10.254.254.100"
  cs_port      = "8080"
  cs_notes     = "tf-acc-test-server"
  weight_factor = "100"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_server.test", "cs_ip_addr", "10.254.254.100"),
					resource.TestCheckResourceAttr("edgeadc_server.test", "cs_port", "8080"),
					resource.TestCheckResourceAttr("edgeadc_server.test", "weight_factor", "100"),
				),
			},
			// Step 2: Update weight_factor and notes
			{
				Config: providerConfig() + `
resource "edgeadc_ip_services" "server_test_vs" {
  ip_addr                   = "10.254.254.4"
  subnet_mask               = "255.255.255.255"
  service_name              = "tf-acc-server-test"
  local_port_enabled_checked = "true"
  primary_checked           = "Active"
  service_type              = "HTTP"
  port                      = "19083"
}

resource "edgeadc_server" "test" {
  ip_service   = edgeadc_ip_services.server_test_vs.id
  cs_activity  = "1"
  cs_ip_addr   = "10.254.254.100"
  cs_port      = "8080"
  cs_notes     = "tf-acc-updated-server"
  weight_factor = "75"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_server.test", "cs_notes", "tf-acc-updated-server"),
					resource.TestCheckResourceAttr("edgeadc_server.test", "weight_factor", "75"),
				),
			},
		},
	})
}

// ============================================================================
// Multi-resource acceptance tests
// ============================================================================

// TestAccMultiResource_FullLifecycle creates monitors, VS, and servers
// together, verifies no drift, then destroys. This is the closest test
// to a real customer workflow.
func TestAccMultiResource_FullLifecycle(t *testing.T) {
	testAccPreCheck(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "edgeadc_realserver_monitor" "acc_monitor" {
  name        = "tf-acc-full-monitor"
  description = "Full lifecycle test monitor"
  type        = "Check200"
  url         = "/health"
  content     = ""
  username    = ""
  password    = ""
  threshold   = "3"
}

resource "edgeadc_ip_services" "acc_vs" {
  ip_addr                   = "10.254.254.5"
  subnet_mask               = "255.255.255.255"
  service_name              = "tf-acc-full-test"
  local_port_enabled_checked = "true"
  primary_checked           = "Active"
  service_type              = "HTTP"
  port                      = "19084"
  server_monitoring         = edgeadc_realserver_monitor.acc_monitor.name
}

resource "edgeadc_server" "acc_server" {
  ip_service   = edgeadc_ip_services.acc_vs.id
  cs_activity  = "1"
  cs_ip_addr   = "10.254.254.200"
  cs_port      = "80"
  cs_notes     = "tf-acc-full-server"
  weight_factor = "100"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.acc_monitor", "type", "Check200"),
					resource.TestCheckResourceAttr("edgeadc_ip_services.acc_vs", "service_name", "tf-acc-full-test"),
					resource.TestCheckResourceAttr("edgeadc_server.acc_server", "cs_ip_addr", "10.254.254.200"),
				),
			},
		},
	})
}
