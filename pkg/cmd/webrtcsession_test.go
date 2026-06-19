// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/andreibesleaga/camara-cli/internal/mocktest"
	"github.com/andreibesleaga/camara-cli/internal/requestflag"
)

func TestWebrtcSessionsCreate(t *testing.T) {
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
			"webrtc:sessions", "create",
			"--registration-id", "registrationId",
			"--answer", "{sdp: sdp}",
			"--call-type", "REGULAR",
			"--location-details", "{confidence: {pdf: normal, value: 0}, coordinates: {latitude: 0, longitude: 0, radius: 0}, method: GPS, shape: Circle, timestamp: '2019-12-27T18:11:19.117Z'}",
			"--body-media-session-id", "0AEE1B58BAEEDA3EABA42B32EBB3DFE07E9CFF402EAF9EED8EF",
			"--offer", `{sdp: "v=0\r\no=- 8066321617929821805 2 IN IP4 127.0.0.1\r\ns=-\r\nt=0 0\r\nm=audio 42988 RTP/SAVPF 102 113\r\nc=IN IP6 2001:e0:410:2448:7a05:9b11:66f2:c9e\r\nb=AS:64\r\na=rtcp:9 IN IP4 0.0.0.0\r\na=candidate:1645903805 1 udp 2122262783 2001:e0:410:2448:7a05:9b11:66f2:c9e 42988 typ host generation 0 network-id 3 network-cost 900\r\na=ice-ufrag:4eKp\r\na=ice-pwd:D4sF5Pv9vx9ggaqxBlHbAFMx\r\na=ice-options:trickle renomination\r\na=mid:audio\r\na=extmap:2 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01\r\na=sendrecv\r\na=rtcp-mux\r\na=crypto:1 AES_CM_128_HMAC_SHA1_80 inline:Xm3YciqVIWFNSwy19e9MvfZ2YOdAZil7oT/tHjdf\r\na=rtpmap:102 AMR-WB/16000\r\na=fmtp:102 octet-align=0; mode-set=0,1,2; mode-change-capability=2\r\na=rtpmap:113 telephone-event/16000\r\n"}`,
			"--originator-address", "tel:+17085852753",
			"--originator-name", "tel:+17085852753",
			"--receiver-address", "tel:+17085854000",
			"--receiver-name", "tel:+17085854000",
			"--status", "Ringing",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(webrtcSessionsCreate)

		// Alternative argument passing style using inner flags
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
			"webrtc:sessions", "create",
			"--registration-id", "registrationId",
			"--answer.sdp", "sdp",
			"--call-type", "REGULAR",
			"--location-details.confidence", "{pdf: normal, value: 0}",
			"--location-details.coordinates", "{latitude: 0, longitude: 0, radius: 0}",
			"--location-details.method", "GPS",
			"--location-details.shape", "Circle",
			"--location-details.timestamp", "2019-12-27T18:11:19.117Z",
			"--body-media-session-id", "0AEE1B58BAEEDA3EABA42B32EBB3DFE07E9CFF402EAF9EED8EF",
			"--offer.sdp", "v=0\r\no=- 8066321617929821805 2 IN IP4 127.0.0.1\r\ns=-\r\nt=0 0\r\nm=audio 42988 RTP/SAVPF 102 113\r\nc=IN IP6 2001:e0:410:2448:7a05:9b11:66f2:c9e\r\nb=AS:64\r\na=rtcp:9 IN IP4 0.0.0.0\r\na=candidate:1645903805 1 udp 2122262783 2001:e0:410:2448:7a05:9b11:66f2:c9e 42988 typ host generation 0 network-id 3 network-cost 900\r\na=ice-ufrag:4eKp\r\na=ice-pwd:D4sF5Pv9vx9ggaqxBlHbAFMx\r\na=ice-options:trickle renomination\r\na=mid:audio\r\na=extmap:2 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01\r\na=sendrecv\r\na=rtcp-mux\r\na=crypto:1 AES_CM_128_HMAC_SHA1_80 inline:Xm3YciqVIWFNSwy19e9MvfZ2YOdAZil7oT/tHjdf\r\na=rtpmap:102 AMR-WB/16000\r\na=fmtp:102 octet-align=0; mode-set=0,1,2; mode-change-capability=2\r\na=rtpmap:113 telephone-event/16000\r\n",
			"--originator-address", "tel:+17085852753",
			"--originator-name", "tel:+17085852753",
			"--receiver-address", "tel:+17085854000",
			"--receiver-name", "tel:+17085854000",
			"--status", "Ringing",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"answer:\n" +
			"  sdp: sdp\n" +
			"callType: REGULAR\n" +
			"locationDetails:\n" +
			"  confidence:\n" +
			"    pdf: normal\n" +
			"    value: 0\n" +
			"  coordinates:\n" +
			"    latitude: 0\n" +
			"    longitude: 0\n" +
			"    radius: 0\n" +
			"  method: GPS\n" +
			"  shape: Circle\n" +
			"  timestamp: '2019-12-27T18:11:19.117Z'\n" +
			"mediaSessionId: 0AEE1B58BAEEDA3EABA42B32EBB3DFE07E9CFF402EAF9EED8EF\n" +
			"offer:\n" +
			"  sdp: \"v=0\\r\\no=- 8066321617929821805 2 IN IP4 127.0.0.1\\r\\ns=-\\r\\nt=0 0\\r\\nm=audio 42988 RTP/SAVPF 102 113\\r\\nc=IN IP6 2001:e0:410:2448:7a05:9b11:66f2:c9e\\r\\nb=AS:64\\r\\na=rtcp:9 IN IP4 0.0.0.0\\r\\na=candidate:1645903805 1 udp 2122262783 2001:e0:410:2448:7a05:9b11:66f2:c9e 42988 typ host generation 0 network-id 3 network-cost 900\\r\\na=ice-ufrag:4eKp\\r\\na=ice-pwd:D4sF5Pv9vx9ggaqxBlHbAFMx\\r\\na=ice-options:trickle renomination\\r\\na=mid:audio\\r\\na=extmap:2 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01\\r\\na=sendrecv\\r\\na=rtcp-mux\\r\\na=crypto:1 AES_CM_128_HMAC_SHA1_80 inline:Xm3YciqVIWFNSwy19e9MvfZ2YOdAZil7oT/tHjdf\\r\\na=rtpmap:102 AMR-WB/16000\\r\\na=fmtp:102 octet-align=0; mode-set=0,1,2; mode-change-capability=2\\r\\na=rtpmap:113 telephone-event/16000\\r\\n\"\n" +
			"originatorAddress: tel:+17085852753\n" +
			"originatorName: tel:+17085852753\n" +
			"receiverAddress: tel:+17085854000\n" +
			"receiverName: tel:+17085854000\n" +
			"status: Ringing\n")
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
			"webrtc:sessions", "create",
			"--registration-id", "registrationId",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})
}

