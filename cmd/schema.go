package cmd

import (
	// 	"github.com/appwrite/sdk-for-go/appwrite"
	// "github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/tablesdb"
	"github.com/spf13/cobra"
)

const (
	dbName                 = "attendexDb"
	membersTable           = "members"
	contestsTable          = "contests"
	eventsTable            = "events"
	regularAttendanceTable = "participants" // for simple events
	contestAttendanceTable = "teams"        // for complex events / contests
)

var (
	Orm *tablesdb.TablesDB // tablesDB API gives SQL-like terminology
	Db  *models.Database

	// Tables
	MembersTable           *models.Table
	EventsTable            *models.Table
	ContestsTable          *models.Table
	RegularAttendanceTable *models.Table // for regular events
	ContestAttendanceTable *models.Table // for contests involving teams
)

// schemaCmd represents the schema command
var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Populate DB migration on AppWrite",
	Long: `Populate your AppWrite database with migration data using this command.

Use this command to apply database schema changes and manage migrations
directly on your AppWrite backend from the CLI. It supports setting up 
collections and other schema related updates to keep your database consistent 
with your application models. As a safety measure, this runs only in Dev mode.
You`,
	RunE: schemaFunc,
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}

func schemaFunc(cmd *cobra.Command, args []string) error {
	initDatabase()
	// if err != nil {
	// 	return err
	// }
	//
	// err = setupMembersTable()
	// if err != nil {
	// 	return err
	// }
	//
	// err = setupEventsTable()
	// if err != nil {
	// 	return err
	// }
	//
	// err = setupContestsTable()
	// if err != nil {
	// 	return err
	// }
	//
	// err = setupRegularAttendanceTable()
	// if err != nil {
	// 	return err
	// }
	//
	// err = setupContestAttendanceTable()
	// if err != nil {
	// 	return err
	// }
	//
	// return nil
	return nil
}

func initDatabase() {
	// 	Orm = appwrite.NewTablesDB(appwriteClient)

	// 	Db, err := Orm.Create(
	// 		id.Unique(),
	// 		dbName,
	// 		Orm.WithCreateEnabled(true))
	// 	if err != nil {
	// 		// err
	// 		panic("Error occured")
	// 	}

	// 	MembersTable, err = Orm.CreateTable(
	// 		Db.Id,
	// 		id.Unique(),
	// 		membersTable,
	// 	)
	// 	if err != nil {
	// 		// TODO: Log the error
	// 		panic("error occured")
	// 	}

	// 	EventsTable, err = Orm.CreateTable(
	// 		Db.Id,
	// 		id.Unique(),
	// 		eventsTable,
	// 	)
	// 	if err != nil {
	// 		// TODO: Log the error
	// 		panic("error occured")
	// 	}

	// 	ContestsTable, err = Orm.CreateTable(
	// 		Db.Id,
	// 		id.Unique(),
	// 		contestsTable,
	// 	)
	// 	if err != nil {
	// 		// TODO: Log the error
	// 		panic("error occured")
	// 	}

	// 	RegularAttendanceTable, err = Orm.CreateTable(
	// 		Db.Id,
	// 		id.Unique(),
	// 		regularAttendanceTable,
	// 	)
	// 	if err != nil {
	// 		// TODO: Log the error
	// 		panic("error occured")
	// 	}

	// 	MembersTable, err = Orm.CreateTable(
	// 		Db.Id,
	// 		id.Unique(),
	// 		contestAttendanceTable,
	// 	)
	// 	if err != nil {
	// 		// TODO: Log the error
	// 		panic("error occured")
	// 	}
	// }

	// func setupMembersTable() error {
	// 	_, err := Orm.CreateEmailColumn(
	// 		Db.Id,
	// 		MembersTable.Id,
	// 		"email",
	// 		true,
	// 	)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	_, err = Orm.CreateStringColumn(
	// 		Db.Id,
	// 		MembersTable.Id,
	// 		"fullName",
	// 		255,
	// 		true,
	// 	)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	_, err = Orm.CreateBooleanColumn(
	// 		Db.Id,
	// 		MembersTable.Id,
	// 		"isPresent",
	// 		true,
	// 	)

	return
}

func setupEventsTable() error {
	return nil
}

func setupContestsTable() error {
	return nil
}

func setupRegularAttendanceTable() error {
	return nil
}

func setupContestAttendanceTable() error {
	return nil
}
