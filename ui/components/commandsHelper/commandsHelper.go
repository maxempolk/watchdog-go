package commandshelper

import (
	"fmt"
	"stat_by_sites/internal/commands"
	"stat_by_sites/ui/components/base"
	"strings"
)

type CommandHelper struct{
	base.Component
	Commands []commands.Command
}

func NewCommandHelper(width int, commands []commands.Command) *CommandHelper{
	return &CommandHelper{
		base.Component{Width: width},
		commands,
	}
}

func (ch *CommandHelper) View() string{
	var sb strings.Builder;
	sb.Grow(100)

	for _, v := range ch.Commands {
		sb.WriteString( fmt.Sprintf("[%s] %s  ", v.Key, v.Description) )
	}

	return sb.String()
}