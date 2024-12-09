package industrial_process

import "fmt"

import "autonomic/iCPSDL"
import "autonomic/iCPSDL/util"

//var eventSD chan int = make(chan int)
//var requestSD chan bool = make(chan bool)
//var replyEngine chan util.JsonMap = make(chan util.JsonMap)
var eventTrigger chan struct {timestamp int; head float64} = make(chan struct {timestamp int; head float64})

// func PseudoSemanticDescription() {
//     scenario := 1
//     for {
//         select {
//             case scenario = <- eventSD:
//             case <- requestSD:
//                 switch scenario {
//                     case 1:
//                         replyEngine <- scenario1
//                     case 2:
//                         replyEngine <- scenario2
//                     case 3:
//                         replyEngine <- scenario3
//                 }
//         }
//     }
// }

func SemanticReasoningEngine(stateServ *stateService, pumpChan chan *labelSignal, iCPSDLChannel chan *iCPSDL.ServiceMessage) {
    fmt.Println("[Semantic Reasoning Engine] Starting.")

    command := "wdn := load example/wdn\n"
    command += "agents := load example/agents\n"
    command += "simple := load example/simple"
    iCPSDLChannel <- &iCPSDL.ServiceMessage{Jmap: nil, Request: "initalise", Command: command, Process: "", State: "", ConfArgs: "", Ack: ""}
    <- iCPSDLChannel

    fmt.Println("[Semantic Reasoning Engine] Request control-loop configuration description from iCPSDL.")
    iCPSDLChannel <- &iCPSDL.ServiceMessage{Jmap: nil, Request: "configure", Command: "", Process: "simple", State: "t.head", ConfArgs: "agents controller u", Ack: ""}
    reply := <-iCPSDLChannel
    fmt.Println("[Semantic Reasoning Engine] Receive control-loop configuration description from iCPSDL.")
    fmt.Println(reply.Ack)

    scenario := reply.Jmap
    fmt.Println("[Semantic Reasoning Engine] Invoking the Control loop configuration deployment module.")
    cancel, wg := deployScenario(stateServ, pumpChan, scenario)

    for {
        select {
            case data := <- eventTrigger:
                //fmt.Println("[Semantic Reasoning Engine] Receiving request for control-loop reconfiguration.")

                fmt.Println("[Semantic Reasoning Engine] Request new control-loop configuration description from iCPSDL.")
                iCPSDLChannel <- &iCPSDL.ServiceMessage{Jmap: nil, Request: "configure", Command: "", Process: "faulty", State: "t.head", ConfArgs: "agents controller u", Ack: ""}
                reply := <-iCPSDLChannel

                fmt.Println("[Semantic Reasoning Engine] Receive new control-loop configuration description from iCPSDL.")
                fmt.Println(reply.Ack)

                scenario := reply.Jmap
                scenario["agents"].(util.JsonMap)["t.tank_mass"].(util.JsonMap)["capacity"] = 100.0
                scenario["agents"].(util.JsonMap)["t.tank_mass"].(util.JsonMap)["height"] = 10.0
                scenario["agents"].(util.JsonMap)["t.tank_mass"].(util.JsonMap)["timestamp"] = data.timestamp
                scenario["agents"].(util.JsonMap)["t.tank_mass"].(util.JsonMap)["volume"] = data.head * 10

                fmt.Println("[Semantic Reasoning Engine] Invoking the Control loop configuration deployment module.")
                cancel()
                wg.Wait()

                cancel, wg = deployScenario(stateServ, pumpChan, scenario)
        }
    }
}

func EventManager(p *plant, stateServ *stateService, iCPSDLChannel chan *iCPSDL.ServiceMessage) {
    fault := 0
    fmt.Println("[Event Manager] Starting.")

    for {
        s := stateServ.getState("tank.level")
        timestamp, head := s.data()
        if timestamp == 20 && fault < 1 {
            // introduce error
            fault = 1
            fmt.Println("[Event Manager] Error is introduced. Sensor s6 presents a fault!")

            command := "faulty := load example/faulty"
            fmt.Printf("[Event Manager] Update the iCPSDL Knowledge Base... Removing sensor 6. Running iCPSDL command: %v\n", command)

            iCPSDLChannel <- &iCPSDL.ServiceMessage{Jmap: nil, Request: "", Command: command, Process: "", State: "", ConfArgs: "", Ack: ""}
            <- iCPSDLChannel

            fmt.Println("[Event Manager] Invoking the Semantic Reasoning Engine.")
            // send also the timestamp and head data
            eventTrigger <- struct {timestamp int; head float64} {timestamp: timestamp, head: head}
        } else if timestamp == 40 && fault < 2 {
            // introduce error
            fault = 2
            fmt.Println("[Event Manager] Error is introduced. Device 2 presents a fault!")

            command := "faulty := load example/faulty2"
            fmt.Printf("[Event Manager] Update the iCPSDL Knowledge Base... Removing Device 2. Running iCPSDL command: %v\n", command)

            iCPSDLChannel <- &iCPSDL.ServiceMessage{Jmap: nil, Request: "", Command: command, Process: "", State: "", ConfArgs: "", Ack: ""}
            <- iCPSDLChannel

            fmt.Println("[Event Manager] Invoking the Semantic Reasoning Engine.")
            // send also the timestamp and head data
            eventTrigger <- struct {timestamp int; head float64} {timestamp: timestamp, head: head}
        } else if timestamp == 60 {
            p.Stop()
        }
    }
}
