package industrial_process

import "time"
import "sync"
//import "context"

type plant struct {
    stateServ *stateService
    pumpChan chan *labelSignal
    timestamp int

    u_pressure_constant float64
    u_flow_constant float64

    r_head float64
    u_flow float64
    j_demand float64
    j_head float64
    p1_flow float64
    tank_level float64
    p2_flow float64
    d_demand float64

    tank_volume float64
    tank_capacity float64
    tank_height float64

    start chan bool
    stop chan bool
}

func NewPlant(u_pressure_constant float64, u_flow_constant float64, tank_capacity float64, tank_height float64) (*plant, *stateService, chan *labelSignal) {
    this := new(plant)

    // Initialise interaction datastructures
    this.stateServ = newStateService()
    this.pumpChan = make(chan *labelSignal)

    // Initialise tank parameters
    this.u_pressure_constant = u_pressure_constant
    this.u_flow_constant = u_flow_constant
    this.tank_capacity = tank_capacity
    this.tank_height = tank_height

    // Initialise operational parameters
    this.timestamp = 0
    this.tank_volume = 0.0

    this.r_head = 0
    this.u_flow = 0
    this.j_demand = 0
    this.j_head = 0
    this.p1_flow = 0
    this.tank_level = 0
    this.p2_flow = 0
    this.d_demand = 0

    this.stateServ.updateState("r.head", newHeadState(0, this.r_head))
    this.stateServ.updateState("u.flow", newFlowState(0, this.u_flow))
    this.stateServ.updateState("j.demand", newFlowState(0, this.j_demand))
    this.stateServ.updateState("j.head", newHeadState(0, this.j_head))
    this.stateServ.updateState("p1.flow", newFlowState(0, this.p1_flow))
    this.stateServ.updateState("tank.level", newHeadState(0, this.tank_level))
    this.stateServ.updateState("p2.flow", newFlowState(0, this.p2_flow))
    this.stateServ.updateState("d.demand", newFlowState(0, this.d_demand))

    return this, this.stateServ, this.pumpChan
}

func (this *plant) Deploy(distribution [][]float64) (wg *sync.WaitGroup){
    this.start = make(chan bool)
    this.stop = make(chan bool)
    wg = new(sync.WaitGroup)
    wg.Add(1)
    go this.run_(wg, distribution)
    return
}

func (this *plant) Start() {
    this.start <- true
}

func (this *plant) Stop() {
    this.stop <- true
}

func (this *plant) run_(wg *sync.WaitGroup, distribution [][]float64) {
    defer wg.Done()
    <- this.start
    states := make(map[string]state)

    // emulate the pump receiving signals
    go func() {
        for {
            sig := <- this.pumpChan
            lab := sig.label()
            if lab == on {
                this.r_head = this.u_pressure_constant
            } else if lab == off {
                this.r_head = 0
            }
        }
    }()

    i := 0
    for {
        select {
            case <- this.stop:
                return
            default:
                if this.r_head - this.j_head > 0 {
                    this.u_flow = this.u_flow_constant
                } else {
                    this.u_flow = 0
                }

                states["r.head"] = newHeadState(this.timestamp, this.r_head)
                states["u.flow"] = newFlowState(this.timestamp, this.u_flow)

                this.j_demand = distribution[0][i]
                this.j_head = 0

                states["j.demand"] = newFlowState(this.timestamp, this.j_demand)
                states["j.head"] = newHeadState(this.timestamp, this.j_head)

                this.p1_flow = this.u_flow - this.j_demand
                if this.p1_flow < 0 {
                    this.p1_flow = 0
                }

                states["p1.flow"] = newFlowState(this.timestamp, this.p1_flow)

                this.d_demand = distribution[1][i]
                this.p2_flow = this.d_demand
                this.tank_volume = this.tank_volume + this.p1_flow - this.p2_flow
                if this.tank_volume > this.tank_capacity {
                    this.tank_volume = this.tank_capacity
                } else if this.tank_volume < 0 {
                    this.tank_volume = 0
                }
                this.tank_level = this.tank_volume / this.tank_height

                states["tank.level"] = newHeadState(this.timestamp, this.tank_level)
                states["p2.flow"] = newFlowState(this.timestamp, this.p2_flow)
                states["d.demand"] = newFlowState(this.timestamp, this.d_demand)

                this.stateServ.updateStates(states)

                //time.Sleep(2 * time.Second)
                time.Sleep(1 * time.Second)
                i = (i + 1) % len(distribution[0])
                this.timestamp = this.timestamp + 1
        }
    }
}

// func plant(stateServ *stateService, pumpChan chan *labelSignal, distribution [][]float64) {
//
//     states := make(map[string]state)
//
//     _, r_head := stateServ.getState("r.head").data()
//     _, u_flow := stateServ.getState("u.flow").data()
//     _, j_demand := stateServ.getState("j.demand").data()
//     _, j_head := stateServ.getState("j.head").data()
//     _, p1_flow := stateServ.getState("p1.flow").data()
//     _, d_demand := stateServ.getState("d.demand").data()
//     _, p2_flow := stateServ.getState("p2.flow").data()
//     _, tank_level := stateServ.getState("tank.level").data()
//
//     tank_volume := 0.0
//     tank_capacity := 100.0
//     tank_height := 10.0
//
//     // emulate the pump receiving signals
//     go func() {
//         for {
//             sig := <- pumpChan
//             lab := sig.label()
//             if lab == on {
//                 r_head = 3
//             } else if lab == off {
//                 r_head = 0
//             }
//         }
//     }()
//
//     for i := range distribution[0] {
//         if r_head - j_head > 0 {
//             u_flow = 15
//         } else {
//             u_flow = 0
//         }
//
//         states["r.head"] = newHeadState(i, r_head)
//         states["u.flow"] = newFlowState(i, u_flow)
//
//         j_demand = distribution[0][i]
//         j_head = 0
//
//         states["j.demand"] = newFlowState(i, j_demand)
//         states["j.head"] = newHeadState(i, j_head)
//
//         p1_flow = u_flow - j_demand
//         if p1_flow < 0 {
//             p1_flow = 0
//         }
//
//         states["p1.flow"] = newFlowState(i, p1_flow)
//
//         d_demand = distribution[1][i]
//         p2_flow = d_demand
//         tank_volume = tank_volume + p1_flow - p2_flow
//         if tank_volume > tank_capacity {
//             tank_volume = tank_capacity
//         } else if tank_volume < 0 {
//             tank_volume = 0
//         }
//         tank_level = tank_volume / tank_height
//
//         states["tank.level"] = newHeadState(i, tank_level)
//         states["p2.flow"] = newFlowState(i, p2_flow)
//         states["d.demand"] = newFlowState(i, d_demand)
//
//         stateServ.updateStates(states)
//
//         time.Sleep(2 * time.Second)
//     }
// }
