package migrations

import "fmt"

// Migration represents a single migration step.
type Migration struct {
	ID   string
	Up   []string
	Down []string
}

// Generate generates the list of migrations using the given table name.
func Generate(tableName string) []Migration {
	return []Migration{
		{
			ID: "001_create_casbin_table",
			Up: []string{fmt.Sprintf(`
				CREATE TABLE %s (
					id    BIGSERIAL PRIMARY KEY,
					ptype VARCHAR(100) INDEX,
					v0    VARCHAR(100) INDEX,
					v1    VARCHAR(100) INDEX,
					v2    VARCHAR(100) INDEX,
					v3    VARCHAR(100) INDEX,
					v4    VARCHAR(100) INDEX,
					v5    VARCHAR(100) INDEX
				)
			`, tableName)},
			Down: []string{fmt.Sprintf(`
				DROP TABLE %s
			`, tableName)},
		},
	}
}
