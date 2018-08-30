package grifts

import (
	"github.com/gianrubio/nerdalize-challenge/api/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
