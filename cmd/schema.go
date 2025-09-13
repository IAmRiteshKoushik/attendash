package cmd

import (
	"fmt"

	"github.com/IAmRiteshKoushik/attendash/utils"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/tablesdb"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

const (
	dbName = "attendexDb"

	studentsTable        = "students"
	eventsTable          = "events"
	soloAttendanceTable  = "solo_participants"
	eventTeamsTable      = "teams"
	teamsAttendanceTable = "teams_participants"
)

var (
	Orm *tablesdb.TablesDB // tablesDB API gives SQL-like terminology

	// Tables
	StudentsTable        *models.Table
	EventsTable          *models.Table
	SoloAttendanceTable  *models.Table
	EventTeamsTable      *models.Table
	TeamsAttendanceTable *models.Table
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
	if err := initDatabase(); err != nil {
		log.Error("Failed to initialize DB", err)
		return err
	}
	log.Info("Initialized DB")

	if err := setupStudentsTable(); err != nil {
		log.Error("Failed to setup `students` table", err)
		return err
	}
	log.Info("Created `students` table")

	if err := setupEventsTable(); err != nil {
		log.Error("Failed to setup `events` table", err)
		return err
	}
	log.Info("Initialized `events` table successfully")

	if err := setupSoloAttendanceTable(); err != nil {
		log.Error("Failed to setup solo_participants table", err)
		return err
	}
	log.Info("Initialized `solo_participants` table successfully")

	if err := setupEventTeamsTable(); err != nil {
		log.Error("Failed to setup `teams` table", err)
		return err
	}
	log.Info("Initialized `teams` table successfully")

	if err := setupTeamAttendanceTable(); err != nil {
		log.Error("Failed to setup `teams_participants` table", err)
		return err
	}
	log.Info("Initialized `teams_participants` table successfully")

	if err := setupRelationships(); err != nil {
		log.Error("Failed to setup relationships", err)
		return err
	}
	log.Info("Initialized relationships successfully")

	return nil
}

func initDatabase() error {

	_, err := Orm.Create(
		dbName,
		dbName,
		Orm.WithCreateEnabled(true))
	if err != nil {
		return utils.ErrorString(fmt.Sprintf("%v", err))
	}

	StudentsTable, err = Orm.CreateTable(
		dbName,
		id.Unique(),
		studentsTable,
	)
	if err != nil {
		return err
	}

	EventsTable, err = Orm.CreateTable(
		dbName,
		id.Unique(),
		eventsTable,
	)
	if err != nil {
		return err
	}

	SoloAttendanceTable, err = Orm.CreateTable(
		dbName,
		id.Unique(),
		soloAttendanceTable,
	)
	if err != nil {
		return err
	}

	EventTeamsTable, err = Orm.CreateTable(
		dbName,
		id.Unique(),
		eventTeamsTable,
	)
	if err != nil {
		return err
	}

	TeamsAttendanceTable, err = Orm.CreateTable(
		dbName,
		id.Unique(),
		teamsAttendanceTable,
	)
	if err != nil {
		return err
	}

	return nil
}

func setupStudentsTable() error {
	if _, err := Orm.CreateEmailColumn(
		dbName,
		StudentsTable.Id,
		"email",
		true,
		Orm.WithCreateEmailColumnArray(false),
	); err != nil {
		return err
	}

	if _, err := Orm.CreateStringColumn(
		dbName,
		StudentsTable.Id,
		"fullName",
		255,
		true,
		Orm.WithCreateStringColumnEncrypt(false),
		Orm.WithCreateStringColumnArray(false),
	); err != nil {
		return err
	}

	if _, err := Orm.CreateBooleanColumn(
		dbName,
		StudentsTable.Id,
		"isPresent",
		true,
		Orm.WithCreateBooleanColumnArray(false),
	); err != nil {
		return err
	}

	return nil
}

func setupEventsTable() error {
	if _, err := Orm.CreateStringColumn(
		dbName,
		EventsTable.Id,
		"eventName",
		255,
		true,
		Orm.WithCreateStringColumnEncrypt(false),
		Orm.WithCreateStringColumnArray(false),
	); err != nil {
		return err
	}

	if _, err := Orm.CreateBooleanColumn(
		dbName,
		EventsTable.Id,
		"isOffline",
		true,
		Orm.WithCreateBooleanColumnArray(false),
	); err != nil {
		return err
	}

	if _, err := Orm.CreateStringColumn(
		dbName,
		EventsTable.Id,
		"eventLocation",
		255,
		true,
		Orm.WithCreateStringColumnEncrypt(false),
		Orm.WithCreateStringColumnArray(false),
	); err != nil {
		return err
	}

	if _, err := Orm.CreateDatetimeColumn(
		dbName,
		EventsTable.Id,
		"dateAndTime",
		true,
		Orm.WithCreateDatetimeColumnArray(false),
	); err != nil {
		return err
	}

	if _, err := Orm.CreateEnumColumn(
		dbName,
		EventsTable.Id,
		"label",
		[]string{"solo", "team"},
		true,
		Orm.WithCreateEnumColumnArray(false),
	); err != nil {
		return err
	}

	return nil
}

func setupSoloAttendanceTable() error {
	if _, err := Orm.CreateStringColumn(
		dbName,
		SoloAttendanceTable.Id,
		"fullName",
		255,
		true,
		Orm.WithCreateStringColumnEncrypt(false),
		Orm.WithCreateStringColumnArray(false),
	); err != nil {
		return err
	}

	if _, err := Orm.CreateEmailColumn(
		dbName,
		SoloAttendanceTable.Id,
		"email",
		true,
		Orm.WithCreateEmailColumnArray(false),
	); err != nil {
		return err
	}

	if _, err := Orm.CreateBooleanColumn(
		dbName,
		SoloAttendanceTable.Id,
		"isPresent",
		true,
		Orm.WithCreateBooleanColumnArray(false),
	); err != nil {
		return err
	}

	return nil
}

func setupEventTeamsTable() error {
	if _, err := Orm.CreateStringColumn(
		dbName,
		EventTeamsTable.Id,
		"teamName",
		255,
		true,
		Orm.WithCreateStringColumnArray(false),
		Orm.WithCreateStringColumnEncrypt(false),
	); err != nil {
		return err
	}

	if _, err := Orm.CreateIntegerColumn(
		dbName,
		EventsTable.Id,
		"teamSize",
		true,
		Orm.WithCreateIntegerColumnMin(1),
		Orm.WithCreateIntegerColumnArray(false),
	); err != nil {
		return err
	}

	return nil
}

func setupTeamAttendanceTable() error {
	if _, err := Orm.CreateStringColumn(
		dbName,
		TeamsAttendanceTable.Id,
		"fullName",
		255,
		true,
		Orm.WithCreateStringColumnEncrypt(false),
		Orm.WithCreateStringColumnArray(false),
	); err != nil {
		return err
	}

	if _, err := Orm.CreateEmailColumn(
		dbName,
		TeamsAttendanceTable.Id,
		"email",
		true,
		Orm.WithCreateEmailColumnArray(false),
	); err != nil {
		return err
	}

	if _, err := Orm.CreateBooleanColumn(
		dbName,
		TeamsAttendanceTable.Id,
		"isPresent",
		true,
		Orm.WithCreateBooleanColumnArray(false),
	); err != nil {
		return err
	}

	return nil
}

// This is not a table, it just sets up all the references
func setupRelationships() error {
	// Events -> SoloAttendance (2 sided)
	if _, err := Orm.CreateRelationshipColumn(
		dbName,
		EventsTable.Id,
		SoloAttendanceTable.Id,
		"oneToMany",
		Orm.WithCreateRelationshipColumnTwoWay(true),
		Orm.WithCreateRelationshipColumnKey(soloAttendanceTable),
		Orm.WithCreateRelationshipColumnTwoWayKey(eventTeamsTable),
		Orm.WithCreateRelationshipColumnOnDelete("cascade"),
	); err != nil {
		return err
	}

	// Events -> TeamsAttendance (2 sided)
	if _, err := Orm.CreateRelationshipColumn(
		dbName,
		EventsTable.Id,
		TeamsAttendanceTable.Id,
		"oneToMany",
		Orm.WithCreateRelationshipColumnTwoWay(true),
		Orm.WithCreateRelationshipColumnKey(teamsAttendanceTable),
		Orm.WithCreateRelationshipColumnTwoWayKey(eventsTable),
		Orm.WithCreateRelationshipColumnOnDelete("cascade"),
	); err != nil {
		return err
	}

	// EventTeams -> TeamsAttendance (2 sided)
	if _, err := Orm.CreateRelationshipColumn(
		dbName,
		EventTeamsTable.Id,
		TeamsAttendanceTable.Id,
		"oneToMany",
		Orm.WithCreateRelationshipColumnTwoWay(true),
		Orm.WithCreateRelationshipColumnKey(teamsAttendanceTable),
		Orm.WithCreateRelationshipColumnTwoWayKey(eventTeamsTable),
		Orm.WithCreateRelationshipColumnOnDelete("cascade"),
	); err != nil {
		return err
	}
	return nil
}
