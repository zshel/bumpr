package engine

import (
	"fmt"
	"github.com/packagrio/bumpr/pkg/config"
	"github.com/packagrio/go-common/errors"
	"github.com/packagrio/go-common/pipeline"
	"github.com/packagrio/go-common/scm"
)

func Create(engineType string, pipelineData *pipeline.Data, configImpl config.Interface, sourceImpl scm.Interface) (Interface, error) {

	var eng Interface

	switch engineType {
	case "chef":
		eng = new(engineChef)
	case "generic":
		eng = new(engineGeneric)
	case "golang":
		eng = new(engineGolang)
	case "node":
		eng = new(engineNode)
	case "python":
		eng = new(enginePython)
	case "ruby":
		eng = new(engineRuby)
	default:
		return nil, errors.EngineUnspecifiedError(fmt.Sprintf("Unknown Engine Type: %s", engineType))
	}

	if err := eng.Init(pipelineData, configImpl, sourceImpl); err != nil {
		return nil, err
	}
	return eng, nil
}
