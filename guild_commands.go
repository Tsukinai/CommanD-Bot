package CommanD_Bot

import (
	"errors"
	"strconv"
)

type GuildCommands struct {
	commands map[string]func(*Root)error
}

func (m *GuildCommands) RunCommand(root *Root) error {
	return m.commands[root.CommandType()](root)
}

// Set guild command structure //
func LoadGuildCommand() *GuildCommands {
	// Create guild command structure //
	g := GuildCommands{}

	// Create sub command map //
	g.commands = make(map[string]func(*Root) error)

	/*
	// Add word filter command to map //
	g.commands["-wordfilter"] = wordFilter
	g.commands["-wf"] = wordFilter
	*/

	// Add ban time command to map //
	g.commands["-bantime"] = BanTime
	g.commands["-bt"] = BanTime

	/*
	// Add word filter threshold command to map //
	g.commands["-wordfilterthreshold"] = wordFilterThreshold
	g.commands["-wft"] = wordFilterThreshold

	// Add spam filter threshold command to map //
	g.commands["-spamfilterthreshold"] = spamFilterThreshold
	g.commands["-sft"] = spamFilterThreshold
	*/

	// Return a reference to the new guild command structure //
	return &g
}

// TODO: Deprecated function needs to be fixed and changed or removed
// Word filter modification function //
// - returns an error (nil if non)
func WordFilter(root *Root) error {
	// Check if user is an admin //
	// - returns an error if err is not nil
	// - returns if user is not an admin
	if admin, err := root.IsAdmin(); err != nil {
		return err
	} else if !admin {
		return root.MessageSend("You do not have permission to change the ban length.")
	}

	// Check to make sure all arguments were given //
	// - returns an error if they were not
	if len(root.Args()) < 4 {
		return errors.New("the number of messages with in the command was to short")
	}

	// Get guild to edit the world filter for //
	// - returns an error if err is not nil
	if g, err := root.GetGuild(); err != nil {
		return err
	} else {
		// Get server info for guild //
		// - returns an error if guild does not exist in list
		if server, ok := serverList[g.ID]; !ok {
			return errors.New("no server with in serverList")
		} else {
			// Change word for given argument //
			// - add: add a word to the list
			// - remove: remove the word from the list
			// - returns an error if it was not add or remove
			switch root.Args()[2] {
			case "add":
				// Add word to list //
				// - returns an error (nil if non)
				return server.EditWordFilter(root.Args()[3], true)
			case "remove":
				// Remove word from list //
				// - returns an error (nil if non)
				return server.EditWordFilter(root.Args()[3], false)
			default:
				// Return an error as argument was not add or remove
				return errors.New("argument given could not be understood")
			}
		}
	}
}

// Ban time change function //
// - returns an error (nil if non)
func BanTime(root *Root) error {
	// Check if user is an admin //
	// - return an error if err is not nil
	// - return if user is not an admin
	if admin, err := root.IsAdmin(); err != nil {
		return err
	} else if !admin {
		return root.MessageSend("You do not have permission to change the ban length.")
	}

	// Gets the guild the messages was created in //
	// - returns an error if err is not nil
	if guild, err := root.GetGuild(); err != nil {
		return err
	} else {
		// Get server info for guild //
		// - returns an error if guild does not exist in list
		server, ok := serverList[guild.ID]
		if !ok {
			return errors.New("guild did not exist in server list")
		}

		// Send the current ban time to the channel //
		// - returns an error if err is not nil
		if err := root.MessageSend("Ban time was "+strconv.Itoa(int(server.BanTime))); err != nil {
			return err
		}

		// Check length of args //
		// - return if no new time was given
		// - change time if a new time was given
		if len(root.CommandArgs()) <= 0 {
			return nil
		} else {
			// Convert the given time to an integer //
			// - return an error if err is not nil
			if time, err := strconv.Atoi(root.CommandArgs()[1]); err != nil {
				return err
			} else {
				// Check if given time is less then or equal to zero //
				// - return an error if true
				if time <= 0 {
					return errors.New("time was under 0 days")
				}
				// Set new ban time //
				server.BanTime = uint(time)

				// Print the new ban time to the server //
				// - return an error if err is not nil
				return root.MessageSend("Ban time has been set to "+strconv.Itoa(int(server.BanTime)))
			}
		}
	}
}

