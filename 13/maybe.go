package main

import "fmt"

type MaybeBool interface {
	Get() (bool, error)
}

type Just struct {
	val bool
}

func (just Just) Get() (bool, error) {
	return just.val, nil
}

type Nothing struct{}

// Nothing still needs a dummy value, I chose true
const NOTHING_VALUE = true

func (Nothing) Get() (bool, error) {
	return NOTHING_VALUE, fmt.Errorf("cannot get a value from Nothing")
}
