package order

import "time"

type Order struct {
	OrderID              int64     `json:"order_id"`
	PackageID            int       `json:"package_id"`
	TransporterID        int       `json:"transporter_id"`
	DistributorID        int       `json:"distributor_id"`
	OrderPrice           float64   `json:"order_price"`
	ModificationDate     time.Time `json:"modification_date"`
	EstimatedArrivalDate time.Time `json:"est_arrival_date"`
	OrderDate            time.Time `json:"order_date"`
	OrderStatus          string    `json:"order_status"`
}

type Tracking struct {
	TrackingID       int       `json:"tracking_id"`
	OrderID          int64     `json:"order_id"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	ModificationDate time.Time `json:"modification_date"`
}
