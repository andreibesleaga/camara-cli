// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
)

func TestDeviceswapCheck(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"deviceswap", "check",
		"--max-age", "120",
		"--phone-number", "+34666111333",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestDeviceswapRetrieveDate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"deviceswap", "retrieve-date",
		"--phone-number", "+34666111333",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
