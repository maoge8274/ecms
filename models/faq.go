package models

type Faq struct {
	Id uint32
	CategoryId uint32
	Question string
	Answer string
}

func (m *Faq) TableName() string {
	return "faqs"
}