package domain

import (
	"strconv"
	"strings"
	"time"
)

type Plan struct {
	ID                  string    `bson:"_id"`
	Image               string    `bson:"image" json:"image"`
	Entrance            int       `bson:"entrance" json:"entrance"`
	Commerce            bool      `bson:"commerce" json:"commerce"`
	Floor               int       `bson:"floor" json:"floor"`
	Number              int       `bson:"number" json:"number"`
	Rooms               int       `bson:"rooms" json:"rooms"`
	Area                float64   `bson:"area" json:"area"`
	LivingArea          float64   `bson:"living_area" json:"living_area"`
	KitchenArea         float64   `bson:"kitchen_area" json:"kitchen_area"`
	Price               int       `bson:"price" json:"price"`
	DiscountPrice       int       `bson:"discount_price" json:"discount_price"`
	SquarePrice         int       `bson:"square_price" json:"square_price"`
	SquareDiscountPrice int       `bson:"square_discount_price" json:"square_discount_price"`
	Discount            float64   `bson:"discount" json:"discount"`
	Status              bool      `bson:"status" json:"status"`
	WindowView          string    `bson:"window_view" json:"window_view"`
	Liter               string    `bson:"liter" json:"liter"`
	CreatedAt           time.Time `bson:"created_at" json:"created_at"`
	UpdateAt            time.Time `bson:"updated_at" json:"updated_at"`
}

func NewPlanFromCsvField(fieldItems []string) (plan Plan, err error) {
	plan = Plan{
		Image:    "",
		Commerce: fieldItems[2] == "Да",
		Status:   fieldItems[14] == "Свободно",
	}
	plan.ID = fieldItems[0]
	plan.Entrance, _ = strconv.Atoi(fieldItems[1])
	plan.Floor, _ = strconv.Atoi(fieldItems[3])
	plan.Number, _ = strconv.Atoi(fieldItems[4])
	plan.Rooms, _ = strconv.Atoi(fieldItems[5][0:1])
	plan.Area, _ = strconv.ParseFloat(strings.Replace(fieldItems[6], ",", ".", 1), 64)
	plan.LivingArea, _ = strconv.ParseFloat(strings.Replace(fieldItems[7], ",", ".", 1), 64)
	plan.KitchenArea, _ = strconv.ParseFloat(strings.Replace(fieldItems[8], ",", ".", 1), 64)
	plan.Price, _ = strconv.Atoi(fieldItems[9])
	plan.DiscountPrice, _ = strconv.Atoi(fieldItems[10])
	plan.SquarePrice, _ = strconv.Atoi(fieldItems[11])
	plan.SquareDiscountPrice, _ = strconv.Atoi(fieldItems[12])
	plan.Discount, _ = strconv.ParseFloat(strings.Replace(fieldItems[13], ",", ".", 1), 64)
	plan.WindowView = fieldItems[15]
	plan.Liter = fieldItems[16]
	if err != nil {
		return
	}

	return
}
