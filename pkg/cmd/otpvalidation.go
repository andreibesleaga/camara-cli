// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/andreibesleaga/camara-cli/internal/apiquery"
	"github.com/andreibesleaga/camara-cli/internal/requestflag"
	"github.com/andreibesleaga/camara-go"
	"github.com/andreibesleaga/camara-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var otpvalidationSendCode = cli.Command{
	Name:    "send-code",
	Usage:   "Sends an SMS with the desired message and an OTP code to the received phone\nnumber.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "message",
			Usage:    "Message template used to compose the content of the SMS sent to the phone number. It must include the following label indicating where to include the short code `{{code}}`",
			Required: true,
			BodyPath: "message",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "A public identifier addressing a telephone subscription. In mobile networks it corresponds to the MSISDN (Mobile Station International Subscriber Directory Number). In order to be globally unique it has to be formatted in international format, according to E.164 standard, prefixed with '+'.",
			Required: true,
			BodyPath: "phoneNumber",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleOtpvalidationSendCode,
	HideHelpCommand: true,
}

var otpvalidationValidateCode = cli.Command{
	Name:    "validate-code",
	Usage:   "Verifies the code is valid for the received authenticationId",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "authentication-id",
			Usage:    "unique id of the verification attempt the code belongs to.",
			Required: true,
			BodyPath: "authenticationId",
		},
		&requestflag.Flag[string]{
			Name:     "code",
			Usage:    "temporal, short code to be validated",
			Required: true,
			BodyPath: "code",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleOtpvalidationValidateCode,
	HideHelpCommand: true,
}

func handleOtpvalidationSendCode(ctx context.Context, cmd *cli.Command) error {
	client := camara.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := camara.OtpvalidationSendCodeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Otpvalidation.SendCode(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "otpvalidation send-code",
		Transform:      transform,
	})
}

func handleOtpvalidationValidateCode(ctx context.Context, cmd *cli.Command) error {
	client := camara.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := camara.OtpvalidationValidateCodeParams{}

	return client.Otpvalidation.ValidateCode(ctx, params, options...)
}
