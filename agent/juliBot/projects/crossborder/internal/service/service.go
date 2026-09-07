package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"

	"github.com/julien/juliBot/projects/crossborder/internal/domain"
)

var (
	ErrNotFound            = errors.New("business resource not found")
	ErrInvalidTransition   = errors.New("invalid business state transition")
	ErrInsufficientStock   = errors.New("insufficient inventory")
	ErrIdempotencyKey      = errors.New("idempotency_key is required")
	ErrIdempotencyConflict = errors.New("idempotency key was already used for a different request")
)

// 仓库调货请求
type TransferRequest struct {
	SKU            string // 调用商品号
	FormWareHouse  string // 调用商品的库房地址
	IdempotencyKey string // 幂等键
	Quantity       int
	DtyRun         bool // 测试时用
	ToWareHouse    string
}

type Service struct {
	mu          sync.RWMutex
	orders      map[string]domain.Order
	inventory   map[string]domain.InventoryBalance
	transfers   map[string]domain.InventoryTransfer
	idempotency map[string]idempotencyRecord
}

type idempotencyRecord struct {
	fingerprint [32]byte
	transfer  domain.InventoryTransfer
}

func NewSeeded() *Service {
	return &Service{
		orders: map[string]domain.Order{
			"TTS-20260801-1001": {
				ID: "TTS-20260801-1001", Market: "US", Currency: "USD",
				Amount: 129.99, Status: domain.OrderAwaitingShipment,
				FulfillmentWH: "WH-CN-SZ", Cancellation: true,
				Items: []domain.OrderItem{{SKU: "SKU-BLACK-M-01", Quantity: 1, Price: 129.99}},
			},
		},
		inventory: map[string]domain.InventoryBalance{
			inventoryKey("WH-CN-SZ", "SKU-BLACK-M-01"): {
				WarehouseID: "WH-CN-SZ", SKU: "SKU-BLACK-M-01", Available: 0,
			},
			inventoryKey("WH-US-LAX", "SKU-BLACK-M-01"): {
				WarehouseID: "WH-US-LAX", SKU: "SKU-BLACK-M-01", Available: 18,
			},
		},
		transfers:   make(map[string]domain.InventoryTransfer),
		idempotency: make(map[string]domain.InventoryTransfer),
	}
}

func (s *Service) GetOrder(id string) (domain.Order, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	order, ok := s.orders[id]
	return order, ok
}

func (s *Service) Inventory(sku string) []domain.InventoryBalance {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]domain.InventoryBalance, 0, len(s.idempotency))
	for _, balance := range s.inventory {
		if balance.SKU == sku {
			result = append(result, balance)
		}
	}
	return result

}

// 交给大模型告诉大模型拥有仓库调拨的能力，交给大模型调用和决策
// 网络抖动会调用多次，如何处理这个幂等问题？？？
func (s *Service) CreateTransfer(req TransferRequest) (domain.InventoryTransfer, error) {
	// 注意如何在业务侧规避大模型重复调用问题

	// 参数校验
	if req.IdempotencyKey == "" {
		return domain.InventoryTransfer{}, ErrIdempotencyKey
	}
	if req.Quantity <= 0 || req.FormWareHouse == req.ToWareHouse {
		return domain.InventoryTransfer{}, ErrInvalidTransition
	}

	// 幂等校验
	s.mu.Lock()
	defer s.mu.Unlock()
	//   请求的指纹
	fingerprint := transferFingerprint(req)
	if cached, ok := s.idempotency[req.IdempotencyKey];ok {
		if cached.fingerprint != fingerprint {

		}
	}
	// 库存校验

	// 创建调库存的单子
	return domain.InventoryTransfer{}, ErrNotImplemented
}

func inventoryKey(warehouse, sku string) string {
	return warehouse + "|" + sku
}

func transferFingerprint(req TransferRequest) [32]byte {
	return sha256.Sum256([]byte(fmt.Sprintf("%s\x00%s\x00%s\x00%d\x00%t", req.SKU, req.FormWareHouse, req.ToWareHouse, req.Quantity, req.DtyRun)))
}