package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
)

type Kpi struct {
	Id              string
	Name            string
	CategoryId      uuid.UUID
	CreatedBy       string
	IsEnabled       bool
	LogicalDelete   bool
	MeasurementType int64
}

func readFile() []Kpi {
	file, err := os.Open("kpis.json")
	var kpis []Kpi
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		decoder.Decode(&kpis)
	}
	return kpis
}

func GetEnumString(num int64) string {
	switch num {
	case 0:
		return "MeasurementType.Numeric"
	case 1:
		return "MeasurementType.Currency"
	case 2:
		return "MeasurementType.Percentage"
	case 3:
		return "MeasurementType.Time"
	default:
		return ""
	}
}

func getCategoryName(id string) string {

	switch id {
	case "1a506c01-b89d-439e-8439-373b347b30bb":
		return "Supply Chain"
	case "1b4ee5dc-d835-492f-841b-94aaaef09e8a":
		return "Azure Devops"
	case "1dfde185-4f91-451e-aa84-a124ef535592":
		return "Customer Support"
	case "39be1c1c-f955-4ac9-ba01-127820312a26":
		return "IT"
	case "4afe93c1-ba6a-45ea-b556-ce3e9f88bf1d":
		return "Procurement"
	case "81f76aea-7c79-4600-b799-94bcd8696853":
		return "Inventory Management"
	case "8fd4d9cc-cb8d-4193-a960-f2227d25ee0a":
		return "HR"
	case "9bcf7c9f-9c7d-477f-8825-1a1af244c41d":
		return "Operations"
	case "a9d1000d-9da4-4fc9-9d59-f915918a9185":
		return "Finance"
	case "df68b3df-2285-4d62-afb4-284fef5d818a":
		return "Sales"
	case "e9cd3afa-0c8d-4da6-95b7-e78d6b050f67":
		return "Professional Services"
	case "f04a727b-d210-447a-9e70-b68f15765037":
		return "Marketing"
	default:
		return ""
	}
}

func main() {
	kpis := readFile()
	file, err := os.OpenFile("seedData.cs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		defer file.Close()
		file.WriteString("{")
		for _, kpi := range kpis {
			dataString := fmt.Sprintf("new KPI{"+
				"Id = Guid.Parse(\"%v\"),\n"+
				"Name = \"%v\",\n"+
				"CategoryId = Guid.Parse(\"%v\"),\n"+
				"CreatedBy = \"%v\",\n"+
				"IsEnabled = %v,\n"+
				"LogicalDelete = %v,\n"+
				"MeasurementType = %v,\n"+
				"CreatedDate = new DateTime(%v, %v, %v),"+
				"CategoryName = \"%v\"},",

				uuid.NewString(), kpi.Name, kpi.CategoryId, "System", kpi.IsEnabled, kpi.LogicalDelete, GetEnumString(kpi.MeasurementType), 2024, 10, 3, getCategoryName(kpi.CategoryId.String()))
			file.WriteString(dataString)
		}
		file.WriteString("}")
	}

}
