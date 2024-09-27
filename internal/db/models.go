package db

type TodoList struct {
	Items []ListItem
}

type ListItem struct {
	ItemID  int    `json:"item_id"`
	Content string `json:"content"`
}
