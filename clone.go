package cogi

import (
	"sync"
)

type CloneResults struct {
	Succeed []string
	Failed  []string
	Errors  []error

	l sync.Mutex
}

func (cr *CloneResults) AddSuccess(cloneURL string) {
	cr.l.Lock()
	cr.Succeed = append(cr.Succeed, cloneURL)
	cr.l.Unlock()
}

func (cr *CloneResults) AddFail(cloneURL string, reason error) {
	cr.l.Lock()
	cr.Failed = append(cr.Failed, cloneURL)
	cr.Errors = append(cr.Errors, reason)
	cr.l.Unlock()
}
