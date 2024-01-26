package hellsgopher

import (
	"bufio"
	"os/user"
	"strings"
)

// return a user.User for the current user
func GetCurrentUser() (*user.User, error) {
	return user.Current()
}

// get the current username
func GetCurrentUsername() (string, error) {
	u, err := GetCurrentUser()
	if err != nil {
		return "", err
	}

	return u.Name, nil
}

// get the current uid
func GetCurrentUid() (string, error) {
	u, err := GetCurrentUser()
	if err != nil {
		return "", err
	}

	return u.Uid, nil
}

// get the main gid for the current user
func GetCurrentGid() (string, error) {
	u, err := GetCurrentUser()
	if err != nil {
		return "", err
	}

	return u.Gid, nil
}

// get all gids for the current user
func GetCurrentGids() ([]string, error) {
	u, err := GetCurrentUser()
	if err != nil {
		return nil, err
	}

	return u.GroupIds()
}

// return a uid from a given username
func GetUidFromName(name string) (string, error) {
	u, err := user.Lookup(name)
	if err != nil {
		return "", err
	}

	return u.Uid, nil
}

// return a username from a given uid
func GetNameFromUid(uid string) (string, error) {
	u, err := user.LookupId(uid)
	if err != nil {
		return "", err
	}

	return u.Name, nil
}

// return a user.User from username
func GetUserFromName(name string) (*user.User, error) {
	u, err := user.Lookup(name)
	return u, err
}

// return a user.User from uid
func GetUserFromUid(uid string) (*user.User, error) {
	u, err := user.LookupId(uid)
	return u, err
}

// return a slice of all users on the machine
func GetAllUsers() ([]*user.User, error) {
	out, err := PsReturn("Get-WmiObject -Class Win32_UserAccount")
	if err != nil {
		return nil, err
	}

	var uNames []string

	scanner := bufio.NewScanner(strings.NewReader(out))
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "Name") {
			uNames = append(uNames, scanner.Text())
		}
	}

	var uNamesRefined []string

	for _, uName := range uNames {
		uNameSplit := strings.Split(uName, ": ")
		uNamesRefined = append(uNamesRefined, uNameSplit[1])
	}

	var users []*user.User
	for _, user := range uNamesRefined {
		u, err := GetUserFromName(user)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

// return a slice of all usernames on the machine
func GetAllUsernames() ([]string, error) {
	out, err := PsReturn("Get-WmiObject -Class Win32_UserAccount")
	if err != nil {
		return nil, err
	}

	var uNames []string

	scanner := bufio.NewScanner(strings.NewReader(out))
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "Name") {
			uNames = append(uNames, scanner.Text())
		}
	}

	var uNamesRefined []string

	for _, uName := range uNames {
		uNameSplit := strings.Split(uName, ": ")
		uNamesRefined = append(uNamesRefined, uNameSplit[1])
	}

	return uNamesRefined, nil
}
