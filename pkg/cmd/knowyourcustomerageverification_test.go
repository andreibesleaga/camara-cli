// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
)

func TestKnowyourcustomerageverificationVerify(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"knowyourcustomerageverification", "verify",
		"--age-threshold", "18",
		"--birthdate", "1978-08-22",
		"--email", "federicaSanchez.Arjona@example.com",
		"--family-name", "Sanchez Arjona",
		"--family-name-at-birth", "YYYY",
		"--given-name", "Federica",
		"--id-document", "66666666q",
		"--include-content-lock=true",
		"--include-parental-control=true",
		"--middle-names", "Sanchez",
		"--name", "Federica Sanchez Arjona",
		"--phone-number", "+34629255833",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
