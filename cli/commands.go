import (

	"fmt"
	"github.com/spf13/cobra"

)

var path string

// function to watch all pods logic
func CreateResourcesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "create k8s resources in the given path",
		Run: func(cmd *cobra.Command, args []string) error {
			createK8sResources(path)
		},		
	}
}

func init() {
	CreateResourcesCmd.Flags().StringVar(&path, "path", "", "Path to the directory containing Kubernetes YAML files")
	CreateResourcesCmd.MarkFlagRequired("path")
}


