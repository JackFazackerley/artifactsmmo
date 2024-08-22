package characters

type InventorySlot struct {
	// Inventory slot identifier.
	Slot int32 `json:"slot"`
	// Item code.
	Code string `json:"code"`
	// Quantity in the slot.
	Quantity int32 `json:"quantity"`
}
