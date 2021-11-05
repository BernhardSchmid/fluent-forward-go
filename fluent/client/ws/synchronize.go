package ws

import "fmt"

type synchronize struct {
	lock   func()
	unlock func()
}

func NewSynchronize(lock, unlock func()) (*synchronize, error) {
	if lock == nil || unlock == nil {
		return nil, fmt.Errorf("illegal argument")
	}
	return &synchronize{
		lock:   lock,
		unlock: unlock,
	}, nil
}

func (s *synchronize) Run(fs ...func() error) (err error) {
	s.lock()
	defer s.unlock()

	for _, f := range fs {
		err = f()
		if err != nil {
			return
		}
	}
	return
}
