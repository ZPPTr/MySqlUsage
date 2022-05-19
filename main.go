package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

func main() {
	initMode := flag.Bool("initmode", false, "init mode")
	flag.Parse()

	db := sqlConnect()
	defer db.Close()

	if *initMode {
		fmt.Println("the app is running in initial mode")
		initData(db)
	}

}

func initData(database *sql.DB) {
	_, err := database.Exec("CREATE TABLE IF NOT EXISTS user (id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, name varchar(32) not null, birthday date not null)")
	if err != nil {
		fmt.Errorf("%v", err)
	}

	limit := 4 * 10 * 1000
	data := make([]string, limit)
	for j := 1; j < 1000; j++ {
		for i, _ := range data {
			data[i] = fmt.Sprintf("('%s', '%s')", randomdata.SillyName(), randomDate())
		}

		_, err = database.Exec("insert into user (name, birthday) values " + strings.Join(data, ", "))
		if err != nil {
			log.Fatal(err)
		}
		data = make([]string, limit)
	}

}

func randomDate() string {
	year := randomdata.Decimal(1990, 2020, 0)
	month := randomdata.Decimal(1, 12, 0)

	maxDay := 30
	if month == 2 {
		maxDay = 28
	}
	day := randomdata.Decimal(1, maxDay, 0)

	return fmt.Sprintf("%d-%d-%d", int(year), int(month), int(day))
}

func sqlConnect() (database *sql.DB) {
	db, err := sql.Open("mysql", "go_test:password@tcp(127.0.0.1:3306)/go_database")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connection successful")

	return db
}
