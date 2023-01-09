package chicken_old

type GetChickenResult struct {
	ID            string `json:"id"`
	DateOfBirth   int    `json:"dateOfBirth"`
	EggsLaid      int    `json:"eggsLaid"`
	GoldEggsLaid  int    `json:"goldEggsLaid"`
	GoldEggChance int    `json:"goldEggChance"`
	IsAlive       bool   `json:"isAlive"`
}

type NewChickenRequest struct {
	FarmID string `json:"farmID" validate:"required"`
}

type NewChickenResult struct {
	ChickenID string `json:"chickenID"`
}
