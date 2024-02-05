package proto2

func (m *Timestamp) SetCreatedAt(v *int64) {
	m.CreatedAt = v
}

func (m *Timestamp) SetUpdatedAt(v *int64) {
	m.UpdatedAt = v
}

func (m *Page) SetTimestamps(v *Timestamp) {
	m.Timestamps = v
}

func (m *Page) SetTypeToText(v *string) {
	m.Type = &Page_Text{Text: v}
}

func (m *Page) SetTypeToNumber(v *int64) {
	m.Type = &Page_Number{Number: v}
}

func (m *Page) SetType(v isPage_Type) {
	m.Type = v
}
