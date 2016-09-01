package main

import (
	"flag"
	"log"

	"github.com/nlopes/slack"
)

var slackAPIKey = "..."

func init() {
	flag.StringVar(&slackAPIKey, "slack", slackAPIKey, "slack bot users api key")
}

func mindSlack() {
	api := slack.New(slackAPIKey)
	api.SetDebug(false)
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	for {
		msg := <-rtm.IncomingEvents
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if ev.User == "" {
				// Known not to have user data sometimes...
				if ev.SubType == "message_changed" {
					break
				}
				// Something new not having user data...
				log.Printf("*slack.MessageEvent missing user data: %#v", *ev)
				break
			}
			seen.saw(ev.User)
		case *slack.PresenceChangeEvent:
			if ev.User == "" {
				log.Printf("*slack.PresenceChangeEvent missing user data: %#v", *ev)
				break
			}
			seen.saw(ev.User)
		case *slack.UserTypingEvent:
			if ev.User == "" {
				log.Printf("*slack.UserTypingEvent missing user data: %#v", *ev)
				break
			}
			seen.saw(ev.User)
		case *slack.UserChangeEvent:
			if ev.User.ID == "" {
				log.Printf("*slack.UserChangeEvent missing user data: %#v", *ev)
				break
			}
			seen.saw(ev.User.ID)
		case *slack.TeamJoinEvent:
			if ev.User.ID == "" {
				log.Printf("*slack.TeamJoinEvent missing user data: %#v", *ev)
				break
			}
			seen.saw(ev.User.ID)
		case *slack.RTMError:
			log.Printf(ev.Error())
		case *slack.InvalidAuthEvent:
			log.Fatal("Invalid credentials")
		}
	}
}
