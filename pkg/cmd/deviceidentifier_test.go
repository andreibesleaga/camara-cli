// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
	"github.com/stainless-sdks/camara-cli/internal/requestflag"
)

func TestDeviceidentifierRetrieveIdentifier(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"deviceidentifier", "retrieve-identifier",
		"--device", "{ipv4Address: {privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344, networkAccessIdentifier: 123456789@example.com, phoneNumber: '+123456789'}",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(deviceidentifierRetrieveIdentifier)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"deviceidentifier", "retrieve-identifier",
		"--device.ipv4Address", "{privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}",
		"--device.ipv6Address", "2001:db8:85a3:8d3:1319:8a2e:370:7344",
		"--device.networkAccessIdentifier", "123456789@example.com",
		"--device.phoneNumber", "+123456789",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestDeviceidentifierRetrievePpid(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"deviceidentifier", "retrieve-ppid",
		"--device", "{ipv4Address: {privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344, networkAccessIdentifier: 123456789@example.com, phoneNumber: '+123456789'}",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(deviceidentifierRetrievePpid)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"deviceidentifier", "retrieve-ppid",
		"--device.ipv4Address", "{privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}",
		"--device.ipv6Address", "2001:db8:85a3:8d3:1319:8a2e:370:7344",
		"--device.networkAccessIdentifier", "123456789@example.com",
		"--device.phoneNumber", "+123456789",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestDeviceidentifierRetrieveType(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"deviceidentifier", "retrieve-type",
		"--device", "{ipv4Address: {privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}, ipv6Address: 2001:db8:85a3:8d3:1319:8a2e:370:7344, networkAccessIdentifier: 123456789@example.com, phoneNumber: '+123456789'}",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(deviceidentifierRetrieveType)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"deviceidentifier", "retrieve-type",
		"--device.ipv4Address", "{privateAddress: 84.125.93.10, publicAddress: 84.125.93.10, publicPort: 59765}",
		"--device.ipv6Address", "2001:db8:85a3:8d3:1319:8a2e:370:7344",
		"--device.networkAccessIdentifier", "123456789@example.com",
		"--device.phoneNumber", "+123456789",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
