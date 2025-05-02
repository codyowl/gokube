package utilities

import (
	"fmt"
	"strings"
	"os/exec"
)


// helper function to read and create resources from a directory
func Createk8sResources(rootDir, namespace string) error {
	
}

// helper function to create a namespace if not exists
func createNamespaceIfNotExist(namespace string) error {
	getNamespacecmd := exec.Command("kubectl", "get", "namespace", namespace)
	if err := getNamespacecmd.Run(); err != nil {
		fmt.Printf("Namespace %s not found", namespace)
		createNamespaceCmd := exec.Command("kubectl", "create", "namespace", namespace)
		createNamespaceCmd.Stdout = os.Stdout
		createNamespaceCmd.Stderr = os.Stderr
		return createNamespaceCmd.Run()
	}
	fmt.Printf("Namespace %s found, using existing namespace", namespace)
	return nil
}