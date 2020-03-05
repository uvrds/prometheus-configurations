package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetEndopoint(landscape string) []string {
	var endpoints []string
	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("RUNCHERURL")+"/v3/clusters/?name="+landscape, nil)
	req.SetBasicAuth(os.Getenv("RUNCHERUSER"), os.Getenv("RUNCHERTOKEN"))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var structRancher structure.RancherClusters
	err = json.Unmarshal(content, &structRancher)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i <= len(structRancher.Data[0].AppliedSpec.RancherKubernetesEngineConfig.Nodes)-1; i++ {
		data := structRancher.Data[0].AppliedSpec.RancherKubernetesEngineConfig.Nodes[i]
		endpoints = append(endpoints, data.Address)
	}
	return endpoints
}
