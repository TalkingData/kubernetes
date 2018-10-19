package priorities

import (
	"fmt"

	"k8s.io/api/core/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	"k8s.io/kubernetes/pkg/scheduler/schedulercache"
)

func FillupMemAllocationPriorityMap(
	pod *v1.Pod,
	meta interface{},
	nodeInfo *schedulercache.NodeInfo) (schedulerapi.HostPriority, error) {
	node := nodeInfo.Node()
	if node == nil {
		return schedulerapi.HostPriority{}, fmt.Errorf("node not found")
	}
	allocatable := nodeInfo.AllocatableResource()
	memAlloc := allocatable.Memory

	return schedulerapi.HostPriority{
		Host:  node.Name,
		Score: int(memAlloc),
	}, nil
}
