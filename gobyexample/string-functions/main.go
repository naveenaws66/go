package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Containes: 	", s.Contains("test", "es"))
	p("Count: 	", s.Count("test", "t"))
	p("HasPrefix: 	", s.HasPrefix("test", "te"))
	p("HasSuffix: 	", s.HasSuffix("test", "st"))
	p("Index: 	", s.Index("test", "e"))
	p("Join: 	", s.Join([]string{"a", "b"}, "_"))
	p("Repeat: 	", s.Repeat("a", 5))
	p("Split: 	", s.Split("a-b-c-d-e", "-"))
	p("ToLower: 	", s.ToLower("TEST"))
	p("ToUpper: 	", s.ToUpper("test"))
}
