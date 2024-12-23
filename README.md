[![Open in Code Ocean](https://codeocean.com/codeocean-assets/badge/open-in-code-ocean.svg)](https://codeocean.com/capsule/1441773/tree/)

# iCPS-DL

The iCPS-DL (industrial Cyber Physical Systems - Description Language) is a description language that maps an industrial process as a knowledge graph. It includes information about

* physical and cyber-physical components,
* a state estimation model, and
* software component interaction.

A novel aspect language is the use of communication semantics to ensure correct interaction among distributed entities. Reasoning on the knowledge graph facilitates the configuration of cyber-physical elements in an industrial system.


## Compile

### Requirements for compilation
* java runtime environment
* Go compiler 1.23

### Compilation commands

		java -jar .\res\antlr-4.10.1-complete.jar -Dlanguage=Go -no-listener -visitor .\iCPSDL\parser\ontology.g4

compiles the antlr4 file containing the language grammar.

		go build main

compiles the Go program.


## Run

* Command

		.\main -sim

	executes a very basic simulation of the running example supervised by a proof-of-concept implementation of the autonomic supervisor.

* The running simulation deploys a set of ditributed agents communicating using Go channels to control a water distribution network simulation.

* The event manager module of the autonomic supervisor introduces failures at specific points of the simulation and invokes the semantic reasoning engine to reconfigure the control loop.

* When invoked, the semantic reasoning engine interacts with the iCPS-DL service and receives a new control loop configuration. It then invokes the control loop configuration deployment module.

* The control loop configuration deployment module will stop all running agents. It will materialise a new set of distributed agents corresponding to the new control loop configuration and deploy them to continue controlling the water network simulator.

## Interactive Run

* Executing:

		.\main -load script

* run the commands in the *script.dss* file.

* Specifically, it will load files */example/wdn.dss*, */example/simple.dss* and the */example/kb.* agent knowledge base. It will then perform a number of commands that find state estimation redundancies and compose agents based on their interactions.

* A number of *.html* files are outputted in the running folder.

* Users can provide their own commands through the script.dss file, as well as create their own modules.

## Modules
The language provides a syntax for defining the following modules:
* *paradigm*
* *process*
* *local configuration*
* *global configuration*
* *knowledge base*.

## Commands
Running the language in interactive mode allows to issue commands that manipulate and return one of the following modules:
* *paradigm*,
* *process*,
* *local configuration*,
* *global configuration*,
* *knowledge base*
* *analytical redundancy graph*,
* *analytical redundancy tree*, and
* *analytical redundancy forest*.

### load a file:
		load <filename>
opens, reads and executes the commands in file  `<filename>.dss`. The command returns the result of the last command in the file.

### assign a module to a name
		<identifier> := <command>
assigns the result of `<command>` to the `<identifier>` name. Returns the result of `command`.

### refer to a name
		<identifier>
referring to a name `<identifier>` returns the corresponding module.

### print a module
		show <command>?
		ls <command>?
prints the module returned by `<command>`. Commands `show` or `ls` without a `<command>` lists all names and their types. Returns the result of `<command>`.

### print metadata for a module
		info <command>
prints metadata on the module returned by `<command>`. Returns the result of `<command>`.

### create a mermaid diagram
		mermaid <command>
creates a mermaid diagram for module returned by `<command>`, only if the module is one of *process, analytical redundancy graph,  analytical redundancy tree, or analytical redundancy forest*. Returns the result of `<command>`.

### access an analytical redundancy forest
		<command> [<integer>]
returns *analytical redundancy tree* numbered as `<integer>`, whenever `<command>` returns an *analytical redundancy forest*.

### remove a name:
		rm <identifier>
removes name `<identifier>` from memory. Returns the result of `<identifier>`.

### translate a process
		translate <command>
translates the *process* returned by `<command>` into an *analytical redundancy graph* and returns the latter.

### traverse an analytical redundancy graph
		translate <node>.<property> <command>
traverses the *analytical redundancy graph* returned by `<command>` and finds all *analytical redundancy trees* rooted on node `<node>.<property>`. It returns an *analytical redundancy forest*.

### configure an analytical redundancy tree
		configure <command1> <command2> <controller> <actuator>
constructs and returns a *local configuration* that implements the *analytical redundancy tree* returned by `<command1>`, and inputs the result to controller <controller> that controls actuator <actuator>. The command takes local sessions from *Knowledge base* `<command2>`.

### compose a local context
		compose <command>
composes and returns a *global configuration* from *local configuration* `<command>`.

### project a global context
		compose <command>
projects and returns a *local configuration* from *global configuration* `<command>`.
