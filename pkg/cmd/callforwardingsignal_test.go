// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
)

func TestCallforwardingsignalCheckActiveForwardings(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"callforwardingsignal", "check-active-forwardings",
		"--phone-number", "+123456789",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestCallforwardingsignalCheckUnconditionalForwarding(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"callforwardingsignal", "check-unconditional-forwarding",
		"--phone-number", "+123456789",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
