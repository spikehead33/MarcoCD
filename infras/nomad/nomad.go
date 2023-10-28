package nomad

import "github.com/hashicorp/nomad/api"

type NomadClient struct {
	client *api.Client
}

func (c *NomadClient) RunJob(job *api.Job) error {
	c.client.Jobs().Register(job, nil)
	return nil
}

func (client *NomadClient) PurgeJob() error {
	return nil
}

func (cliet *NomadClient) ListJob() (*api.Job, error) {
	return nil, nil
}

func New(addr string) (NomadClient, error) {
	config := api.Config{
		Address: addr,
	}

	client, err := api.NewClient(&config)
	if err != nil {
		return NomadClient{}, err
	}

	return NomadClient{client: client}, nil
}
