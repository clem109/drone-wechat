package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "WeChat work plugin"
	app.Usage = "Wechat Work plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.0+%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "method",
			Usage:  "webhook method",
			EnvVar: "PLUGIN_METHOD",
			Value:  "POST",
		},
		cli.StringFlag{
			Name:   "access-token",
			Usage:  "The access token for authorization",
			EnvVar: "PLUGIN_ACCESS_TOKEN,WEBHOOK_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:   "msgtype",
			Usage:  "The type of message, either text, textcard",
			EnvVar: "PLUGIN_MSGTYPE,WEBHOOK_MSGTYPE",
		},
		cli.StringFlag{
			Name:   "content-type",
			Usage:  "content type",
			EnvVar: "PLUGIN_CONTENT_TYPE",
			Value:  "application/json",
		},
		cli.StringFlag{
			Name:   "touser",
			Usage:  "The users to send the message to, @all for all users",
			EnvVar: "PLUGIN_TO_USER",
			Value:  "@all",
		},
		cli.StringFlag{
			Name:   "toparty",
			Usage:  "Party ID to send messages to",
			EnvVar: "PLUGIN_TO_PARTY",
		},
		cli.StringFlag{
			Name:   "safe",
			Usage:  "Whether to make this message confidential or not, 0 is false, 1 is true. Defaults to false",
			EnvVar: "PLUGIN_SAFE",
			Value:  "0",
		},
		cli.StringFlag{
			Name:   "content",
			Usage:  "custom template for webhook",
			EnvVar: "PLUGIN_CONTENT",
		},
		cli.StringSliceFlag{
			Name:   "headers",
			Usage:  "custom headers key map",
			EnvVar: "PLUGIN_HEADERS",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "enable debug information",
			EnvVar: "PLUGIN_DEBUG",
		},
		cli.BoolFlag{
			Name:   "skip-verify",
			Usage:  "skip ssl verification",
			EnvVar: "PLUGIN_SKIP_VERIFY",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Tag:     c.String("build.tag"),
			Number:  c.Int("build.number"),
			Event:   c.String("build.event"),
			Status:  c.String("build.status"),
			Commit:  c.String("commit.sha"),
			Ref:     c.String("commit.ref"),
			Branch:  c.String("commit.branch"),
			Author:  c.String("commit.author"),
			Message: c.String("commit.message"),
			Link:    c.String("build.link"),
			Started: c.Int64("build.started"),
			Created: c.Int64("build.created"),
		},
		Job: Job{
			Started: c.Int64("job.started"),
		},
		Config: Config{
			Method:      c.String("method"),
			AccessToken: c.String("access-token"),
			MsgType:     c.String("msgtype"),
			ToUser:      c.StringSlice("touser"),
			ToParty:     c.StringSlice("toparty"),
			Safe:        c.Bool("safe"),
			Content:     c.String("content"),
			ContentType: c.String("content-type"),
			Template:    c.String("template"),
			Headers:     c.StringSlice("headers"),
			Debug:       c.Bool("debug"),
			SkipVerify:  c.Bool("skip-verify"),
		},
	}
	return plugin.Exec()
}
