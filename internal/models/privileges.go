// internal/models/privileges.go
package models

type AdminPrivileges struct {
	CanEditEvents     bool `bson:"can_edit_events"`
	CanManageMembers  bool `bson:"can_manage_members"`
	CanArchive        bool `bson:"can_archive"`
	CanDeleteMessages bool `bson:"can_delete_messages"`
	//...
}
