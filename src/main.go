package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

func main() {
	initMode := flag.Bool("initmode", false, "init mode")
	ips := flag.Int("ips", 500, "inserts per second")
	count := flag.Int("count", 10000, "rows count")
	flag.Parse()

	db := sqlConnect()
	defer db.Close()

	if *initMode {
		initData(db)
	} else {
		insertData(db, *ips, *count)
	}
}

func insertData(database *sql.DB, ips int, count int) {
	fmt.Printf("The app is running with insert mode... \n Inserts per second: %d, Inserts count: %d", ips, count)

	start := time.Now()

	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan int)
	insertsCount := 0
	go func() {
		for {
			select {
			case <-ticker.C:
				startIteration := time.Now()
				for i := 0; i < ips; i++ {
					if time.Since(startIteration) > time.Second {
						fmt.Printf("\n%d rows\n", i)
						break
					} else {
						fmt.Print(".")
					}

					_, err := database.Exec("insert into user (name, birthday) values " + fakeData())
					handleError(err)

					insertsCount++
					if insertsCount >= count {
						close(quit)
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	fmt.Println(<-quit)

	elapsed := time.Since(start)
	log.Printf("Insert took %s", elapsed)
}

func initData(database *sql.DB) int {
	fmt.Println("The app is running with init mode...")

	_, err := database.Exec("CREATE TABLE IF NOT EXISTS user (id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, name varchar(32) not null, birthday date not null)")
	handleError(err)

	insertCount := 0
	limit := 40 * 1000
	iterationLimit := 1000

	for j := 0; j < limit; j++ {
		data := make([]string, iterationLimit)
		for i := range data {
			data[i] = fakeData()
		}

		_, err = database.Exec("INSERT INTO user (name, birthday) VALUES " + strings.Join(data, ", "))

		fmt.Print(".")
		insertCount += 1000
		handleError(err)

		//data = make([]string, iterationLimit)
	}
	return insertCount
}

func fakeData() string {
	return fmt.Sprintf("('%s', '%s')", randomdata.SillyName(), randomDate())
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
	db, err := sql.Open("mysql", "go_test:password@tcp(db_mysql:3306)/go_database")
	handleError(err)

	fmt.Println("DB connection successful")

	return db
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
