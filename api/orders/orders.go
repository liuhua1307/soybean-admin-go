package orders

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"soybean-admin-go/config"
	"soybean-admin-go/db"
	"soybean-admin-go/db/gen"
	"soybean-admin-go/db/model"
	"soybean-admin-go/utils/log"
	"strconv"
	"sync"
	"time"
)

func GetOrderList(ctx *gin.Context) {
	var (
		req          OrderQuery
		orders       = gen.Q.Order
		customerInfo = gen.Q.CustomerInfo
	)
	err := ctx.BindQuery(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid query params"})
		return
	}
	customerInfoQuery := customerInfo.WithContext(ctx)
	if req.Name != "" {
		customerInfoQuery = customerInfoQuery.Where(customerInfo.Name.Like("%" + req.Name + "%"))
	}
	if req.Phone != "" {
		customerInfoQuery = customerInfoQuery.Where(customerInfo.Phone.Like("%" + req.Phone + "%"))
	}
	if req.Address != "" {
		customerInfoQuery = customerInfoQuery.Where(customerInfo.Address.Like("%" + req.Address + "%"))
	}
	var orderIds []int64
	err = customerInfoQuery.Select(customerInfo.OrderID).Scan(&orderIds)
	if err != nil {
		config.Logger.Error("Failed to get order list", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "Failed to get order list"})
		return
	}
	find, err := orders.
		WithContext(ctx).
		Preload(orders.CustomerInfo).
		Preload(orders.Goods).
		Where(orders.ID.In(orderIds...)).
		Offset((req.Current - 1) * req.Size).
		Limit(req.Size).
		Find()
	if err != nil {
		config.Logger.Error("Failed to get order list", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "Failed to get order list"})
		return
	}
	var orderInfos []OrderInfo

	for _, order := range find {
		var itemsInfo []ItemsInfo
		err = db.DB.Table("goods").Select("goods.id as 'id', goods.name as 'name', goods.weight as 'weight', good_orders.quantity as 'quantity'").
			Joins("left join good_orders on goods.id = good_orders.good_id").
			Where("good_orders.order_id = ?", order.ID).
			Scan(&itemsInfo).Error
		if err != nil {
			config.Logger.Error("Failed to get good order", log.Field{Key: "error", Value: err})
			ctx.JSON(500, gin.H{"error": "Failed to get good order"})
			return
		}
		fmt.Println(itemsInfo)
		orderInfos = append(orderInfos, OrderInfo{
			ID:           order.ID,
			Name:         order.Name,
			Price:        order.Price,
			CreateTime:   order.CreatedAt.Unix(),
			SentOutTime:  order.SentOutTime.Unix(),
			DeliveryTime: order.DeliveryTime.Unix(),
			CustomerInfo: *order.CustomerInfo,
			Items:        itemsInfo,
		})
	}
	ctx.JSON(200, gin.H{
		"data": map[string]interface{}{
			"total":   len(find),
			"records": orderInfos,
			"size":    req.Size,
			"current": req.Current,
		},
		"code": "0000",
		"msg":  "success",
	})

}

func AddOrder(ctx *gin.Context) {
	var (
		tx                = gen.Q.Begin()
		orderInfo         OrderInfo
		modelOrder        = model.Order{}
		modelCustomerInfo = model.CustomerInfo{}
	)
	err := ctx.BindJSON(&orderInfo)
	fmt.Println(orderInfo.SentOutTime)
	if err != nil {
		config.Logger.Error("Failed to insert order", log.Field{Key: "error", Value: err})
		ctx.JSON(400, gin.H{"error": "Invalid json params"})
		return
	}

	modelOrder.Name = orderInfo.Name
	modelOrder.Price = orderInfo.Price
	modelOrder.SentOutTime = time.UnixMilli(orderInfo.SentOutTime)
	modelOrder.DeliveryTime = time.UnixMilli(orderInfo.DeliveryTime)
	modelCustomerInfo.Name = orderInfo.CustomerInfo.Name
	modelCustomerInfo.Phone = orderInfo.CustomerInfo.Phone
	modelCustomerInfo.Address = orderInfo.CustomerInfo.Address

	err = tx.Order.WithContext(ctx).Create(&modelOrder)
	if err != nil {
		tx.Rollback()
		config.Logger.Error("Failed to insert order", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "Failed to insert order"})
		return
	}
	modelCustomerInfo.OrderID = modelOrder.ID
	err = tx.CustomerInfo.WithContext(ctx).Create(&modelCustomerInfo)
	if err != nil {
		tx.Rollback()
		config.Logger.Error("Failed to insert customer info", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "Failed to insert customer info"})
		return
	}
	wg := &sync.WaitGroup{}
	once := &sync.Once{}
	errChan := make(chan error, len(orderInfo.Items))
	cancel, cancelFunc := context.WithCancel(context.Background())
	wg.Add(len(orderInfo.Items))
	for _, good := range orderInfo.Items {
		go func(good ItemsInfo) {
			defer wg.Done()
			select {
			case <-cancel.Done():
				return
			default:
				err = tx.GoodOrder.WithContext(cancel).Create(&model.GoodOrder{
					OrderID:  modelOrder.ID,
					GoodID:   good.ID,
					Quantity: good.Quantity,
				})
				if err != nil {
					once.Do(func() {
						cancelFunc()
						tx.Rollback()
					})
					errChan <- err
					return
				}
			}
		}(good)
	}
	wg.Wait()
	close(errChan)
	for err := range errChan {
		config.Logger.Error("Failed to insert good order", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "Failed to insert good order"})
		return
	}

	tx.Commit()
	ctx.JSON(200, gin.H{"code": "0000", "msg": "success"})
}

