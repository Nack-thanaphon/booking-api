package models

import (
    "encoding/json"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Appointment struct {
    ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    PatientID string             `json:"patient_id" bson:"patient_id"`
    DoctorID  string             `json:"doctor_id" bson:"doctor_id"`
    Date      time.Time          `json:"date" bson:"date"`
    CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

func ToJSON(a Appointment) ([]byte, error) {
    return json.Marshal(a)
}