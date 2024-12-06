package industrial_process
import "fmt"
import "sync"
import "context"

import "autonomic/iCPSDL/util"

/******************************************************************************
 * signal
 ******************************************************************************/

type label string
const (
    on label = "ON"
    off label = "OFF"
    nval label = "NVAL"
    start label = "start"
    stop label = "stop"
)

type labelSignal struct {
    lab label
}

func (this *labelSignal) label() label {
    return this.lab
}

/******************************************************************************
 * concurrent component interface
 ******************************************************************************/

type concurrentComponent interface {
    run(ctx context.Context, wg *sync.WaitGroup)
}

/******************************************************************************
 * sensor component
 ******************************************************************************/

type sensor struct {
    service *stateService
    stateName string
}

/************************
 * flow sensor
 ************************/

type flowSensor struct {
    sensor
    output chan *flowState
}

func newFlowSensor(service *stateService, name string, output chan *flowState) (this *flowSensor) {
    this = new(flowSensor)
    this.service = service
    this.stateName = name
    this.output = output
    return
}

func (this *flowSensor) run(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()
    defer fmt.Println("[Control loop configuration deployment module] Stopping flow sensor reading state:", this.stateName)
    fmt.Println("[Control loop configuration deployment module] Starting flow sensor reading state:", this.stateName)
    for {
        s := this.service.getState(this.stateName).(*flowState)
        select {
            case <- ctx.Done():
                return
            case this.output <- s:
        }
    }
}

/************************
 * head sensor
 ************************/

type headSensor struct {
    sensor
    output chan *headState
}

func newHeadSensor(service *stateService, name string, output chan *headState) (this *headSensor) {
    this = new(headSensor)
    this.service = service
    this.stateName = name
    this.output = output
    return this
}

func (this *headSensor) run(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()
    defer fmt.Println("[Control loop configuration deployment module] Stopping head sensor reading state:", this.stateName)
    fmt.Println("[Control loop configuration deployment module] Starting head sensor reading state:", this.stateName)
    for {
        s := this.service.getState(this.stateName).(*headState)
        select {
            case <- ctx.Done():
                return
            case this.output <- s:
        }
    }
}

/******************************************************************************
 * controller
 ******************************************************************************/

type controller struct {
    pumpChan chan *labelSignal
    inp chan *headState
}

func newController(pumpChan chan *labelSignal, inp chan *headState) (this *controller) {
    this = new(controller)
    this.pumpChan = pumpChan
    this.inp = inp
    return
}

func (this *controller) run(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()
    defer fmt.Println("[Control loop configuration deployment module] Stopping controller.")
    fmt.Println("[Control loop configuration deployment module] Starting controller.")
    for {
        select {
            case <- ctx.Done():
                return
            case  s := <-this.inp:
                 _, head := s.data()
                if head > 7 {
                    select {
                        case <- ctx.Done():
                            return
                        case this.pumpChan <- &labelSignal{lab: off}:
                    }
                } else if head < 3 {
                    select {
                        case <- ctx.Done():
                            return
                        case this.pumpChan <- &labelSignal{lab: on}:
                    }
                }
        }
    }
}

/******************************************************************************
 * junction estimator
 ******************************************************************************/

type junctionEstimator struct {
    inflow chan *flowState
    demand chan *flowState
    outflow chan *flowState
}

func newJunctionEstimator(inflow chan *flowState, demand chan *flowState, outflow chan *flowState) (this *junctionEstimator) {
    this = new(junctionEstimator)
    this.inflow = inflow
    this.demand = demand
    this.outflow = outflow
    return this
}

func (this *junctionEstimator) run(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()
    defer fmt.Println("[Control loop configuration deployment module] Stopping junction estimator.")
    fmt.Println("[Control loop configuration deployment module] Starting juntion estimator.")
    for {
        select {
            case <- ctx.Done():
                return
            case  inflowState := <- this.inflow:
                  select {
                      case <- ctx.Done():
                          return
                      case demandState := <- this.demand:
                          in_timestamp, inflow := inflowState.data()
                          demand_timestamp, demand := demandState.data()
                          if in_timestamp != demand_timestamp {
                              continue
                          }

                          outflow := inflow - demand

                          if outflow < 0 {
                              outflow = 0
                          }

                          select {
                              case <- ctx.Done():
                                  return
                              case this.outflow <- newFlowState(in_timestamp, outflow):
                          }
                }
        }
    }
}

