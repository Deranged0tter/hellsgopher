package hellsgopher

import (
	"os/user"
)

// returns the username of the user the process is running under
func GetUsername() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return currentUser.Username, nil
}

// returns the display name of the user the process is running under
func GetDisplayName() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return currentUser.Name, nil
}

// returns the user id of the user the process is running under
func GetUID() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return currentUser.Uid, nil
}

// returns the group id of the user the process is running under
func GetGID() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	return currentUser.Gid, nil
}

// list all the users on the systen
func ListUsers() ([]string, error) {
	return nil, ErrFuncNotSupported
}
