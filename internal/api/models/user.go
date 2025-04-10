package orm_models

type ControlVersion struct {
	ID           uint   `gorm:"primaryKey;autoIncrement;column:id"`
	IssueId      string `gorm:"column:issue_id"`
	RepoOwner    string `gorm:"column:repo_owner"`
	RepoName     string `gorm:"column:repo_name"`
	TargetBranch string `gorm:"column:target_branch"`
	PrID         string `gorm:"column:pr_id"`
	Title        string `gorm:"column:title"`
	Description  string `gorm:"column:description"`
	Version      string `gorm:"column:version"`
	AppLayer     string `gorm:"column:app_layer"`
	IssueType    string `gorm:"column:issue_type"`
	RawMessage   string `gorm:"column:raw_message"`
	MergedAt     string `gorm:"column:merged_at"`
}

func (ControlVersion) TableName() string {
	return "amsys.control_versions"
}
