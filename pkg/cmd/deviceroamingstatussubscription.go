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

var deviceroamingstatusSubscriptionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a device roaming status event subscription for a device",
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
			Usage:    "Camara Event types eligible to be delivered by this subscription.\nNote: for the current Commonalities version (v0.5) only one event type per subscription is allowed, yet in the following releases use of array of event types SHALL be specified without changing this definition.\n",
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
	Action:          handleDeviceroamingstatusSubscriptionsCreate,
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
			Usage:      "Set to `true` by API consumer if consumer wants to get an event as soon as the subscription is created and current situation reflects event request.\nExample: Consumer request Roaming event. If consumer sets initialEvent to true and device is in roaming situation, an event is triggered.\n",
			InnerField: "initialEvent",
		},
		&requestflag.InnerFlag[any]{
			Name:       "config.subscription-expire-time",
			Usage:      "The subscription expiration time (in date-time format) requested by the API consumer. It must follow [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must have time zone.",
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

var deviceroamingstatusSubscriptionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "retrieve device roaming status subscription information for a given\nsubscription.",
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
	Action:          handleDeviceroamingstatusSubscriptionsRetrieve,
	HideHelpCommand: true,
}

var deviceroamingstatusSubscriptionsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a list of device roaming status event subscription(s)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleDeviceroamingstatusSubscriptionsList,
	HideHelpCommand: true,
}

var deviceroamingstatusSubscriptionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a given device-roaming-status subscription by ID",
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
	Action:          handleDeviceroamingstatusSubscriptionsDelete,
	HideHelpCommand: true,
}

func handleDeviceroamingstatusSubscriptionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.DeviceroamingstatusSubscriptionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Deviceroamingstatus.Subscriptions.New(ctx, params, options...)
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
		Title:          "deviceroamingstatus:subscriptions create",
		Transform:      transform,
	})
}

func handleDeviceroamingstatusSubscriptionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.DeviceroamingstatusSubscriptionGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Deviceroamingstatus.Subscriptions.Get(
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
		Title:          "deviceroamingstatus:subscriptions retrieve",
		Transform:      transform,
	})
}

func handleDeviceroamingstatusSubscriptionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.DeviceroamingstatusSubscriptionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Deviceroamingstatus.Subscriptions.List(ctx, params, options...)
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
		Title:          "deviceroamingstatus:subscriptions list",
		Transform:      transform,
	})
}

func handleDeviceroamingstatusSubscriptionsDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.DeviceroamingstatusSubscriptionDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Deviceroamingstatus.Subscriptions.Delete(
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
		Title:          "deviceroamingstatus:subscriptions delete",
		Transform:      transform,
	})
}
