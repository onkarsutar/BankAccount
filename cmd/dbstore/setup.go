package dbstore

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/onkarsutar/BankAccount/helper/confighelper"
	"github.com/onkarsutar/BankAccount/helper/loghelper"
	"github.com/onkarsutar/BankAccount/helper/requesthelper"
	"github.com/onkarsutar/BankAccount/model"
)

func InitData() {
	GetDB()
	doneDBWrite := make(chan struct{})
	JSONPipe := make(chan model.AccountJSON, 10)
	accountPipe := make(chan model.Account, 10)

	go initAccountProcessor(JSONPipe, accountPipe)
	go dbWriter(accountPipe, doneDBWrite)
	dataGenerator(JSONPipe)
	<-doneDBWrite

}

func initAccountProcessor(JSONPipe <-chan model.AccountJSON, accountPipe chan<- model.Account) {
	var wg sync.WaitGroup
	for i := 1; i <= numOfWorkers; i++ {
		wg.Add(1)
		go processAccount(&wg, i, JSONPipe, accountPipe)
	}
	wg.Wait()
	close(accountPipe)
}

func processAccount(wg *sync.WaitGroup, id int, JSONPipe <-chan model.AccountJSON, accountPipe chan<- model.Account) {
	for accountJSON := range JSONPipe {
		// fmt.Println("worker", id, "processed  account", accountJSON.ID)

		now := time.Now()
		accountObj := model.Account{}
		accountObj.ID = accountJSON.ID
		accountObj.AccountNumber = generateAccountNumber()
		accountObj.Name = accountJSON.Name
		bal, err := strconv.ParseFloat(accountJSON.Balance, 32)
		if err != nil {
			loghelper.LogError("processAccount Error: ", err)
			return
		}
		accountObj.Balance = bal
		accountObj.CreatedOn = &now
		accountObj.ModifiedOn = &now
		accountPipe <- accountObj
	}
	wg.Done()
}

func dataGenerator(JSONPipe chan<- model.AccountJSON) {
	jsonDataURL := confighelper.GetConfig("jsonDataURL")
	accountJSONData, err := loadData(jsonDataURL)
	if err != nil {
		loghelper.LogError("dataGenerator Error: ", err)
		return
	}

	for _, accountJSONObj := range accountJSONData {
		JSONPipe <- accountJSONObj
	}

	close(JSONPipe)
}
func dbWriter(accountPipe chan model.Account, doneDBWrite chan<- struct{}) {
	store.Lock.Lock()
	defer store.Lock.Unlock()
	for accountObj := range accountPipe {
		// if store.Records == nil {
		// 	store.Records = make(map[int64]model.Account)
		// }
		store.Records[accountObj.AccountNumber] = accountObj
	}

	doneDBWrite <- struct{}{}
}

func generateAccountNumber() int64 {
	atomic.AddInt64(&accountNumberCounter, 1)
	return accountNumberCounter
}

func loadData(jsonDataURL string) ([]model.AccountJSON, error) {
	accountData := []model.AccountJSON{}

	resp, err := requesthelper.NewRequest(jsonDataURL, "GET", "application/json", "", nil, 60)

	if resp.StatusCode != 200 {
		loghelper.LogError("loadData Error: ", resp.StatusCode)
		return accountData, errors.New("JSON_URL_ERROR")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		loghelper.LogError("loadData Error: ", err)
		return accountData, err

	}
	err = json.Unmarshal(body, &accountData)
	if err != nil {
		loghelper.LogError("loadData Error: ", err)
		return accountData, err
	}

	return accountData, nil
}
