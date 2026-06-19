// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/andreibesleaga/camara-go"
	"github.com/andreibesleaga/camara-go/option"
	"github.com/stainless-sdks/camara-cli/internal/apiquery"
	"github.com/stainless-sdks/camara-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var populationdensitydataRetrieve = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves population density estimation together with the estimation range\nrelated for a time slot for a given area (described as a polygon) as a data set\nconsisting of a sequence of equally-sized objects covering the input polygon\narea.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "area",
			Usage:    "Base schema for all areas",
			Required: true,
			BodyPath: "area",
		},
		&requestflag.Flag[any]{
			Name:     "end-time",
			Usage:    "End date time. It must follow [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must have time zone. Recommended format is yyyy-MM-dd'T'HH:mm:ss.SSSZ (i.e. which allows 2023-07-03T14:27:08.312+02:00 or 2023-07-03T12:27:08.312Z) The maximum endTime allowed is 3 months from the time of the request.",
			Required: true,
			BodyPath: "endTime",
		},
		&requestflag.Flag[any]{
			Name:     "start-time",
			Usage:    "Start date time. It must follow [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must have time zone. Recommended format is yyyy-MM-dd'T'HH:mm:ss.SSSZ",
			Required: true,
			BodyPath: "startTime",
		},
		&requestflag.Flag[int64]{
			Name:     "precision",
			Usage:    "Precision required of response cells. Precision defines a geohash level and corresponds to the length of the geohash for each cell. More information at [Geohash system](https://en.wikipedia.org/wiki/Geohash)\" If not included the default precision level 7 is used by default. In case of using a not supported level by the MNO, the API returns the error response `POPULATION_DENSITY_DATA.UNSUPPORTED_PRECISION`.",
			Default:  7,
			BodyPath: "precision",
		},
		&requestflag.Flag[string]{
			Name:     "sink",
			Usage:    "The address where the API response will be asynchronously delivered, using the HTTP protocol.",
			BodyPath: "sink",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sink-credential",
			Usage:    "A sink credential provides authentication or authorization information necessary to enable delivery of events to a target.",
			BodyPath: "sinkCredential",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handlePopulationdensitydataRetrieve,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"area": {
		&requestflag.InnerFlag[string]{
			Name:       "area.area-type",
			Usage:      "Type of this area.\nPOLYGON - The area is defined as a polygon.\n",
			InnerField: "areaType",
		},
	},
	"sink-credential": {
		&requestflag.InnerFlag[string]{
			Name:       "sink-credential.credential-type",
			Usage:      "The type of the credential.\nNote: Type of the credential - MUST be set to ACCESSTOKEN for now\n",
			InnerField: "credentialType",
		},
	},
})

func handlePopulationdensitydataRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.PopulationdensitydataGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Populationdensitydata.Get(ctx, params, options...)
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
		Title:          "populationdensitydata retrieve",
		Transform:      transform,
	})
}
