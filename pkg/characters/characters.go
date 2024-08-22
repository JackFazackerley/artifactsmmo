package characters

import "time"

type Character struct {
	// Name of the character.
	Name string `json:"name"`
	// Character skin code.
	Skin string `json:"skin"`
	// Combat level.
	Level int32 `json:"level"`
	// The current xp level of the combat level.
	Xp int32 `json:"xp"`
	// XP required to level up the character.
	MaxXp int32 `json:"max_xp"`
	// Total XP of your character.
	TotalXp int32 `json:"total_xp"`
	// The numbers of golds on this character.
	Gold int32 `json:"gold"`
	// *Not available, on the roadmap. Character movement speed.
	Speed int32 `json:"speed"`
	// Mining level.
	MiningLevel int32 `json:"mining_level"`
	// The current xp level of the Mining skill.
	MiningXp int32 `json:"mining_xp"`
	// Mining XP required to level up the skill.
	MiningMaxXp int32 `json:"mining_max_xp"`
	// Woodcutting level.
	WoodcuttingLevel int32 `json:"woodcutting_level"`
	// The current xp level of the Woodcutting skill.
	WoodcuttingXp int32 `json:"woodcutting_xp"`
	// Woodcutting XP required to level up the skill.
	WoodcuttingMaxXp int32 `json:"woodcutting_max_xp"`
	// Fishing level.
	FishingLevel int32 `json:"fishing_level"`
	// The current xp level of the Fishing skill.
	FishingXp int32 `json:"fishing_xp"`
	// Fishing XP required to level up the skill.
	FishingMaxXp int32 `json:"fishing_max_xp"`
	// Weaponcrafting level.
	WeaponcraftingLevel int32 `json:"weaponcrafting_level"`
	// The current xp level of the Weaponcrafting skill.
	WeaponcraftingXp int32 `json:"weaponcrafting_xp"`
	// Weaponcrafting XP required to level up the skill.
	WeaponcraftingMaxXp int32 `json:"weaponcrafting_max_xp"`
	// Gearcrafting level.
	GearcraftingLevel int32 `json:"gearcrafting_level"`
	// The current xp level of the Gearcrafting skill.
	GearcraftingXp int32 `json:"gearcrafting_xp"`
	// Gearcrafting XP required to level up the skill.
	GearcraftingMaxXp int32 `json:"gearcrafting_max_xp"`
	// Jewelrycrafting level.
	JewelrycraftingLevel int32 `json:"jewelrycrafting_level"`
	// The current xp level of the Jewelrycrafting skill.
	JewelrycraftingXp int32 `json:"jewelrycrafting_xp"`
	// Jewelrycrafting XP required to level up the skill.
	JewelrycraftingMaxXp int32 `json:"jewelrycrafting_max_xp"`
	// The current xp level of the Cooking skill.
	CookingLevel int32 `json:"cooking_level"`
	// Cooking XP.
	CookingXp int32 `json:"cooking_xp"`
	// Cooking XP required to level up the skill.
	CookingMaxXp int32 `json:"cooking_max_xp"`
	// Character HP.
	Hp int32 `json:"hp"`
	// *Character Haste. Increase speed attack (reduce fight cooldown)
	Haste int32 `json:"haste"`
	// *Not available, on the roadmap. Character Critical   Strike. Critical strikes increase the attack's damage.
	CriticalStrike int32 `json:"critical_strike"`
	// *Not available, on the roadmap. Regenerates life at the start of each turn.
	Stamina int32 `json:"stamina"`
	// Fire attack.
	AttackFire int32 `json:"attack_fire"`
	// Earth attack.
	AttackEarth int32 `json:"attack_earth"`
	// Water attack.
	AttackWater int32 `json:"attack_water"`
	// Air attack.
	AttackAir int32 `json:"attack_air"`
	// % Fire damage.
	DmgFire int32 `json:"dmg_fire"`
	// % Earth damage.
	DmgEarth int32 `json:"dmg_earth"`
	// % Water damage.
	DmgWater int32 `json:"dmg_water"`
	// % Air damage.
	DmgAir int32 `json:"dmg_air"`
	// % Fire resistance.
	ResFire int32 `json:"res_fire"`
	// % Earth resistance.
	ResEarth int32 `json:"res_earth"`
	// % Water resistance.
	ResWater int32 `json:"res_water"`
	// % Air resistance.
	ResAir int32 `json:"res_air"`
	// Character x coordinate.
	X int32 `json:"x"`
	// Character y coordinate.
	Y int32 `json:"y"`
	// Cooldown in seconds.
	Cooldown int32 `json:"cooldown"`
	// Datetime Cooldown expiration.
	CooldownExpiration *time.Time `json:"cooldown_expiration,omitempty"`
	// Weapon slot.
	WeaponSlot string `json:"weapon_slot"`
	// Shield slot.
	ShieldSlot string `json:"shield_slot"`
	// Helmet slot.
	HelmetSlot string `json:"helmet_slot"`
	// Body armor slot.
	BodyArmorSlot string `json:"body_armor_slot"`
	// Leg armor slot.
	LegArmorSlot string `json:"leg_armor_slot"`
	// Boots slot.
	BootsSlot string `json:"boots_slot"`
	// Ring 1 slot.
	Ring1Slot string `json:"ring1_slot"`
	// Ring 2 slot.
	Ring2Slot string `json:"ring2_slot"`
	// Amulet slot.
	AmuletSlot string `json:"amulet_slot"`
	// Artifact 1 slot.
	Artifact1Slot string `json:"artifact1_slot"`
	// Artifact 2 slot.
	Artifact2Slot string `json:"artifact2_slot"`
	// Artifact 3 slot.
	Artifact3Slot string `json:"artifact3_slot"`
	// Consumable 1 slot.
	Consumable1Slot string `json:"consumable1_slot"`
	// Consumable 1 quantity.
	Consumable1SlotQuantity int32 `json:"consumable1_slot_quantity"`
	// Consumable 2 slot.
	Consumable2Slot string `json:"consumable2_slot"`
	// Consumable 2 quantity.
	Consumable2SlotQuantity int32 `json:"consumable2_slot_quantity"`
	// Task in progress.
	Task string `json:"task"`
	// Task type.
	TaskType string `json:"task_type"`
	// Task progression.
	TaskProgress int32 `json:"task_progress"`
	// Task total objective.
	TaskTotal int32 `json:"task_total"`
	// Inventory max items.
	InventoryMaxItems int32 `json:"inventory_max_items"`
	// List of inventory slots.
	Inventory []InventorySlot `json:"inventory,omitempty"`
}
