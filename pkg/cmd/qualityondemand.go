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

var qualityondemandRetrieveQosProfile = cli.Command{
	Name:    "retrieve-qos-profile",
	Usage:   "Returns a QoS Profile that matches the given name.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "A unique name for identifying a specific QoS profile.\nThis may follow different formats depending on the service providers implementation.\nSome options addresses:\n  - A UUID style string\n  - Support for predefined profile names like `QOS_E`, `QOS_S`, `QOS_M`, and `QOS_L`\n  - A searchable descriptive name\n",
			Required:  true,
			PathParam: "name",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			Usage:      "Value for the x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleQualityondemandRetrieveQosProfile,
	HideHelpCommand: true,
}

var qualityondemandRetrieveQosProfiles = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve-qos-profiles",
	Usage:   "Returns all QoS Profiles that match the given criteria. **NOTES:**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "device",
			Usage:    "End-user equipment able to connect to a mobile network. Examples of devices include smartphones or IoT sensors/actuators.\n\nThe developer can choose to provide the below specified device identifiers:\n\n* `ipv4Address`\n* `ipv6Address`\n* `phoneNumber`\nNOTE1: the network operator might support only a subset of these options. The API consumer can provide multiple identifiers to be compatible across different operators. In this case the identifiers MUST belong to the same device.\nNOTE2: as for this Commonalities release, we are enforcing that the networkAccessIdentifier is only part of the schema for future-proofing, and CAMARA does not currently allow its use. After the CAMARA meta-release work is concluded and the relevant issues are resolved, its use will need to be explicitly documented in the guidelines.\n",
			BodyPath: "device",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "A unique name for identifying a specific QoS profile.\nThis may follow different formats depending on the service providers implementation.\nSome options addresses:\n  - A UUID style string\n  - Support for predefined profile names like `QOS_E`, `QOS_S`, `QOS_M`, and `QOS_L`\n  - A searchable descriptive name\n",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The current status of the QoS Profile\n- `ACTIVE`- QoS Profile is available to be used\n- `INACTIVE`- QoS Profile is not currently available to be deployed\n- `DEPRECATED`- QoS profile is actively being used in a QoD session, but can not be deployed in new QoD sessions\n",
			BodyPath: "status",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			Usage:      "Value for the x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleQualityondemandRetrieveQosProfiles,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"device": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "device.ipv4-address",
			Usage:      "The device should be identified by either the public (observed) IP address and port as seen by the application server, or the private (local) and any public (observed) IP addresses in use by the device (this information can be obtained by various means, for example from some DNS servers).\n\nIf the allocated and observed IP addresses are the same (i.e. NAT is not in use) then  the same address should be specified for both publicAddress and privateAddress.\n\nIf NAT64 is in use, the device should be identified by its publicAddress and publicPort, or separately by its allocated IPv6 address (field ipv6Address of the Device object)\n\nIn all cases, publicAddress must be specified, along with at least one of either privateAddress or publicPort, dependent upon which is known. In general, mobile devices cannot be identified by their public IPv4 address alone.\n",
			InnerField: "ipv4Address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "device.ipv6-address",
			Usage:      "The device should be identified by the observed IPv6 address, or by any single IPv6 address from within the subnet allocated to the device (e.g. adding ::0 to the /64 prefix).\n\nThe session shall apply to all IP flows between the device subnet and the specified application server, unless further restricted by the optional parameters devicePorts or applicationServerPorts.\n",
			InnerField: "ipv6Address",
		},
		&requestflag.InnerFlag[string]{
			Name:       "device.network-access-identifier",
			Usage:      "A public identifier addressing a subscription in a mobile network. In 3GPP terminology, it corresponds to the GPSI formatted with the External Identifier ({Local Identifier}@{Domain Identifier}). Unlike the telephone number, the network access identifier is not subjected to portability ruling in force, and is individually managed by each operator.",
			InnerField: "networkAccessIdentifier",
		},
		&requestflag.InnerFlag[string]{
			Name:       "device.phone-number",
			Usage:      "A public identifier addressing a telephone subscription. In mobile networks it corresponds to the MSISDN (Mobile Station International Subscriber Directory Number). In order to be globally unique it has to be formatted in international format, according to E.164 standard, prefixed with '+'.",
			InnerField: "phoneNumber",
		},
	},
})

func handleQualityondemandRetrieveQosProfile(ctx context.Context, cmd *cli.Command) error {
	client := camara.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("name") && len(unusedArgs) > 0 {
		cmd.Set("name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := camara.QualityondemandGetQosProfileParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Qualityondemand.GetQosProfile(
		ctx,
		cmd.Value("name").(string),
		params,
		options...,
	)
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
		Title:          "qualityondemand retrieve-qos-profile",
		Transform:      transform,
	})
}

func handleQualityondemandRetrieveQosProfiles(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.QualityondemandGetQosProfilesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Qualityondemand.GetQosProfiles(ctx, params, options...)
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
		Title:          "qualityondemand retrieve-qos-profiles",
		Transform:      transform,
	})
}
