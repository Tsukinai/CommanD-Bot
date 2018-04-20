package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
)

//type utilityCommand *commands

/*
func (uc *utilityCommand) command(s *discordgo.Session, m *discordgo.Message) error {
	args := ParceInput(m.Content)
	if len(args) < 2 {
		_, err := s.ChannelMessageSend(m.ChannelID, uc.commandInfo.Help())
		return err
	} else {
		return uc.subCommands[args[1]](s, m)
	}
}*/

func loadUtilityCommand() *commands {
	u := commands{}
	u.commandInfo = loadUtilityCommandInfo()
	u.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
	u.subCommands["-dice"] = diceRole
	u.subCommands["-d"] = diceRole

	return &u
}

// Create CommandInfo struct data //
func loadUtilityCommandInfo() *CommandInfo {

	u := &CommandInfo{}
	u.detail = "**!utility** or **!util** : Extra fun commands for utility with in your server."
	u.commands = make(map[string]string)
	u.commands["-dice"] = "**-dice** or **-d**.\n**Info**: Roles a dice.\n" +
		"**Arguments:**\n		**<Number of sides>**: Roles a dice for the given number of sides (number will be between 1 and the number you gave)."
	return u
}

// Roles a dice and prints the results //
// Returns an error (nil if non)
func diceRole(s *discordgo.Session, m *discordgo.Message) error {
	// Parce messages on a space //
	args := ParceInput(m.Content)

	// Convert the third argument to an int //
	// - Returns an error if err is not nil
	if val, err := StrToInt(args[2]); err != nil {
		return err
	} else {
		// Get a random number from 0 to the given value //
		rand.Seed(time.Now().Unix())
		rng := rand.Intn(val - 1)
		rng++

		// Print random number to channel //
		// - Returns an error if err is not nil
		if _, err := s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" got "+IntToStr(rng)); err != nil {
			return err
		}
	}
	return nil
}

// TODO - Implement IGN
func ign(s *discordgo.Session, m *discordgo.Message) error {
	return nil
}

// TODO - Implement Trinity
func trinity(s *discordgo.Session, m *discordgo.Message) error {
	return nil
}