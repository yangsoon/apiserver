package options

import (
	controlplane "github.com/yangsoon/apiserver/pkg/controlplane/apiserver/options"
)

// completedOptions is a private wrapper that enforces a call of Complete() before Run can be invoked.
type completedOptions struct {
	controlplane.CompletedOptions
}

type CompletedOptions struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedOptions
}

func (s *ServerRunOptions) Complete() (CompletedOptions, error) {
	co, err := s.Options.Complete()
	if err != nil {
		return CompletedOptions{}, err
	}
	return CompletedOptions{
		completedOptions: &completedOptions{
			CompletedOptions: co,
		},
	}, nil
}
