package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// GeoJSON Point
type Point struct {
	Type        string    `bson:"type" json:"type"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}

// GeoJSON Polygon
type Polygon struct {
	Type        string        `bson:"type" json:"type"`
	Coordinates [][][]float64 `bson:"coordinates" json:"coordinates"`
}

// Struct vùng đất
type Land struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Location Polygon            `bson:"location" json:"location"`
	Type     string             `bson:"type" json:"type"`
	Area     float64            `bson:"area" json:"area"`
}
