# EdgeADC Terraform Provider - Testing Guide

## Test Structure

The provider has two categories of tests:

| Category | Count | ADC Required | Run Time | Purpose |
|----------|-------|-------------|----------|---------|
| **Unit tests** | 43 | No | < 1 second | Test logic: type mapping, merge functions, error handling, round-trips |
| **Acceptance tests** | 8 | Yes | ~2-5 minutes | Test against a real ADC: create, read, update, destroy all resource types |

Unit tests run automatically in CI on every push and PR. Acceptance tests are
skipped in CI and must be run manually against a real EdgeADC appliance.

---

## Running Unit Tests (no ADC needed)

```bash
# Run all unit tests
go test ./internal/provider/... -v -count=1

# Run a specific test
go test ./internal/provider/... -v -run "TestNormalizeMonitorType"

# Run only the integration-style mock tests
go test ./internal/provider/... -v -run "IntegrationStyle"
```

These tests use mock HTTP responses and never make network calls.

---

## Running Acceptance Tests (real ADC)

### Prerequisites

1. A test/staging EdgeADC appliance (do **not** use production)
2. Admin credentials for the appliance
3. Terraform CLI installed (the test framework downloads it automatically, but
   having it locally helps for debugging)
4. Go 1.24+ installed

### Configuration

Set these environment variables:

```bash
export EDGEADC_ENDPOINT=https://192.168.1.1    # Your ADC URL (with port if needed)
export EDGEADC_USERNAME=admin                    # ADC admin username
export EDGEADC_PASSWORD=jetnexus                 # ADC admin password
export TF_ACC=1                                  # Enable acceptance tests
export EDGEADC_SKIP_CERT_VERIFY=true             # Optional: skip TLS cert verification
```

### Running

```bash
# Run ALL acceptance tests
go test ./internal/provider/... -v -run "TestAcc" -timeout 300s

# Run only monitor tests
go test ./internal/provider/... -v -run "TestAccRealserverMonitor" -timeout 300s

# Run only IP services tests
go test ./internal/provider/... -v -run "TestAccIpServices" -timeout 300s

# Run only server tests
go test ./internal/provider/... -v -run "TestAccServer" -timeout 300s

# Run the full multi-resource lifecycle test
go test ./internal/provider/... -v -run "TestAccMultiResource" -timeout 300s

# Run everything (unit + acceptance)
go test ./internal/provider/... -v -count=1 -timeout 300s
```

### What the acceptance tests do

Each test creates real resources on the ADC, verifies the attributes, and
destroys them when finished (even if the test fails). The test framework
handles cleanup automatically.

| Test | What It Does |
|------|-------------|
| `TestAccRealserverMonitor_CreateAndRead` | Create a Check200 monitor, verify attributes, destroy |
| `TestAccRealserverMonitor_FriendlyTypeName` | Create a monitor using `type = "HTTP Head"`, verify it's normalised to `CheckHead` |
| `TestAccRealserverMonitor_Update` | Create a monitor, update description and threshold, verify no errors |
| `TestAccIpServices_CreateAndRead` | Create a VS with `primary_checked = "Active"`, verify attributes, destroy |
| `TestAccIpServices_PrimaryCheckedEmpty` | Create a VS with `primary_checked = ""`, verify no inconsistency error |
| `TestAccIpServices_Update` | Create a VS, update service_name, verify no errors |
| `TestAccServer_CreateUpdateDestroy` | Create VS + server, update weight_factor and notes, verify values persisted |
| `TestAccMultiResource_FullLifecycle` | Create monitor + VS (with server_monitoring) + server together, verify all attributes |

### Test IP addresses

The acceptance tests use IPs in the `10.254.254.0/24` range and ports
`19080-19084` to avoid conflicts with existing ADC configuration. If these
conflict with your environment, edit the test configs in
`internal/provider/acceptance_test.go`.

### Troubleshooting

**Tests hang or time out:** The ADC may be unreachable. Check
`EDGEADC_ENDPOINT` and ensure the appliance is running.

**"Session Expired" errors:** The provider handles session re-authentication
automatically. If tests still fail with session errors, restart the ADC
management interface.

**Resources left behind after a test crash:** If a test is interrupted (Ctrl+C)
before cleanup, you may have leftover resources on the ADC. Look for resources
with names starting with `tf-acc-` and delete them manually.

**TLS certificate errors:** Set `EDGEADC_SKIP_CERT_VERIFY=true` if the ADC
uses a self-signed certificate.

---

## Unit Test Coverage

The unit tests cover every customer-reported bug:

| Bug | Test(s) |
|-----|---------|
| `"HTTP Head"` creates non-functional monitor | `TestNormalizeMonitorType_FriendlyToBackend`, `TestToRealserverMonitorOpt_NormalizesType` |
| `primary_checked: was "" but now "Passive"/"Active"` | `TestPrimaryChecked_EmptyStringBecomesUnknown`, `TestMergeBasicTabs_SkipsNullValues` |
| Server update sends empty fields | `TestToUpdateServer_AllFieldsMapped` |
| `server_monitoring` conversion silently fails | `TestConvertComboOptionsToIds_UnknownName`, `TestConvertComboOptionsToNames_UnknownId` |
| Monitor ID shifts between applies | `TestRealserverMonitorSchema_IdIsComputedWithoutUseStateForUnknown` |
| `terraform destroy` fails after failed apply | `TestDeleteServer_NotFound_IntegrationStyle`, `TestDeleteIPService_NotFound_IntegrationStyle`, `TestDeleteServer_ErrorPrefixMatching` |
| Server with hostname not found on delete | `TestGetServerByAddressAndPortFromIpServices_HostnameNotFound` |

---

## CI Configuration

The `.github/workflows/test.yml` workflow runs on every push to `main` and
every PR:

```yaml
steps:
  - go build ./...      # Compile check
  - go vet ./...        # Static analysis
  - go test ./... -v    # Unit tests (acceptance tests auto-skip)
```

Acceptance tests are **not** run in CI because they require a real ADC. They
must be run manually before each release.
