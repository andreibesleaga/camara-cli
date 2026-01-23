// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/camara-cli/internal/apiquery"
	"github.com/stainless-sdks/camara-cli/internal/requestflag"
	"github.com/stainless-sdks/camara-go"
	"github.com/stainless-sdks/camara-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var webrtcSessionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a voice and/or video session",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:       "registration-id",
			Required:   true,
			HeaderPath: "registrationId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "answer",
			Usage:    "**OFFER**: An inlined session description in SDP format [RFC4566].If XML syntax\nis used, the content of this element SHALL be embedded in a CDATA\nsection.\n\n**ANSWER**: This type represents an answer in WebRTC Signaling. This element is not\npresent in case there is no answer yet, or the session invitation has\nbeen declined by the Terminating Participant.This element MUST NOT be\npresent in a request from the application to the server to create a\nsession.",
			BodyPath: "answer",
		},
		&requestflag.Flag[string]{
			Name:     "media-session-id",
			Usage:    "The media session ID created by the network. The mediaSessionId shall not be included in POST requests by the client, but must be included in the notifications from the network to the client device.",
			BodyPath: "mediaSessionId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "offer",
			Usage:    "**OFFER**: An inlined session description in SDP format [RFC4566].If XML syntax\nis used, the content of this element SHALL be embedded in a CDATA\nsection.\n\n**ANSWER**: This type represents an answer in WebRTC Signaling. This element is not\npresent in case there is no answer yet, or the session invitation has\nbeen declined by the Terminating Participant.This element MUST NOT be\npresent in a request from the application to the server to create a\nsession.",
			BodyPath: "offer",
		},
		&requestflag.Flag[string]{
			Name:     "originator-address",
			Usage:    "Subscriber address (Sender or Receiver)",
			BodyPath: "originatorAddress",
		},
		&requestflag.Flag[string]{
			Name:     "originator-name",
			Usage:    "Friendly name of the call originator",
			BodyPath: "originatorName",
		},
		&requestflag.Flag[string]{
			Name:     "receiver-address",
			Usage:    "Subscriber address (Sender or Receiver)",
			BodyPath: "receiverAddress",
		},
		&requestflag.Flag[string]{
			Name:     "receiver-name",
			Usage:    "Friendly name of the call terminator",
			BodyPath: "receiverName",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Provides the status of the media session. During the session creation, this attribute SHALL NOT be included in the request.",
			BodyPath: "status",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleWebrtcSessionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"answer": {
		&requestflag.InnerFlag[string]{
			Name:       "answer.sdp",
			Usage:      "An inlined session description in SDP format [RFC4566].If XML syntax is used, the content of this element SHALL be embedded in a CDATA section",
			InnerField: "sdp",
		},
	},
	"offer": {
		&requestflag.InnerFlag[string]{
			Name:       "offer.sdp",
			Usage:      "An inlined session description in SDP format [RFC4566].If XML syntax is used, the content of this element SHALL be embedded in a CDATA section",
			InnerField: "sdp",
		},
	},
})

var webrtcSessionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get the media Session description based on `mediaSessionId`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "media-session-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleWebrtcSessionsRetrieve,
	HideHelpCommand: true,
}

var webrtcSessionsUpdateStatus = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-status",
	Usage:   "Update the status of the media session, this may include updating SDP media",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "media-session-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "answer",
			Usage:    "**OFFER**: An inlined session description in SDP format [RFC4566].If XML syntax\nis used, the content of this element SHALL be embedded in a CDATA\nsection.\n\n**ANSWER**: This type represents an answer in WebRTC Signaling. This element is not\npresent in case there is no answer yet, or the session invitation has\nbeen declined by the Terminating Participant.This element MUST NOT be\npresent in a request from the application to the server to create a\nsession.",
			BodyPath: "answer",
		},
		&requestflag.Flag[string]{
			Name:     "media-session-id",
			Usage:    "The media session ID created by the network. The mediaSessionId shall not be included in POST requests by the client, but must be included in the notifications from the network to the client device.",
			BodyPath: "mediaSessionId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "offer",
			Usage:    "**OFFER**: An inlined session description in SDP format [RFC4566].If XML syntax\nis used, the content of this element SHALL be embedded in a CDATA\nsection.\n\n**ANSWER**: This type represents an answer in WebRTC Signaling. This element is not\npresent in case there is no answer yet, or the session invitation has\nbeen declined by the Terminating Participant.This element MUST NOT be\npresent in a request from the application to the server to create a\nsession.",
			BodyPath: "offer",
		},
		&requestflag.Flag[string]{
			Name:     "originator-address",
			Usage:    "Subscriber address (Sender or Receiver)",
			BodyPath: "originatorAddress",
		},
		&requestflag.Flag[string]{
			Name:     "originator-name",
			Usage:    "Friendly name of the call originator",
			BodyPath: "originatorName",
		},
		&requestflag.Flag[string]{
			Name:     "receiver-address",
			Usage:    "Subscriber address (Sender or Receiver)",
			BodyPath: "receiverAddress",
		},
		&requestflag.Flag[string]{
			Name:     "receiver-name",
			Usage:    "Friendly name of the call terminator",
			BodyPath: "receiverName",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Provides the status of the media session. During the session creation, this attribute SHALL NOT be included in the request.",
			BodyPath: "status",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleWebrtcSessionsUpdateStatus,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"answer": {
		&requestflag.InnerFlag[string]{
			Name:       "answer.sdp",
			Usage:      "An inlined session description in SDP format [RFC4566].If XML syntax is used, the content of this element SHALL be embedded in a CDATA section",
			InnerField: "sdp",
		},
	},
	"offer": {
		&requestflag.InnerFlag[string]{
			Name:       "offer.sdp",
			Usage:      "An inlined session description in SDP format [RFC4566].If XML syntax is used, the content of this element SHALL be embedded in a CDATA section",
			InnerField: "sdp",
		},
	},
})

func handleWebrtcSessionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := camara.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := camara.WebrtcSessionNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Webrtc.Sessions.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webrtc:sessions create", obj, format, transform)
}

func handleWebrtcSessionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := camara.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("media-session-id") && len(unusedArgs) > 0 {
		cmd.Set("media-session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := camara.WebrtcSessionGetParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Webrtc.Sessions.Get(
		ctx,
		cmd.Value("media-session-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webrtc:sessions retrieve", obj, format, transform)
}

func handleWebrtcSessionsUpdateStatus(ctx context.Context, cmd *cli.Command) error {
	client := camara.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("media-session-id") && len(unusedArgs) > 0 {
		cmd.Set("media-session-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := camara.WebrtcSessionUpdateStatusParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Webrtc.Sessions.UpdateStatus(
		ctx,
		cmd.Value("media-session-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webrtc:sessions update-status", obj, format, transform)
}
