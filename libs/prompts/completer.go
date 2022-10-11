package prompts

import (
	prompt "github.com/c-bata/go-prompt"
)

func Chain(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "bitcoin", Description: "bitcoin"},
		{Text: "ethereum", Description: "ethereum classic"},
		{Text: "end", Description: "end chain select"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func ScanType(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "blocks", Description: "Search blocks date after YY-MM-DD"},
		{Text: "transactions", Description: "Search transactions by block num"},
		{Text: "transaction by hash", Description: "Search transaction by hash"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
