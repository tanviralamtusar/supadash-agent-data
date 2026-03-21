package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"supadash/database"
	"supadash/provisioner"
)

type RotateSecretRequest struct {
	Type string `json:"type" binding:"required"` // "jwt" or "database"
}

// rotateProjectSecret rotates the JWT secret or Database password for a project
func (a *Api) rotateProjectSecret(c *gin.Context) {
	_, err := a.GetAccountFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req RotateSecretRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload: " + err.Error()})
		return
	}

	projectRef := c.Param("ref")
	project, err := a.queries.GetProjectByRef(c.Request.Context(), projectRef)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Add audit log
	// We should ideally extract AccountID to put into ActorID.
	// We will skip inserting ActorID for now if not available, or we can parse from token.

	switch req.Type {
	case "jwt":
		newJwtSecret, _ := provisioner.GenerateRandomString(32) // Generate a strong random JWT secret
		
		err = a.queries.UpdateProjectJwtSecret(c.Request.Context(), database.UpdateProjectJwtSecretParams{
			JwtSecret:  newJwtSecret,
			ProjectRef: projectRef,
		})
		if err != nil {
			a.logger.Error("Failed to update JWT Secret: " + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rotate JWT secret"})
			return
		}

	case "database":
		// Database password rotation means we must change the password in the database 
		// as well as in the configuration. 
		// Actually, in SupaDash, Postgres runs in a Docker container with POSTGRES_PASSWORD environment variable.
		// To rotate it, we update the env var and recreate the container.
		// This requires integration with Provisioner.
		c.JSON(http.StatusNotImplemented, gin.H{"error": "Database password rotation is not yet implemented"})
		return
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown secret type to rotate"})
		return
	}

	// Trigger async restart for changes to apply
	if a.provisioner != nil {
		go func() {
			ctx := context.Background()
			projectIDStr := strconv.Itoa(int(project.ID))
			a.provisioner.PauseProject(ctx, projectIDStr)
			a.provisioner.ResumeProject(ctx, projectIDStr)
		}()
	}

	c.JSON(http.StatusOK, gin.H{"message": "Secret rotated successfully. Project will restart to apply changes."})
}
