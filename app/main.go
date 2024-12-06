package main

//import "fmt"
import "os"
import "autonomic/industrial_process"
import "autonomic/iCPSDL"

/******************************************************************************
 * main
 ******************************************************************************/

func main() {
    if len(os.Args) > 1 && os.Args[1] == "-sim" {
        simulation()
    } else {
      iCPSDL.Execute(os.Args[1:])
    }
}

func simulation() {
    plant, stateServ, pumpChan := industrial_process.NewPlant(3, 15, 100, 10)

    distribution := [][] float64 {
        {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5},
    }

    go stateServ.Monitor([]string {"tank.level"})

    waitPlant := plant.Deploy(distribution)

    eventChannel := make(chan *iCPSDL.ServiceMessage)
    engineChannel := make(chan *iCPSDL.ServiceMessage)

    go iCPSDL.Service(eventChannel, engineChannel)

    //go industrial_process.PseudoSemanticDescription()
    go industrial_process.SemanticReasoningEngine(stateServ, pumpChan, engineChannel)
    go industrial_process.EventManager(plant, stateServ, eventChannel)

    plant.Start()
    waitPlant.Wait()
}
