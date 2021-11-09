package paperspace

import (
	"context"
	"fmt"
	"time"

	"github.com/Paperspace/go-graphql-client"
)

// SSHKey defines an SSH Key registered with Paperspace.
type SSHKey struct {
	ID         string
	Name       string
	PublicKey  string
	DtCreated  time.Time
	DtModified time.Time
	DtDeleted  *time.Time
}

// CreateSSHKey calls the createSSHKey mutation.
func (c Client) CreateSSHKey(ctx context.Context, x CreateSSHKeyInput) (*SSHKey, error) {
	m := new(sshKeyCreateMutation)
	if err := c.graphql.Mutate(ctx, m, x.vars()); err != nil {
		return nil, fmt.Errorf("createSSHKey: request failed %v", err)
	}
	return m.sshKey(x.Name, x.PublicKey), nil
}

// CreateSSHKeyInput is the arg to CreateSSHKey.
type CreateSSHKeyInput struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
}

func (x CreateSSHKeyInput) vars() map[string]interface{} {
	return inputvars(x)
}

type sshKeyCreateMutation struct {
	CreateSSHKey struct {
		SSHKey struct {
			ID         graphql.String
			DtCreated  graphql.String
			DtModified graphql.String
		} `graphql:"sshKey"`
	} `graphql:"createSSHKey(input: $input)"`
}

func (m *sshKeyCreateMutation) sshKey(name string, publicKey string) *SSHKey {
	x := m.CreateSSHKey.SSHKey
	dtc, _ := time.Parse(time.RFC3339, string(x.DtCreated))
	dtm, _ := time.Parse(time.RFC3339, string(x.DtModified))
	return &SSHKey{
		ID:         string(x.ID),
		Name:       name,
		PublicKey:  publicKey,
		DtCreated:  dtc,
		DtModified: dtm,
	}
}

// DeleteSSHKey calls the deleteSSHKey mutation.
func (c Client) DeleteSSHKey(ctx context.Context, x DeleteSSHKeyInput) (*SSHKey, error) {
	m := new(sshKeyDeleteMutation)
	if err := c.graphql.Mutate(ctx, m, x.vars()); err != nil {
		return nil, fmt.Errorf("deleteSSHKey: request failed %v", err)
	}
	return m.sshKey(x.ID), nil
}

// DeleteSSHKeyInput is the arg to DeleteSSHKey.
type DeleteSSHKeyInput struct {
	ID string `json:"id"`
}

func (x DeleteSSHKeyInput) vars() map[string]interface{} {
	return inputvars(x)
}

type sshKeyDeleteMutation struct {
	DeleteSSHKey struct {
		SSHKey struct {
			Name       graphql.String
			PublicKey  graphql.String
			DtCreated  graphql.String
			DtModified graphql.String
			DtDeleted  graphql.String
		} `graphql:"sshKey"`
	} `graphql:"deleteSSHKey(input: $input)"`
}

func (m *sshKeyDeleteMutation) sshKey(id string) *SSHKey {
	x := m.DeleteSSHKey.SSHKey
	dtc, _ := time.Parse(time.RFC3339, string(x.DtCreated))
	dtm, _ := time.Parse(time.RFC3339, string(x.DtModified))
	dtd, _ := time.Parse(time.RFC3339, string(x.DtDeleted))
	return &SSHKey{
		ID:         id,
		Name:       string(x.Name),
		PublicKey:  string(x.PublicKey),
		DtCreated:  dtc,
		DtModified: dtm,
		DtDeleted:  &dtd,
	}
}

// GetSSHKey retrieves a single ssh key.
func (c Client) GetSSHKey(ctx context.Context, name string) (*SSHKey, error) {
	q := new(sshKeyQuery)
	if err := c.graphql.Query(ctx, q, inputvars(graphql.String(name))); err != nil {
		return nil, fmt.Errorf("sshKey: request failed: %v", err)
	}
	return q.sshKey(name), nil
}

type sshKeyQuery struct {
	SSHKey struct {
		ID         graphql.String
		PublicKey  graphql.String
		DtCreated  graphql.String
		DtModified graphql.String
		DtDeleted  graphql.String
	} `graphql:"sshKey(name: $input)"`
}

func (q *sshKeyQuery) sshKey(name string) *SSHKey {
	x := q.SSHKey
	dtc, _ := time.Parse(time.RFC3339, string(x.DtCreated))
	dtm, _ := time.Parse(time.RFC3339, string(x.DtModified))
	dtd, _ := time.Parse(time.RFC3339, string(x.DtDeleted))
	return &SSHKey{
		ID:         string(x.ID),
		Name:       name,
		PublicKey:  string(x.PublicKey),
		DtCreated:  dtc,
		DtModified: dtm,
		DtDeleted:  &dtd,
	}
}

// ListSSHKeys retrieves ssh keys for a given user.
func (c Client) ListSSHKeys(ctx context.Context) ([]*SSHKey, error) {
	q := new(sshKeysQuery)
	if err := c.graphql.Query(ctx, q, nil); err != nil {
		return nil, fmt.Errorf("sshKeys: request failed: %v", err)
	}
	return q.sshKeys(), nil
}

type sshKeysQuery struct {
	SSHKeys struct {
		Nodes []struct {
			ID         graphql.String
			Name       graphql.String
			PublicKey  graphql.String
			DtCreated  graphql.String
			DtModified graphql.String
			DtDeleted  graphql.String
		} `graphql:"nodes"`
	} `graphql:"sshKeys(first: 100)"`
}

func (q *sshKeysQuery) sshKeys() []*SSHKey {
	nodes := q.SSHKeys.Nodes
	y := make([]*SSHKey, len(nodes))
	for i, x := range nodes {
		dtc, _ := time.Parse(time.RFC3339, string(x.DtCreated))
		dtm, _ := time.Parse(time.RFC3339, string(x.DtModified))
		dtd, _ := time.Parse(time.RFC3339, string(x.DtDeleted))
		y[i] = &SSHKey{
			ID:         string(x.ID),
			Name:       string(x.Name),
			PublicKey:  string(x.PublicKey),
			DtCreated:  dtc,
			DtModified: dtm,
			DtDeleted:  &dtd,
		}
	}
	return y
}

func inputvars(x interface{}) map[string]interface{} {
	return map[string]interface{}{"input": x}
}
