package main

import (
	"flag"
	"git.crptech.ru/prometheus-operator-configurations/pkg/help"
	"git.crptech.ru/prometheus-operator-configurations/pkg/services"
	"git.crptech.ru/prometheus-operator-configurations/pkg/structure"
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
