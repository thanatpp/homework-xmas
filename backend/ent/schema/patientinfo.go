package schema

import "github.com/facebookincubator/ent"

// PatientInfo holds the schema definition for the PatientInfo entity.
type PatientInfo struct {
	ent.Schema
}

// Fields of the PatientInfo.
func (PatientInfo) Fields() []ent.Field {
	return nil
}

// Edges of the PatientInfo.
func (PatientInfo) Edges() []ent.Edge {
	return nil
}
