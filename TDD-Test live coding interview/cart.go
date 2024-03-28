package cart

type Cart struct {
	LineItems     []LineItem
	CashPromotion CashPromotion
}

type CashPromotion struct {
	Name  string
	Price int
}

type LineItem struct {
	Name      string
	UnitPrice int
	Quantity  int
}

func (c *Cart) TotalPrice() int {
	totalPrice := 0
	for _, v := range c.LineItems {
		if v.Quantity == 0 {
			v.Quantity = 1
		}

		if c.CashPromotion.Name == v.Name {
			totalPrice += v.Quantity * (v.UnitPrice - c.CashPromotion.Price)
		} else {
			totalPrice += v.Quantity * v.UnitPrice
		}
	}

	return totalPrice
}

func (c *Cart) AddItems(item LineItem) {
	c.LineItems = append(c.LineItems, item)
}

func (c *Cart) AddPromotion(promotion CashPromotion) {
	c.CashPromotion = promotion
}
