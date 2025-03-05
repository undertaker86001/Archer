package cluster

import (
	"cluster_manager/health_check"
	"cluster_manager/model"
	"fmt"
	"sync"
	"sync/atomic"
)

type Cluster struct {
	Config      *model.ClusterConfig
	HealthCheck *health_check.HealthChecker
}

var clusterIndex int32 = 1

type (
	ClusterManager struct {
		rw sync.RWMutex

		store *ClusterStore
	}

	// ClusterStore store for cluster array
	ClusterStore struct {
		Config      []*model.ClusterConfig `yaml:"config" json:"config"`
		Version     int32                  `yaml:"version" json:"version"`
		clustersMap map[string]*Cluster
	}

	// xdsControlStore help convert ClusterStore to controls.ClusterStore interface
	xdsControlStore struct {
		*ClusterStore
	}
)

func (cm *ClusterManager) SetEndpoint(clusterName string, endpoint *model.Endpoint) {
	cm.rw.Lock()
	defer cm.rw.Unlock()

	cm.store.IncreaseVersion()
	cm.store.SetEndpoint(clusterName, endpoint)
}

func (cm *ClusterManager) AddCluster(config *model.ClusterConfig) {
	cm.rw.Lock()
	defer cm.rw.Unlock()

	cm.store.AddCluster(config)
}

func NewCluster(clusterConfig *model.ClusterConfig) *Cluster {
	c := &Cluster{
		Config: clusterConfig,
	}

	//todo implements health check

	return c
}

func CreateDefaultClusterManager(bs *model.Bootstrap) *ClusterManager {
	return &ClusterManager{store: newClusterStore(bs)}
}

func (s *ClusterStore) AddCluster(c *model.ClusterConfig) {
	if c.Name == "" {
		index := atomic.AddInt32(&clusterIndex, 1)
		c.Name = fmt.Sprintf("cluster%d", index)
	}
	s.Config = append(s.Config, c)
}

func (s *ClusterStore) SetEndpoint(clusterName string, endpoint *model.Endpoint) {
	cluster := s.clustersMap[clusterName]
	if cluster == nil {
		c := &model.ClusterConfig{Name: clusterName, Lbstr: model.LoadBalancerRoundRobin, Endpoints: []*model.Endpoint{}}
		s.AddCluster(c)
		cluster = s.clustersMap[clusterName]
	}

	for _, c := range s.Config {
		if c.Name == clusterName {
			for _, e := range c.Endpoints {
				if e.ID == endpoint.ID {
					//轮询当前worker端点列表 是否发生断连
					cluster.RemoveEndpoint(endpoint)
					e.Name = endpoint.Name
					e.Metadata = endpoint.Metadata
					e.Address = endpoint.Address
					cluster.AddEndpoint(endpoint)
				}
			}
		}
	}
}

func newClusterStore(bs *model.Bootstrap) *ClusterStore {
	store := &ClusterStore{
		clustersMap: map[string]*Cluster{},
	}
	for _, cluster := range bs.StaticResources.Clusters {
		store.AddCluster(cluster)
	}
	return store
}

func (s *ClusterStore) IncreaseVersion() {
	atomic.AddInt32(&s.Version, 1)
}

func (c *Cluster) RemoveEndpoint(endpoint *model.Endpoint) {
	if c.HealthCheck != nil {
		c.HealthCheck.StopOne(endpoint)
	}
}

func (c *Cluster) AddEndpoint(endpoint *model.Endpoint) {
	if c.HealthCheck != nil {
		c.HealthCheck.StartOne(endpoint)
	}
}
