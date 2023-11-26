package opa

import (
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/plugins"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/rs/xid"
)

func SetRego() (*plugins.Manager, *ast.Compiler) {
	// Create a new plugin manager without any configuration.
	pluginManager, err := plugins.New([]byte{}, xid.New().String(), inmem.New())
	if err != nil {
		return nil, nil
	}

	// Compile the RBAC module.
	module := `
		package rbac

		default allow = false

		allow {
			# Look up the list of projects the user has access too.
			project_roles := data.roles[input.user_id]

			# For each of the roles held by the user for the named project.
			project_role := project_roles[input.project]
			pr := project_role[_]

			# Lookup the permissions for the roles.
			permissions := data.permissions[pr]

			# For each role permission, check if there is a match.
			p := permissions[_]
			p == concat("", [input.action, ":", input.resource])
		}
	`
	compiler, err := ast.CompileModules(map[string]string{"rbac": module})
	if err != nil {
		return nil, nil
	}

	return pluginManager, compiler

}
