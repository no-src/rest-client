package internal

import (
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

var config map[string]string

func initConfig(configFile string) error {
	if len(configFile) == 0 {
		return nil
	}
	data, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &config)
}

func parseVariables(text string) string {
	re := regexp.MustCompile(`\{\{[^\}]*\}\}`)
	matches := re.FindAllString(text, -1)
	for _, match := range matches {
		field := strings.TrimLeft(match, "{{")
		field = strings.TrimRight(field, "}}")
		if value, ok := config[field]; ok {
			text = strings.ReplaceAll(text, match, value)
		}
	}
	return text
}
