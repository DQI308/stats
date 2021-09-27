package stats

import (
	"fmt"
	"github.com/DQI308/bank/v2/pkg/types"
)

func ExampleAvg() {
	payment := []types.Payment{
		{
			ID: 1,
			Amount: 10_000,
			Category: "mobile",
			Status: types.StatusOk,
		},{
			ID: 2,
			Amount: 10_000,
			Category: "mobile",
			Status: types.StatusOk,
		},
	}
	fmt.Println(Avg(payment))
	//Output:
	//10000
}


func ExampleTotalInCategory() {
	payment := []types.Payment{
		{
			ID: 1,
			Amount: 10_000,
			Category: "mobile",
			Status: types.StatusOk,
		},{
			ID: 2,
			Amount: 10_000,
			Category: "mobile",
			Status: types.StatusOk,
		},
	}
	fmt.Println(TotalInCategory(payment,"mobile"))
	//Output:
	//20000
	
}