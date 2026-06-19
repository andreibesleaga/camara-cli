// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
)

func TestKnowyourcustomerageverificationVerify(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--bearer-token", "string",
			"--customer-insights-token", "string",
			"--device-swap-token", "string",
			"--kyc-age-verification-token", "string",
			"--kyc-fill-in-token", "string",
			"--kyc-match-token", "string",
			"--tenure-token", "string",
			"--number-recycling-token", "string",
			"--otp-validation-token", "string",
			"--call-forwarding-signal-token", "string",
			"--device-location-token", "string",
			"--population-density-data-token", "string",
			"--region-device-count-token", "string",
			"--web-rtc-token", "string",
			"--connectivity-insights-token", "string",
			"--quality-on-demand-token", "string",
			"--device-identifier-token", "string",
			"--sim-swap-token", "string",
			"--device-roaming-status-token", "string",
			"--device-reachability-status-token", "string",
			"--connected-network-type-token", "string",
			"--device-location-notifications-api-key", "string",
			"--notifications-api-key", "string",
			"--population-density-data-notifications-api-key", "string",
			"--region-device-count-notifications-api-key", "string",
			"--connectivity-insights-notifications-api-key", "string",
			"--sim-swap-notifications-api-key", "string",
			"--device-roaming-status-notifications-api-key", "string",
			"--device-reachability-status-notifications-api-key", "string",
			"--connected-network-type-notifications-api-key", "string",
			"knowyourcustomerageverification", "verify",
			"--age-threshold", "18",
			"--birthdate", "'1978-08-22'",
			"--email", "federicaSanchez.Arjona@example.com",
			"--family-name", "Sanchez Arjona",
			"--family-name-at-birth", "YYYY",
			"--given-name", "Federica",
			"--id-document", "66666666q",
			"--include-content-lock=true",
			"--include-parental-control=true",
			"--middle-names", "Sanchez",
			"--name", "Federica Sanchez Arjona",
			"--phone-number", "+34629255833",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"ageThreshold: 18\n" +
			"birthdate: '1978-08-22'\n" +
			"email: federicaSanchez.Arjona@example.com\n" +
			"familyName: Sanchez Arjona\n" +
			"familyNameAtBirth: YYYY\n" +
			"givenName: Federica\n" +
			"idDocument: 66666666q\n" +
			"includeContentLock: true\n" +
			"includeParentalControl: true\n" +
			"middleNames: Sanchez\n" +
			"name: Federica Sanchez Arjona\n" +
			"phoneNumber: '+34629255833'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--bearer-token", "string",
			"--customer-insights-token", "string",
			"--device-swap-token", "string",
			"--kyc-age-verification-token", "string",
			"--kyc-fill-in-token", "string",
			"--kyc-match-token", "string",
			"--tenure-token", "string",
			"--number-recycling-token", "string",
			"--otp-validation-token", "string",
			"--call-forwarding-signal-token", "string",
			"--device-location-token", "string",
			"--population-density-data-token", "string",
			"--region-device-count-token", "string",
			"--web-rtc-token", "string",
			"--connectivity-insights-token", "string",
			"--quality-on-demand-token", "string",
			"--device-identifier-token", "string",
			"--sim-swap-token", "string",
			"--device-roaming-status-token", "string",
			"--device-reachability-status-token", "string",
			"--connected-network-type-token", "string",
			"--device-location-notifications-api-key", "string",
			"--notifications-api-key", "string",
			"--population-density-data-notifications-api-key", "string",
			"--region-device-count-notifications-api-key", "string",
			"--connectivity-insights-notifications-api-key", "string",
			"--sim-swap-notifications-api-key", "string",
			"--device-roaming-status-notifications-api-key", "string",
			"--device-reachability-status-notifications-api-key", "string",
			"--connected-network-type-notifications-api-key", "string",
			"knowyourcustomerageverification", "verify",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})
}
