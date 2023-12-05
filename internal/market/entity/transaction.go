package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

func NewTransaction(sellingOrder *Order, buingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buingOrder,
		Shares:       shares,
		Price:        price,
		Total:        total,
		DateTime:     time.Now(),
	}
}

func (t *Transaction) AddBuyOrderPendingShares(id string, shares int) {
	t.BuyingOrder.Investor.UpdateAssetPosition(id, shares)
	t.BuyingOrder.PedingShares -= shares
}

func (t *Transaction) AddSellOrderPendingShares(id string, shares int) {
	t.SellingOrder.Investor.UpdateAssetPosition(id, -shares)
	t.SellingOrder.PedingShares -= shares
}

func (t *Transaction) CalculateTotal(shares int, price float64) {
	t.Total = float64(t.Shares) * t.Price
}

func (t *Transaction) CloseBuyOrder() {
	if t.BuyingOrder.PedingShares == 0 {
		t.BuyingOrder.Status = "CLOSED"
	}
}
func (t *Transaction) CloseSellOrder() {
	if t.SellingOrder.PedingShares == 0 {
		t.SellingOrder.Status = "CLOSED"
	}
}
