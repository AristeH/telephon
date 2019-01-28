package document



// Chet -  Счет за выставленные услуги  за телефона.
type Chet struct {
	Doc string
	ID       string       `sql:"AUTO_INCREMENT" gorm:"primary_key" json:"id,omitempty"`
	IDTelephon string `json:"telephon,omitempty"`
	IDZvonki   []string `gorm:"ForeignKey:IDchet;AssociationForeignKey:Refer" json:"zvonki,omitempty"`
	IDRegion   string   `json:"region,omitempty"`
}
