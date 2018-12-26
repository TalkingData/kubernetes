package cache

type ResourceCache interface {
	Get(*NodeInfo) (*Resource, error)
}

type FakeResourceMonitor struct {
}

func (f *FakeResourceMonitor) Get(n *NodeInfo) (*Resource, error) {
	res := n.AllocatableResource()
	return &res, nil
}
