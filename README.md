# prometheus-operator-configurations

simple example auto-configurations values for prometheus-operator

In then Example rancher api for get ip address extarnal container kubeControllerManager and kubeScheduler


How to start 

--landscape kafka-prod

env:
TPL_PROM= prom.tpl
RUNCHERURL= rancher.com
RUNCHERUSER= token user
RUNCHERTOKEN= rancher token
VMURL= url VictoriaMetrics for prometheus remote write

Also based on our environment by  template prom.ptl you can get configuration:

 ---
 kubeControllerManager:
   endpoints:
     - 72.29.17.90
     - 72.29.17.91
     - 72.29.17.92
     
 kubeScheduler:
   endpoints:
     - 72.29.17.90
     - 72.29.17.91
     - 72.29.17.92
     
 prometheus:
   ingress:
     hosts:
       - prometheus.kafka.prod.example.com
   prometheusSpec:
     externalLabels:
       landscape: prod-kafka
     remoteWrite:
       - url: <victoriaMetrics>
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
       - alertmanager.kafka.prod.example.com

 