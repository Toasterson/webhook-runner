//go:generate goexports gopkg.in/go-playground/webhooks.v5/github
//go:generate goexports os/exec

package interpreter

import (
	"github.com/containous/yaegi/interp"
)

var Symbols = interp.Exports{}
