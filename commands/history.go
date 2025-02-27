package commands

import (
	"Go-Shell/models"
	"errors"
	"fmt"
	"sort"
)

func History(args []string, user *models.User) error {
	if len(args) > 1 {
		return errors.New("usage: history [clean]")
	}
	if len(args) == 1 && args[0] == "clean" {
		user.ClearHistory()
		fmt.Println("empty command history")
		return nil
	}

	history := user.GetHistory()
	commandCount := make(map[string]int)
	for _, cmd := range history {
		commandCount[cmd.Name]++
	}

	type countCommand struct {
		Command string
		Count   int
	}

	var countCommands []countCommand
	for cmd, count := range commandCount {
		countCommands = append(countCommands, countCommand{cmd, count})
	}

	sort.Slice(countCommands, func(i, j int) bool {
		if countCommands[i].Count == countCommands[j].Count {
			return history[len(history)-1].Timestamp.After(history[len(history)-1].Timestamp)
		}
		return countCommands[i].Count > countCommands[j].Count
	})

	for _, cc := range countCommands {
		fmt.Printf("| %s | %d |\n", cc.Command, cc.Count)
	}
	return nil
}
