package data_store

import (
	"sync"

	"github.com/litmuschaos/litmus/litmus-portal/backend/graphql-server/graph/model"
)

//Application state, contains channels and mutexes used for subscriptions
type StateData struct {
	ClusterEventPublish  map[string][]chan *model.ClusterEvent
	ConnectedCluster     map[string]chan *model.ClusterAction
	WorkflowEventPublish map[string][]chan *model.WorkflowRun
	WorkflowLog          map[string]chan *model.PodLogResponse
	Mutex                *sync.Mutex
}

func New() *StateData {
	return &StateData{
		ClusterEventPublish:  make(map[string][]chan *model.ClusterEvent),
		ConnectedCluster:     make(map[string]chan *model.ClusterAction),
		WorkflowEventPublish: make(map[string][]chan *model.WorkflowRun),
		WorkflowLog:          make(map[string]chan *model.PodLogResponse),
		Mutex:                &sync.Mutex{},
	}
}
