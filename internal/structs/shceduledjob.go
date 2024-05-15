package structs

type ScheduledJob struct {
	Id          uint
	CronExp     string
	CronDesc    string
	BotName     string
	BotVersion  string
	TargetAgent string
	Params      string
	Uuid        string
}
