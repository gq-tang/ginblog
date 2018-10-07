//go:generate go-bindata -prefix ../migrations/ -pkg migrations -o ../migrations/migrations_gen.go ../migrations
package main

import (
	"github.com/gq-tang/ginblog/blog/cmd"
)

var version string = "v1.0" // set by the complier

func main() {
	cmd.Execute(version)
}
