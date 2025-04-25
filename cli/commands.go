import (

	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"

)

var (
	k8sRootDirectory string
	namespace string
)

// command for create resources
func CreateResourcesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <root_directory>",
		Short: "Create Kubernetes resources based on YAML filenames and folder names",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("please provide a root directory path to your k8s yaml files")
			}
			k8sRootDirectory := args[0]

			// namespace flag or fall back to root directoryname 
			ns := namespace
			if ns == "" {
				ns = filepath.Base(k8sRootDirectory)
				fmt.Printf("Namespace not provided. Using directory name as namespace: %s\n", ns)
			}

			return createK8sResources(k8sRootDirectory, ns)
		},
	}
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Kubernetes namespace to use")
	return cmd
}




