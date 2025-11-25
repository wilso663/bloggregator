package main

import (
	"fmt"
	"github.com/wilso663/go-blog/internal/config"
)

type state struct {
	Cfg *config.Config
}

func NewState(cfg *config.Config) *state{
	return &state {
		Cfg: cfg,
	}
}

type Command struct {
	Name string
	Args	[]string
}

type Commands struct {
	handlers map[string]func(*state, Command) error
}

func NewCommands() *Commands {
	return &Commands {
		handlers: make(map[string]func(*state, Command)error),
	}
}

func (c *Commands) run(s *state, cmd Command) error {
	valFunc, exists := c.handlers[cmd.Name];
	if !exists {
		return fmt.Errorf("run command not found %s", cmd.Name);
	}
	err := valFunc(s, cmd); if err != nil {
		return fmt.Errorf("error running command %s", err);
	}
	return nil
}

func (c *Commands) register(name string, f func(*state, Command) error) {
	_, exists := c.handlers[name];
	if exists {
		fmt.Printf("%s command already registered", name);
	}
	c.handlers[name] = f;
}
