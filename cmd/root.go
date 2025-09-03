package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

var extractKinds string
var extractName string

func init() {
	rootCmd.Flags().StringVar(&extractKinds, "kind", "", "the kind of object you want to extract")
	_ = rootCmd.MarkFlagRequired("kind")
	rootCmd.Flags().StringVar(&extractName, "name", "", "the name of object you want to extract")
}

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "k8s-yaml-extract --kind=Deployment [--name=deploy-name]",
	Short: "This CLI tool allows extracting Kubernetes YAMLs from YAML lists based on specified criteria",
	Run: func(cmd *cobra.Command, args []string) {

		// catch any panics
		defer func() {
			if err := recover(); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "unexpected panic: \n\t%v\n", err)
				os.Exit(1)
			}
		}()

		var inputReader = cmd.InOrStdin()

		// the argument received looks like a file, we try to open it
		if len(args) > 0 {
			if args[0] != "-" && args[0] != "" {
				file, err := os.Open(args[0])
				if err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "failed open file: %v", err)
					os.Exit(1)
				}
				inputReader = file
			}
		}

		rw := kio.ByteReadWriter{
			Reader:            inputReader,
			Writer:            os.Stdout,
			WrapBareSeqNode:   true,
			PreserveSeqIndent: true,
		}

		err := kio.Pipeline{
			Inputs: []kio.Reader{&rw},
			Filters: []kio.Filter{
				ExtractFilter{},
			},
			Outputs: []kio.Writer{&rw},
		}.Execute()

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "[ERROR] while parsing file: %s\n", err)
			os.Exit(1)
		}

		os.Exit(0)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

type ExtractFilter struct{}

func (f ExtractFilter) Filter(slice []*yaml.RNode) ([]*yaml.RNode, error) {
	// only add to slice what you want to change
	var newSlices []*yaml.RNode

	for i := range slice {
		kindNode, err := slice[i].Pipe(yaml.Get("kind"))
		if err != nil {
			return nil, fmt.Errorf("error while getting RNode kind: %w", err)
		}
		if kindNode == nil {
			continue
		}

		kind := kindNode.YNode().Value
		if !strings.EqualFold(kind, extractKinds) {
			continue
		}

		if extractName != "" && !strings.EqualFold(slice[i].GetName(), extractName) {
			continue
		}

		newSlices = append(newSlices, slice[i])
	}

	return newSlices, nil
}
