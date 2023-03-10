// Code generated by ent, DO NOT EDIT.

package ent

import (
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent/form"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	formFields := schema.Form{}.Fields()
	_ = formFields
	// formDescIsDeleted is the schema descriptor for is_deleted field.
	formDescIsDeleted := formFields[4].Descriptor()
	// form.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	form.DefaultIsDeleted = formDescIsDeleted.Default.(bool)
}
