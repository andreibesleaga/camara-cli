// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
	"github.com/stainless-sdks/camara-cli/internal/requestflag"
)

func TestPopulationdensitydataRetrieve(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"populationdensitydata", "retrieve",
		"--area", "{areaType: POLYGON}",
		"--end-time", "2024-04-23T14:44:18.165Z",
		"--start-time", "2024-04-23T14:44:18.165Z",
		"--precision", "7",
		"--sink", "https://endpoint.example.com/sink",
		"--sink-credential", "{credentialType: PLAIN}",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(populationdensitydataRetrieve)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"populationdensitydata", "retrieve",
		"--area.areaType", "POLYGON",
		"--end-time", "2024-04-23T14:44:18.165Z",
		"--start-time", "2024-04-23T14:44:18.165Z",
		"--precision", "7",
		"--sink", "https://endpoint.example.com/sink",
		"--sink-credential.credentialType", "PLAIN",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
