// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
)

func TestKnowyourcustomermatchMatch(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"knowyourcustomermatch", "match",
		"--address", "Tokyo-to Chiyoda-ku Iidabashi 3-10-10",
		"--birthdate", "1978-08-22",
		"--city-of-birth", "Madrid",
		"--country", "JP",
		"--country-of-birth", "ES",
		"--email", "abc@example.com",
		"--family-name", "Sanchez Arjona",
		"--family-name-at-birth", "YYYY",
		"--gender", "OTHER",
		"--given-name", "Federica",
		"--house-number-extension", "VVVV",
		"--id-document", "66666666q",
		"--id-document-expiry-date", "2027-07-12",
		"--id-document-type", "passport",
		"--locality", "ZZZZ",
		"--middle-names", "Sanchez",
		"--name", "Federica Sanchez Arjona",
		"--name-kana-hankaku", "federica",
		"--name-kana-zenkaku", "Ｆｅｄｅｒｉｃａ",
		"--nationality", "ES",
		"--phone-number", "+34629255833",
		"--postal-code", "1028460",
		"--region", "Tokyo",
		"--street-name", "Nicolas Salmeron",
		"--street-number", "4",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
