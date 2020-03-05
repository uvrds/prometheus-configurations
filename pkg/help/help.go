package help

import (
	"strings"
)

func LandscapeHostIng(landscape string) string {

	land := strings.Split(landscape, "-")
	hosts := "prometheus." + land[1] + "." + land[0] + ".example.com"
	return hosts
}

func LandscapeHostAlertmanager(landscape string) string {

	land := strings.Split(landscape, "-")
	hosts := "alertmanager." + land[1] + "." + land[0] + ".example.com"
	return hosts
}

