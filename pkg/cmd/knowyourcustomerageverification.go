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

var knowyourcustomerageverificationVerify = cli.Command{
	Name:    "verify",
	Usage:   "Verify that the age of the subscriber associated with a phone number is equal to\nor greater than the specified age threshold value.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "age-threshold",
			Usage:    "The age to be verified. The indicated range is a global definition of maximum and minimum values allowed to be requested. It is important to note that this range might be more restrictive in some implementations due to local regulations of a country i.e. A country does not allow to request for an age under 18. This limitation must be informed during the onboarding process.",
			Required: true,
			BodyPath: "ageThreshold",
		},
		&requestflag.Flag[any]{
			Name:     "birthdate",
			Usage:    "The birthdate of the customer, in RFC 3339 / ISO 8601 calendar date format (YYYY-MM-DD).",
			BodyPath: "birthdate",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "Email address of the customer in the RFC specified format (local-part@domain).",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "family-name",
			Usage:    "Last name, family name, or surname of the customer.",
			BodyPath: "familyName",
		},
		&requestflag.Flag[string]{
			Name:     "family-name-at-birth",
			Usage:    "Last/family/sur- name at birth of the customer.",
			BodyPath: "familyNameAtBirth",
		},
		&requestflag.Flag[string]{
			Name:     "given-name",
			Usage:    "First/given name or compound first/given name of the customer.",
			BodyPath: "givenName",
		},
		&requestflag.Flag[string]{
			Name:     "id-document",
			Usage:    "Id number associated to the official identity document in the country. It may contain alphanumeric characters.",
			BodyPath: "idDocument",
		},
		&requestflag.Flag[bool]{
			Name:     "include-content-lock",
			Usage:    "If this parameter is included in the request with value `true`, the response property `contentLock` will be returned. If it is not included or its value is `false`, the response property will not be returned.",
			Default:  false,
			BodyPath: "includeContentLock",
		},
		&requestflag.Flag[bool]{
			Name:     "include-parental-control",
			Usage:    "If this parameter is included in the request with value `true`, the response property `parentalControl` will be returned. If it is not included or its value is `false`, the response property will not be returned.",
			Default:  false,
			BodyPath: "includeParentalControl",
		},
		&requestflag.Flag[string]{
			Name:     "middle-names",
			Usage:    "Middle name/s of the customer.",
			BodyPath: "middleNames",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Complete name of the customer, usually composed of first/given name and last/family/sur- name in a country.  Depending on the country, the order of first/give name and last/family/sur- name varies, and middle name could be included.  It can use givenName, middleNames, familyName and/or familyNameAtBirth. For example, in ESP, name+familyName; in NLD, it can be name+middleNames+familyName or name+middleNames+familyNameAtBirth, etc.",
			BodyPath: "name",
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
	Action:          handleKnowyourcustomerageverificationVerify,
	HideHelpCommand: true,
}

func handleKnowyourcustomerageverificationVerify(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.KnowyourcustomerageverificationVerifyParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Knowyourcustomerageverification.Verify(ctx, params, options...)
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
		Title:          "knowyourcustomerageverification verify",
		Transform:      transform,
	})
}
