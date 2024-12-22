package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func LimitByIPMiddleware(allowedIPs []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		if !isAllowedIP(clientIP, allowedIPs) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			return
		}
		c.Next()
	}
}

func isAllowedIP(ip string, allowedIPs []string) bool {
	for _, allowedIP := range allowedIPs {
		if strings.TrimSpace(ip) == allowedIP {
			return true
		}
	}
	return false
}
