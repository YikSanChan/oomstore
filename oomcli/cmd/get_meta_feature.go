package cmd

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/oom-ai/oomstore/pkg/errdefs"
	"github.com/spf13/cobra"

	"github.com/oom-ai/oomstore/pkg/oomstore"
	"github.com/oom-ai/oomstore/pkg/oomstore/types"
	"github.com/oom-ai/oomstore/pkg/oomstore/types/apply"
)

var getMetaFeatureOpt types.ListFeatureOpt
var getMetaFeatureEntityName, getMetaFeatureGroupName *string

var getMetaFeatureCmd = &cobra.Command{
	Use:   "feature [feature_name]",
	Short: "Get existing features given specific conditions",
	Args:  cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("entity") {
			getMetaFeatureOpt.EntityNames = &[]string{*getMetaFeatureEntityName}
		}
		if cmd.Flags().Changed("group") {
			getMetaFeatureOpt.GroupNames = &[]string{*getMetaFeatureGroupName}
		}
		if len(args) == 1 {
			getMetaFeatureOpt.FeatureNames = &[]string{args[0]}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		oomStore := mustOpenOomStore(ctx, oomStoreCfg)
		defer oomStore.Close()

		features, err := queryFeatures(ctx, oomStore, getMetaFeatureOpt)
		if err != nil {
			exit(err)
		}

		if err := outputFeature(os.Stdout, features, *getMetaOutput); err != nil {
			exitf("failed printing features: %+v\n", err)
		}
	},
}

func init() {
	getMetaCmd.AddCommand(getMetaFeatureCmd)

	flags := getMetaFeatureCmd.Flags()
	getMetaFeatureEntityName = flags.StringP("entity", "e", "", "entity")
	getMetaFeatureGroupName = flags.StringP("group", "g", "", "feature group")
}

func queryFeatures(ctx context.Context, oomStore *oomstore.OomStore, opt types.ListFeatureOpt) (types.FeatureList, error) {
	features, err := oomStore.ListFeature(ctx, opt)
	if err != nil {
		return nil, fmt.Errorf("failed getting features, error %v\n", err)
	}

	if opt.FeatureNames != nil && len(features) == 0 {
		return nil, errdefs.Errorf("feature '%s' not found", (*opt.FeatureNames)[0])
	}

	return features, nil
}

func outputFeature(w io.Writer, features types.FeatureList, outputOpt string) error {
	switch outputOpt {
	case YAML:
		return serializeInYaml(w, apply.BuildFeatureItems(features))
	default:
		return serializeMetadata(w, features, outputOpt, *getMetaWide)
	}
}
