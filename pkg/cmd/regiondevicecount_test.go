// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
	"github.com/stainless-sdks/camara-cli/internal/requestflag"
)

func TestRegiondevicecountGetCount(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"regiondevicecount", "get-count",
		"--area", "{areaType: CIRCLE}",
		"--endtime", "2023-07-04T14:27:08.312+02:00",
		"--filter", "{deviceType: [human device, IoT device], roamingStatus: [roaming]}",
		"--sink", "https://endpoint.example.com/sink",
		"--sink-credential", "{credentialType: ACCESSTOKEN}",
		"--starttime", "2023-07-03T14:27:08.312+02:00",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(regiondevicecountGetCount)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"regiondevicecount", "get-count",
		"--area.areaType", "CIRCLE",
		"--endtime", "2023-07-04T14:27:08.312+02:00",
		"--filter.deviceType", "[human device, IoT device]",
		"--filter.roamingStatus", "[roaming]",
		"--sink", "https://endpoint.example.com/sink",
		"--sink-credential.credentialType", "ACCESSTOKEN",
		"--starttime", "2023-07-03T14:27:08.312+02:00",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
