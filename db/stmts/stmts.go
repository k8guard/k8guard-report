package stmts

/*
   This package we need to store CQL statements
*/

const (
	SELECT_ACTIONS = `
		SELECT * FROM %s.alog_namespace_type LIMIT ?;

	`

	SELECT_ACTIONS_BY_NAMESPACE = `
		SELECT * FROM %s.alog_namespace_type WHERE namespace = ? AND type = ? LIMIT ?;

	`
)
