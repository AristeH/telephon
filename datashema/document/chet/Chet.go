package document

import (
	ref "Telephon/datashema/reference"
)

// Chet -  Счет за выставленные услуги  за телефона.
type Chet struct {
	Doc
	ID       uint64       `sql:"AUTO_INCREMENT" gorm:"primary_key" json:"id,omitempty"`
	Telephon ref.Telephon `json:"telephon,omitempty"`
	Zvonki   []ChetZvonki `gorm:"ForeignKey:IDchet;AssociationForeignKey:Refer" json:"zvonki,omitempty"`
	Region   ref.Region   `json:"region,omitempty"`
}
