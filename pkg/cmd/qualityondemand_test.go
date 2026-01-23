// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
	"github.com/stainless-sdks/camara-cli/internal/requestflag"
)

func TestQualityondemandRetrieveQosProfile(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"qualityondemand", "retrieve-qos-profile",
		"--name", "voice",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestQualityondemandRetrieveQosProfiles(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"qualityondemand", "retrieve-qos-profiles",
		"--device", "{ipv4Address: {privateAddress: 203.0.113.0, publicAddress: 203.0.113.0, publicPort: 59765}, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344, networkAccessIdentifier: 123456789@domain.com, phoneNumber: '+123456789'}",
		"--name", "voice",
		"--status", "ACTIVE",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(qualityondemandRetrieveQosProfiles)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"qualityondemand", "retrieve-qos-profiles",
		"--device.ipv4Address", "{privateAddress: 203.0.113.0, publicAddress: 203.0.113.0, publicPort: 59765}",
		"--device.ipv6Address", "2001:db8:85a3:8d3:1319:8a2e:370:7344",
		"--device.networkAccessIdentifier", "123456789@domain.com",
		"--device.phoneNumber", "+123456789",
		"--name", "voice",
		"--status", "ACTIVE",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
