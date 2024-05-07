package main
import "main/ast"

 /*********************
   main
  *********************/

func main() {
    program := ast.NewProgram()
    program.RunTerminal()
}
