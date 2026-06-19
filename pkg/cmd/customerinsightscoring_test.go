// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
)

func TestCustomerinsightsScoringRetrieve(t *testing.T) {
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
			"customerinsights:scoring", "retrieve",
			"--id-document", "idDocument",
			"--phone-number", "+4960513",
			"--scoring-type", "gaugeMetric",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"idDocument: idDocument\n" +
			"phoneNumber: '+4960513'\n" +
			"scoringType: gaugeMetric\n")
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
			"customerinsights:scoring", "retrieve",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})
}
