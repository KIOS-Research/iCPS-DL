################################################################################
                   Execution of the Autonomic Supervisor
################################################################################

The following execution demonstrates a basic simulation of a simple water
distribution network, supervised by a proof-of-concept implementation of an
autonomic supervisor that uses the functionalities of iCPS-DL.

The execution runs concurrently:

- The water distribution network simulation:
  The simulation is controlled by a pump, which receives ON or OFF signals via a
  Go channel.

- A set of distributed agents:
  This includes sensor, estimator, and control agents. The agents communicate
  using Go channels to measure and estimate values, and perform control sending
  signals to the pump actuator in the water distribution network simulation.

- The modules of the autonomic supervisor:
    - An iCPS-DL service: Maintains a knowledge base with the iCPS-DL description
      of the running simulation.
    - The event manager: Introduces CPS failures at specific time intervals and
      updates the knowledge base with changes to the simulation description.
    - The semantic reasoning engine: Triggered by the event manager when an
      event occurs. It interacts with the iCPS-DL service to request and obtain
      a control loop configuration.
    - The control loop configuration deployment module: Activated by the
      reasoning engine. This module stops all running agents and materialises
      the new control loop configuration by deploying a set of agents
      concurrently.

- A monitoring service:
  Outputs the tank water level at each simulation time interval, verifying the
  control law.

Each autonomic supervisor module outputs its activity. Additionally, whenever a
new control loop is selected for implementation, the iCPS-DL service outputs a
Mermaid diagram representing the corresponding state estimation tree.

###############################################################################
