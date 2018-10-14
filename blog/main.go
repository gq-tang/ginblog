//go:generate go-bindata -prefix ../migrations/ -pkg migrations -o ../migrations/migrations_gen.go ../migrations
//go:generate go-bindata -prefix ../static/ -pkg static -o ../static/static_gen.go ../static/...
//go:generate go-bindata -prefix ../views/ -pkg views -o ../views/views_gen.go ../views/...

package main

import (
	"github.com/gq-tang/ginblog/blog/cmd"
)

var version string = "v1.0" // set by the complier

func main() {
	cmd.Execute(version)
}
