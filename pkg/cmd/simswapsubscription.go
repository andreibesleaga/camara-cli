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

var simswapSubscriptionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a sim swap event subscription for a phone number",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Implementation-specific configuration parameters needed by the subscription manager for acquiring events.\nIn CAMARA we have predefined attributes like `subscriptionExpireTime` or `subscriptionMaxEvents` to limit subscription lifetime.\nEvent type attributes must be defined in `subscriptionDetail`\n",
			Required: true,
			BodyPath: "config",
		},
		&requestflag.Flag[string]{
			Name:     "protocol",
			Usage:    "Identifier of a delivery protocol. Only HTTP is allowed for now",
			Required: true,
			BodyPath: "protocol",
		},
		&requestflag.Flag[string]{
			Name:     "sink",
			Usage:    "The address to which events shall be delivered using the selected protocol.",
			Required: true,
			BodyPath: "sink",
		},
		&requestflag.Flag[[]string]{
			Name:     "type",
			Usage:    "Camara Event types eligible for subscription:\n- org.camaraproject.sim-swap-subscriptions.v0.swapped: receive a notification when a sim swap is performed on the line.\n",
			Required: true,
			BodyPath: "types",
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
	Action:          handleSimswapSubscriptionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"config": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "config.subscription-detail",
			Usage:      "The detail of the requested event subscription",
			InnerField: "subscriptionDetail",
		},
		&requestflag.InnerFlag[any]{
			Name:       "config.subscription-expire-time",
			Usage:      "The subscription expiration time (in date-time format) requested by the API consumer. It must follow [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must have time zone.",
			InnerField: "subscriptionExpireTime",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "config.subscription-max-events",
			Usage:      "Identifies the maximum number of event reports to be generated (>=1) requested by the API consumer - Once this number is reached, the subscription ends.\n",
			InnerField: "subscriptionMaxEvents",
		},
	},
	"sink-credential": {
		&requestflag.InnerFlag[string]{
			Name:       "sink-credential.credential-type",
			Usage:      "The type of the credential. With the current API version the type MUST be set to ACCESSTOKEN.",
			InnerField: "credentialType",
		},
	},
})

var simswapSubscriptionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "retrieve event subscription information for a given subscription.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "subscription-id",
			Usage:     "The unique identifier of the subscription in the scope of the subscription manager. When this information is contained within an event notification, this concept SHALL be referred as subscriptionId as per Commonalities Event Notification Model.",
			Required:  true,
			PathParam: "subscriptionId",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleSimswapSubscriptionsRetrieve,
	HideHelpCommand: true,
}

var simswapSubscriptionsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a list of sim swap event subscription(s)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleSimswapSubscriptionsList,
	HideHelpCommand: true,
}

var simswapSubscriptionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "delete a given event subscription.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "subscription-id",
			Usage:     "The unique identifier of the subscription in the scope of the subscription manager. When this information is contained within an event notification, this concept SHALL be referred as subscriptionId as per Commonalities Event Notification Model.",
			Required:  true,
			PathParam: "subscriptionId",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleSimswapSubscriptionsDelete,
	HideHelpCommand: true,
}

func handleSimswapSubscriptionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.SimswapSubscriptionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simswap.Subscriptions.New(ctx, params, options...)
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
		Title:          "simswap:subscriptions create",
		Transform:      transform,
	})
}

func handleSimswapSubscriptionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := camara.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
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

	params := camara.SimswapSubscriptionGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simswap.Subscriptions.Get(
		ctx,
		cmd.Value("subscription-id").(string),
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
		Title:          "simswap:subscriptions retrieve",
		Transform:      transform,
	})
}

func handleSimswapSubscriptionsList(ctx context.Context, cmd *cli.Command) error {
	client := camara.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := camara.SimswapSubscriptionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simswap.Subscriptions.List(ctx, params, options...)
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
		Title:          "simswap:subscriptions list",
		Transform:      transform,
	})
}

func handleSimswapSubscriptionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := camara.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
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

	params := camara.SimswapSubscriptionDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Simswap.Subscriptions.Delete(
		ctx,
		cmd.Value("subscription-id").(string),
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
		Title:          "simswap:subscriptions delete",
		Transform:      transform,
	})
}
