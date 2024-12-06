package industrial_process

import "fmt"
import "time"
import "sync"

/******************************************************************************
 * state
 ******************************************************************************/

type state interface {
    data() (int, float64)
    toString() string
}

type stateImpl struct {
    timestamp int
    val float64
}

func (this *stateImpl) init(timestamp int, val float64) {
    this.timestamp = timestamp
    this.val = val
}

func (this *stateImpl) data() (int, float64) {
    return this.timestamp, this.val
}

func (this *stateImpl) toString() string {
    return fmt.Sprintf("[Timestamp: %v, Value: %v]", this.timestamp, this.val)
}

////// flow state

type flowState struct {
    stateImpl
}

func newFlowState(timestamp int, val float64) (this *flowState) {
    this = new(flowState)
    this.stateImpl.init(timestamp, val)
    return
}

func (this *flowState) toString() string {
    timestamp, val := this.data()
    return fmt.Sprintf("[Timestamp: %v, Flow: %v m/s]", timestamp, val)
}

/////// head state

type headState struct {
    stateImpl
}

func newHeadState(timestamp int, val float64) (this *headState) {
    this = new(headState)
    this.stateImpl.init(timestamp, val)
    return
}

func (this *headState) toString() string {
    timestamp, val := this.data()
    return fmt.Sprintf("[Timestamp: %v, Head: %v m]", timestamp, val)
}

/******************************************************************************
 * state service
 ******************************************************************************/

type stateService struct {
    mu sync.Mutex
    states map[string]state
}

func newStateService() (this *stateService) {
    this = new(stateService)
    this.states = make(map[string]state)
    return
}

func (this *stateService) updateStates(states map[string]state) {
    this.mu.Lock()
    defer this.mu.Unlock()
    for name, msg := range states {
        this.states[name] = msg
    }
}

func (this *stateService) updateState(name string, st state) {
    this.mu.Lock()
    defer this.mu.Unlock()
    this.states[name] = st
}

func (this *stateService) getState(name string) state {
    this.mu.Lock()
    defer this.mu.Unlock()
    return this.states[name]
}

func (this *stateService) Monitor(stateNames []string) {
    //i := 0
    fmt.Println("[Monitoring Serice] Starting.")
    for {
        this.mu.Lock()
        for _, name := range stateNames {
            fmt.Printf("[Monitoring Service] state: %s, data: %v\n", name, this.states[name].toString())
        }
        this.mu.Unlock()
        time.Sleep(1 * time.Second)
        //i = i + 1
    }
}
