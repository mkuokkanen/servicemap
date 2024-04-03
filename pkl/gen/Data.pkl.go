// Code generated from Pkl module `Data`. DO NOT EDIT.
package gen

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Data struct {
	Title string `pkl:"title"`

	Env string `pkl:"env"`

	Groups []*Group `pkl:"groups"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Data
func LoadFromPath(ctx context.Context, path string) (ret *Data, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Data
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Data, error) {
	var ret Data
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
