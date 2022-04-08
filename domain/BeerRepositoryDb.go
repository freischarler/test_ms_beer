package domain

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	// registering database driver
	_ "github.com/lib/pq"

	errorsCode "github.com/freischarler/hexpattern/internal"
)

type BeerRepositoryDb struct {
	client *sql.DB
}

func (d BeerRepositoryDb) GetAll() ([]Beer, error) {
	findAllSql := "SELECT beer_id, name, brewery, country, price, currency FROM beers"
	beers := make([]Beer, 0)

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Print("Error while querying customer table" + err.Error())
		return beers, errorsCode.BadRequest
	}

	for rows.Next() {
		var c Beer
		rows.Scan(&c.Id, &c.Name, &c.Brewery, &c.Country, &c.Price, &c.Currency)

		if err != nil {
			log.Print("Error while scanning beer" + err.Error())
			return beers, errorsCode.BadRequest
		}
		beers = append(beers, c)
	}
	return beers, nil
}

func (d BeerRepositoryDb) PostOne(b Beer) error {
	sqlInsert := "INSERT INTO beers (beer_id, name, brewery, country, price, currency) VALUES ($1,$2,$3,$4,$5,$6)"

	_, error := d.client.Exec(sqlInsert, b.Id, b.Name, b.Brewery, b.Country, b.Price, b.Currency)
	if error != nil {
		//err.Error("El ID de la cerveza ya existe")
		return errorsCode.ID
	}
	return nil
}

func (d BeerRepositoryDb) GetOneByID(id string) (Beer, error) {
	beerSql := "SELECT beer_id, name, brewery, country, price, currency FROM beers WHERE beer_id = $1"
	log.Println(beerSql)
	var beer Beer

	beer_id, err := strconv.Atoi(id)
	if err != nil {
		log.Print("Error while querying customer table" + err.Error())
		//return beer, err
	}

	err = d.client.QueryRow(beerSql, &beer_id).Scan(&beer.Id, &beer.Name, &beer.Brewery, &beer.Country, &beer.Price, &beer.Currency)
	if err != nil {
		return beer, errorsCode.ID
	}
	return beer, nil
}

func (d BeerRepositoryDb) GetBoxPrice(id string, currency string, count int) (float64, error) {
	var beer Beer

	beer, err := d.GetOneByID(id)
	if err != nil {
		return 0, errorsCode.ID
	}

	factor, err := getPrice(beer.Currency, currency)
	if err != nil {
		return 0, errorsCode.BadRequest
	}

	price := beer.Price * float64(count) * factor
	fmt.Println(price)
	return price, nil
}

func NewBeerRepositoryDb() BeerRepositoryDb {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDBNAME"))
	log.Print(psqlInfo)

	client, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxIdleTime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxOpenConns(10)

	err = MakeMigration(client)
	if err != nil {
		log.Panic(err)
	}

	return BeerRepositoryDb{client}
}

func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}

type Data struct {
	Host     string `json:"pgHost"`
	Port     string `json:"pgPort"`
	User     string `json:"pgUser"`
	Password string `json:"pgPassword"`
	DBName   string `json:"pgDBname"`
}

func getPrice(currency_source string, currency_destiny string) (float64, error) {
	currencyApiKey := os.Getenv("CURRENCY_LAYER_KEY")
	if currencyApiKey == "" {
		log.Fatal("Error: Can't get CURRENCY ENV")
	}

	//No valid in free subscription
	/*r, err := http.Get("http://api.currencylayer.com/convert?access_key=" + currencyApiKey + "&from=" + currency_source + "&to=" + currency_destine + "&amount=" + fmt.Sprintf("%v", amount))
	if err != nil {
		log.Fatal(err)
	}*/

	r, err := http.Get("http://api.currencylayer.com/live?access_key=" + currencyApiKey)
	if err != nil {
		return 0, errorsCode.NotFound
	}

	var d interface{}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return 0, errorsCode.Decode
	}

	defer recoverWeb()

	factor1 := (d.(map[string]interface{})["quotes"].(map[string]interface{})[("USD" + currency_source)]).(float64)

	factor2 := (d.(map[string]interface{})["quotes"].(map[string]interface{})[("USD" + currency_destiny)]).(float64)

	value := 1 / factor1 * factor2

	return value, nil
}

func recoverWeb() {
	if r := recover(); r != nil {
		fmt.Println("recovered from bad currency request", r)
	}
}
