package main

import (
	"fmt"
	"regexp"

	phonedb "github.com/bradsec/gophercises/phone/db"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "max1mus45"
	dbname   = "gophercises_phone"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	err := phonedb.Reset("postgres", psqlInfo, dbname)
	if err != nil {
		panic(err)
	}

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	err = phonedb.Migrate("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	db, err := phonedb.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Seed(); err != nil {
		panic(err)
	}

	phones, err := db.AllPhones()
	if err != nil {
		panic(err)
	}
	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating or removing...", number)
			existing, err := db.FindPhone(number)
			if err != nil {
				panic(err)
			}
			if existing != nil {
				err = db.DeletePhone(p.ID)
			} else {
				p.Number = number
				err = db.UpdatePhone(&p)
				if err != nil {
					panic(err)
				}
			}

		} else {
			fmt.Println("No changes required.")
		}
	}

}

func normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}

// func normalize(phone string) string {
// 	var buf bytes.Buffer
// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			buf.WriteRune(ch)
// 		}
// 	}
// 	return buf.String()
// }
