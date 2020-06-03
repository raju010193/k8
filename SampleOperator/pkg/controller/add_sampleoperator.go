package controller

import (
	"github.com/martin-helmich/SampleOperator/pkg/controller/sampleoperator"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, sampleoperator.Add)
}
