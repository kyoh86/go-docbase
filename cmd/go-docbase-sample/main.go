// +build sample

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/kyoh86/go-docbase/docbase"
	"github.com/kyoh86/go-docbase/docbase/postquery"
)

func main() {
	var (
		token  string
		domain string
	)
	flag.StringVar(&token, "t", "", "DocBase API Token")
	flag.StringVar(&domain, "d", "", "DocBase Domain")
	flag.Parse()

	fmt.Printf("A version of the Package %s is %s\n", "go-docbase", docbase.Version())

	client := docbase.NewAuthClient(domain, token)

	var samplePost docbase.Post
	var sampleComment docbase.Comment
	var sampleAttachment docbase.Attachment
	var sampleGroup docbase.Group

	{
		users, res, err := client.
			User.
			List().
			Query("Kyo").
			IncludeUserGroups(true).
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
		fmt.Println(jsonify(users[0]))
	}

	{
		post, res, err := client.
			Post.
			Create("testTitle", "testBody").
			Scope(docbase.ScopePrivate).
			Notice(false).
			Tags([]string{"go-docbase-test"}).
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
		fmt.Println(jsonify(post))
		samplePost = *post
	}
	{
		posts, res, err := client.
			Post.
			Get(samplePost.ID).
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
		fmt.Println(jsonify(posts))
	}
	{
		posts, res, err := client.
			Post.
			Edit(samplePost.ID).
			Body("testBody - updated").
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
		fmt.Println(jsonify(posts))
	}
	{
		posts, res, err := client.
			Post.
			List().
			Query(postquery.Title("testTitle")).
			Do(context.Background())
		fmt.Println(res.Response.StatusCode)
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(jsonify(posts))
	}
	{
		res, err := client.
			Post.
			Archive(samplePost.ID).
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
	}
	{
		tags, res, err := client.
			Tag.
			List().
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
		fmt.Println(jsonify(tags))
	}
	{
		res, err := client.
			Post.
			Unarchive(samplePost.ID).
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
	}
	{
		group, res, err := client.
			Group.
			Create("go-docbase-test").
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
		fmt.Println(jsonify(group))
		sampleGroup = *group
	}
	{
		group, res, err := client.
			Group.
			Get(sampleGroup.ID).
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
		fmt.Println(jsonify(group))
	}
	{
		res, err := client.
			Group.
			AddUsers(sampleGroup.ID, []docbase.UserID{6425}).
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
	}
	{
		res, err := client.
			Group.
			RemoveUsers(sampleGroup.ID, []docbase.UserID{6425}).
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
	}
	{
		attachments, res, err := client.
			Attachment.
			Upload().
			AddPayload("go-docbase-test.gif", []byte{
				0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x01, 0x00, 0x01, 0x00,
				0x80, 0x01, 0x00, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x21,
				0xf9, 0x04, 0x01, 0x0a, 0x00, 0x01, 0x00, 0x2c, 0x00, 0x00,
				0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x02, 0x02, 0x4c, 0x01, 0x00, 0x3b}).
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
		for _, attachment := range attachments {
			sampleAttachment = attachment
		}
	}
	{
		comment, res, err := client.
			Comment.
			Create(samplePost.ID, "test comment\n"+sampleAttachment.Markdown).
			Notice(false).
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
		fmt.Println(jsonify(comment))
		sampleComment = *comment
	}
	{
		res, err := client.
			Comment.
			Delete(sampleComment.ID).
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
	}
	{
		res, err := client.
			Post.
			Delete(samplePost.ID).
			Do(context.Background())
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(res.Response.StatusCode)
	}
}

func jsonify(o interface{}) string {
	buf, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(buf)
}
