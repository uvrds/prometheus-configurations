---

kubeControllerManager:
  endpoints:
    {{range .Endpoints}}{{println "-" .}}{{"    "}}{{end}}
kubeScheduler:
  endpoints:
    {{range .Endpoints}}{{println "-" .}}{{"    "}}{{end}}
prometheus:
  ingress:
    hosts:
      - {{.HostsIng}}
  prometheusSpec:
    externalLabels:
      landscape: {{.Landscape}}
    remoteWrite:
      - url: {{.UrlVictoria}}
        queue_config:
          max_samples_per_send: 10000
          max_shards: 30
        writeRelabelConfigs:
          - separator: ;
            regex: prometheus_replica
            replacement: $1
            action: labeldrop

alertmanager:
  ingress:
    hosts:
      - {{.HostAlertmanager}}