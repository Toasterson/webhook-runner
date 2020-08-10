//go:generate go run generate.go

package main

import (
	"os"
	"path/filepath"

	"github.com/toasterson/svcgen"
)

func main() {
	if err := os.MkdirAll(filepath.Join("manifest"), 0755); err != nil {
		panic(err)
	}

	bundle := svcgen.NewManifestWithParams("application/webhooked", "application/webhooked", "", svcgen.Params{
		WorkingDirectory: "/var/www",
		UserName:         "root",
		GroupName:        "root",
		Env: []string{
			"CONFIG=/etc/webhooked.hcl",
		},
		StartCommand: []string{
			"/usr/bin/webhooked",
		},
		Type: svcgen.ServiceTypeChild,
	})

	if err := svcgen.WriteManifest(filepath.Join("manifest", "webhooked.xml"), bundle); err != nil {
		panic(err)
	}
}
