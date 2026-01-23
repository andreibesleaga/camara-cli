// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
	"github.com/stainless-sdks/camara-cli/internal/requestflag"
)

func TestConnectednetworktypeSubscriptionsCreate(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"connectednetworktype:subscriptions", "create",
		"--config", "{subscriptionDetail: {device: {ipv4Address: {privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344, networkAccessIdentifier: 123456789@domain.com, phoneNumber: '+123456789'}}, initialEvent: true, subscriptionExpireTime: '2023-01-17T13:18:23.682Z', subscriptionMaxEvents: 5}",
		"--protocol", "HTTP",
		"--sink", "https://endpoint.example.com/sink",
		"--type", "org.camaraproject.connected-network-type-subscriptions.v0.network-type-changed",
		"--sink-credential", "{credentialType: ACCESSTOKEN}",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(connectednetworktypeSubscriptionsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"connectednetworktype:subscriptions", "create",
		"--config.subscriptionDetail", "{device: {ipv4Address: {privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344, networkAccessIdentifier: 123456789@domain.com, phoneNumber: '+123456789'}}",
		"--config.initialEvent=true",
		"--config.subscriptionExpireTime", "2023-01-17T13:18:23.682Z",
		"--config.subscriptionMaxEvents", "5",
		"--protocol", "HTTP",
		"--sink", "https://endpoint.example.com/sink",
		"--type", "org.camaraproject.connected-network-type-subscriptions.v0.network-type-changed",
		"--sink-credential.credentialType", "ACCESSTOKEN",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestConnectednetworktypeSubscriptionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"connectednetworktype:subscriptions", "retrieve",
		"--subscription-id", "qs15-h556-rt89-1298",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestConnectednetworktypeSubscriptionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"connectednetworktype:subscriptions", "list",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestConnectednetworktypeSubscriptionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"connectednetworktype:subscriptions", "delete",
		"--subscription-id", "qs15-h556-rt89-1298",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
