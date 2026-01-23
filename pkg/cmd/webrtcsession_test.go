// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/camara-cli/internal/mocktest"
	"github.com/stainless-sdks/camara-cli/internal/requestflag"
)

func TestWebrtcSessionsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webrtc:sessions", "create",
		"--registration-id", "registrationId",
		"--answer", "{sdp: sdp}",
		"--media-session-id", "0AEE1B58BAEEDA3EABA42B32EBB3DFE07E9CFF402EAF9EED8EF",
		"--offer", `{sdp: "v=0\r\no=- 8066321617929821805 2 IN IP4 127.0.0.1\r\ns=-\r\nt=0 0\r\nm=audio 42988 RTP/SAVPF 102 113\r\nc=IN IP6 2001:e0:410:2448:7a05:9b11:66f2:c9e\r\nb=AS:64\r\na=rtcp:9 IN IP4 0.0.0.0\r\na=candidate:1645903805 1 udp 2122262783 2001:e0:410:2448:7a05:9b11:66f2:c9e 42988 typ host generation 0 network-id 3 network-cost 900\r\na=ice-ufrag:4eKp\r\na=ice-pwd:D4sF5Pv9vx9ggaqxBlHbAFMx\r\na=ice-options:trickle renomination\r\na=mid:audio\r\na=extmap:2 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01\r\na=sendrecv\r\na=rtcp-mux\r\na=crypto:1 AES_CM_128_HMAC_SHA1_80 inline:Xm3YciqVIWFNSwy19e9MvfZ2YOdAZil7oT/tHjdf\r\na=rtpmap:102 AMR-WB/16000\r\na=fmtp:102 octet-align=0; mode-set=0,1,2; mode-change-capability=2\r\na=rtpmap:113 telephone-event/16000\r\n"}`,
		"--originator-address", "tel:+17085852753",
		"--originator-name", "tel:+17085852753",
		"--receiver-address", "tel:+17085854000",
		"--receiver-name", "tel:+17085854000",
		"--status", "Ringing",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(webrtcSessionsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"webrtc:sessions", "create",
		"--registration-id", "registrationId",
		"--answer.sdp", "sdp",
		"--media-session-id", "0AEE1B58BAEEDA3EABA42B32EBB3DFE07E9CFF402EAF9EED8EF",
		"--offer.sdp", "v=0\r\no=- 8066321617929821805 2 IN IP4 127.0.0.1\r\ns=-\r\nt=0 0\r\nm=audio 42988 RTP/SAVPF 102 113\r\nc=IN IP6 2001:e0:410:2448:7a05:9b11:66f2:c9e\r\nb=AS:64\r\na=rtcp:9 IN IP4 0.0.0.0\r\na=candidate:1645903805 1 udp 2122262783 2001:e0:410:2448:7a05:9b11:66f2:c9e 42988 typ host generation 0 network-id 3 network-cost 900\r\na=ice-ufrag:4eKp\r\na=ice-pwd:D4sF5Pv9vx9ggaqxBlHbAFMx\r\na=ice-options:trickle renomination\r\na=mid:audio\r\na=extmap:2 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01\r\na=sendrecv\r\na=rtcp-mux\r\na=crypto:1 AES_CM_128_HMAC_SHA1_80 inline:Xm3YciqVIWFNSwy19e9MvfZ2YOdAZil7oT/tHjdf\r\na=rtpmap:102 AMR-WB/16000\r\na=fmtp:102 octet-align=0; mode-set=0,1,2; mode-change-capability=2\r\na=rtpmap:113 telephone-event/16000\r\n",
		"--originator-address", "tel:+17085852753",
		"--originator-name", "tel:+17085852753",
		"--receiver-address", "tel:+17085854000",
		"--receiver-name", "tel:+17085854000",
		"--status", "Ringing",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestWebrtcSessionsRetrieve(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webrtc:sessions", "retrieve",
		"--media-session-id", "mediaSessionId",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}

func TestWebrtcSessionsUpdateStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"webrtc:sessions", "update-status",
		"--media-session-id", "mediaSessionId",
		"--answer", "{sdp: sdp}",
		"--media-session-id", "0AEE1B58BAEEDA3EABA42B32EBB3DFE07E9CFF402EAF9EED8EF",
		"--offer", "{sdp: sdp}",
		"--originator-address", "tel:+11234567899",
		"--originator-name", "Alice",
		"--receiver-address", "tel:+11234567899",
		"--receiver-name", "Bob",
		"--status", "Ringing",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(webrtcSessionsUpdateStatus)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"webrtc:sessions", "update-status",
		"--media-session-id", "mediaSessionId",
		"--answer.sdp", "sdp",
		"--media-session-id", "0AEE1B58BAEEDA3EABA42B32EBB3DFE07E9CFF402EAF9EED8EF",
		"--offer.sdp", "sdp",
		"--originator-address", "tel:+11234567899",
		"--originator-name", "Alice",
		"--receiver-address", "tel:+11234567899",
		"--receiver-name", "Bob",
		"--status", "Ringing",
		"--x-correlator", "b4333c46-49c0-4f62-80d7-f0ef930f1c46",
	)
}