func TestWebrtcSessionsRetrieve(t *testing.T) {
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
			"webrtc:sessions", "retrieve",
			"--media-session-id", "mediaSessionId",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})
}

func TestWebrtcSessionsDelete(t *testing.T) {
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
			"webrtc:sessions", "delete",
			"--media-session-id", "mediaSessionId",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})
}

func TestWebrtcSessionsUpdateStatus(t *testing.T) {
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
			"webrtc:sessions", "update-status",
			"--media-session-id", "mediaSessionId",
			"--answer", "{sdp: sdp}",
			"--call-type", "REGULAR",
			"--location-details", "{confidence: {pdf: normal, value: 0}, coordinates: {latitude: 0, longitude: 0, radius: 0}, method: GPS, shape: Circle, timestamp: '2019-12-27T18:11:19.117Z'}",
			"--body-media-session-id", "0AEE1B58BAEEDA3EABA42B32EBB3DFE07E9CFF402EAF9EED8EF",
			"--offer", "{sdp: sdp}",
			"--originator-address", "tel:+11234567899",
			"--originator-name", "Alice",
			"--receiver-address", "tel:+11234567899",
			"--receiver-name", "Bob",
			"--status", "Ringing",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(webrtcSessionsUpdateStatus)

		// Alternative argument passing style using inner flags
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
			"webrtc:sessions", "update-status",
			"--media-session-id", "mediaSessionId",
			"--answer.sdp", "sdp",
			"--call-type", "REGULAR",
			"--location-details.confidence", "{pdf: normal, value: 0}",
			"--location-details.coordinates", "{latitude: 0, longitude: 0, radius: 0}",
			"--location-details.method", "GPS",
			"--location-details.shape", "Circle",
			"--location-details.timestamp", "2019-12-27T18:11:19.117Z",
			"--body-media-session-id", "0AEE1B58BAEEDA3EABA42B32EBB3DFE07E9CFF402EAF9EED8EF",
			"--offer.sdp", "sdp",
			"--originator-address", "tel:+11234567899",
			"--originator-name", "Alice",
			"--receiver-address", "tel:+11234567899",
			"--receiver-name", "Bob",
			"--status", "Ringing",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"answer:\n" +
			"  sdp: sdp\n" +
			"callType: REGULAR\n" +
			"locationDetails:\n" +
			"  confidence:\n" +
			"    pdf: normal\n" +
			"    value: 0\n" +
			"  coordinates:\n" +
			"    latitude: 0\n" +
			"    longitude: 0\n" +
			"    radius: 0\n" +
			"  method: GPS\n" +
			"  shape: Circle\n" +
			"  timestamp: '2019-12-27T18:11:19.117Z'\n" +
			"mediaSessionId: 0AEE1B58BAEEDA3EABA42B32EBB3DFE07E9CFF402EAF9EED8EF\n" +
			"offer:\n" +
			"  sdp: sdp\n" +
			"originatorAddress: tel:+11234567899\n" +
			"originatorName: Alice\n" +
			"receiverAddress: tel:+11234567899\n" +
			"receiverName: Bob\n" +
			"status: Ringing\n")
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
			"webrtc:sessions", "update-status",
			"--media-session-id", "mediaSessionId",
			"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
		)
	})
}