/******************************************************************************
 * demand estimator
 ******************************************************************************/

type demandEstimator struct {
    demand chan *flowState
    inflow chan *flowState
}

func newDemandEstimator(demand chan *flowState, inflow chan *flowState) (this *demandEstimator) {
    this = new(demandEstimator)
    this.demand = demand
    this.inflow = inflow
    return this
}

func (this *demandEstimator) run(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()
    defer fmt.Println("[Control loop configuration deployment module] Stopping demand estimator.")
    fmt.Println("[Control loop configuration deployment module] Starting demand estimator.")
    for {
        select {
            case <- ctx.Done():
                return
            case  demandState := <- this.demand:
                  select {
                      case <- ctx.Done():
                          return
                      case this.inflow <- demandState:

                  }
        }
    }
}

/******************************************************************************
 * tank estimator
 ******************************************************************************/

type tankEstimator struct {
    capacity float64
    height float64
    volume float64
    timestamp int
    inflow chan *flowState
    outflow chan *flowState
    output chan *headState
}

func newTankEstimator(capacity float64, height float64, timestamp int, volume float64, inflow chan *flowState, outflow chan *flowState, output chan *headState) (this *tankEstimator) {
    this = new(tankEstimator)
    this.capacity = capacity
    this.height = height
    this.timestamp = timestamp
    this.volume = volume
    this.inflow = inflow
    this.outflow = outflow
    this.output = output
    return this
}

func (this *tankEstimator) run(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()
    defer fmt.Println("[Control loop configuration deployment module] Stopping tank estimator.")
    fmt.Println("[Control loop configuration deployment module] Starting tank estimator.")
    for {
        select {
            case <- ctx.Done():
                return
            case  inflowState := <- this.inflow:
                  select {
                      case <- ctx.Done():
                          return
                      case outflowState := <- this.outflow:
                          next_in_timestamp, inflow := inflowState.data()
                          next_out_timestamp, outflow := outflowState.data()
                          inVol := 0.0
                          outVol := 0.0
                          if next_in_timestamp == next_out_timestamp && this.timestamp < next_in_timestamp {
                              inVol = float64(next_in_timestamp - this.timestamp) * inflow
                              outVol = float64(next_out_timestamp - this.timestamp) * outflow
                              this.timestamp = next_in_timestamp
                          }

                          this.volume = this.volume + inVol - outVol
                          if this.volume < 0 {
                              this.volume = 0
                          } else if this.volume > this.capacity {
                              this.volume = this.capacity
                          }
                          level := this.volume / this.height
                          select {
                              case <- ctx.Done():
                                  return
                              case this.output <- newHeadState(this.timestamp, level):
                          }
                }
        }
    }
}

/******************************************************************************
 * pump estimator
 ******************************************************************************/

type pumpEstimator struct {
    pumpPressure float64

    inphead chan *headState
    outhead chan *headState
    outflow chan *flowState
}

func newPumpEstimator(pumpPressure float64, inphead chan *headState, outhead chan *headState, outflow chan *flowState) (this *pumpEstimator) {
    this = new(pumpEstimator)
    this.pumpPressure = pumpPressure
    this.inphead = inphead
    this.outhead = outhead
    this.outflow = outflow
    return this
}

func (this *pumpEstimator) run(ctx context.Context, wg *sync.WaitGroup) {
    defer wg.Done()
    defer fmt.Println("[Control loop configuration deployment module] Stopping pump flow estimator.")
    fmt.Println("[Control loop configuration deployment module] Starting pump flow estimator.")
    for {
        select {
            case <- ctx.Done():
                return
            case  inpheadState := <- this.inphead:
                  select {
                      case <- ctx.Done():
                          return
                      case outheadState := <- this.outhead:
                          in_timestamp, inhead := inpheadState.data()
                          out_timestamp, outhead := outheadState.data()

                          if in_timestamp != out_timestamp {
                              continue
                          }

                          flow := 0.0
                          if inhead != outhead {
                              flow = this.pumpPressure
                          }

                          select {
                              case <- ctx.Done():
                                  return
                              case this.outflow <- newFlowState(in_timestamp, flow):
                          }
                }
        }
    }
}


