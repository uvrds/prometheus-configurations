package main

import (
	"flag"
	"github.com/pkg/help"
	"github.com/pkg/services"
	"github.com/pkg/structure"
	"log"
	"os"
)

func main() {

	landscape := flag.String("landscape", "nil",
		"Please input landscape")
	flag.Parse()

	if *landscape == "nil" {
		log.Fatal("landscape: " + *landscape)
	}

	obj := structure.Prom{
		services.GetEndopoint(*landscape),
		help.LandscapeHostIng(*landscape),
		*landscape, os.Getenv("VMURL"),
		help.LandscapeHostAlertmanager(*landscape),
	}
	services.RenderYml(obj)

}
