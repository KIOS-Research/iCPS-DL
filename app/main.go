package main
import "os"
import "main/ast"

 /*********************
   main
  *********************/

func main() {
    program := ast.NewProgram()
    //program.RunTerminal()
    program.Run(os.Args[1:])
}
