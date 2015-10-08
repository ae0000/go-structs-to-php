package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gocraft/web"
)

const (
	defaultRawStruct = `
type Ha struct {
	ID          int64
	Name        string
	Amount      float32
	Version     int64
	Alive       bool
}`
)

func (c *Context) convertStruct(rw web.ResponseWriter, req *web.Request) {

	// Check for POST
	if req.Request.Method == "POST" {
		req.ParseForm()
		c.RawStruct = req.Request.FormValue("clstruct")
		c.ConvertedPHP = Convert(c.RawStruct)
	} else {
		// Add default
		c.RawStruct = defaultRawStruct
	}

	rend.HTML(rw, http.StatusOK, "index", c)
}

// Convert takes a struct and spits back PHP, the unholy blah blah
func Convert(raw string) string {
	var out, name string

	lines := strings.Split(raw, "\n")

	if len(lines) < 3 {
		return "(╯°□°）╯︵ ┻━┻"
	}

	for _, l := range lines {
		l = strings.TrimSpace(l)
		words := strings.Fields(l)

		if len(words) > 1 {
			if words[0] == "type" {
				// The name
				name = words[1]
				out += fmt.Sprintf("$%s = array();\n", name)
			} else {
				words[1] = strings.Replace(words[1], "64", "", -1)
				words[1] = strings.Replace(words[1], "32", "", -1)
				out += fmt.Sprintf("$%s['%s'] = (%s) $data['%s_%s'];\n", name, words[0], words[1], name, words[0])
			}
		}
	}

	return out
}
