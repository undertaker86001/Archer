package cluster

import (
	"cluster_manager/model"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestClusterManager(t *testing.T) {

	bs := &model.Bootstrap{
		StaticResources: model.StaticResources{
			Clusters: []*model.ClusterConfig{
				{
					Name: "test",
					Endpoints: []*model.Endpoint{
						{
							Address: model.SocketAddress{},
							ID:      "1",
						},
					},
				},
			},
		},
	}

	cm := CreateDefaultClusterManager(bs)

	assert.Equal(t, len(cm.store.Config), 1)

	cm.AddCluster(&model.ClusterConfig{
		Name: "test2",
		Endpoints: []*model.Endpoint{
			{
				Address: model.SocketAddress{},
				ID:      "1",
			},
		},
	})

	assert.Equal(t, len(cm.store.Config), 2)

	cm.SetEndpoint("test2", &model.Endpoint{
		Address: model.SocketAddress{},
		ID:      "2",
	})

}
