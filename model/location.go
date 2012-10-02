package model

// Location for user or feed
type Location struct {
	Ctime int64 // location create time
	Loc   LatLon
}

type LatLon [2]float64 // [lat, lon]
