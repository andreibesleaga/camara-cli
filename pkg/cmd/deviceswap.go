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

var deviceswapCheck = cli.Command{
	Name:    "check",
	Usage:   "Check if device swap has been performed during a past period",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "max-age",
			Usage:    "Period in hours to be checked for device swap.\n",
			Default:  240,
			BodyPath: "maxAge",
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
	Action:          handleDeviceswapCheck,
	HideHelpCommand: true,
}

var deviceswapRetrieveDate = cli.Command{
	Name:    "retrieve-date",
	Usage:   "Get timestamp of last device swap for a mobile user account provided with phone\nnumber.",
	Suggest: true,
	Flags: []cli.Flag{
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
	Action:          handleDeviceswapRetrieveDate,
	HideHelpCommand: true,
}

func handleDeviceswapCheck(ctx context.Context, cmd *cli.Command) error {
	client := camara.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := camara.DeviceswapCheckParams{}

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
	_, err = client.Deviceswap.Check(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "deviceswap check", obj, format, transform)
}

func handleDeviceswapRetrieveDate(ctx context.Context, cmd *cli.Command) error {
	client := camara.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := camara.DeviceswapGetDateParams{}

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
	_, err = client.Deviceswap.GetDate(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "deviceswap retrieve-date", obj, format, transform)
}
