package characters

type CharacterFightDataSchema struct {
	// Cooldown details.
	Cooldown CooldownSchema `json:"cooldown"`
	// Fight details.
	Fight FightSchema `json:"fight"`
	// Player details.
	Character Character `json:"character"`
}

type FightSchema struct {
	// The amount of xp gained by the fight.
	Xp int32 `json:"xp"`
	// The amount of gold gained by the fight.
	Gold int32 `json:"gold"`
	// The items dropped by the fight.
	Drops []DropSchema `json:"drops"`
	// Numbers of the turns of the combat.
	Turns int32 `json:"turns"`
	// The amount of blocked hits by the monster.
	MonsterBlockedHits BlockedHitsSchema `json:"monster_blocked_hits"`
	// The amount of blocked hits by the player.
	PlayerBlockedHits BlockedHitsSchema `json:"player_blocked_hits"`
	// The fight logs.
	Logs []string `json:"logs"`
	// The result of the fight.
	Result string `json:"result"`
}

type DropSchema struct {
	// The code of the item.
	Code string `json:"code"`
	// The quantity of the item.
	Quantity int32 `json:"quantity"`
}

type BlockedHitsSchema struct {
	// The amount of fire hits blocked.
	Fire int32 `json:"fire"`
	// The amount of earth hits blocked.
	Earth int32 `json:"earth"`
	// The amount of water hits blocked.
	Water int32 `json:"water"`
	// The amount of air hits blocked.
	Air int32 `json:"air"`
	// The amount of total hits blocked.
	Total int32 `json:"total"`
}
