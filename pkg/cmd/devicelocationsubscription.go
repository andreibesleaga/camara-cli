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

var devicelocationSubscriptionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a subscription for a device to receive notifications when the device\nenters or exits a specified area.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Implementation-specific configuration parameters are needed by the subscription manager for acquiring events.\nIn CAMARA we have predefined attributes like `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent`.\n",
			Required: true,
			BodyPath: "config",
		},
		&requestflag.Flag[string]{
			Name:     "protocol",
			Usage:    "Identifier of a delivery protocol. Only HTTP is allowed for now.",
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
			Usage:    "Camara Event types which are eligible to be delivered by this subscription.\nNote: As of now we enforce to have only event type per subscription.\n",
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
	Action:          handleDevicelocationSubscriptionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"sink-credential": {
		&requestflag.InnerFlag[string]{
			Name:       "sink-credential.credential-type",
			Usage:      "The type of the credential.\nNote: Type of the credential - MUST be set to ACCESSTOKEN for now\n",
			InnerField: "credentialType",
		},
	},
})

var devicelocationSubscriptionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve Geofencing subscription information for a given subscription ID.",
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
	Action:          handleDevicelocationSubscriptionsRetrieve,
	HideHelpCommand: true,
}

var devicelocationSubscriptionsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieve a list of geofencing event subscription(s).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:       "x-correlator",
			HeaderPath: "x-correlator",
		},
	},
	Action:          handleDevicelocationSubscriptionsList,
	HideHelpCommand: true,
}

var devicelocationSubscriptionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a given Geofencing subscription.",
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
	Action:          handleDevicelocationSubscriptionsDelete,
	HideHelpCommand: true,
}

func handleDevicelocationSubscriptionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.DevicelocationSubscriptionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Devicelocation.Subscriptions.New(ctx, params, options...)
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
		Title:          "devicelocation:subscriptions create",
		Transform:      transform,
	})
}

func handleDevicelocationSubscriptionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.DevicelocationSubscriptionGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Devicelocation.Subscriptions.Get(
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
		Title:          "devicelocation:subscriptions retrieve",
		Transform:      transform,
	})
}

func handleDevicelocationSubscriptionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.DevicelocationSubscriptionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Devicelocation.Subscriptions.List(ctx, params, options...)
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
		Title:          "devicelocation:subscriptions list",
		Transform:      transform,
	})
}

func handleDevicelocationSubscriptionsDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := camara.DevicelocationSubscriptionDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Devicelocation.Subscriptions.Delete(
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
		Title:          "devicelocation:subscriptions delete",
		Transform:      transform,
	})
}
