package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetCommunities list all communities
func GetCommunities(c echo.Context) error {
	return c.String(http.StatusOK, "LISTING COMMUNITIES")
}

// GetCommunity returns a single community
func GetCommunity(c echo.Context) error {
	return c.String(http.StatusOK, "GETTING COMMUNITIES")
}

// CreateCommunity returns a single community
func CreateCommunity(c echo.Context) error {
	return c.String(http.StatusOK, "CREATING COMMUNITIES")
}

// UpdateCommunity returns a single community
func UpdateCommunity(c echo.Context) error {
	return c.String(http.StatusOK, "UPDATING COMMUNITIES")
}

// DeleteCommunity returns a single community
func DeleteCommunity(c echo.Context) error {
	return c.String(http.StatusOK, "DELETING COMMUNITIES")
}
