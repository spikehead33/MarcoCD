package infras

import "marcocd/infras/env"

func New() (env.Env, error) {
	e, err := env.New()
	if err != nil {
		return e, err
	}

	return e, nil
}
