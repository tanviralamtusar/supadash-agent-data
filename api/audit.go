package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"strconv"
	"supadash/database"
)

// getProjectAuditLogs returns the last N mutative actions on a project
func (a *Api) getProjectAuditLogs(c *gin.Context) {
	_, err := a.GetAccountFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	projectRef := c.Param("ref")
	
	limitStr := c.DefaultQuery("limit", "50")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 50
	}

	// Currently the SQL query GetProjectAuditLogs does not support limit/offset
	// but we can slice the array or update the query later.
	logs, err := a.queries.GetProjectAuditLogs(c.Request.Context(), database.GetProjectAuditLogsParams{
		TargetProject: pgtype.Text{String: projectRef, Valid: true},
		Limit:         int32(limit),
		Offset:        0,
	})
	if err != nil {
		a.logger.Error("Failed to fetch audit logs: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch audit logs"})
		return
	}

	if len(logs) > limit {
		logs = logs[:limit]
	}

	c.JSON(http.StatusOK, logs)
}
