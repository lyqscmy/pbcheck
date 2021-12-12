package test

import cmd "github.com/lyqscmy/pbcheck/proto"

func foo(p *cmd.Person) string {
	s := 18
	p.Age = &s
	return *p.Name
}
