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

var connectivityinsightsSubscriptionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Connectivity insights subscription for a device",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Implementation-specific configuration parameters needed by the\nsubscription manager for acquiring events.\nIn CAMARA we have predefined attributes like `subscriptionExpireTime`,\n`subscriptionMaxEvents`, `initialEvent`\nSpecific event type attributes must be defined in `subscriptionDetail`\nNote: if a request is performed for several event type, all subscribed\nevent will use same `config` parameters.\n",
			Required: true,
			BodyPath: "config",
		},
		&requestflag.Flag[string]{
			Name:     "protocol",
			Usage:    "Identifier of a delivery protocol. Only HTTP is allowed for now\n",
			Required: true,
			BodyPath: "protocol",
		},
		&requestflag.Flag[string]{
			Name:     "sink",
			Usage:    "The address to which events shall be delivered using the selected\nprotocol.\n",
			Required: true,
			BodyPath: "sink",
		},
		&requestflag.Flag[[]string]{
			Name:     "type",
			Usage:    "Camara Event types eligible to be delivered by this subscription.\n",
			Required: true,
			BodyPath: "types",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sink-credential",
			Usage:    "A sink credential provides authentication or authorization information\n",
			BodyPath: "sinkCredential",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleConnectivityinsightsSubscriptionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"config": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "config.subscription-detail",
			Usage:      "The detail of the requested event subscription",
			InnerField: "subscriptionDetail",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "config.initial-event",
			Usage:      "Set to `true` by API consumer if consumer wants to get an event as\nsoon as the subscription is created and current situation reflects\nevent request.\n",
			InnerField: "initialEvent",
		},
		&requestflag.InnerFlag[any]{
			Name:       "config.subscription-expire-time",
			Usage:      "The subscription expiration time (in date-time format) requested by\nthe API consumer. Up to API project decision to keep it.\nIt must follow [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must have time zone.\n",
			InnerField: "subscriptionExpireTime",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "config.subscription-max-events",
			Usage:      "Identifies the maximum number of event reports to be generated\n(>=1) requested by the API consumer - Once this number is reached,\nthe subscription ends.\nNote on combined usage of `initialEvent` and\n`subscriptionMaxEvents`: If an event is triggered following\n`initialEvent` set to `true`, this event will be counted towards\n`subscriptionMaxEvents`.\n",
			InnerField: "subscriptionMaxEvents",
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

var connectivityinsightsSubscriptionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a given subscription by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "subscription-id",
			Usage:     "When this information is contained within an event notification, it SHALL be referred to as `subscriptionId` as per the Commonalities Event Notification Model.\n",
			Required:  true,
			PathParam: "subscriptionId",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleConnectivityinsightsSubscriptionsRetrieve,
	HideHelpCommand: true,
}

var connectivityinsightsSubscriptionsList = cli.Command{
	Name:    "list",
	Usage:   "Operation to list subscriptions authorized to be retrieved by the provided\naccess token.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleConnectivityinsightsSubscriptionsList,
	HideHelpCommand: true,
}

var connectivityinsightsSubscriptionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a given subscription by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "subscription-id",
			Usage:     "When this information is contained within an event notification, it SHALL be referred to as `subscriptionId` as per the Commonalities Event Notification Model.\n",
			Required:  true,
			PathParam: "subscriptionId",
		},
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleConnectivityinsightsSubscriptionsDelete,
	HideHelpCommand: true,
}

func handleConnectivityinsightsSubscriptionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.ConnectivityinsightSubscriptionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectivityinsights.Subscriptions.New(ctx, params, options...)
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
		Title:          "connectivityinsights:subscriptions create",
		Transform:      transform,
	})
}

func handleConnectivityinsightsSubscriptionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.ConnectivityinsightSubscriptionGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectivityinsights.Subscriptions.Get(
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
		Title:          "connectivityinsights:subscriptions retrieve",
		Transform:      transform,
	})
}

func handleConnectivityinsightsSubscriptionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.ConnectivityinsightSubscriptionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectivityinsights.Subscriptions.List(ctx, params, options...)
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
		Title:          "connectivityinsights:subscriptions list",
		Transform:      transform,
	})
}

func handleConnectivityinsightsSubscriptionsDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.ConnectivityinsightSubscriptionDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectivityinsights.Subscriptions.Delete(
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
		Title:          "connectivityinsights:subscriptions delete",
		Transform:      transform,
	})
}
