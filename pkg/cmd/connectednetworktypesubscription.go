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

var connectednetworktypeSubscriptionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a subscription for receiving notifications on changes to the connected\nnetwork type of a device.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Implementation-specific configuration parameters needed by the subscription manager for acquiring events.\nIn CAMARA we have predefined attributes like `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent`\nSpecific event type attributes must be defined in `subscriptionDetail`\nNote: if a request is performed for several event type, all subscribed event will use same `config` parameters.\n",
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
			Usage:    "Camara Event types eligible to be delivered by this subscription.\nNote: As of now we enforce to have only event type per subscription.\n",
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
	Action:          handleConnectednetworktypeSubscriptionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"config": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "config.subscription-detail",
			Usage:      "The detail of the requested event subscription.",
			InnerField: "subscriptionDetail",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "config.initial-event",
			Usage:      "Set to `true` by API consumer if consumer wants to get an event as soon as the subscription is created and current situation reflects event request.\nExample: Consumer request area entered event. If consumer sets initialEvent to true and device is already in the geofence, an event is triggered\n",
			InnerField: "initialEvent",
		},
		&requestflag.InnerFlag[any]{
			Name:       "config.subscription-expire-time",
			Usage:      "The subscription expiration time (in date-time format) requested by the API consumer.",
			InnerField: "subscriptionExpireTime",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "config.subscription-max-events",
			Usage:      "Identifies the maximum number of event reports to be generated (>=1) requested by the API consumer - Once this number is reached, the subscription ends.",
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

var connectednetworktypeSubscriptionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "retrieve ConnectedNetworkType subscription information for a given subscription\nID.",
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
	Action:          handleConnectednetworktypeSubscriptionsRetrieve,
	HideHelpCommand: true,
}

var connectednetworktypeSubscriptionsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a list of device connected network type event subscription(s)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleConnectednetworktypeSubscriptionsList,
	HideHelpCommand: true,
}

var connectednetworktypeSubscriptionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "delete a given ConnectedNetworkType subscription.",
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
	Action:          handleConnectednetworktypeSubscriptionsDelete,
	HideHelpCommand: true,
}

func handleConnectednetworktypeSubscriptionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.ConnectednetworktypeSubscriptionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectednetworktype.Subscriptions.New(ctx, params, options...)
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
		Title:          "connectednetworktype:subscriptions create",
		Transform:      transform,
	})
}

func handleConnectednetworktypeSubscriptionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.ConnectednetworktypeSubscriptionGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectednetworktype.Subscriptions.Get(
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
		Title:          "connectednetworktype:subscriptions retrieve",
		Transform:      transform,
	})
}

func handleConnectednetworktypeSubscriptionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.ConnectednetworktypeSubscriptionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectednetworktype.Subscriptions.List(ctx, params, options...)
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
		Title:          "connectednetworktype:subscriptions list",
		Transform:      transform,
	})
}

func handleConnectednetworktypeSubscriptionsDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.ConnectednetworktypeSubscriptionDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectednetworktype.Subscriptions.Delete(
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
		Title:          "connectednetworktype:subscriptions delete",
		Transform:      transform,
	})
}
