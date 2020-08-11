package interpreter

import (
	"fmt"
	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
	"github.com/sirupsen/logrus"

	"io/ioutil"
	"os"
	"path/filepath"
)

type FunctionSignature func(payload interface{}, params map[string]string) error

type Interpreter struct {
	path        string
	interpreter *interp.Interpreter
	loadedHooks map[string]FunctionSignature
}

func New(path string) (*Interpreter, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not find working directorz of process: %w", err)
	}

	i := interp.New(interp.Options{
		GoPath: filepath.Join(wd, path),
	})

	i.Use(stdlib.Symbols)
	i.Use(Symbols)

	files, err := filepath.Glob(filepath.Join(path, "src", "*", "*.go"))
	if err != nil {
		return nil, fmt.Errorf("error while scanning for source files: %w", err)
	}

	for _, f := range files {
		logrus.Infof("loading hook from file %s", f)
		content, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("could not load source files: %w", err)
		}
		_, err = i.Eval(string(content))
		if err != nil {
			return nil, fmt.Errorf("could not eval source file %s: %w", f, err)
		}
	}

	return &Interpreter{
		interpreter: i,
		path:        path,
		loadedHooks: make(map[string]FunctionSignature),
	}, nil
}

func (interp *Interpreter) LoadFunction(name string) error {
	logrus.Infof("attempting to load hook function %s", name)
	v, err := interp.interpreter.Eval(name)
	if err != nil {
		return err
	}

	// Interface must be equivalent but cannot be of type FunctionSignature\
	// Update both accordingly
	hook, ok := v.Interface().(func(interface{}, map[string]string) error)
	if !ok {
		return fmt.Errorf("function is of type %T and not of expected type func(interface {}) error", v.Interface())
	}
	interp.loadedHooks[name] = hook
	logrus.Infof("function %s successfully loaded", name)
	return nil
}

func (interp *Interpreter) RunLoadedHook(name string, payload interface{}, params map[string]string) error {
	if payload == nil {
		return fmt.Errorf("empty payload passed to runner")
	}

	hook, loaded := interp.loadedHooks[name]
	if !loaded {
		return fmt.Errorf("no hook with name %s has been loaded", name)
	}

	if err := hook(payload, params); err != nil {
		return fmt.Errorf("error while running webhook %s: %w", name, err)
	}

	return nil
}
