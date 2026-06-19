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

var customerinsightsScoringRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves Scoring information, for the user associated with the provided\n`idDocument`, `phoneNumber` or the combination of both parameters. It also\nallows to select the type of the Scoring scale measurement.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id-document",
			Usage:    "Identification number associated to the official identity document in the country. It may contain alphanumeric characters.",
			BodyPath: "idDocument",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "A public identifier addressing a telephone subscription. In mobile networks it corresponds to the MSISDN (Mobile Station International Subscriber Directory Number). In order to be globally unique it has to be formatted in international format, according to E.164 standard, prefixed with '+'.",
			BodyPath: "phoneNumber",
		},
		&requestflag.Flag[string]{
			Name:     "scoring-type",
			Usage:    "Scoring type, i.e.: scale. API Client may use this field to indicate the Scoring in one of the defined scales; if this field is not informed, the API will return the Scoring in the scale configured by default in the system.\n\nAllowed values are:\n\n* `gaugeMetric`: ranges from index 850 (lowest risk) to index 300 (highest risk)\n* `veritasIndex`: ranges from index 0 (lowest risk) to index 19 (highest risk)",
			BodyPath: "scoringType",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleCustomerinsightsScoringRetrieve,
	HideHelpCommand: true,
}

func handleCustomerinsightsScoringRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.CustomerinsightScoringGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Customerinsights.Scoring.Get(ctx, params, options...)
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
		Title:          "customerinsights:scoring retrieve",
		Transform:      transform,
	})
}
