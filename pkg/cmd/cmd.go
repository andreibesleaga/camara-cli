// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command *cli.Command
)

func init() {
	Command = &cli.Command{
		Name:    "camara",
		Usage:   "CLI for the camara API",
		Suggest: true,
		Version: Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "customerinsights:scoring",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&customerinsightsScoringRetrieve,
				},
			},
			{
				Name:     "deviceswap",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&deviceswapCheck,
					&deviceswapRetrieveDate,
				},
			},
			{
				Name:     "knowyourcustomerageverification",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&knowyourcustomerageverificationVerify,
				},
			},
			{
				Name:     "knowyourcustomerfill-in",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&knowyourcustomerfillInCreate,
				},
			},
			{
				Name:     "knowyourcustomermatch",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&knowyourcustomermatchMatch,
				},
			},
			{
				Name:     "tenure",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&tenureVerify,
				},
			},
			{
				Name:     "numberrecycling",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&numberrecyclingCheckSubscriberChange,
				},
			},
			{
				Name:     "otpvalidation",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&otpvalidationSendCode,
				},
			},
			{
				Name:     "callforwardingsignal",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&callforwardingsignalCheckActiveForwardings,
					&callforwardingsignalCheckUnconditionalForwarding,
				},
			},
			{
				Name:     "devicelocation:subscriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&devicelocationSubscriptionsCreate,
					&devicelocationSubscriptionsRetrieve,
					&devicelocationSubscriptionsList,
					&devicelocationSubscriptionsDelete,
				},
			},
			{
				Name:     "populationdensitydata",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&populationdensitydataRetrieve,
				},
			},
			{
				Name:     "regiondevicecount",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&regiondevicecountGetCount,
				},
			},
			{
				Name:     "webrtc:sessions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&webrtcSessionsCreate,
					&webrtcSessionsRetrieve,
					&webrtcSessionsUpdateStatus,
				},
			},
			{
				Name:     "connectivityinsights:subscriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&connectivityinsightsSubscriptionsCreate,
					&connectivityinsightsSubscriptionsRetrieve,
					&connectivityinsightsSubscriptionsList,
					&connectivityinsightsSubscriptionsDelete,
				},
			},
			{
				Name:     "qualityondemand",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&qualityondemandRetrieveQosProfile,
					&qualityondemandRetrieveQosProfiles,
				},
			},
			{
				Name:     "deviceidentifier",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&deviceidentifierRetrieveIdentifier,
					&deviceidentifierRetrievePpid,
					&deviceidentifierRetrieveType,
				},
			},
			{
				Name:     "simswap:subscriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&simswapSubscriptionsCreate,
					&simswapSubscriptionsRetrieve,
					&simswapSubscriptionsList,
					&simswapSubscriptionsDelete,
				},
			},
			{
				Name:     "deviceroamingstatus:subscriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&deviceroamingstatusSubscriptionsCreate,
					&deviceroamingstatusSubscriptionsRetrieve,
					&deviceroamingstatusSubscriptionsList,
					&deviceroamingstatusSubscriptionsDelete,
				},
			},
			{
				Name:     "devicereachabilitystatus:subscriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&devicereachabilitystatusSubscriptionsCreate,
					&devicereachabilitystatusSubscriptionsRetrieve,
					&devicereachabilitystatusSubscriptionsList,
					&devicereachabilitystatusSubscriptionsDelete,
				},
			},
			{
				Name:     "connectednetworktype:subscriptions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&connectednetworktypeSubscriptionsCreate,
					&connectednetworktypeSubscriptionsRetrieve,
					&connectednetworktypeSubscriptionsList,
					&connectednetworktypeSubscriptionsDelete,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "camara @manpages [-o camara.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
		},
		EnableShellCompletion:      true,
		ShellCompletionCommandName: "@completion",
		HideHelpCommand:            true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "camara.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "camara.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
