package pdf_generator

import (
	"context"
	"fmt"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/services/plans"
	"strconv"
	"strings"

	"github.com/go-pdf/fpdf"
	"github.com/leekchan/accounting"
)

const DEFAULT_PRICE = 15200
const SETTING_NAME = "remont_square_price"

type PdfGenerator interface {
	GeneratePdfs(ctx context.Context, price int) (err error)
}

type pdfGenerator struct {
	plansService plans.Plans
}

func NewPdfGenerator(plansService plans.Plans) PdfGenerator {
	return &pdfGenerator{
		plansService: plansService,
	}
}

func (g *pdfGenerator) GeneratePdfs(ctx context.Context, price int) (err error) {
	remontPrice := DEFAULT_PRICE
	if price != 0 {
		remontPrice = price
	}
	plans, err := g.plansService.GetActivePlans(context.Background())
	if err != nil {
		return err
	}
	for _, plan := range plans {
		if plan.Status {
			err = generatePdfPlan(plan, remontPrice)
			if err != nil {
				fmt.Println(plan.Image)
			}
		}
	}
	return nil
}

func generatePdfPlan(plan *domain.Plan, remontPrice int) (err error) {

	pdf := fpdf.NewCustom(&fpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		Size: fpdf.SizeType{
			Wd: 912,
			Ht: 1623,
		},
	})

	pdf.AddPage()
	pdf.Image("./cmd/generatePlanPdfs/pages/1.jpg", 0, 0, 912, 1623, false, "", 0, "")

	pdf.AddPage()
	pdf.Image("./cmd/generatePlanPdfs/pages/2.jpg", 0, 0, 912, 1623, false, "", 0, "")
	pdf.Image("./file_store/"+plan.Image, 0, 150, 900, 0, false, "", 0, "")
	//pdf.AddFont("Geometria", "B", "./cmd/generatePlanPdfs/font/Geometria.ttf")
	pdf.AddFont("Helvetica", "", "./cmd/generatePlanPdfs/font/helvetica_1251.json")
	tr := pdf.UnicodeTranslatorFromDescriptor("./cmd/generatePlanPdfs/font/cp1251")
	pdf.SetFont("Helvetica", "", 55)
	pdf.Text(245, 1210, tr(strings.Replace(plan.Liter, "ЖК ВЫСОКИЙ БЕРЕГ Литер", "", 1)))
	pdf.Text(280, 1265, tr("1 квартал 2023"))
	pdf.Text(275, 1315, tr(strconv.Itoa(plan.Floor)))
	pdf.Text(315, 1370, fmt.Sprintf("%.2f", plan.LivingArea))
	pdf.Text(325, 1420, fmt.Sprintf("%.2f", plan.Area))
	ac := accounting.Accounting{Symbol: "", Precision: 0, Thousand: " "}
	pdf.Text(285, 1470, tr(ac.FormatMoney(plan.Price)+" руб."))
	pdf.Text(435, 1520, tr(ac.FormatMoney(float64(plan.Price)+float64(remontPrice)*plan.Area)+" руб."))

	pdf.AddPage()
	pdf.Image("./cmd/generatePlanPdfs/pages/3.jpg", 0, 0, 912, 1623, false, "", 0, "")
	pdf.Image("./file_store/flats_on_floor/"+plan.ID+".png", 50, 250, 800, 0, false, "", 0, "")

	pdf.AddPage()
	pdf.Image("./cmd/generatePlanPdfs/pages/4.jpg", 0, 0, 912, 1623, false, "", 0, "")

	pdf.AddPage()
	pdf.Image("./cmd/generatePlanPdfs/pages/5.jpg", 0, 0, 912, 1623, false, "", 0, "")

	pdf.AddPage()
	pdf.Image("./cmd/generatePlanPdfs/pages/6.jpg", 0, 0, 912, 1623, false, "", 0, "")

	pdf.AddPage()
	pdf.Image("./cmd/generatePlanPdfs/pages/7.jpg", 0, 0, 912, 1623, false, "", 0, "")

	err = pdf.OutputFileAndClose("./file_store/pdfs/" + plan.ID + ".pdf")
	return
}
