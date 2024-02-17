// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/schema"
	"app/ent/todo"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescStatus is the schema descriptor for status field.
	todoDescStatus := todoFields[2].Descriptor()
	// todo.DefaultStatus holds the default value on creation for the status field.
	todo.DefaultStatus = todoDescStatus.Default.(bool)
}
