package executers

import (
	gorountinesshow "goloveu/goinactioncode/goroutinesshow"
	"goloveu/goinactioncode/matcherssearch"
)

// Executer performs to run all the samples
type Executer interface {
	Run()
}

// executers represents A map of registered goinaction samples for running.
var executers = make(map[string]Executer)

func init() {
	var msexecuter matcherssearch.MatchersSearch
	executers["MatchersSearch"] = msexecuter

	var grshowexecuter gorountinesshow.GoRountinesShow
	executers["grShow"] = grshowexecuter
}

// RetrieveExecuter performs to a specified executer
func RetrieveExecuter(name string) Executer {
	executer, exists := executers[name]
	if !exists {
		executer = executers["MatchersSearch"]
	}
	return executer
}
