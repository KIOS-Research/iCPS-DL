# iCPS-DL

# Compile and Run

## Requirements for compilation
* java runtime environment
* Go compiler 1.21

### compilation
		java -jar .\res\antlr-4.10.1-complete.jar -Dlanguage=Go -no-listener -visitor .\ast\parser\ontology.g4
compiles the antlr4 file containing the language grammar.

		go build main
compiles the Go program.

## Requirements for running
* mermaid command line interface

		.\main
runs the language.

# Modules
The language provides a syntax for defining the following modules:
* *paradigm*
* *process*
* *local configuration*
* *global configuration*
* *knowledge base*.

# Commands
Running the language in interactive mode allows to issue commands that manipulate and return one of the following modules:
* *paradigm*,
* *process*,
* *local configuration*,
* *global configuration*,
* *knowledge base*
* *analytical redundancy graph*,
* *analytical redundancy tree*, and
* *analytical redundancy forest*.

## load a file:
		load [filename]
opens, reads and executes the commands in file  `[filename].dss`. The command returns the result of the last command in the file.

## assign a module to a name
		[identifier] := [command]
assigns the result of `[command]` to the `[identifier]` name. Returns the result of `command`.

## refer to a name
		[identifier]
referring to a name `[identifier]` returns the corresponding module.

## print a module
		show [command]?
		ls [command]?
prints the module returned by `[command]`. Commands `show` or `ls` without a `[command]` lists all names and their types. Returns the result of `[command]`.

## print metadata for a module
		info [command]
prints metadata on the module returned by `[command]`. Returns the result of `[command]`.

## create a mermaid diagram
		mermaid [command]
creates a mermaid diagram for module returned by `[command]`, only if the module is one of *process, analytical redundancy graph,  analytical redundancy tree, or analytical redundancy forest*. Returns the result of `[command]`.

## access an analytical redundancy forest
		[command] [ integer ]
returns *analytical redundancy tree* numbered as `integer`, whenever `[command]` returns an *analytical redundancy forest*.

## remove a name:
		rm [identifier]
removes name `[identifier]` from memory. Returns the result of `[identifier]`.

## translate a process
		translate [command]
translates the *process* returned by `[command]` into an *analytical redundancy graph* and returns the latter.

## traverse an analytical redundancy graph
		translate [node].[property] [command]
traverses the *analytical redundancy graph* returned by `[command]` and finds all *analytical redundancy trees* rooted on node `[node].[property]`. It returns an *analytical redundancy forest*.

## configure an analytical redundancy tree
		configure [command1] [command2]
constructs and returns a *local configuration* that implements the *analytical redundancy tree* returned by `[command1]`.  The command takes local sessions from *Knowledge base* `[command2]`.

## compose a local context
		compose [command]
composes and returns a *global configuration* from *local configuration* `[command]`.

## project a global context
		compose [command]
projects and returns a *local configuration* from *global configuration* `[command]`.
