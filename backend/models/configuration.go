package models

type Configuration struct {
	ID				uint64	`json:"id"`
	ProjectId		uint64	`json:"project_id"`
	Project			Project `json:"project"`
	Key				string	`json:"key"`
	Val				string	`json:"val"`
	ParentConfigId	uint64	`json:"parent_config_id"`
	HashedKey		string	`json:"hashed_key"`
}