/*********************
 * deploy components
 ********************/

func deployScenario(stateServ *stateService, pumpChan chan *labelSignal, spec util.JsonMap) (context.CancelFunc, *sync.WaitGroup) {
    fmt.Println("[Control loop configuration deployment module] Deploying control loop configuration.")
    ctx, cancel := context.WithCancel(context.Background())
    wg := new(sync.WaitGroup)

    flowChannels := make(map[string]chan *flowState)
    headChannels := make(map[string]chan *headState)

    channels := spec.GetMap("channels")
    for name, ch := range channels {
        channel := ch.(string)
        if channel == "flowChannel" {
            flowChannels[name] = make(chan *flowState)
        } else if channel == "headChannel" {
            headChannels[name] = make(chan *headState)
        }
    }
    concurrentComponents := make(map[string]concurrentComponent)

    components := spec.GetMap("agents")
    wg.Add(len(components))

    for name, component := range components {
        comp := component.(util.JsonMap)
        if comp.GetString("type") == "flowSensor" {
            channel := comp.GetString("channel")
            state := comp.GetString("state")
            ch := flowChannels[channel]
            concurrentComponents[name] = newFlowSensor(stateServ, state, ch)
        } else if comp.GetString("type") == "headSensor" {
            channel := comp.GetString("channel")
            state := comp.GetString("state")
            ch := headChannels[channel]
            concurrentComponents[name] = newHeadSensor(stateServ, state, ch)
        } else if comp.GetString("type") == "controller" {
            headChannel := comp.GetString("inp")
            ch := headChannels[headChannel]
            concurrentComponents[name] = newController(pumpChan, ch)
        } else if comp.GetString("type") == "tankEstimator" {
            capacity := comp.GetFloat("capacity")
            height := comp.GetFloat("height")
            volume := comp.GetFloat("volume")
            timestamp := comp.GetInt("timestamp")
            inflowChannel := comp.GetString("inflowChannel")
            outflowChannel := comp.GetString("outflowChannel")
            headChannel := comp.GetString("headChannel")
            inflowC := flowChannels[inflowChannel]
            outflowC := flowChannels[outflowChannel]
            headC := headChannels[headChannel]

            concurrentComponents[name] = newTankEstimator(capacity, height, timestamp, volume, inflowC, outflowC, headC)
        } else if comp.GetString("type") == "junctionEstimator" {
            inflowChannel := comp.GetString("inflowChannel")
            demandChannel := comp.GetString("demandChannel")
            outflowChannel := comp.GetString("outflowChannel")
            inflowC := flowChannels[inflowChannel]
            demandC := flowChannels[demandChannel]
            outflowC := flowChannels[outflowChannel]

            concurrentComponents[name] = newJunctionEstimator(inflowC, demandC, outflowC)
        } else if comp.GetString("type") == "pumpEstimator" {
            pumpPressure := comp.GetFloat("pumpPressure")
            headChannel1 := comp.GetString("outpressure")
            headChannel2 := comp.GetString("inpressure")
            outflowChannel := comp.GetString("outflow")
            headC1 := headChannels[headChannel1]
            headC2 := headChannels[headChannel2]
            outflowC := flowChannels[outflowChannel]

            concurrentComponents[name] = newPumpEstimator(pumpPressure, headC1, headC2, outflowC)
        } else if comp.GetString("type") == "demandEstimator" {
            demandChannel := comp.GetString("demandChannel")
            inflowChannel := comp.GetString("inflowChannel")
            demandC := flowChannels[demandChannel]
            inflowC := flowChannels[inflowChannel]
            concurrentComponents[name] = newDemandEstimator(demandC, inflowC)
         } //else {
        //     fmt.Println("yeay!")
        // }
    }

    for _, cc := range concurrentComponents {
        go cc.run(ctx, wg)
    }

   return cancel, wg
}
