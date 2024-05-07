package ast

import "fmt"
import "os"
import "bufio"

import "main/ast/util"

type program struct {
    expressions map[string]expression
    suffix string
    infoFile string
    quit bool
}

func NewProgram() (this *program) {
    this = new(program)
    this.expressions = make(map[string]expression)
    this.suffix = ".dss"
    this.infoFile = "./res/help.res"
    this.quit = false
    return
}

func (this *program) RunTerminal() {
    reader := bufio.NewReader(os.Stdin)
    for this.quit == false {
        fmt.Print(">: ")
        input, _ := reader.ReadString('\n')

        log := newCustomErrorListener("terminal")
        expressions := parseString(input, log)

        for _, expr := range expressions {
            expr.setFilename("terminal")
        }

        if log.HasErrors() == true {
            log.PrintErrors()
        } else {
            _, log := this.executeStack(expressions)
            if log.HasErrors() {
                log.PrintErrors()
            }
        }
    }
}

func (this *program) executeStack(expressions []expression) (expression, util.ErrorLog) {
    var expr expression = nil
    for _, stmt := range expressions {
        var log util.ErrorLog
        expr, log = stmt.execute(this)
        if log.HasErrors() == true {
            return nil, log
        }
    }
    return expr, util.NewErrorLog()
}
