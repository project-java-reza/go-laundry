package model

//type Product struct {
//	Id    string
//	Name  string
//	Price int
//	UomId string
//}

type Product struct {
	Id    string
	Name  string
	Price int
	Uom   Uom // supaya dapat uom beserta isi yang lain tidak hanya id
}

// Uom di Emmbeddable dari struct untuk mempermudah ketika select
// SELECT p.id, p.name, p.price, u.id, u.name FROM product p
// JOIN uom u ON p.uom_id = u.id
