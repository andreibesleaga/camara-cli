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

var regiondevicecountGetCount = requestflag.WithInnerFlags(cli.Command{
	Name:    "get-count",
	Usage:   "Get the number of devices in the specified area during a certain time interval.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "area",
			BodyPath: "area",
		},
		&requestflag.Flag[any]{
			Name:     "endtime",
			Usage:    "Ending timestamp for counting the number of devices in the area. It must follow [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must have time zone.",
			BodyPath: "endtime",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filter",
			Usage:    "This parameter is used to filter devices. Currently, two filtering criteria are defined, `roamingStatus` and `deviceType`, which can be expanded in the future. `IN` logic is used used for multiple filtering items within a single filtering criterion, `AND` logic is used between multiple filtering criteria.\n- If a filtering critera is not provided, it means that there is no need to filter this item.\n- At least one of the criteria must be provided,a filter without any criteria is not allowed.\n- If no filtering is required, this parameter does not need to be provided.\nFor example ,`\"filter\":{\"roamingStatus\": [\"roaming\"],\"deviceType\": [\"human device\",\"IoT device\"]}` means the API need to return the count of human network devices and IoT devices that are in roaming mode.`\"filter\":{\"roamingStatus\": [\"non-roaming\"]}` means that the API need to return the count of all devices that are not in roaming mode.\n",
			BodyPath: "filter",
		},
		&requestflag.Flag[string]{
			Name:     "sink",
			Usage:    "The URL where the API response will be asynchronously delivered, using the HTTP protocol.",
			BodyPath: "sink",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sink-credential",
			Usage:    "A sink credential provides authentication or authorization information necessary to enable delivery of events to a target.",
			BodyPath: "sinkCredential",
		},
		&requestflag.Flag[any]{
			Name:     "starttime",
			Usage:    "Starting timestamp for counting the number of devices in the area. It must follow [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must have time zone.",
			BodyPath: "starttime",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleRegiondevicecountGetCount,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"area": {
		&requestflag.InnerFlag[string]{
			Name:       "area.area-type",
			Usage:      "Type of this area.\nCIRCLE - The area is defined as a circle.\nPOLYGON - The area is defined as a polygon.\n",
			InnerField: "areaType",
		},
	},
	"filter": {
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.device-type",
			Usage:      "Filtering by device type, 'human device' represents the need to filter for human network devices, 'IoT device' represents the need to filter for IoT devices, and 'other' represents the need to filter for other types of devices.",
			InnerField: "deviceType",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "filter.roaming-status",
			Usage:      "Filter whether the device is in roaming mode,'roaming' represents the need to filter devices that are in roaming mode,'non-roaming' represents the need to filter devices that are not roaming.",
			InnerField: "roamingStatus",
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

func handleRegiondevicecountGetCount(ctx context.Context, cmd *cli.Command) error {
	client := camara.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := camara.RegiondevicecountGetCountParams{}

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
	_, err = client.Regiondevicecount.GetCount(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "regiondevicecount get-count", obj, format, transform)
}
