package ast

import "fmt"
import "main/ast/util"

/******************************************************************************
 * knowledge base
 ******************************************************************************/

type knowledge_base struct {
    baseNode
    pdgmName string
    pdgm *paradigm
    roles []irole
    estimateAgents map[string]*estimate
    senseAgents map[string]*sense
    controlAgents map[string]*control
    actuateAgents map[string]*actuate
    arg *arGraph
}

func newKnowledgeBase(pdgmName string, line int) (this *knowledge_base) {
    this = new(knowledge_base)
    this.init(line)
    this.pdgmName = pdgmName
    this.roles = make([]irole, 0)
    this.estimateAgents = make(map[string]*estimate)
    this.senseAgents = make(map[string]*sense)
    this.controlAgents = make(map[string]*control)
    this.actuateAgents = make(map[string]*actuate)
    return
}

func (this *knowledge_base) getController(name string) session {
    controller, ok := this.controlAgents[name]
    if ok == false {
        return nil
    }
    return controller.getSession()
}

func (this *knowledge_base) getActuatorClass(class string) session {
    for _, r := range this.actuateAgents {
        if r.getClass() == class { return r.getSession() }
    }
    return nil
}
func (this *knowledge_base) getEstimatorClass(class string) session {
    for _, r := range this.estimateAgents {
        if r.getClass() == class { return r.getSession() }
    }
    return nil
}
func (this *knowledge_base) getSensorClass(class string) session {
    for _, r := range this.senseAgents {
        if r.getClass() == class { return r.getSession() }
    }
    return nil
}

func (this *knowledge_base) setFilename(filename string) {
    this.baseNode.setFilename(filename)
    for _, r := range this.roles {
        r.setFilename(filename)
    }
}

func (this *knowledge_base) add(r irole) {
    this.roles = append(this.roles, r)
}

func (this *knowledge_base) typeCheck(context *program, log util.ErrorLog) {
    pdgm, ok := context.expressions[this.pdgmName]
    if ok == false {
        error := fmt.Sprintf("Paradigm %v is not defined", this.pdgmName)
        this.reportError(error, log)
        return
    }
    this.pdgm, ok = pdgm.(*paradigm)
    if ok == false {
        error := fmt.Sprintf("Expression %v is not a paradigm", this.pdgmName)
        this.reportError(error, log)
        return
    }

    size := len(this.roles)
    for i := 0; i < size; i++ {
        for j := i + 1; j < size; j++ {
            if this.roles[i].getName() == this.roles[j].getName() {
                error := fmt.Sprintf(   "Duplicate definition of %v. Originally defined in %v:%v ",
                                        this.roles[j].getName(), this.roles[i].getFilename(), this.roles[i].getLine())
                this.roles[j].reportError(error, log)
            }
        }
    }

    for _, role := range this.roles {
        switch r := role.(type) {
            case *estimate:
                this.estimateAgents[r.getName()] = r
            case *sense:
                this.senseAgents[r.getName()] = r
            case *control:
                this.controlAgents[r.getName()] = r
            case *actuate:
                this.actuateAgents[r.getName()] = r
      }
  }

  for _, r := range this.roles {
      r.typeCheck(this, log)
  }
  return
}

func (this *knowledge_base) execute(context *program) (expression, util.ErrorLog) {
    _, ok := context.expressions[this.pdgmName]
    if ok == false {
        expr, log := newLoad(this.pdgmName, 0).execute(context)
        if log.HasErrors() {
            return nil, log
        }
        context.expressions[this.pdgmName] = expr
    }

    var log util.ErrorLog = util.NewErrorLog()

    this.typeCheck(context, log)
    return this, log
}

func (this *knowledge_base) render() {}

func (this *knowledge_base) prettyPrint(stream *util.Stream) {
    stream.Println("knowledge base " + this.pdgmName + " {" )
    stream.Inc()
    for _, r := range this.roles {
        r.prettyPrint(stream)
        stream.Println("")
    }
    stream.Dec()
    stream.Println("}")
}

func (this *knowledge_base) getType() string {
    return "knowledge base"
}

func (this *knowledge_base) getModule() string {
    return "*"
}

