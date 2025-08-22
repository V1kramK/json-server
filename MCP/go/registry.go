package main

import (
	"github.com/json-server-api/mcp-server/config"
	"github.com/json-server-api/mcp-server/models"
	tools_name "github.com/json-server-api/mcp-server/tools/name"
	tools_general "github.com/json-server-api/mcp-server/tools/general"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_name.CreateGet_nameTool(cfg),
		tools_name.CreatePatch_nameTool(cfg),
		tools_name.CreatePost_nameTool(cfg),
		tools_name.CreatePut_nameTool(cfg),
		tools_name.CreateDelete_name_idTool(cfg),
		tools_name.CreateGet_name_idTool(cfg),
		tools_name.CreatePatch_name_idTool(cfg),
		tools_name.CreatePut_name_idTool(cfg),
		tools_general.CreateGetTool(cfg),
	}
}
