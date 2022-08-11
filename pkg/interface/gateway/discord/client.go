package discord

import "github.com/bwmarrin/discordgo"

type Client interface {
	ChannelMessageSend(channelID string, content string) error
}

type client struct {
	session *discordgo.Session
}

func NewClient(token string) (Client, error) {
	session, err := discordgo.New(token)
	if err != nil {
		return nil, err
	}
	return &client{
		session: session,
	}, nil
}

func (c client) ChannelMessageSend(channelID string, content string) error {
	if _, err := c.session.ChannelMessageSend(channelID, content); err != nil {
		return err
	}
	return nil
}
