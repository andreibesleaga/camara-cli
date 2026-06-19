// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/andreibesleaga/camara-cli/internal/mocktest"
	"github.com/andreibesleaga/camara-cli/internal/requestflag"
)

func TestConnectivityinsightsSubscriptionsCreate(t *testing.T) {
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
			"connectivityinsights:subscriptions", "create",
			"--config", "{subscriptionDetail: {applicationProfileId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, device: {ipv4Address: {privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344, networkAccessIdentifier: 123456789@domain.com, phoneNumber: '+123456789'}, applicationServer: {ipv4Address: 192.168.0.1/24, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344}, applicationServerPorts: {ports: [5060, 5070], ranges: [{from: 5010, to: 5020}]}}, initialEvent: true, subscriptionExpireTime: '2023-07-03T12:27:08.312Z', subscriptionMaxEvents: 5}",
			"--protocol", "HTTP",
			"--sink", "https://endpoint.example.com/sink",
			"--type", "org.camaraproject.connectivity-insights-subscriptions.v0.network-quality",
			"--sink-credential", "{credentialType: PLAIN}",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(connectivityinsightsSubscriptionsCreate)

		// Alternative argument passing style using inner flags
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
			"connectivityinsights:subscriptions", "create",
			"--config.subscription-detail", "{applicationProfileId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, device: {ipv4Address: {privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344, networkAccessIdentifier: 123456789@domain.com, phoneNumber: '+123456789'}, applicationServer: {ipv4Address: 192.168.0.1/24, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344}, applicationServerPorts: {ports: [5060, 5070], ranges: [{from: 5010, to: 5020}]}}",
			"--config.initial-event=true",
			"--config.subscription-expire-time", "2023-07-03T12:27:08.312Z",
			"--config.subscription-max-events", "5",
			"--protocol", "HTTP",
			"--sink", "https://endpoint.example.com/sink",
			"--type", "org.camaraproject.connectivity-insights-subscriptions.v0.network-quality",
			"--sink-credential.credential-type", "PLAIN",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"config:\n" +
			"  subscriptionDetail:\n" +
			"    applicationProfileId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"    device:\n" +
			"      ipv4Address:\n" +
			"        privateAddress: 84.125.93.10\n" +
			"        publicAddress: 84.125.93.10\n" +
			"        publicPort: 59765\n" +
			"      ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344\n" +
			"      networkAccessIdentifier: 123456789@domain.com\n" +
			"      phoneNumber: '+123456789'\n" +
			"    applicationServer:\n" +
			"      ipv4Address: 192.168.0.1/24\n" +
			"      ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344\n" +
			"    applicationServerPorts:\n" +
			"      ports:\n" +
			"        - 5060\n" +
			"        - 5070\n" +
			"      ranges:\n" +
			"        - from: 5010\n" +
			"          to: 5020\n" +
			"  initialEvent: true\n" +
			"  subscriptionExpireTime: '2023-07-03T12:27:08.312Z'\n" +
			"  subscriptionMaxEvents: 5\n" +
			"protocol: HTTP\n" +
			"sink: https://endpoint.example.com/sink\n" +
			"types:\n" +
			"  - org.camaraproject.connectivity-insights-subscriptions.v0.network-quality\n" +
			"sinkCredential:\n" +
			"  credentialType: PLAIN\n")
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
			"connectivityinsights:subscriptions", "create",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})
}

func TestConnectivityinsightsSubscriptionsRetrieve(t *testing.T) {
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
			"connectivityinsights:subscriptions", "retrieve",
			"--subscription-id", "qs15-h556-rt89-1298",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})
}

func TestConnectivityinsightsSubscriptionsList(t *testing.T) {
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
			"connectivityinsights:subscriptions", "list",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})
}

func TestConnectivityinsightsSubscriptionsDelete(t *testing.T) {
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
			"connectivityinsights:subscriptions", "delete",
			"--subscription-id", "qs15-h556-rt89-1298",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})
}
