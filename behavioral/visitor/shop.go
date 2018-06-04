package visitor

import "fmt"

type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

type VisitorShop interface {
	Visit(ProductInfoRetriever)
}

type VisitableShop interface {
	Accept(VisitorShop)
}

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) Accept(v VisitorShop) {
	v.Visit(p)
}

type Rice struct {
	Product
}

type Pasta struct {
	Product
}

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NamePrinter struct {
	ProductList string
}

func (np *NamePrinter) Visit(p ProductInfoRetriever) {
	np.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), np.ProductList)
}

type Fridge struct {
	Product
}

func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 20
}

func (f *Fridge) Accept(v VisitorShop) {
	v.Visit(f)
}
