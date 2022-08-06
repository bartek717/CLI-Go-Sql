package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	user := "postgres"
	password := "Bartek2004"
	host := "host.docker.internal"
	port := "5432"
	dbname := "postgres"
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// get commands
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAllCmd := getCmd.Bool("all", false, "Get all users")
	getIdCmd := getCmd.String("id", "", "Get infp by user id")

	// add commands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addMarkCmd := addCmd.String("mark", "", "Add mark by user id")
	addNameCmd := addCmd.String("name", "", "Add name by user id")
	addSurnameCmd := addCmd.String("surname", "", "Add surname by user id")
	addBirthdayCmd := addCmd.String("birthday", "", "Add birthday by user id")

	// delete commands
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteidCmd := addCmd.String("id", "", "find user id to deleete from")
	deleteMarkCmd := deleteCmd.String("mark", "", "Delete mark by user id")
	deleteNameCmd := deleteCmd.String("name", "", "Delete name by user id")
	deleteSurnameCmd := deleteCmd.String("surname", "", "Delete surname by user id")
	deleteBirthdayCmd := deleteCmd.String("birthday", "", "Delete birthday by user id")

	if len(os.Args) < 2 {
		fmt.Printf("Expected command\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		// get command
		getHandler(getCmd, getAllCmd, getIdCmd, db)
	case "add":
		// add command
		addHandler(addCmd, addMarkCmd, addNameCmd, addSurnameCmd, addBirthdayCmd, db)
	case "delete":
		// delete command
		deleteHandler(deleteCmd, deleteidCmd, deleteMarkCmd, deleteNameCmd, deleteSurnameCmd, deleteBirthdayCmd, db)
	default:
		fmt.Printf("Expected command\n")

	}
}

func getHandler(getCmd *flag.FlagSet, all *bool, id *string, db *sql.DB) {

	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Printf("Expected id or all\n")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all == true {
		rows, err := db.Query("SELECT id, firstName, lastName, birthday, mark FROM students")
		if err == nil {
			for rows.Next() {
				var id int
				var mark string
				var name string
				var surname string
				var birthday string
				rows.Scan(&id, &name, &surname, &birthday, &mark)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%d %s %s %s %s\n", id, name, surname, birthday, mark)
			}

		}
	}

	if *id != "" {
		find := *id
		fmt.Printf("id TRUE")
		rows, err := db.Query("SELECT id, firstName, lastName, birthday, mark FROM students")
		if err == nil {
			for rows.Next() {
				var id int
				var mark string
				var name string
				var surname string
				var birthday string
				rows.Scan(&id, &name, &surname, &birthday, &mark)
				if err != nil {
					log.Fatal(err)
				}
				if find == strconv.Itoa(id) {
					fmt.Printf("%d %s %s %s %s\n", id, name, surname, birthday, mark)
				}
			}

		}
	}
}

func addHandler(addCmd *flag.FlagSet, mark *string, name *string, surname *string, birthday *string, db *sql.DB) {

	addCmd.Parse(os.Args[2:])

	if *name == "" || *surname == "" {
		fmt.Printf("Minimum requirments are first and lastname\n")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
	if *name != "" && *surname != "" {
		fmt.Printf("MARK")
		db.Query("INSERT INTO students (firstName, lastName, birthday, mark) VALUES ($1, $2, $3, $4)", *name, *surname, *birthday, *mark)
	}

}

func deleteHandler(deleteCmd *flag.FlagSet, id *string, mark *string, name *string, surname *string, birthday *string, db *sql.DB) {
	deleteCmd.Parse(os.Args[2:])

	if *name == "" || *surname == "" {
		fmt.Printf("Expected mark or name or surname or birthday\n")
		deleteCmd.PrintDefaults()
		os.Exit(1)
	}
	if *name != "" && *surname != "" {
		fmt.Printf("MARK")
		db.Query("DELETE FROM students WHERE firstName = $1 AND lastName = $2 ", *name, *surname)
	}
}
