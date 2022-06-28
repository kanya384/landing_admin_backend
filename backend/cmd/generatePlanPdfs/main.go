package main

import (
	"context"
	"fmt"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/internal/services/memcache"
	"landing_admin_backend/pkg/helpers"
	"os"
	"strconv"
	"strings"

	mng "landing_admin_backend/internal/repository/mongo"

	"github.com/go-pdf/fpdf"
	"github.com/leekchan/accounting"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.InitConfig("APP")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	file, err := os.OpenFile(cfg.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening log file config %s", err))
	}
	defer file.Close()

	logger, err := helpers.ConfigureLogger(cfg.ServiceName, cfg.LogLevel, file)
	if err != nil {
		panic(fmt.Sprintf("error initializing logger %s", err))
	}

	clientOptions := options.Client().ApplyURI(cfg.DSN)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Panic("error connection to mongo", err, nil)
	}

	cache := memcache.New()

	repository := mng.Setup(client.Database("public"))
	services := services.Setup(cfg, repository, logger, cache)

	plans, err := services.Plans.GetPlans(context.Background())
	if err != nil {
		logger.Panic("ошибка получения планировок", err, nil)
	}

	f, err := os.Open("./cmd/generatePlanPdfs/template.pdf")
	if err != nil {
		fmt.Println(err)
		logger.Panic("ошибка открытия pdf", err, nil)
	}
	defer f.Close()

	for _, plan := range plans {
		if plan.Status {
			generatePdfPlan(plan, nil)
		}
		/*if index > 30 {
			break
		}*/
	}

}

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(`bb44cf593cc9e457b4cc178f9baf9e44f78543eb37963db15ce4144e7cf3884a`)
	if err != nil {
		panic(err)
	}
}

func generatePdfPlan(plan *domain.Plan, template *model.PdfReader) (err error) {

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
	pdf.Text(435, 1520, tr(ac.FormatMoney(plan.Price)+" руб."))

	pdf.AddPage()
	pdf.Image("./cmd/generatePlanPdfs/pages/3.jpg", 0, 0, 912, 1623, false, "", 0, "")
	planImg := "./cmd/generatePlanPdfs/entrance/" + strconv.Itoa(plan.Entrance) + "/1/floor.png"
	if plan.Floor > 1 {
		planImg = "./cmd/generatePlanPdfs/entrance/" + strconv.Itoa(plan.Entrance) + "/2-9/floor.png"
	}
	pdf.Image(planImg, 50, 400, 800, 0, false, "", 0, "")

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
