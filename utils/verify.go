package utils

var (
	RegisterVerify   = Rules{"Phone": {NotEmpty()}, "Password": {NotEmpty()}, "Code": {NotEmpty()}}
	LoginVerify      = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	TaskCreateVerify = Rules{"TaskTitle": {NotEmpty()}, "CreatorId": {NotEmpty()}, "ListId": {NotEmpty()}, "ProjectId": {NotEmpty()}}
	TaskListVerify   = Rules{"ListTitle": {NotEmpty()}, "ProjectId": {NotEmpty()}}
	ProjectVerify    = Rules{"ProjectName": {NotEmpty()}}
)
