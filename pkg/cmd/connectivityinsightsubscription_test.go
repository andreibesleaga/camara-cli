// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
	"github.com/stainless-sdks/camara-cli/internal/requestflag"
)

func TestConnectivityinsightsSubscriptionsCreate(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"connectivityinsights:subscriptions", "create",
		"--config", "{subscriptionDetail: {applicationProfileId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, device: {ipv4Address: {privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344, networkAccessIdentifier: 123456789@domain.com, phoneNumber: '+123456789'}, applicationServer: {ipv4Address: 192.168.0.1/24, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344}, applicationServerPorts: {ports: [5060, 5070], ranges: [{from: 5010, to: 5020}]}}, initialEvent: true, subscriptionExpireTime: '2023-07-03T12:27:08.312Z', subscriptionMaxEvents: 5}",
		"--protocol", "HTTP",
		"--sink", "https://endpoint.example.com/sink",
		"--type", "org.camaraproject.connectivity-insights-subscriptions.v0.network-quality",
		"--sink-credential", "{credentialType: PLAIN}",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(connectivityinsightsSubscriptionsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"connectivityinsights:subscriptions", "create",
		"--config.subscriptionDetail", "{applicationProfileId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, device: {ipv4Address: {privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344, networkAccessIdentifier: 123456789@domain.com, phoneNumber: '+123456789'}, applicationServer: {ipv4Address: 192.168.0.1/24, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344}, applicationServerPorts: {ports: [5060, 5070], ranges: [{from: 5010, to: 5020}]}}",
		"--config.initialEvent=true",
		"--config.subscriptionExpireTime", "2023-07-03T12:27:08.312Z",
		"--config.subscriptionMaxEvents", "5",
		"--protocol", "HTTP",
		"--sink", "https://endpoint.example.com/sink",
		"--type", "org.camaraproject.connectivity-insights-subscriptions.v0.network-quality",
		"--sink-credential.credentialType", "PLAIN",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestConnectivityinsightsSubscriptionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"connectivityinsights:subscriptions", "retrieve",
		"--subscription-id", "qs15-h556-rt89-1298",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestConnectivityinsightsSubscriptionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"connectivityinsights:subscriptions", "list",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestConnectivityinsightsSubscriptionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"connectivityinsights:subscriptions", "delete",
		"--subscription-id", "qs15-h556-rt89-1298",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
