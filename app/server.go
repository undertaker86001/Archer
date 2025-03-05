package app

import (
	config2 "cluster_manager/config"
	"cluster_manager/resource"
)

func CreateUnixServerChain(runOptions *resource.ComputeOptions, stopCh <-chan struct{}, finishCh chan struct{}) {

	config := &config2.Config{}

	real_config, err := config.Complete().New("real_config")
	if real_config != nil && err != nil {

	}

}