func UpdateOrder(ctx *gin.Context) {
	var (
		tx                = gen.Q.Begin()
		orderInfo         OrderInfo
		modelOrder        = model.Order{}
		modelCustomerInfo = model.CustomerInfo{}
	)
	err := ctx.BindJSON(&orderInfo)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid json params"})
		return
	}

	modelOrder.Name = orderInfo.Name
	modelOrder.Price = orderInfo.Price
	modelOrder.SentOutTime = time.Unix(orderInfo.SentOutTime, 0)
	modelOrder.DeliveryTime = time.Unix(orderInfo.DeliveryTime, 0)
	modelCustomerInfo.Name = orderInfo.CustomerInfo.Name
	modelCustomerInfo.Phone = orderInfo.CustomerInfo.Phone
	modelCustomerInfo.Address = orderInfo.CustomerInfo.Address

	_, err = tx.Order.WithContext(ctx).Where(tx.Order.ID.Eq(orderInfo.ID)).Updates(&modelOrder)
	if err != nil {
		tx.Rollback()
		config.Logger.Error("Failed to update order", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "Failed to update order"})
		return
	}
	_, err = tx.CustomerInfo.WithContext(ctx).Where(tx.CustomerInfo.OrderID.Eq(orderInfo.ID)).Updates(&modelCustomerInfo)
	if err != nil {
		tx.Rollback()
		config.Logger.Error("Failed to update customer info", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "Failed to update customer info"})
		return
	}
	tx.Commit()
	ctx.JSON(200, gin.H{"code": "0000", "msg": "success"})
}

func DeleteOrder(ctx *gin.Context) {
	var (
		tx = gen.Q.Begin()
	)
	IDI64, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	_, err = tx.Order.WithContext(ctx).Where(tx.Order.ID.Eq(IDI64)).Delete()
	if err != nil {
		tx.Rollback()
		config.Logger.Error("Failed to delete order", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "Failed to delete order"})
		return
	}
	_, err = tx.CustomerInfo.WithContext(ctx).Where(tx.CustomerInfo.OrderID.Eq(IDI64)).Delete()
	if err != nil {
		tx.Rollback()
		config.Logger.Error("Failed to delete customer info", log.Field{Key: "error", Value: err})
		ctx.JSON(500, gin.H{"error": "Failed to delete customer info"})
		return
	}
	tx.Commit()
	ctx.JSON(200, gin.H{"code": "0000", "msg": "success"})
}

type OrderQuery struct {
	Name    string `form:"name"`
	Current int    `form:"current"`
	Size    int    `form:"size"`
	Phone   string `form:"phone"`
	Address string `form:"address"`
}

type OrderInfo struct {
	ID           int64              `json:"id"`
	Name         string             `json:"name"`
	Price        float64            `json:"price"`
	CreateTime   int64              `json:"createTime"`
	SentOutTime  int64              `json:"sentOutTime"`
	DeliveryTime int64              `json:"deliveryTime"`
	CustomerInfo model.CustomerInfo `json:"customerInfo"`
	Items        []ItemsInfo        `json:"items"`
}

type ItemsInfo struct {
	ID       int64  `json:"id"`
	Quantity int32  `gorm:"column:quantity" json:"quantity"`
	Name     string `gorm:"column:name;not null" json:"name"`
	Weight   string `gorm:"column:weight;not null" json:"weight"`
}

func transferTimeStringToTime(timeString string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
