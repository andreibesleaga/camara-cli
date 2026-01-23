// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
	"github.com/stainless-sdks/camara-cli/internal/requestflag"
)

func TestDevicelocationSubscriptionsCreate(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"devicelocation:subscriptions", "create",
		"--config", "{initialEvent: true, subscriptionExpireTime: '2024-03-22T05:40:58.469Z', subscriptionMaxEvents: 10, subscriptionDetail: {area: {areaType: CIRCLE}, device: {ipv4Address: {privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344, networkAccessIdentifier: 123456789@domain.com, phoneNumber: '+12345678912'}}}",
		"--protocol", "HTTP",
		"--sink", "https://notificationSendServer12.supertelco.com",
		"--type", "org.camaraproject.geofencing-subscriptions.v0.area-entered",
		"--sink-credential", "{credentialType: PLAIN}",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(devicelocationSubscriptionsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"devicelocation:subscriptions", "create",
		"--config", "{initialEvent: true, subscriptionExpireTime: '2024-03-22T05:40:58.469Z', subscriptionMaxEvents: 10, subscriptionDetail: {area: {areaType: CIRCLE}, device: {ipv4Address: {privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344, networkAccessIdentifier: 123456789@domain.com, phoneNumber: '+12345678912'}}}",
		"--protocol", "HTTP",
		"--sink", "https://notificationSendServer12.supertelco.com",
		"--type", "org.camaraproject.geofencing-subscriptions.v0.area-entered",
		"--sink-credential.credentialType", "PLAIN",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestDevicelocationSubscriptionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"devicelocation:subscriptions", "retrieve",
		"--subscription-id", "qs15-h556-rt89-1298",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestDevicelocationSubscriptionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"devicelocation:subscriptions", "list",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestDevicelocationSubscriptionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"devicelocation:subscriptions", "delete",
		"--subscription-id", "qs15-h556-rt89-1298",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
