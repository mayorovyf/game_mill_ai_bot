// internal/models/admin.go
package models

type GroupAdmin struct {
	UserID     int64           `bson:"user_id"`
	Title      string          `bson:"title"`
	Privileges AdminPrivileges `bson:"privileges"`
}
