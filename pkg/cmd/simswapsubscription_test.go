// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
	"github.com/stainless-sdks/camara-cli/internal/requestflag"
)

func TestSimswapSubscriptionsCreate(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"simswap:subscriptions", "create",
		"--config", "{subscriptionDetail: {phoneNumber: '+123456789'}, subscriptionExpireTime: '2025-01-17T13:18:23.682Z', subscriptionMaxEvents: 10}",
		"--protocol", "HTTP",
		"--sink", "https://endpoint.example.com/sink",
		"--type", "org.camaraproject.sim-swap-subscriptions.v0.swapped",
		"--sink-credential", "{credentialType: ACCESSTOKEN}",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(simswapSubscriptionsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"simswap:subscriptions", "create",
		"--config.subscriptionDetail", "{phoneNumber: '+123456789'}",
		"--config.subscriptionExpireTime", "2025-01-17T13:18:23.682Z",
		"--config.subscriptionMaxEvents", "10",
		"--protocol", "HTTP",
		"--sink", "https://endpoint.example.com/sink",
		"--type", "org.camaraproject.sim-swap-subscriptions.v0.swapped",
		"--sink-credential.credentialType", "ACCESSTOKEN",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestSimswapSubscriptionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"simswap:subscriptions", "retrieve",
		"--subscription-id", "qs15-h556-rt89-1298",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestSimswapSubscriptionsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"simswap:subscriptions", "list",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestSimswapSubscriptionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"simswap:subscriptions", "delete",
		"--subscription-id", "qs15-h556-rt89-1298",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
