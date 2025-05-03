package utilities

import (
	"os"
	"gopkg.in/yaml.v3"
)

commandsYamlPath = "./commands.yaml"


// for getting the command string out of commands.yaml
func GetCommand(key string) string {
	cmd := CmdsYamlParser[key]
	return cmd
}