func (this *knowledge_base) info() string {
    s := "Type: knowledge base"
    s += fmt.Sprintf(", Paradigm: %v", this.pdgmName)
    s += fmt.Sprintf(", Estimator agents: %v", len(this.estimateAgents))
    s += fmt.Sprintf(", Sensor agents: %v", len(this.senseAgents))
    s += fmt.Sprintf(", Actuator agents: %v", len(this.actuateAgents))
    s += fmt.Sprintf(", Control agents: %v.", len(this.controlAgents))
    return s
}

/******************************************************************************
 * role
 ******************************************************************************/

type irole interface {
    getName() string
    getClass() string
    getSession() session
    getFilename() string
    setFilename(string)
    getLine() int
    reportError(string, util.ErrorLog)
    typeCheck(*knowledge_base,util.ErrorLog)
    prettyPrint(stream *util.Stream)
}


type role struct {
    baseNode
    class string
    name string
    sess session
}

func  (this *role) newRole(class string, name string, sess session, line int) {
    this.class = class
    this.name = name
    this.sess = sess
    this.init(line)
    return
}

func (this *role) getName() string {
    return this.name
}

func (this *role) getClass() string {
    return this.class
}

func (this *role) getSession() session {
    return this.sess
}

func (this *role) setFilename(filename string) {
    this.baseNode.setFilename(filename)
    this.sess.setFilename(filename)
}

func (this *role) typeCheck(context *knowledge_base, log util.ErrorLog) {
    this.sess.typeCheck(newSessionContext(context.pdgm), log)
}

func (this *role) prettyPrint(stream *util.Stream) {
    stream.Println(this.class + " using " + this.name + " = ")
    stream.Inc()
    this.sess.prettyPrint(stream)
    stream.Dec()
}

/******************************************************************************
 * estimator agent
 ******************************************************************************/

type estimate struct {
    role
}

func newEstimate(class string, name string, sess session, line int) (this *estimate) {
    this = new(estimate)
    this.newRole(class, name, sess, line)
    return
}

func (this *estimate) typeCheck(context *knowledge_base, log util.ErrorLog) {
    _, ok := context.pdgm.models[this.class]
    if ok == false {
        error := fmt.Sprintf("Model %v is undefined", this.class)
        this.reportError(error, log)
    }

    this.role.typeCheck(context, log)
}

func (this *estimate) prettyPrint(stream *util.Stream) {
    stream.Print("estimate ")
    this.role.prettyPrint(stream)
}

/******************************************************************************
 * sense agent
 ******************************************************************************/

type sense struct {
    role
}

func newSense(class string, name string, sess session, line int) (this *sense) {
    this = new(sense)
    this.newRole(class, name, sess, line)
    return
}

func (this *sense) typeCheck(context *knowledge_base, log util.ErrorLog) {
    _, ok := context.pdgm.properties[this.class]
    if ok == false {
        error := fmt.Sprintf("Property %v is undefined", this.class)
        this.reportError(error, log)
    }
    this.role.typeCheck(context, log)
}

func (this *sense) prettyPrint(stream *util.Stream) {
    stream.Print("sense ")
    this.role.prettyPrint(stream)
}

/******************************************************************************
 * control agent
 ******************************************************************************/

type control struct {
    role
}

func newControl(class string, name string, sess session, line int) (this *control) {
    this = new(control)
    this.newRole(class, name, sess, line)
    return
}

func (this *control) typeCheck(context *knowledge_base, log util.ErrorLog) {
    _, ok := context.pdgm.actuators[this.class]
    if ok == false {
        error := fmt.Sprintf("Actuator class %v is undefined", this.class)
        this.reportError(error, log)
    }
    this.role.typeCheck(context, log)
}

func (this *control) prettyPrint(stream *util.Stream) {
    stream.Print("control ")
    this.role.prettyPrint(stream)
}

/******************************************************************************
 * actuate agent
 ******************************************************************************/

type actuate struct {
    role
}

func newActuate(class string, name string, sess session, line int) (this *actuate) {
    this = new(actuate)
    this.newRole(class, name, sess, line)
    return
}

func (this *actuate) typeCheck(context *knowledge_base, log util.ErrorLog) {
    _, ok := context.pdgm.actuators[this.class]
    if ok == false {
        error := fmt.Sprintf("Actuator class %v is undefined", this.class)
        this.reportError(error, log)
    }
    this.role.typeCheck(context, log)
}

func (this *actuate) prettyPrint(stream *util.Stream) {
    stream.Print("actuate ")
    this.role.prettyPrint(stream)
}
