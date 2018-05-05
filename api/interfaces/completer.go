package interfaces

import (
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

var commands = []prompt.Suggest{
	// cbt commands
	{Text: "ls", Description: "List tables"},
	{Text: "lookup", Description: "Read from a single row"},
	{Text: "read", Description: "Read from a multi rows"},

	// btcli commands
	{Text: "exit", Description: "Exit this prompt"},
	{Text: "quit", Description: "Exit this prompt"},
}

var tables = []prompt.Suggest{
	{Text: "users", Description: "users"},
	{Text: "articles", Description: "articles"},
}

// Completer provide completion to prompt
func Completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")

	return completeWithArguments(args...)
}

func completeWithArguments(args ...string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]

	second := args[1]
	switch first {
	case "lookup", "read":
		return prompt.FilterHasPrefix(getTableSuggestions(), second, true)
	}

	return []prompt.Suggest{}
}

func getTableSuggestions() []prompt.Suggest {
	return tables
}