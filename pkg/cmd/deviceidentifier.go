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

var deviceidentifierRetrieveIdentifier = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve-identifier",
	Usage:   "Get details about the specific device being used by a given mobile subscriber",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "device",
			Usage:    "End-user equipment able to connect to a mobile network. Examples of devices include smartphones or IoT sensors/actuators.\nThe developer can choose to provide the below specified device identifiers:\n* `ipv4Address`\n* `ipv6Address`\n* `phoneNumber`\n* `networkAccessIdentifier`\nNOTE 1: The MNO might support only a subset of these options. The API invoker can provide multiple identifiers to be compatible across different MNOs. In this case the identifiers MUST belong to the same device.\nNOTE 2: For the current Commonalities release, we are enforcing that the networkAccessIdentifier is only part of the schema for future-proofing, and CAMARA does not currently allow its use. After the CAMARA meta-release work is concluded and the relevant issues are resolved, its use will need to be explicitly documented in the guidelines.\n",
			BodyPath: "device",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleDeviceidentifierRetrieveIdentifier,
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
			Usage:      "The device should be identified by the observed IPv6 address, or by any single IPv6 address from within the subnet allocated to the device (e.g. adding ::0 to the /64 prefix).\n",
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

var deviceidentifierRetrievePpid = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve-ppid",
	Usage:   "Get a pseudonymous identifier for device being used by a given mobile subscriber",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "device",
			Usage:    "End-user equipment able to connect to a mobile network. Examples of devices include smartphones or IoT sensors/actuators.\nThe developer can choose to provide the below specified device identifiers:\n* `ipv4Address`\n* `ipv6Address`\n* `phoneNumber`\n* `networkAccessIdentifier`\nNOTE 1: The MNO might support only a subset of these options. The API invoker can provide multiple identifiers to be compatible across different MNOs. In this case the identifiers MUST belong to the same device.\nNOTE 2: For the current Commonalities release, we are enforcing that the networkAccessIdentifier is only part of the schema for future-proofing, and CAMARA does not currently allow its use. After the CAMARA meta-release work is concluded and the relevant issues are resolved, its use will need to be explicitly documented in the guidelines.\n",
			BodyPath: "device",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleDeviceidentifierRetrievePpid,
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
			Usage:      "The device should be identified by the observed IPv6 address, or by any single IPv6 address from within the subnet allocated to the device (e.g. adding ::0 to the /64 prefix).\n",
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

var deviceidentifierRetrieveType = requestflag.WithInnerFlags(cli.Command{
	Name:    "retrieve-type",
	Usage:   "Get details about the type of device being used by a given mobile subscriber",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "device",
			Usage:    "End-user equipment able to connect to a mobile network. Examples of devices include smartphones or IoT sensors/actuators.\nThe developer can choose to provide the below specified device identifiers:\n* `ipv4Address`\n* `ipv6Address`\n* `phoneNumber`\n* `networkAccessIdentifier`\nNOTE 1: The MNO might support only a subset of these options. The API invoker can provide multiple identifiers to be compatible across different MNOs. In this case the identifiers MUST belong to the same device.\nNOTE 2: For the current Commonalities release, we are enforcing that the networkAccessIdentifier is only part of the schema for future-proofing, and CAMARA does not currently allow its use. After the CAMARA meta-release work is concluded and the relevant issues are resolved, its use will need to be explicitly documented in the guidelines.\n",
			BodyPath: "device",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleDeviceidentifierRetrieveType,
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
			Usage:      "The device should be identified by the observed IPv6 address, or by any single IPv6 address from within the subnet allocated to the device (e.g. adding ::0 to the /64 prefix).\n",
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

func handleDeviceidentifierRetrieveIdentifier(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.DeviceidentifierGetIdentifierParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Deviceidentifier.GetIdentifier(ctx, params, options...)
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
		Title:          "deviceidentifier retrieve-identifier",
		Transform:      transform,
	})
}

func handleDeviceidentifierRetrievePpid(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.DeviceidentifierGetPpidParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Deviceidentifier.GetPpid(ctx, params, options...)
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
		Title:          "deviceidentifier retrieve-ppid",
		Transform:      transform,
	})
}

func handleDeviceidentifierRetrieveType(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.DeviceidentifierGetTypeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Deviceidentifier.GetType(ctx, params, options...)
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
		Title:          "deviceidentifier retrieve-type",
		Transform:      transform,
	})
}
