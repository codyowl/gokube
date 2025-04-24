import (

	"fmt"
	"github.com/spf13/cobra"

)

var path string

// command for create resources
func CreateResourcesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "-create [path]",
		Short: "Create Kubernetes resources based on YAML filenames and folder names",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("please provide a path to your k8s yaml files")
			}
			path := args[0]

			return createK8sResources(path)
		},
	}
	return cmd
}