// TODO: Deprecated function needs to be fixed and changed or removed
// World filter threshold function //
// - return an error (nil if non)
func WordFilterThreshold(root *Root) error {
	// Check if user is an admin //
	// - returns an error if err is not nil
	// - returns if user is not an admin
	if admin, err := root.IsAdmin(); err != nil {
		return err
	} else if !admin {
		_, err := root.ChannelMessageSend(root.ChannelID, "You do not have permition to change word filter settings.")
		return err
	}

	// Get guild the message was sent in //
	// - returns an error if err is not nil
	if guild, err := root.GetGuild(); err != nil {
		return err
	} else {
		// Get server info for guild //
		// - returns an error if guild is not in server list
		server, ok := serverList[guild.ID]
		if !ok {
			return errors.New("guild was not in the serverList")
		}

		// Send message for current word filter threshold //
		// - returns an error if err is not nil
		if _, err := root.ChannelMessageSend(root.ChannelID, "Current word filter threshold is "+strconv.FormatFloat(server.wordFilterThresh, 'f', 2, 64)); err != nil {
			return err
		}

		// Check if new threshold was given //
		// - returns if false
		if len(root.Args()) < 3 {
			return nil
		} else {
			// Convert argument to float for new threshold //
			// - return an error if err is not nil
			if thresh, err := strconv.ParseFloat(root.Args()[2], 64); err != nil {
				return err
			} else {
				// Check if threshold given is with in range //
				// - returns an error if its not
				if thresh < 0.0 || thresh > 1.0 {
					return errors.New("given threshold was ether less then 0.0 or greater then 1.0")
				}

				// Set new threshold to server info //
				server.wordFilterThresh = thresh

				// Send message displaying changed word filter threshold //
				// - returns an error (nil if non)
				_, err := root.ChannelMessageSend(root.ChannelID, "New word filter threshold is "+strconv.FormatFloat(server.wordFilterThresh, 'f', 2, 64))
				return err
			}
		}
	}
}

// TODO: Deprecated function needs to be fixed and changed or removed
// Spam filter threshold function //
// - returns an error (nil if non)
func SpamFilterThreshold(root *Root) error {
	// Check if user is an admin //
	// - returns an error if err is not nil
	// - returns if user is not an admin
	if admin, err := root.IsAdmin(); err != nil {
		return err
	} else if !admin {
		return root.MessageSend("You do not have permition to change word filter settings.")
	}

	// Get guild the message was sent in //
	// - returns an error if err is not nil
	if guild, err := root.GetGuild(); err != nil {
		return err
	} else {
		// Get server info for guild //
		// - returns an error if guild is not in server list
		server, ok := serverList[guild.ID]
		if !ok {
			return errors.New("guild was not in the serverList")
		}

		// Send message for current spam filter threshold //
		// - returns an error if err is not nil
		if err := root.MessageSend("Current spam filter threshold is "+strconv.FormatFloat(server.spamFilterThresh, 'f', 2, 64)); err != nil {
			return err
		}

		// Check if new threshold was given //
		// - returns if false
		if len(root.Args()) < 3 {
			return nil
		} else {
			// Convert argument to float for new threshold //
			// - return an error if err is not nil
			if thresh, err := strconv.ParseFloat(root.Args()[2], 64); err != nil {
				return err
			} else {
				// Check if threshold given is with in range //
				// - returns an error if its not
				if thresh < 0.0 || thresh > 1.0 {
					return errors.New("given threshold was ether less then 0.0 or greater then 1.0")
				}

				// Set new threshold to server info //
				server.spamFilterThresh = thresh

				// Send message displaying changed spam filter threshold //
				// - returns an error (nil if non)
				return root.MessageSend("New word filter threshold is "+strconv.FormatFloat(server.spamFilterThresh, 'f', 2, 64))
			}
		}
	}
}
