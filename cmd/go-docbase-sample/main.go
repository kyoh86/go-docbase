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

	transport := docbase.TokenTransport{Token: token}
	client := docbase.NewClient(domain, transport.Client())

	{
		users, res, err := client.
			User.
			List().
			Query("Kyo").
			IncludeUserGroups(true).
			Do(context.Background())
		fmt.Println(res.Response.StatusCode)
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(jsonify(users[0]))
	}

	var samplePost docbase.Post
	{
		post, res, err := client.
			Post.
			Create("testTitle", "testBody").
			Scope(docbase.ScopePrivate).
			Notice(false).
			Tags([]string{"go-docbase-test"}).
			Do(context.Background())
		fmt.Println(res.Response.StatusCode)
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(jsonify(post))
		samplePost = *post
	}
	{
		posts, res, err := client.
			Post.
			Get(samplePost.ID).
			Do(context.Background())
		fmt.Println(res.Response.StatusCode)
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(jsonify(posts))
	}
	{
		posts, res, err := client.
			Post.
			Edit(samplePost.ID).
			Body("testBody - updated").
			Do(context.Background())
		fmt.Println(res.Response.StatusCode)
		if err != nil {
			_ = log.Output(0, err.Error())
		}
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
		fmt.Println(res.Response.StatusCode)
		if err != nil {
			_ = log.Output(0, err.Error())
		}
	}
	{
		res, err := client.
			Post.
			Unarchive(samplePost.ID).
			Do(context.Background())
		fmt.Println(res.Response.StatusCode)
		if err != nil {
			_ = log.Output(0, err.Error())
		}
	}
	{
		res, err := client.
			Post.
			Delete(samplePost.ID).
			Do(context.Background())
		fmt.Println(res.Response.StatusCode)
		if err != nil {
			_ = log.Output(0, err.Error())
		}
	}

}

func jsonify(o interface{}) string {
	buf, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(buf)
}
