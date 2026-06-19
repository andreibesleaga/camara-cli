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

var knowyourcustomermatchMatch = cli.Command{
	Name:    "match",
	Usage:   "Verify matching of a number of attributes related to a customer identity against\nthe verified data bound to their phone number in the Operator systems.\nRegardless of whether the `phoneNumber` is explicitly stated in the request\nbody, at least one of the other fields must be provided, otherwise a\n`HTTP 400 - KNOW_YOUR_CUSTOMER.INVALID_PARAM_COMBINATION` error will be\nreturned.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "address",
			Usage:    "Complete address of the customer.  For some countries, it is built following the usual concatenation of parameters in a country, but for other countries, this is not the case.  For some countries, it can use streetName, streetNumber and/or houseNumberExtension. For example, in ESP, streetName+streetNumber; in NLD, it can be streetName+streetNumber or streetName+streetNumber+houseNumberExtension.",
			BodyPath: "address",
		},
		&requestflag.Flag[any]{
			Name:     "birthdate",
			Usage:    "The birthdate of the customer, in RFC 3339 / ISO 8601 calendar date format (YYYY-MM-DD).",
			BodyPath: "birthdate",
		},
		&requestflag.Flag[string]{
			Name:     "city-of-birth",
			Usage:    "City where the customer was born.",
			BodyPath: "cityOfBirth",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "Country of the customer's address. Format ISO 3166-1 alpha-2",
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "country-of-birth",
			Usage:    "Country where the customer was born. Format ISO 3166-1 alpha-2.",
			BodyPath: "countryOfBirth",
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
			Name:     "gender",
			Usage:    "Gender of the customer (Male/Female/Other).",
			BodyPath: "gender",
		},
		&requestflag.Flag[string]{
			Name:     "given-name",
			Usage:    "First/given name or compound first/given name of the customer.",
			BodyPath: "givenName",
		},
		&requestflag.Flag[string]{
			Name:     "house-number-extension",
			Usage:    "Specific identifier of the house needed depending on the property type. For example, number of apartment in an apartment building.",
			BodyPath: "houseNumberExtension",
		},
		&requestflag.Flag[string]{
			Name:     "id-document",
			Usage:    "Id number associated to the official identity document in the country. It may contain alphanumeric characters.",
			BodyPath: "idDocument",
		},
		&requestflag.Flag[any]{
			Name:     "id-document-expiry-date",
			Usage:    "Expiration date of the identity document (ISO 8601).",
			BodyPath: "idDocumentExpiryDate",
		},
		&requestflag.Flag[string]{
			Name:     "id-document-type",
			Usage:    "Type of the official identity document provided.",
			BodyPath: "idDocumentType",
		},
		&requestflag.Flag[string]{
			Name:     "locality",
			Usage:    "Locality of the customer's address",
			BodyPath: "locality",
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
			Name:     "name-kana-hankaku",
			Usage:    "Complete name of the customer in Hankaku-Kana format (reading of name) for Japan.",
			BodyPath: "nameKanaHankaku",
		},
		&requestflag.Flag[string]{
			Name:     "name-kana-zenkaku",
			Usage:    "Complete name of the customer in Zenkaku-Kana format (reading of name) for Japan.",
			BodyPath: "nameKanaZenkaku",
		},
		&requestflag.Flag[string]{
			Name:     "nationality",
			Usage:    "ISO 3166-1 alpha-2 code of the customer’s nationality. In the case a customer has more than one nationality, it is supposed to be the nationality related to the ID document provided in the match request.",
			BodyPath: "nationality",
		},
		&requestflag.Flag[string]{
			Name:     "phone-number",
			Usage:    "A public identifier addressing a telephone subscription. In mobile networks it corresponds to the MSISDN (Mobile Station International Subscriber Directory Number). In order to be globally unique it has to be formatted in international format, according to E.164 standard, prefixed with '+'.",
			BodyPath: "phoneNumber",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			Usage:    "Zip code or postal code",
			BodyPath: "postalCode",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region/prefecture of the customer's address",
			BodyPath: "region",
		},
		&requestflag.Flag[string]{
			Name:     "street-name",
			Usage:    "Name of the street of the customer's address. It should not include the type of the street.",
			BodyPath: "streetName",
		},
		&requestflag.Flag[string]{
			Name:     "street-number",
			Usage:    "The street number of the customer's address.  Number identifying a specific property on the 'streetName'.",
			BodyPath: "streetNumber",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleKnowyourcustomermatchMatch,
	HideHelpCommand: true,
}

func handleKnowyourcustomermatchMatch(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.KnowyourcustomermatchMatchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Knowyourcustomermatch.Match(ctx, params, options...)
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
		Title:          "knowyourcustomermatch match",
		Transform:      transform,
	})
}
