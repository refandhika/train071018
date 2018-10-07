package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/refandhika/train071018/app/models"
)

type (
	TaxController struct{}
	Data          struct {
		Output []models.TaxData
		Totals models.TaxTotal
	}
)

func NewTaxController() *TaxController {
	return &TaxController{}
}

func (taxController TaxController) TaxShow(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data Data
	data.Output, data.Totals = models.GetTaxData()

	t, _ := template.ParseFiles("templates/tax.html")
	t.Execute(w, data)

}

func (taxController TaxController) TaxSubmit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	var form models.TaxData
	form.Name = r.Form["name"][0]
	form.TaxCode, _ = strconv.Atoi(r.Form["code"][0])
	form.Price, _ = strconv.ParseFloat(r.Form["price"][0], 64)

	var data Data
	data.Output, data.Totals = models.SaveTaxData(&form)

	t, _ := template.ParseFiles("templates/tax.html")
	t.Execute(w, data)
}
