package model

import (
	"fmt"
	"strconv"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Item struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Unit  string `json:"Unit"`
	Size  string `json:"Size"`
}
type Order struct {
	Id       string        `json:"id"`
	Buyer    string        `json:"buyer"`
	Date     time.Time     `json:"date"`
	Discount int           `json:"discount"`
	Price    int           `json:"price"`
	Status1  string        `json:"status1"`
	Status2  string        `json:"status2"`
	ItemList []ItemInOrder `json:"itemList"`
}
type ItemInOrder struct {
	Item_id string `json:"item_id"`
	Amount  int    `json:"amount"`
}

func MapToItem(m_item map[string]interface{}, id string) Item {
	var item Item
	item.Id = id
	item.Name = fmt.Sprintf("%v", m_item["name"])
	item.Price, _ = strconv.Atoi(fmt.Sprintf("%v", m_item["price"]))
	item.Size = fmt.Sprintf("%v", m_item["size"])
	item.Unit = fmt.Sprintf("%v", m_item["unit"])
	return item
}
func ItemToMap(item Item) (map[string]interface{}, string) {
	map_return := map[string]interface{}{
		"name":  item.Name,
		"price": item.Price,
		"unit":  item.Unit,
		"size":  item.Size,
	}
	return map_return, item.Id
}
func MapToOrder(m_order map[string]interface{}, id string) Order {
	var order Order
	var timeError bool
	order.Id = id
	order.Buyer = fmt.Sprintf("%v", m_order["buyer"])
	order.Date, timeError = m_order["date"].(time.Time)
	if !timeError {
		fmt.Print("Bug at convert date time")
	}
	order.Price, _ = strconv.Atoi(fmt.Sprintf("%v", m_order["price"]))
	order.Discount, _ = strconv.Atoi(fmt.Sprintf("%v", m_order["discount"]))
	itemListMap, _ := m_order["itemlist"].([]interface{})
	length := len(itemListMap)
	var itemList []ItemInOrder
	for i := 0; i < length; i++ {
		var item ItemInOrder
		item_map, _ := itemListMap[i].(map[string]interface{})
		item.Item_id = fmt.Sprintf("%v", item_map["item_id"])
		item.Amount, _ = strconv.Atoi(fmt.Sprintf("%v", item_map["amount"]))
		itemList = append(itemList, item)
	}
	order.ItemList = itemList
	order.Status1 = fmt.Sprintf("%v", m_order["status1"])
	order.Status2 = fmt.Sprintf("%v", m_order["status2"])

	// Type assertion to convert interface{} to time.Time
	return order
}
func OrderToMap(order Order) (map[string]interface{}, string) {

}
