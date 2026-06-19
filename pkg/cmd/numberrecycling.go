// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/camara-cli/internal/apiquery"
	"github.com/stainless-sdks/camara-cli/internal/requestflag"
	"github.com/stainless-sdks/camara-go"
	"github.com/stainless-sdks/camara-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var numberrecyclingCheckSubscriberChange = cli.Command{
	Name:    "check-subscriber-change",
	Usage:   "Check whether the subscriber of the phone number has changed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "specified-date",
			Usage:    "Specified date to check whether there has been a change in the subscriber associated with the specific phone number, in RFC 3339 calendar date format (YYYY-MM-DD).",
			Required: true,
			BodyPath: "specifiedDate",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "A public identifier addressing a telephone subscription. In mobile networks it corresponds to the MSISDN (Mobile Station International Subscriber Directory Number). In order to be globally unique it has to be formatted in international format, according to E.164 standard, prefixed with '+'.",
			BodyPath: "phoneNumber",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleNumberrecyclingCheckSubscriberChange,
	HideHelpCommand: true,
}

func handleNumberrecyclingCheckSubscriberChange(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.NumberrecyclingCheckSubscriberChangeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Numberrecycling.CheckSubscriberChange(ctx, params, options...)
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
		Title:          "numberrecycling check-subscriber-change",
		Transform:      transform,
	})
}
