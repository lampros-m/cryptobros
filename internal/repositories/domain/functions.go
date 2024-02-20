package domain

import (
	"encoding/json"
	"errors"
	"repositories/cryptobros/internal/helpers"
	"time"
)

var (
	DateFormatFileName = "20060102"
	DBPath             = "db/"
	ResultsPath        = "results/"
	CoinsFileName      = "coins.json"
	ResultsFileName    = "results.json"
)

func SaveCoins(coins Coins) error {
	timeStamp := time.Now().Format(DateFormatFileName)
	pathName := DBPath + timeStamp + CoinsFileName
	helpers.WriteFileJSON(coins, pathName)

	return nil
}

func SaveReults(coins Coins) error {
	pathName := ResultsPath + ResultsFileName
	helpers.WriteFileJSON(coins, pathName)

	return nil
}

func GetCoinsForSpecificDay(date string) (Coins, error) {
	_, err := time.Parse(DateFormatFileName, date)
	if err != nil {
		return Coins{}, errors.New("error GetCoinsForSpecificDay date has not the right format: " + err.Error())
	}

	path := DBPath + date + CoinsFileName

	bytes, err := helpers.ReadFile(path)
	if err != nil {
		return Coins{}, errors.New("error GetCoinsForSpecificDay cannot ReadFile: " + err.Error())
	}

	var coins Coins
	err = json.Unmarshal(bytes, &coins)
	if err != nil {
		return Coins{}, errors.New("error GetCoinsForSpecificDay cannot Unmarshal: " + err.Error())
	}

	return coins, nil
}
