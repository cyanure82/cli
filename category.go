package cli

type CommandCategories interface {
	// AddCommand adds a command to a category, creating a new category if necessary.
	AddCommand(category string, command *Command)
	// Categories returns a copy of the category slice
	Categories() []CommandCategory
}

type commandCategories []*commandCategory

func newCommandCategories() CommandCategories {
	ret := commandCategories([]*commandCategory{})
	return &ret
}

func (c *commandCategories) Less(i, j int) bool {
	return (*c)[i].Name() < (*c)[j].Name()
}

func (c *commandCategories) Len() int {
	return len(*c)
}

func (c *commandCategories) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *commandCategories) AddCommand(category string, command *Command) {
	for _, commandCategory := range []*commandCategory(*c) {
		if commandCategory.name == category {
			commandCategory.commands = append(commandCategory.commands, command)
			return
		}
	}
	newVal := commandCategories(append(*c,
		&commandCategory{name: category, commands: []*Command{command}}))
	(*c) = newVal
}

func (c *commandCategories) Categories() []CommandCategory {
	ret := []CommandCategory{}
	for _, cat := range *c {
		ret = append(ret, cat)
	}
	return ret
}

// CommandCategory is a category containing commands.
type CommandCategory interface {
	// Name returns the category name string
	Name() string
	// VisibleCommands returns a slice of the Commands with Hidden=false
	VisibleCommands() []*Command
}

type commandCategory struct {
	name     string
	commands []*Command
}

func newCommandCategory(name string) *commandCategory {
	return &commandCategory{
		name:     name,
		commands: []*Command{},
	}
}

func (c *commandCategory) Name() string {
	return c.name
}

func (c *commandCategory) VisibleCommands() []*Command {
	if c.commands == nil {
		c.commands = []*Command{}
	}

	ret := []*Command{}
	for _, command := range c.commands {
		if !command.Hidden {
			ret = append(ret, command)
		}
	}
	return ret
}
