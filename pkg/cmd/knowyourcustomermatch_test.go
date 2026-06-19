// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
)

func TestKnowyourcustomermatchMatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--bearer-token", "string",
			"--customer-insights-token", "string",
			"--device-swap-token", "string",
			"--kyc-age-verification-token", "string",
			"--kyc-fill-in-token", "string",
			"--kyc-match-token", "string",
			"--tenure-token", "string",
			"--number-recycling-token", "string",
			"--otp-validation-token", "string",
			"--call-forwarding-signal-token", "string",
			"--device-location-token", "string",
			"--population-density-data-token", "string",
			"--region-device-count-token", "string",
			"--web-rtc-token", "string",
			"--connectivity-insights-token", "string",
			"--quality-on-demand-token", "string",
			"--device-identifier-token", "string",
			"--sim-swap-token", "string",
			"--device-roaming-status-token", "string",
			"--device-reachability-status-token", "string",
			"--connected-network-type-token", "string",
			"--device-location-notifications-api-key", "string",
			"--notifications-api-key", "string",
			"--population-density-data-notifications-api-key", "string",
			"--region-device-count-notifications-api-key", "string",
			"--connectivity-insights-notifications-api-key", "string",
			"--sim-swap-notifications-api-key", "string",
			"--device-roaming-status-notifications-api-key", "string",
			"--device-reachability-status-notifications-api-key", "string",
			"--connected-network-type-notifications-api-key", "string",
			"knowyourcustomermatch", "match",
			"--address", "Tokyo-to Chiyoda-ku Iidabashi 3-10-10",
			"--birthdate", "'1978-08-22'",
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
			"--id-document-expiry-date", "'2027-07-12'",
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
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"address: Tokyo-to Chiyoda-ku Iidabashi 3-10-10\n" +
			"birthdate: '1978-08-22'\n" +
			"cityOfBirth: Madrid\n" +
			"country: JP\n" +
			"countryOfBirth: ES\n" +
			"email: abc@example.com\n" +
			"familyName: Sanchez Arjona\n" +
			"familyNameAtBirth: YYYY\n" +
			"gender: OTHER\n" +
			"givenName: Federica\n" +
			"houseNumberExtension: VVVV\n" +
			"idDocument: 66666666q\n" +
			"idDocumentExpiryDate: '2027-07-12'\n" +
			"idDocumentType: passport\n" +
			"locality: ZZZZ\n" +
			"middleNames: Sanchez\n" +
			"name: Federica Sanchez Arjona\n" +
			"nameKanaHankaku: federica\n" +
			"nameKanaZenkaku: Ｆｅｄｅｒｉｃａ\n" +
			"nationality: ES\n" +
			"phoneNumber: '+34629255833'\n" +
			"postalCode: '1028460'\n" +
			"region: Tokyo\n" +
			"streetName: Nicolas Salmeron\n" +
			"streetNumber: '4'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--bearer-token", "string",
			"--customer-insights-token", "string",
			"--device-swap-token", "string",
			"--kyc-age-verification-token", "string",
			"--kyc-fill-in-token", "string",
			"--kyc-match-token", "string",
			"--tenure-token", "string",
			"--number-recycling-token", "string",
			"--otp-validation-token", "string",
			"--call-forwarding-signal-token", "string",
			"--device-location-token", "string",
			"--population-density-data-token", "string",
			"--region-device-count-token", "string",
			"--web-rtc-token", "string",
			"--connectivity-insights-token", "string",
			"--quality-on-demand-token", "string",
			"--device-identifier-token", "string",
			"--sim-swap-token", "string",
			"--device-roaming-status-token", "string",
			"--device-reachability-status-token", "string",
			"--connected-network-type-token", "string",
			"--device-location-notifications-api-key", "string",
			"--notifications-api-key", "string",
			"--population-density-data-notifications-api-key", "string",
			"--region-device-count-notifications-api-key", "string",
			"--connectivity-insights-notifications-api-key", "string",
			"--sim-swap-notifications-api-key", "string",
			"--device-roaming-status-notifications-api-key", "string",
			"--device-reachability-status-notifications-api-key", "string",
			"--connected-network-type-notifications-api-key", "string",
			"knowyourcustomermatch", "match",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})
}
