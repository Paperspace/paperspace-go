package paperspace_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Paperspace/paperspace-go"
)

func mockGQL(req, resp string) (*Client, *httptest.Server) {
	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		b, err := ioutil.ReadAll(r.Body)
		Expect(err).ToNot(HaveOccurred())
		Expect(b).Should(MatchJSON(req))

		w.Write([]byte(resp))
	})
	server := httptest.NewServer(mux)
	os.Setenv("PAPERSPACE_BASEURL", server.URL)
	return NewClient(), server
}

var _ = Describe("SSHKeys", func() {
	Context("CreateSSHKey", func() {
		It("should work", func() {
			client, server := mockGQL(
				`{
	"query": "mutation($input:CreateSSHKeyInput!){createSSHKey(input: $input){sshKey{id,name,publicKey,dtCreated,dtModified,dtDeleted}}}",
	"variables": { 
		"input": { 
			"name": "name", 
			"publicKey": "public-ssh-key"}}}`,
				`{
	"data":{
		"createSSHKey": {
			"sshKey": {
				"id": "some-uuid-v4",
				"name": "name",
				"publicKey": "public-ssh-key",
				"dtCreated": "2021-11-09T17:43:48.102Z",
				"dtModified": "2021-11-09T17:43:48.105Z",
				"dtDeleted": null}}}}`,
			)
			defer server.Close()

			ctx := context.Background()
			resp, err := client.CreateSSHKey(ctx, CreateSSHKeyInput{Name: "name", PublicKey: "public-ssh-key"})
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.ID).To(Equal("some-uuid-v4"))
		})
	})

	Context("DeleteSSHKey", func() {
		It("should work", func() {
			client, server := mockGQL(
				`{
	"query": "mutation($input:DeleteSSHKeyInput!){deleteSSHKey(input: $input){sshKey{id,name,publicKey,dtCreated,dtModified,dtDeleted}}}",
	"variables": { 
		"input": { 
			"id": "some-uuid-v4"}}}`,
				`{
	"data":{
		"deleteSSHKey":{
			"sshKey": {
				"id": "some-uuid-v4",
				"name": "name",
				"publicKey": "public-ssh-key",
				"dtCreated": "2021-11-09T17:43:48.102Z",
				"dtModified": "2021-11-09T17:43:48.105Z",
				"dtDeleted": "2021-11-09T17:43:48.105Z"}}}}`,
			)
			defer server.Close()

			ctx := context.Background()
			resp, err := client.DeleteSSHKey(ctx, DeleteSSHKeyInput{ID: "some-uuid-v4"})
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.DtDeleted).ToNot(BeNil())
		})
	})

	Context("GetSSHKey", func() {
		It("should work", func() {
			client, server := mockGQL(
				`{
	"query": "query($input:String!){sshKey(name: $input){id,name,publicKey,dtCreated,dtModified,dtDeleted}}",
	"variables": {
		"input": "name"}}`,
				`{
	"data":{
		"sshKey": {
			"id": "some-uuid-v4",
			"name": "name",
			"publicKey": "public-ssh-key",
			"dtCreated": "2021-11-09T17:43:48.102Z",
			"dtModified": "2021-11-09T17:43:48.105Z",
			"dtDeleted": null}}}}`,
			)
			defer server.Close()

			ctx := context.Background()
			resp, err := client.GetSSHKey(ctx, "name")
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.ID).To(Equal("some-uuid-v4"))
		})
	})

	Context("ListSSHKeys", func() {
		It("should work", func() {
			client, server := mockGQL(
				`{
	"query": "{sshKeys(first: 100){nodes{id,name,publicKey,dtCreated,dtModified,dtDeleted}}}"}`,
				`{
	"data":{
		"sshKeys": {
			"nodes": [{
			  "id": "some-uuid-v4",
			  "name": "name",
			  "publicKey": "public-ssh-key",
			  "dtCreated": "2021-11-09T17:43:48.102Z",
			  "dtModified": "2021-11-09T17:43:48.105Z",
			  "dtDeleted": null
			}, {
			  "id": "some-uuid-v4-1",
			  "name": "name-1",
			  "publicKey": "public-ssh-key",
			  "dtCreated": "2021-11-09T17:43:48.102Z",
			  "dtModified": "2021-11-09T17:43:48.105Z",
			  "dtDeleted": null
			}]}}}`,
			)
			defer server.Close()

			ctx := context.Background()
			resp, err := client.ListSSHKeys(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(resp)).To(Equal(2))
			Expect(resp[0].Name).To(Equal("name"))
			Expect(resp[1].Name).To(Equal("name-1"))
		})
	})
})
