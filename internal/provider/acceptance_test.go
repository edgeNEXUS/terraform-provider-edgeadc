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
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// testdataDir returns the absolute path to the testdata directory.
func testdataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..", "..", "testdata")
}

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
// the backend name on the API, but the user's friendly name is preserved
// in state to prevent perpetual diffs.
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
					// The user's friendly name is preserved in state
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.friendly", "type", "HTTP Head"),
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

// TestAccIpServices_PrimaryCheckedOmitted verifies that omitting primary_checked
// does NOT cause an inconsistent state error. The API will return either
// "Active" or "Passive", and the provider must accept it because the field
// is Optional+Computed with UseStateForUnknown.
func TestAccIpServices_PrimaryCheckedOmitted(t *testing.T) {
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

// TestAccIpServices_PrimaryCheckedEmptyString reproduces the exact customer
// error: setting primary_checked = "" causes "Provider produced inconsistent
// result after apply" because the API returns "Active". The
// emptyStringToActiveModifier should convert "" to "Active" during planning.
func TestAccIpServices_PrimaryCheckedEmptyString(t *testing.T) {
	testAccPreCheck(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + `
resource "edgeadc_ip_services" "pc_empty" {
  ip_addr                   = "10.254.254.6"
  subnet_mask               = "255.255.255.255"
  service_name              = "tf-acc-test-pc-empty"
  local_port_enabled_checked = "true"
  service_type              = "HTTP"
  port                      = "19086"
  primary_checked           = ""
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_ip_services.pc_empty", "ip_addr", "10.254.254.6"),
					// The CRUD handler preserves "" in state even though the API
					// returns "Active". This matches the plan value, avoiding
					// inconsistency errors.
					resource.TestCheckResourceAttr("edgeadc_ip_services.pc_empty", "primary_checked", ""),
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

// ============================================================================
// Users acceptance tests
// ============================================================================

// TestAccUsers_CreateReadUpdateDelete exercises the full CRUD lifecycle
// for the edgeadc_users resource.
func TestAccUsers_CreateReadUpdateDelete(t *testing.T) {
	testAccPreCheck(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create user
			{
				Config: providerConfig() + `
resource "edgeadc_users" "test" {
  user_name    = "tf_acc_test_user"
  new_password = "TestPass123!"
  is_admin     = "0"
  is_api       = "1"
  is_guir      = "0"
  is_guiw      = "0"
  is_ssh       = "0"
  is_add_on    = "0"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_users.test", "user_name", "tf_acc_test_user"),
					resource.TestCheckResourceAttr("edgeadc_users.test", "is_api", "1"),
					resource.TestCheckResourceAttr("edgeadc_users.test", "is_admin", "0"),
				),
			},
			// Step 2: Update permissions
			{
				Config: providerConfig() + `
resource "edgeadc_users" "test" {
  user_name    = "tf_acc_test_user"
  new_password = "TestPass123!"
  is_admin     = "0"
  is_api       = "1"
  is_guir      = "1"
  is_guiw      = "1"
  is_ssh       = "0"
  is_add_on    = "0"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_users.test", "is_guir", "1"),
					resource.TestCheckResourceAttr("edgeadc_users.test", "is_guiw", "1"),
				),
			},
		},
	})
}

// ============================================================================
// SSL Certificates acceptance tests
// ============================================================================

// TestAccSslCertificates_CreateAndDelete creates an SSL certificate from a
// PFX file, verifies it exists, then destroys it.
func TestAccSslCertificates_CreateAndDelete(t *testing.T) {
	testAccPreCheck(t)

	pfxPath := filepath.Join(testdataDir(), "test-cert.pfx")
	if _, err := os.Stat(pfxPath); os.IsNotExist(err) {
		t.Skipf("test PFX file not found at %s — skipping", pfxPath)
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + fmt.Sprintf(`
resource "edgeadc_ssl_certificates" "test" {
  id        = "tf-acc-test-cert"
  file_path = %q
  password  = "testpass123"
}
`, pfxPath),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_ssl_certificates.test", "id", "tf-acc-test-cert"),
				),
			},
		},
	})
}

// ============================================================================
// Custom Monitor acceptance tests
// ============================================================================

// TestAccCustomMonitor_CreateAndDelete creates a custom monitor from a
// Perl script, verifies it exists, then destroys it.
func TestAccCustomMonitor_CreateAndDelete(t *testing.T) {
	testAccPreCheck(t)

	plPath := filepath.Join(testdataDir(), "test-monitor.pl")
	if _, err := os.Stat(plPath); os.IsNotExist(err) {
		t.Skipf("test Perl file not found at %s — skipping", plPath)
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig() + fmt.Sprintf(`
resource "edgeadc_custom_monitor" "test" {
  name      = "tf-acc-test-csm"
  file_path = %q
}
`, plPath),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_custom_monitor.test", "name", "tf-acc-test-csm"),
					resource.TestCheckResourceAttrSet("edgeadc_custom_monitor.test", "id"),
				),
			},
		},
	})
}


// ============================================================================
// Multi-monitor tests (reproduces customer ID renumbering bug)
// ============================================================================

// TestAccMultiMonitor_CreateDeleteNoDrift creates 3 monitors, then removes
// the middle one. This reproduces the customer bug where the ADC renumbers
// monitor IDs after deletion, causing "was cty.StringVal("6"), but now
// cty.StringVal("5")" errors.
func TestAccMultiMonitor_CreateDeleteNoDrift(t *testing.T) {
	testAccPreCheck(t)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create 3 monitors
			{
				Config: providerConfig() + `
resource "edgeadc_realserver_monitor" "mon_a" {
  name        = "tf-acc-multi-mon-a"
  description = "Monitor A"
  type        = "Check200"
  url         = "/a"
  content     = ""
  username    = ""
  password    = ""
  threshold   = "3"
}

resource "edgeadc_realserver_monitor" "mon_b" {
  name        = "tf-acc-multi-mon-b"
  description = "Monitor B"
  type        = "CheckHead"
  url         = "/b"
  content     = ""
  username    = ""
  password    = ""
  threshold   = "3"
}

resource "edgeadc_realserver_monitor" "mon_c" {
  name        = "tf-acc-multi-mon-c"
  description = "Monitor C"
  type        = "Connect"
  url         = ""
  content     = ""
  username    = ""
  password    = ""
  threshold   = "3"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.mon_a", "name", "tf-acc-multi-mon-a"),
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.mon_b", "name", "tf-acc-multi-mon-b"),
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.mon_c", "name", "tf-acc-multi-mon-c"),
				),
			},
			// Step 2: Remove monitor B — this triggers ID renumbering on the ADC
			{
				Config: providerConfig() + `
resource "edgeadc_realserver_monitor" "mon_a" {
  name        = "tf-acc-multi-mon-a"
  description = "Monitor A"
  type        = "Check200"
  url         = "/a"
  content     = ""
  username    = ""
  password    = ""
  threshold   = "3"
}

resource "edgeadc_realserver_monitor" "mon_c" {
  name        = "tf-acc-multi-mon-c"
  description = "Monitor C"
  type        = "Connect"
  url         = ""
  content     = ""
  username    = ""
  password    = ""
  threshold   = "3"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.mon_a", "name", "tf-acc-multi-mon-a"),
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.mon_c", "name", "tf-acc-multi-mon-c"),
				),
			},
		},
	})
}

// ============================================================================
// No-drift re-plan tests
// ============================================================================

// TestAccNoDrift_ReplanAfterApply applies a config with multiple resources,
// then re-applies the EXACT same config. The second step must produce zero
// changes — any drift means the provider is not correctly reading back state.
func TestAccNoDrift_ReplanAfterApply(t *testing.T) {
	testAccPreCheck(t)

	config := providerConfig() + `
resource "edgeadc_realserver_monitor" "drift_mon" {
  name        = "tf-acc-drift-monitor"
  description = "Drift test monitor"
  type        = "HTTP 200 OK"
  url         = "/drift"
  content     = ""
  username    = ""
  password    = ""
  threshold   = "3"
}

resource "edgeadc_ip_services" "drift_vs" {
  ip_addr                   = "10.254.254.20"
  subnet_mask               = "255.255.255.255"
  service_name              = "tf-acc-drift-test"
  local_port_enabled_checked = "true"
  primary_checked           = "Active"
  service_type              = "HTTP"
  port                      = "19095"
  server_monitoring         = edgeadc_realserver_monitor.drift_mon.name
}

resource "edgeadc_server" "drift_srv" {
  ip_service   = edgeadc_ip_services.drift_vs.id
  cs_activity  = "1"
  cs_ip_addr   = "10.254.254.220"
  cs_port      = "80"
  cs_notes     = "tf-acc-drift-server"
  weight_factor = "100"
}
`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Apply
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.drift_mon", "type", "HTTP 200 OK"),
					resource.TestCheckResourceAttr("edgeadc_ip_services.drift_vs", "service_name", "tf-acc-drift-test"),
					resource.TestCheckResourceAttr("edgeadc_server.drift_srv", "cs_ip_addr", "10.254.254.220"),
				),
			},
			// Step 2: Re-apply same config — must produce zero changes
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("edgeadc_realserver_monitor.drift_mon", "type", "HTTP 200 OK"),
					resource.TestCheckResourceAttr("edgeadc_ip_services.drift_vs", "service_name", "tf-acc-drift-test"),
					resource.TestCheckResourceAttr("edgeadc_server.drift_srv", "cs_ip_addr", "10.254.254.220"),
				),
			},
		},
	})
}