package configcatpreevaluate

// from: https://github.com/configcat/go-sdk/blob/218279137924bf88947c8cc8120aec6287b33104/user.go

import "strings"

// User is an object containing attributes to properly identify a given user for rollout evaluation.
type User struct {
	identifier string
	attributes map[string]string
}

// NewUser creates a new user object. The identifier argument is mandatory.
func NewUser(identifier string) *User {
	return NewUserWithAdditionalAttributes(identifier, "", "", map[string]string{})
}

// NewUserWithAdditionalAttributes creates a new user object with additional attributes. The identifier argument is mandatory.
func NewUserWithAdditionalAttributes(identifier string, email string, country string, custom map[string]string) *User {
	if len(identifier) == 0 {
		panic("identifier cannot be empty")
	}

	user := &User{identifier: identifier, attributes: map[string]string{}}
	user.attributes["identifier"] = identifier

	if len(email) > 0 {
		user.attributes["email"] = email
	}

	if len(country) > 0 {
		user.attributes["country"] = country
	}

	if len(custom) > 0 {
		for k, v := range custom {
			user.attributes[strings.ToLower(k)] = v
		}
	}

	return user
}

// GetAttribute retrieves a user attribute identified by a key.
func (user *User) GetAttribute(key string) string {
	val := user.attributes[strings.ToLower(key)]
	if len(val) > 0 {
		return val
	}

	return ""
}
