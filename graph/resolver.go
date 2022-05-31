package graph

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/rai-kargo/kargo-trucks/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Trucks    []*model.Truck
	Shipments []*model.Shipment
}

func (r *Resolver) Init() {
	for i := 0; i < 20; i++ {
		truck := &model.Truck{
			ID:      fmt.Sprintf("TRUCK-%d", len(r.Trucks)+1),
			PlateNo: fmt.Sprintf("B %d CD", len(r.Trucks)+1),
		}
		r.Trucks = append(r.Trucks, truck)
	}
}

func generateCSVData(data [][]string) (string, error) {
	filename := fmt.Sprintf("truck_%d.csv", time.Now().Unix())
	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {

		return "", errors.New(fmt.Sprintf("failed to open file %v", err))
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(data) // calls Flush internally

	if err != nil {
		return "", err
	}

	return filename, nil
}

func (r *mutationResolver) generateTruckData() [][]string {
	trucksData := make([][]string, 10)

	for _, truck := range r.Trucks {
		if len(trucksData) == 10 {
			break
		}
		trucksData = append(trucksData, []string{
			truck.ID,
			truck.PlateNo,
			fmt.Sprintf("%v", truck.Isdeleted),
		})
	}
	return trucksData
}
