package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/kozhamseitova/balance-service/internal/models"
	"github.com/stretchr/testify/require"
)

func TestReserveFunds(t *testing.T) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	numofGoRoutines := 10

	respGetBalance, err := http.Get("http://localhost:8080/api/v1/balance/1")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, respGetBalance.StatusCode)

	var getBalance models.BalanceResponse
	err = json.NewDecoder(respGetBalance.Body).Decode(&getBalance)
	require.NoError(t, err)

	amount := getBalance.Data.Balance

	fmt.Println(amount)


	for i := 0; i < numofGoRoutines; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			data := models.ReserveInput{
				UserID: 1,
				ServiceID: i+9,
				OrderID: 1,
				Amount: 200,
			}

			json_data, err := json.Marshal(data)
			require.NoError(t, err)

			resp, err := http.Post("http://localhost:8080/api/v1/balance/reserve", "application/json",
			bytes.NewReader(json_data))
			require.NoError(t, err)
			amount -= data.Amount

			fmt.Println(amount)

			defer resp.Body.Close()
	
			if(amount < 0) {	
				require.Equal(t, http.StatusInternalServerError, respGetBalance.StatusCode)
				
			}else{
				require.Equal(t, http.StatusOK, respGetBalance.StatusCode)
			}


			
		}()
	}

	wg.Wait()
}
