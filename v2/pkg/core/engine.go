package core

import (
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols"
	"github.com/projectdiscovery/nuclei/v2/pkg/types"
)

// Engine is an engine for running Nuclei Templates/Workflows.
//
// The engine contains multiple thread pools which allow using different
// concurrency values per protocol executed.
//
// The engine does most of the heavy lifting of execution, from clustering
// templates to leading to the final execution by the workpool, it is
// handled by the engine.
type Engine struct {
	workPool     *WorkPool
	options      *types.Options
	executerOpts protocols.ExecuterOptions
}

// InputProvider is an input provider interface for the nuclei execution
// engine.
//
// An example InputProvider is provided in form of hmap input provider.
type InputProvider interface {
	// Count returns the number of items for input provider
	Count() int64
	// Scan calls a callback function till the input provider is exhausted
	Scan(callback func(value string))
}

// New returns a new Engine instance
func New(options *types.Options) *Engine {
	workPool := NewWorkPool(WorkPoolConfig{
		InputConcurrency:         options.BulkSize,
		TypeConcurrency:          options.TemplateThreads,
		HeadlessInputConcurrency: options.HeadlessBulkSize,
		HeadlessTypeConcurrency:  options.HeadlessTemplateThreads,
	})
	engine := &Engine{
		options:  options,
		workPool: workPool,
	}
	return engine
}

// SetExecuterOptions sets the executer options for the engine. This is required
// before using the engine to perform any execution.
func (e *Engine) SetExecuterOptions(options protocols.ExecuterOptions) {
	e.executerOpts = options
}

// ExecuterOptions returns protocols.ExecuterOptions for nuclei engine.
func (e *Engine) ExecuterOptions() protocols.ExecuterOptions {
	return e.executerOpts
}
