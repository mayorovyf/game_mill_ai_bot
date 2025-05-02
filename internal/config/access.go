package config

type AccessLevel string

const (
	AccessUser       AccessLevel = "user"
	AccessModerator  AccessLevel = "moderator"
	AccessAdmin      AccessLevel = "admin"
	AccessSuperAdmin AccessLevel = "superadmin"
)
