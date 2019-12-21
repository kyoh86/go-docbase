// +build sample

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/kyoh86/go-docbase/docbase"
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
			Do(context.Background())
		fmt.Println(res.Response.StatusCode)
		if err != nil {
			_ = log.Output(0, err.Error())
		}
		fmt.Println(jsonify(post))
		samplePost = *post
	}
	{
		res, err := client.Post.Delete(context.Background(), samplePost.ID)
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
