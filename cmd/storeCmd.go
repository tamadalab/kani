package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/tamada/kani/utils"
)

var storeCmd = &cobra.Command{
	Use:    "store",
	Short:  "kani store",
	Long:   "store data to the specified database",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 5 {
			return fmt.Errorf("%v: required argument missing", args)
		}
		return storeImpl(cmd, args)
	},
}

func createDatabase(filePath string) (*sql.DB, error) {
	db, err := openDatabase(filePath)
	if err != nil {
		return nil, err
	}
	_, err2 := db.Exec(`CREATE TABLE histories (
   id          INTEGER PRIMARY KEY,
   datetime    TEXT    DEFAULT CURRENT_TIMESTAMP,
   command     TEXT    NOT NULL,
   status_code INTEGER NOT NULL,
   branch      TEXT,
   revision    TEXT
);`)
	if err2 != nil {
		return nil, err2
	}
	return db, nil
}

func openDatabase(filePath string) (*sql.DB, error) {
	return sql.Open("sqlite3", filePath)
}

func openOrCreateDatabase(filePath string) (*sql.DB, error) {
	if !utils.ExistFile(filePath) {
		return createDatabase(filePath)
	}
	return openDatabase(filePath)
}

func storeData(db *sql.DB, args []string) error {
	insert := `INSERT INTO histories (command, status_code, branch, revision) VALUES (?, ?, ?, ?)`
	r, err := db.Exec(insert, args[0], args[1], args[2], args[3])
	if err != nil {
		return err
	}
	rows, err2 := r.RowsAffected()
	if err2 == nil && rows != 1 {
		return fmt.Errorf("affected rows should be 1, but %d", rows)
	}
	return err2
}

// args[0]: file name of database.
// args[1]: executed command
// args[2]: status code
// args[3]: branch name
// args[4]: revision
func storeImpl(cmd *cobra.Command, args []string) error {
	db, err := openOrCreateDatabase(args[0])
	if err != nil {
		return err
	}
	defer db.Close()
	return storeData(db, args[1:])
}

func init() {
	RootCmd.AddCommand(storeCmd)
}
