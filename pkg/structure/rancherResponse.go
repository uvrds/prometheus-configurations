package structure

import "time"

//struct for rancher api. Generic https://mholt.github.io/json-to-go/
type RancherClusters struct {
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	CreateTypes struct {
		Cluster string `json:"cluster"`
	} `json:"createTypes"`
	Actions struct {
		CreateFromTemplate string `json:"createFromTemplate"`
	} `json:"actions"`
	Pagination struct {
		Limit int `json:"limit"`
		Total int `json:"total"`
	} `json:"pagination"`
	Sort struct {
		Order   string `json:"order"`
		Reverse string `json:"reverse"`
		Links   struct {
			AgentImage                         string `json:"agentImage"`
			APIEndpoint                        string `json:"apiEndpoint"`
			AppliedPodSecurityPolicyTemplateID string `json:"appliedPodSecurityPolicyTemplateId"`
			AuthImage                          string `json:"authImage"`
			CaCert                             string `json:"caCert"`
			Description                        string `json:"description"`
			DesiredAgentImage                  string `json:"desiredAgentImage"`
			DesiredAuthImage                   string `json:"desiredAuthImage"`
			DockerRootDir                      string `json:"dockerRootDir"`
			Driver                             string `json:"driver"`
			State                              string `json:"state"`
			Transitioning                      string `json:"transitioning"`
			TransitioningMessage               string `json:"transitioningMessage"`
			UUID                               string `json:"uuid"`
		} `json:"links"`
	} `json:"sort"`
	Filters struct {
		AgentImage                          interface{} `json:"agentImage"`
		APIEndpoint                         interface{} `json:"apiEndpoint"`
		AppliedEnableNetworkPolicy          interface{} `json:"appliedEnableNetworkPolicy"`
		AppliedPodSecurityPolicyTemplateID  interface{} `json:"appliedPodSecurityPolicyTemplateId"`
		AuthImage                           interface{} `json:"authImage"`
		CaCert                              interface{} `json:"caCert"`
		ClusterTemplateID                   interface{} `json:"clusterTemplateId"`
		ClusterTemplateRevisionID           interface{} `json:"clusterTemplateRevisionId"`
		Created                             interface{} `json:"created"`
		CreatorID                           interface{} `json:"creatorId"`
		DefaultClusterRoleForProjectMembers interface{} `json:"defaultClusterRoleForProjectMembers"`
		DefaultPodSecurityPolicyTemplateID  interface{} `json:"defaultPodSecurityPolicyTemplateId"`
		Description                         interface{} `json:"description"`
		DesiredAgentImage                   interface{} `json:"desiredAgentImage"`
		DesiredAuthImage                    interface{} `json:"desiredAuthImage"`
		DockerRootDir                       interface{} `json:"dockerRootDir"`
		Driver                              interface{} `json:"driver"`
		EnableClusterAlerting               interface{} `json:"enableClusterAlerting"`
		EnableClusterMonitoring             interface{} `json:"enableClusterMonitoring"`
		EnableNetworkPolicy                 interface{} `json:"enableNetworkPolicy"`
		ID                                  interface{} `json:"id"`
		Internal                            interface{} `json:"internal"`
		IstioEnabled                        interface{} `json:"istioEnabled"`
		Name                                []struct {
			Modifier string `json:"modifier"`
			Value    string `json:"value"`
		} `json:"name"`
		Removed                interface{} `json:"removed"`
		State                  interface{} `json:"state"`
		Transitioning          interface{} `json:"transitioning"`
		TransitioningMessage   interface{} `json:"transitioningMessage"`
		UUID                   interface{} `json:"uuid"`
		WindowsPreferedCluster interface{} `json:"windowsPreferedCluster"`
	} `json:"filters"`
	ResourceType string `json:"resourceType"`
	Data         []struct {
		Actions struct {
			BackupEtcd            string `json:"backupEtcd"`
			EnableMonitoring      string `json:"enableMonitoring"`
			ExportYaml            string `json:"exportYaml"`
			GenerateKubeconfig    string `json:"generateKubeconfig"`
			ImportYaml            string `json:"importYaml"`
			RestoreFromEtcdBackup string `json:"restoreFromEtcdBackup"`
			RotateCertificates    string `json:"rotateCertificates"`
			RunSecurityScan       string `json:"runSecurityScan"`
		} `json:"actions"`
		AgentImage  string `json:"agentImage"`
		Allocatable struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
			Pods   string `json:"pods"`
		} `json:"allocatable"`
		Annotations struct {
			AuthzManagementCattleIoCreatorRoleBindings           string `json:"authz.management.cattle.io/creator-role-bindings"`
			LifecycleCattleIoCreateClusterAgentControllerCleanup string `json:"lifecycle.cattle.io/create.cluster-agent-controller-cleanup"`
			LifecycleCattleIoCreateClusterProvisionerController  string `json:"lifecycle.cattle.io/create.cluster-provisioner-controller"`
			LifecycleCattleIoCreateClusterScopedGc               string `json:"lifecycle.cattle.io/create.cluster-scoped-gc"`
			LifecycleCattleIoCreateMgmtClusterRbacRemove         string `json:"lifecycle.cattle.io/create.mgmt-cluster-rbac-remove"`
			ProvisionerCattleIoKeDriverUpdate                    string `json:"provisioner.cattle.io/ke-driver-update"`
		} `json:"annotations"`
		Answers struct {
			ClusterID interface{} `json:"clusterId"`
			ProjectID interface{} `json:"projectId"`
			Type      string      `json:"type"`
		} `json:"answers"`
		APIEndpoint                        string `json:"apiEndpoint"`
		AppliedEnableNetworkPolicy         bool   `json:"appliedEnableNetworkPolicy"`
		AppliedPodSecurityPolicyTemplateID string `json:"appliedPodSecurityPolicyTemplateId"`
		AppliedSpec                        struct {
			Answers struct {
				ClusterID interface{} `json:"clusterId"`
				ProjectID interface{} `json:"projectId"`
				Type      string      `json:"type"`
			} `json:"answers"`
			ClusterTemplateID                   interface{} `json:"clusterTemplateId"`
			ClusterTemplateRevisionID           interface{} `json:"clusterTemplateRevisionId"`
			DefaultClusterRoleForProjectMembers interface{} `json:"defaultClusterRoleForProjectMembers"`
			DefaultPodSecurityPolicyTemplateID  interface{} `json:"defaultPodSecurityPolicyTemplateId"`
			Description                         string      `json:"description"`
			DesiredAgentImage                   string      `json:"desiredAgentImage"`
			DesiredAuthImage                    string      `json:"desiredAuthImage"`
			DisplayName                         string      `json:"displayName"`
			DockerRootDir                       string      `json:"dockerRootDir"`
			EnableClusterAlerting               bool        `json:"enableClusterAlerting"`
			EnableClusterMonitoring             bool        `json:"enableClusterMonitoring"`
			EnableNetworkPolicy                 bool        `json:"enableNetworkPolicy"`
			Internal                            bool        `json:"internal"`
			LocalClusterAuthEndpoint            struct {
				Enabled bool   `json:"enabled"`
				Type    string `json:"type"`
			} `json:"localClusterAuthEndpoint"`
			RancherKubernetesEngineConfig struct {
				AddonJobTimeout int `json:"addonJobTimeout"`
				Authentication  struct {
					Strategy string `json:"strategy"`
					Type     string `json:"type"`
				} `json:"authentication"`
				Authorization struct {
					Type string `json:"type"`
				} `json:"authorization"`
				BastionHost struct {
					SSHAgentAuth bool   `json:"sshAgentAuth"`
					Type         string `json:"type"`
				} `json:"bastionHost"`
				CloudProvider struct {
					Type string `json:"type"`
				} `json:"cloudProvider"`
				IgnoreDockerVersion bool `json:"ignoreDockerVersion"`
				Ingress             struct {
					Provider string `json:"provider"`
					Type     string `json:"type"`
				} `json:"ingress"`
				KubernetesVersion string `json:"kubernetesVersion"`
				Monitoring        struct {
					Provider string `json:"provider"`
					Type     string `json:"type"`
				} `json:"monitoring"`
				Network struct {
					Options struct {
						FlannelBackendType string `json:"flannel_backend_type"`
					} `json:"options"`
					Plugin string `json:"plugin"`
					Type   string `json:"type"`
				} `json:"network"`
				Nodes []struct {
					Address          string   `json:"address"`
					HostnameOverride string   `json:"hostnameOverride"`
					InternalAddress  string   `json:"internalAddress"`
					NodeID           string   `json:"nodeId"`
					Port             string   `json:"port"`
					Role             []string `json:"role"`
					SSHAgentAuth     bool     `json:"sshAgentAuth"`
					Type             string   `json:"type"`
					User             string   `json:"user"`
				} `json:"nodes"`
				Restore struct {
					Restore bool   `json:"restore"`
					Type    string `json:"type"`
				} `json:"restore"`
				Services struct {
					Etcd struct {
						BackupConfig struct {
							Enabled        bool        `json:"enabled"`
							IntervalHours  int         `json:"intervalHours"`
							Retention      int         `json:"retention"`
							S3BackupConfig interface{} `json:"s3BackupConfig"`
							SafeTimestamp  bool        `json:"safeTimestamp"`
							Type           string      `json:"type"`
						} `json:"backupConfig"`
						Creation  string `json:"creation"`
						ExtraArgs struct {
							ElectionTimeout   string `json:"election-timeout"`
							HeartbeatInterval string `json:"heartbeat-interval"`
						} `json:"extraArgs"`
						Gid       int    `json:"gid"`
						Retention string `json:"retention"`
						Snapshot  bool   `json:"snapshot"`
						Type      string `json:"type"`
						UID       int    `json:"uid"`
					} `json:"etcd"`
					KubeAPI struct {
						AlwaysPullImages     bool   `json:"alwaysPullImages"`
						PodSecurityPolicy    bool   `json:"podSecurityPolicy"`
						ServiceNodePortRange string `json:"serviceNodePortRange"`
						Type                 string `json:"type"`
					} `json:"kubeApi"`
					KubeController struct {
						Type string `json:"type"`
					} `json:"kubeController"`
					Kubelet struct {
						FailSwapOn bool   `json:"failSwapOn"`
						Type       string `json:"type"`
					} `json:"kubelet"`
					Kubeproxy struct {
						Type string `json:"type"`
					} `json:"kubeproxy"`
					Scheduler struct {
						Type string `json:"type"`
					} `json:"scheduler"`
					Type string `json:"type"`
				} `json:"services"`
				SSHAgentAuth bool   `json:"sshAgentAuth"`
				Type         string `json:"type"`
			} `json:"rancherKubernetesEngineConfig"`
			Type                   string `json:"type"`
			WindowsPreferedCluster bool   `json:"windowsPreferedCluster"`
		} `json:"appliedSpec"`
		AuthImage    string `json:"authImage"`
		BaseType     string `json:"baseType"`
		CaCert       string `json:"caCert"`
		Capabilities struct {
			IngressCapabilities []struct {
				CustomDefaultBackend bool   `json:"customDefaultBackend"`
				IngressProvider      string `json:"ingressProvider"`
				Type                 string `json:"type"`
			} `json:"ingressCapabilities"`
			LoadBalancerCapabilities struct {
				HealthCheckSupported bool   `json:"healthCheckSupported"`
				Type                 string `json:"type"`
			} `json:"loadBalancerCapabilities"`
			NodePoolScalingSupported bool   `json:"nodePoolScalingSupported"`
			NodePortRange            string `json:"nodePortRange"`
			TaintSupport             bool   `json:"taintSupport"`
			Type                     string `json:"type"`
		} `json:"capabilities"`
		Capacity struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
			Pods   string `json:"pods"`
		} `json:"capacity"`
		CertificatesExpiration struct {
			KubeApiserver struct {
				ExpirationDate time.Time `json:"expirationDate"`
				Type           string    `json:"type"`
			} `json:"kube-apiserver"`
			KubeCa struct {
				ExpirationDate time.Time `json:"expirationDate"`
				Type           string    `json:"type"`
			} `json:"kube-ca"`
			KubeControllerManager struct {
				ExpirationDate time.Time `json:"expirationDate"`
				Type           string    `json:"type"`
			} `json:"kube-controller-manager"`
			KubeEtcd172301646 struct {
				ExpirationDate time.Time `json:"expirationDate"`
				Type           string    `json:"type"`
			} `json:"kube-etcd-172-30-16-46"`
			KubeEtcd172301648 struct {
				ExpirationDate time.Time `json:"expirationDate"`
				Type           string    `json:"type"`
			} `json:"kube-etcd-172-30-16-48"`
			KubeNode struct {
				ExpirationDate time.Time `json:"expirationDate"`
				Type           string    `json:"type"`
			} `json:"kube-node"`
			KubeProxy struct {
				ExpirationDate time.Time `json:"expirationDate"`
				Type           string    `json:"type"`
			} `json:"kube-proxy"`
			KubeScheduler struct {
				ExpirationDate time.Time `json:"expirationDate"`
				Type           string    `json:"type"`
			} `json:"kube-scheduler"`
		} `json:"certificatesExpiration"`
		ClusterTemplateID         interface{} `json:"clusterTemplateId"`
		ClusterTemplateRevisionID interface{} `json:"clusterTemplateRevisionId"`
		ComponentStatuses         []struct {
			Conditions []struct {
				Message string `json:"message"`
				Status  string `json:"status"`
				Type    string `json:"type"`
			} `json:"conditions"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"componentStatuses"`
		Conditions []struct {
			Status         string    `json:"status"`
			Type           string    `json:"type"`
			LastUpdateTime time.Time `json:"lastUpdateTime,omitempty"`
		} `json:"conditions"`
		Created                             time.Time   `json:"created"`
		CreatedTS                           int64       `json:"createdTS"`
		CreatorID                           string      `json:"creatorId"`
		DefaultClusterRoleForProjectMembers interface{} `json:"defaultClusterRoleForProjectMembers"`
		DefaultPodSecurityPolicyTemplateID  interface{} `json:"defaultPodSecurityPolicyTemplateId"`
		Description                         string      `json:"description"`
		DesiredAgentImage                   string      `json:"desiredAgentImage"`
		DesiredAuthImage                    string      `json:"desiredAuthImage"`
		DockerRootDir                       string      `json:"dockerRootDir"`
		Driver                              string      `json:"driver"`
		EnableClusterAlerting               bool        `json:"enableClusterAlerting"`
		EnableClusterMonitoring             bool        `json:"enableClusterMonitoring"`
		EnableNetworkPolicy                 bool        `json:"enableNetworkPolicy"`
		ID                                  string      `json:"id"`
		Internal                            bool        `json:"internal"`
		IstioEnabled                        bool        `json:"istioEnabled"`
		Labels                              struct {
			CattleIoCreator string `json:"cattle.io/creator"`
		} `json:"labels"`
		Limits struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
			Pods   string `json:"pods"`
		} `json:"limits"`
		Links struct {
			APIServices                 string `json:"apiServices"`
			ClusterAlertGroups          string `json:"clusterAlertGroups"`
			ClusterAlertRules           string `json:"clusterAlertRules"`
			ClusterAlerts               string `json:"clusterAlerts"`
			ClusterCatalogs             string `json:"clusterCatalogs"`
			ClusterLoggings             string `json:"clusterLoggings"`
			ClusterMonitorGraphs        string `json:"clusterMonitorGraphs"`
			ClusterRegistrationTokens   string `json:"clusterRegistrationTokens"`
			ClusterRoleTemplateBindings string `json:"clusterRoleTemplateBindings"`
			ClusterScans                string `json:"clusterScans"`
			EtcdBackups                 string `json:"etcdBackups"`
			Namespaces                  string `json:"namespaces"`
			NodePools                   string `json:"nodePools"`
			Nodes                       string `json:"nodes"`
			Notifiers                   string `json:"notifiers"`
			PersistentVolumes           string `json:"persistentVolumes"`
			Projects                    string `json:"projects"`
			Remove                      string `json:"remove"`
			Self                        string `json:"self"`
			Shell                       string `json:"shell"`
			StorageClasses              string `json:"storageClasses"`
			Subscribe                   string `json:"subscribe"`
			Templates                   string `json:"templates"`
			Tokens                      string `json:"tokens"`
			Update                      string `json:"update"`
		} `json:"links"`
		LocalClusterAuthEndpoint struct {
			Enabled bool   `json:"enabled"`
			Type    string `json:"type"`
		} `json:"localClusterAuthEndpoint"`
		Name                          string `json:"name"`
		RancherKubernetesEngineConfig struct {
			AddonJobTimeout int `json:"addonJobTimeout"`
			Authentication  struct {
				Strategy string `json:"strategy"`
				Type     string `json:"type"`
			} `json:"authentication"`
			Authorization struct {
				Type string `json:"type"`
			} `json:"authorization"`
			BastionHost struct {
				SSHAgentAuth bool   `json:"sshAgentAuth"`
				Type         string `json:"type"`
			} `json:"bastionHost"`
			CloudProvider struct {
				Type string `json:"type"`
			} `json:"cloudProvider"`
			IgnoreDockerVersion bool `json:"ignoreDockerVersion"`
			Ingress             struct {
				Provider string `json:"provider"`
				Type     string `json:"type"`
			} `json:"ingress"`
			KubernetesVersion string `json:"kubernetesVersion"`
			Monitoring        struct {
				Provider string `json:"provider"`
				Type     string `json:"type"`
			} `json:"monitoring"`
			Network struct {
				Options struct {
					FlannelBackendType string `json:"flannel_backend_type"`
				} `json:"options"`
				Plugin string `json:"plugin"`
				Type   string `json:"type"`
			} `json:"network"`
			Restore struct {
				Restore bool   `json:"restore"`
				Type    string `json:"type"`
			} `json:"restore"`
			Services struct {
				Etcd struct {
					BackupConfig struct {
						Enabled        bool        `json:"enabled"`
						IntervalHours  int         `json:"intervalHours"`
						Retention      int         `json:"retention"`
						S3BackupConfig interface{} `json:"s3BackupConfig"`
						SafeTimestamp  bool        `json:"safeTimestamp"`
						Type           string      `json:"type"`
					} `json:"backupConfig"`
					Creation  string `json:"creation"`
					ExtraArgs struct {
						ElectionTimeout   string `json:"election-timeout"`
						HeartbeatInterval string `json:"heartbeat-interval"`
					} `json:"extraArgs"`
					Gid       int    `json:"gid"`
					Retention string `json:"retention"`
					Snapshot  bool   `json:"snapshot"`
					Type      string `json:"type"`
					UID       int    `json:"uid"`
				} `json:"etcd"`
				KubeAPI struct {
					AlwaysPullImages     bool   `json:"alwaysPullImages"`
					PodSecurityPolicy    bool   `json:"podSecurityPolicy"`
					ServiceNodePortRange string `json:"serviceNodePortRange"`
					Type                 string `json:"type"`
				} `json:"kubeApi"`
				KubeController struct {
					Type string `json:"type"`
				} `json:"kubeController"`
				Kubelet struct {
					FailSwapOn bool   `json:"failSwapOn"`
					Type       string `json:"type"`
				} `json:"kubelet"`
				Kubeproxy struct {
					Type string `json:"type"`
				} `json:"kubeproxy"`
				Scheduler struct {
					Type string `json:"type"`
				} `json:"scheduler"`
				Type string `json:"type"`
			} `json:"services"`
			SSHAgentAuth bool   `json:"sshAgentAuth"`
			Type         string `json:"type"`
		} `json:"rancherKubernetesEngineConfig"`
		Requested struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
			Pods   string `json:"pods"`
		} `json:"requested"`
		State                string `json:"state"`
		Transitioning        string `json:"transitioning"`
		TransitioningMessage string `json:"transitioningMessage"`
		Type                 string `json:"type"`
		UUID                 string `json:"uuid"`
		Version              struct {
			BuildDate    time.Time `json:"buildDate"`
			Compiler     string    `json:"compiler"`
			GitCommit    string    `json:"gitCommit"`
			GitTreeState string    `json:"gitTreeState"`
			GitVersion   string    `json:"gitVersion"`
			GoVersion    string    `json:"goVersion"`
			Major        string    `json:"major"`
			Minor        string    `json:"minor"`
			Platform     string    `json:"platform"`
			Type         string    `json:"type"`
		} `json:"version"`
		WindowsPreferedCluster bool `json:"windowsPreferedCluster"`
	} `json:"data"`
}
