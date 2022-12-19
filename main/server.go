package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ServerHost string
var ServerPort string
var RedisHost string
var RedisPort string
var RedisDB string
var KeyExp int
var CoinApi string

type MyJSONSerializer struct {
	echo.DefaultJSONSerializer
}

func (d MyJSONSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	c.Response().Header().Set(echo.HeaderContentType, "application/vnd.siren+json")
	return d.DefaultJSONSerializer.Serialize(c, i, indent)
}

func getEnv(key string) string {
	value := os.Getenv(key)
	log.Printf("- [ENV] %s:%s\n", key, value)

	return value
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file\n")
	}

	ServerHost = getEnv("HOST")
	ServerPort = getEnv("SERVER_PORT")
	RedisHost = getEnv("REDIS_HOST")
	RedisPort = getEnv("REDIS_PORT")
	RedisDB = getEnv("REDIS_DB")

	s := getEnv("KEY_EX")

	ct, err := strconv.Atoi(s)
	if err != nil {
		log.Printf(" - [FATAL] %v\n", err)
		ct = 5
	}
	KeyExp = ct * 60

	CoinApi = getEnv("API_KEY")
}

func main() {
	loadEnv()
	initRedis()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	e.POST("/get", getPrice)

	e.Logger.Fatal(e.Start(ServerHost + ":" + ServerPort))
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"Body": "Hello, World!",
		"Date": time.Now().String(),
	})
}

func getPrice(c echo.Context) error {
	coinName := c.QueryParam("name")
	var currentRate string

	val, err := get(coinName)
	if err != nil {
		coin, rate, _ := getExchangeRate(coinName)

		err := set(coin, fmt.Sprintf("%v", rate))
		if err != nil {
			log.Printf(" - [FATAL] Fail set %v\n", err)
		}

		currentRate = rate
		log.Println("\n - [INFO] Exchange rate by https://www.coinapi.io/")
	} else {
		currentRate = val
		log.Println("\n - [INFO] Exchange rate by Cache")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Name": coinName,
		"Rate": fmt.Sprintf("%v", currentRate),
		"Date": time.Now().String(),
	})
}

func getExchangeRate(name string) (string, string, error) {
	type ExchangeRate struct {
		Time         string  `json:"time"`
		AssetIdBase  string  `json:"asset_id_base"`
		AssetIdQuote string  `json:"asset_id_quote"`
		Rate         float64 `json:"rate"`
	}

	client := resty.New()

	rates := ExchangeRate{}
	_, err := client.R().
		SetHeader("X-CoinAPI-Key", CoinApi).
		SetResult(&rates).
		Get("https://rest.coinapi.io/v1/exchangerate/" + name + "/USD")

	if err != nil {
		log.Printf(" - [ERROR] Fail to get %s exchange rate, Error %s\n", name, err.Error())
		return "", "", err
	}
	return rates.AssetIdBase, fmt.Sprintf("%v", rates.Rate), err
}
