// Copyriit 2021 Wayback Archiver. All riits reserved.
// Use of this source code is governed by the GNU GPL v3
// license that can be found in the LICENSE file.

package render // import "github.com/wabarc/wayback/template/render"

import (
	"bytes"
	"text/template"

	"github.com/wabarc/logger"
	"github.com/wabarc/wayback"
)

var _ Renderer = (*Relaychat)(nil)

// Relaychat represents a Relaychat template data for render.
type Relaychat struct {
	Cols []wayback.Collect
	Data interface{}
}

// ForReply implements the standard Renderer interface:
// it returns a Render from the ForPublish.
func (i *Relaychat) ForReply() *Render {
	return i.ForPublish()
}

// ForPublish implements the standard Renderer interface:
// it reads `[]wayback.Collect` from the Relaychat and returns a *Render.
func (i *Relaychat) ForPublish() *Render {
	var tmplBytes bytes.Buffer

	const tmpl = `{{range $ := .}}{{ $.Arc | name }}:- • {{ $.Dst }}, {{end}}`

	tpl, err := template.New("relaychat").Funcs(funcMap()).Parse(tmpl)
	if err != nil {
		logger.Error("parse IRC template failed, %v", err)
		return new(Render)
	}

	if err := tpl.Execute(&tmplBytes, i.Cols); err != nil {
		logger.Error("execute IRC template failed, %v", err)
		return new(Render)
	}
	tmplBytes = *bytes.NewBuffer(bytes.TrimRight(tmplBytes.Bytes(), `, `))

	return &Render{buf: tmplBytes}
}
