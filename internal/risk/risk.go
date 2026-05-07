package risk

type Level string

const (
	LevelLow    Level = "low"
	LevelMedium Level = "medium"
	LevelHigh   Level = "high"
)

type Finding struct {
	Level   Level
	Address string
	Action  string
	Message string
}
