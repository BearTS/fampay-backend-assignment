package cmd

import (
	"fmt"

	"github.com/BearTS/fampay-backend-assignment/dbapp/pkg"
	"github.com/BearTS/fampay-backend-assignment/pkg/db"
	"github.com/spf13/cobra"
)

func Migrate() *cobra.Command {
	return &cobra.Command{
		Use: "migrate",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbConnection, sqlConnection := db.Connection()
			defer sqlConnection.Close()

			begin := dbConnection.Begin()

			for i, migrate := range pkg.AutoMigrate(begin) {
				if err := migrate.Run(begin); err != nil {
					begin.Rollback()
					fmt.Println("[Migrate] Running raw sql schema creation failed")
					panic(err)
				}
				fmt.Println("[", i, "]: ", "Migrate table: ", migrate.TableName)
			}
			begin.Commit()
			fmt.Println("Migration Completed")
			return nil
		},
	}
}
