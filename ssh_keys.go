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

type gqlSSHKey struct {
	ID         graphql.String
	Name       graphql.String
	PublicKey  graphql.String
	DtCreated  graphql.String
	DtModified graphql.String
	DtDeleted  *graphql.String
}

func (s gqlSSHKey) sshKey() *SSHKey {
	dtc, _ := time.Parse(time.RFC3339, string(s.DtCreated))
	dtm, _ := time.Parse(time.RFC3339, string(s.DtModified))
	key := &SSHKey{
		ID:         string(s.ID),
		Name:       string(s.Name),
		PublicKey:  string(s.PublicKey),
		DtCreated:  dtc,
		DtModified: dtm,
	}
	if s.DtDeleted != nil {
		dtd, _ := time.Parse(time.RFC3339, string(*s.DtDeleted))
		key.DtDeleted = &dtd
	}
	return key
}

// CreateSSHKey calls the createSSHKey mutation.
func (c Client) CreateSSHKey(ctx context.Context, x CreateSSHKeyInput) (*SSHKey, error) {
	m := new(sshKeyCreateMutation)
	if err := c.graphql.Mutate(ctx, m, inputvars(x)); err != nil {
		return nil, fmt.Errorf("createSSHKey: request failed %v", err)
	}
	return m.CreateSSHKey.SSHKey.sshKey(), nil
}

// CreateSSHKeyInput is the arg to CreateSSHKey.
type CreateSSHKeyInput struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
}

type sshKeyCreateMutation struct {
	CreateSSHKey struct {
		SSHKey gqlSSHKey `graphql:"sshKey"`
	} `graphql:"createSSHKey(input: $input)"`
}

// DeleteSSHKey calls the deleteSSHKey mutation.
func (c Client) DeleteSSHKey(ctx context.Context, x DeleteSSHKeyInput) (*SSHKey, error) {
	m := new(sshKeyDeleteMutation)
	if err := c.graphql.Mutate(ctx, m, x.vars()); err != nil {
		return nil, fmt.Errorf("deleteSSHKey: request failed %v", err)
	}
	return m.DeleteSSHKey.SSHKey.sshKey(), nil
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
		SSHKey gqlSSHKey `graphql:"sshKey"`
	} `graphql:"deleteSSHKey(input: $input)"`
}

// GetSSHKey retrieves a single ssh key.
func (c Client) GetSSHKey(ctx context.Context, name string) (*SSHKey, error) {
	q := new(sshKeyQuery)
	if err := c.graphql.Query(ctx, q, inputvars(graphql.String(name))); err != nil {
		return nil, fmt.Errorf("sshKey: request failed: %v", err)
	}
	return q.SSHKey.sshKey(), nil
}

type sshKeyQuery struct {
	SSHKey gqlSSHKey `graphql:"sshKey(name: $input)"`
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
		Nodes []gqlSSHKey `graphql:"nodes"`
	} `graphql:"sshKeys(first: 100)"`
}

func (q *sshKeysQuery) sshKeys() []*SSHKey {
	nodes := q.SSHKeys.Nodes
	y := make([]*SSHKey, len(nodes))
	for i, x := range nodes {
		y[i] = x.sshKey()
	}
	return y
}

func inputvars(x interface{}) map[string]interface{} {
	return map[string]interface{}{"input": x}
}
