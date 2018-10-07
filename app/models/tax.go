package models

import (
	"log"
	"time"

	"github.com/refandhika/train071018/app/database"
	"github.com/refandhika/train071018/app/modules"
)

type (
	TaxData struct {
		ID         int64
		Name       string
		TaxCode    int
		Type       string
		Refundable bool
		Price      float64
		Tax        float64
		Amount     float64
		CreatedAt  time.Time
		UpdatedAt  time.Time
		DeletedAt  *time.Time
	}
	TaxTotal struct {
		PriceTotal float64
		TaxTotal   float64
		GrandTotal float64
	}
)

func GetTaxData() ([]TaxData, TaxTotal) {
	db := database.Init()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM taxdata WHERE deleted_at IS NULL")
	if err != nil {
		log.Panic(err)
	}

	var output []TaxData
	var totals TaxTotal
	for rows.Next() {
		var row TaxData

		err = rows.Scan(&row.ID,
			&row.Name,
			&row.TaxCode,
			&row.Price,
			&row.CreatedAt,
			&row.UpdatedAt,
			&row.DeletedAt)
		if err != nil {
			log.Panic(err)
		}

		switch row.TaxCode {
		case 1:
			row.Type = "Food & Beverage"
			row.Refundable = true
			row.Tax = modules.FoodAndBeverageTax(row.Price)
		case 2:
			row.Type = "Tobacco"
			row.Refundable = false
			row.Tax = modules.TobaccoTax(row.Price)
		case 3:
			row.Type = "Entertainment"
			row.Refundable = false
			row.Tax = modules.EntertainmentTax(row.Price)
		default:
			log.Panic("Tax Code not defined!")
		}

		row.Amount = row.Price + row.Tax
		output = append(output, row)

		totals.PriceTotal += row.Price
		totals.TaxTotal += row.Tax
		totals.GrandTotal += row.Amount
	}

	return output, totals

}

func SaveTaxData(input *TaxData) ([]TaxData, TaxTotal) {
	db := database.Init()
	defer db.Close()

	stmt, err := db.Prepare("INSERT taxdata SET " +
		"name=?," +
		"code=?," +
		"price=?")
	if err != nil {
		log.Panic(err)
	}

	_, err = stmt.Exec(input.Name, input.TaxCode, input.Price)
	if err != nil {
		log.Panic(err)
	}

	return GetTaxData()
}
