// Code generated (@generated) by entc, DO NOT EDIT.

package group

import (
	"github.com/facebookincubator/ent/entc/integration/ent/schema"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldActive holds the string denoting the active vertex property in the database.
	FieldActive = "active"
	// FieldExpire holds the string denoting the expire vertex property in the database.
	FieldExpire = "expire"
	// FieldType holds the string denoting the type vertex property in the database.
	FieldType = "type"
	// FieldMaxUsers holds the string denoting the max_users vertex property in the database.
	FieldMaxUsers = "max_users"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// Table holds the table name of the group in the database.
	Table = "groups"
	// FilesTable is the table the holds the files relation/edge.
	FilesTable = "files"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "files"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "group_file_id"
	// BlockedTable is the table the holds the blocked relation/edge.
	BlockedTable = "users"
	// BlockedInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	BlockedInverseTable = "users"
	// BlockedColumn is the table column denoting the blocked relation/edge.
	BlockedColumn = "group_blocked_id"
	// UsersTable is the table the holds the users relation/edge. The primary key declared below.
	UsersTable = "user_groups"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// InfoTable is the table the holds the info relation/edge.
	InfoTable = "groups"
	// InfoInverseTable is the table name for the GroupInfo entity.
	// It exists in this package in order to avoid circular dependency with the "groupinfo" package.
	InfoInverseTable = "group_infos"
	// InfoColumn is the table column denoting the info relation/edge.
	InfoColumn = "info_id"

	// FilesLabel holds the string label denoting the files edge type in the database.
	FilesLabel = "group_files"
	// BlockedLabel holds the string label denoting the blocked edge type in the database.
	BlockedLabel = "group_blocked"
	// UsersInverseLabel holds the string label denoting the users inverse edge type in the database.
	UsersInverseLabel = "user_groups"
	// InfoLabel holds the string label denoting the info edge type in the database.
	InfoLabel = "group_info"
)

// Columns holds all SQL columns are group fields.
var Columns = []string{
	FieldID,
	FieldActive,
	FieldExpire,
	FieldType,
	FieldMaxUsers,
	FieldName,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "group_id"}
)

var (
	fields = schema.Group{}.Fields()
	// descActive is the schema descriptor for active field.
	descActive = fields[0].Descriptor()
	// DefaultActive holds the default value for the active field.
	DefaultActive = descActive.Default.(bool)
	// descType is the schema descriptor for type field.
	descType = fields[2].Descriptor()
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator = descType.Validators[0].(func(string) error)
	// descMaxUsers is the schema descriptor for max_users field.
	descMaxUsers = fields[3].Descriptor()
	// DefaultMaxUsers holds the default value for the max_users field.
	DefaultMaxUsers = descMaxUsers.Default.(int)
	// MaxUsersValidator is a validator for the "max_users" field. It is called by the builders before save.
	MaxUsersValidator = descMaxUsers.Validators[0].(func(int) error)
	// descName is the schema descriptor for name field.
	descName = fields[4].Descriptor()
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator = func() func(string) error {
		validators := descName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
)
