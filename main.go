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
			Name:   "corpid",
			Usage:  "The corpid to get the access token",
			EnvVar: "PLUGIN_CORPID,WEBHOOK_CORPID",
		},
		cli.StringFlag{
			Name:   "corp-secret",
			Usage:  "The corpsecret to get the access token",
			EnvVar: "PLUGIN_CORP_SECRET,WEBHOOK_CORP_Secret",
		},
		cli.StringFlag{
			Name:   "agentid",
			Usage:  "Agent ID",
			EnvVar: "PLUGIN_AGENT_ID,WEBHOOK_AGENT_ID",
		},
		cli.StringFlag{
			Name:   "msgtype",
			Usage:  "The type of message, either text, textcard",
			EnvVar: "PLUGIN_MSGTYPE,WEBHOOK_MSGTYPE",
			Value:  "textcard",
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
			Name:   "totag",
			Usage:  "Tag ID to send messages to",
			EnvVar: "PLUGIN_TO_TAG",
		},
		cli.StringFlag{
			Name:   "title",
			Usage:  "Message title",
			EnvVar: "PLUGIN_TITLE",
		},
		cli.StringFlag{
			Name:   "description",
			Usage:  "Description ",
			EnvVar: "PLUGIN_DESCRIPTION",
		},
		cli.StringFlag{
			Name:   "msgurl",
			Usage:  "message url ",
			EnvVar: "PLUGIN_MSG_URL",
		},
		cli.StringFlag{
			Name:   "btntxt",
			Usage:  "Button text ",
			EnvVar: "PLUGIN_BTN_TXT",
		},
		cli.StringFlag{
			Name:   "safe",
			Usage:  "Whether to make this message confidential or not, 0 is false, 1 is true. Defaults to false",
			EnvVar: "PLUGIN_SAFE",
			Value:  "0",
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
			CorpID:      c.String("corpid"),
			CorpSecret:  c.String("corp-secret"),
			Agentid:     c.Int("agentid"),
			MsgType:     c.String("msgtype"),
			MsgURL:      c.String("msgurl"),
			BtnTxt:      c.String("btntxt"),
			ToUser:      c.String("touser"),
			ToParty:     c.String("toparty"),
			ToTag:       c.String("totag"),
			Title:       c.String("title"),
			Description: c.String("description"),
			Safe:        c.Int("safe"),
			ContentType: c.String("content-type"),
			Debug:       c.Bool("debug"),
			SkipVerify:  c.Bool("skip-verify"),
		},
	}
	return plugin.Exec()
}
