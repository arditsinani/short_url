// Redirect controller
package controllers

import (
	"context"
	"net/url"
	"short_url/app/config"
	"short_url/app/internal/db"
	"short_url/app/internal/models"
	"short_url/app/internal/services"
	"time"

	"github.com/gin-gonic/gin"
)

// Redirect Controller
type RedirCtrl struct {
	Config       *config.Config
	DB           *db.DB
	RedirService services.RedirService
}

// short url redirect method
func (ctrl *RedirCtrl) Redirect(c *gin.Context) {
	if c.Param("redirect") == "" {
		c.JSON(ctrl.Config.STATUS_CODES["BAD_REQUEST"], gin.H{
			"error": "Bad Request",
		})
		return
	}
	redir := c.Param("redirect")
	redirect, err := ctrl.RedirService.FindByRedirect(context.TODO(), redir, models.REDIRECTSCOLLECTION)
	if err != nil {
		c.JSON(ctrl.Config.STATUS_CODES["INTERNAL_SERVER_ERROR"], gin.H{
			"error": err.Error(),
		})
		return
	}
	redirect.Count = redirect.Count + 1
	redirUpdateForm := models.UpdateRedirectForm{
		Count:     redirect.Count,
		UpdatedAt: time.Now(),
	}
	_, err = ctrl.RedirService.Update(context.TODO(), redirect.ID, redirUpdateForm, models.REDIRECTSCOLLECTION)
	if err != nil {
		c.JSON(ctrl.Config.STATUS_CODES["INTERNAL_SERVER_ERROR"], gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Redirect(ctrl.Config.STATUS_CODES["MOVED_PERMANENTLY"], redirect.Url)
	c.Abort()
}

// short url read method
func (ctrl *RedirCtrl) ReadByUrl(c *gin.Context) {
	if c.Param("redirect") != "api" || c.Param("url") == "" {
		c.JSON(ctrl.Config.STATUS_CODES["BAD_REQUEST"], gin.H{
			"error": "Bad Request",
		})
		return
	}
	redir := c.Param("url")
	redirect, err := ctrl.RedirService.FindByRedirect(context.TODO(), redir, models.REDIRECTSCOLLECTION)
	if err != nil {
		c.JSON(ctrl.Config.STATUS_CODES["INTERNAL_SERVER_ERROR"], gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(ctrl.Config.STATUS_CODES["OK"], redirect)
}

// short url create method
func (ctrl *RedirCtrl) Create(c *gin.Context) {
	if c.Param("redirect") != "api" {
		c.JSON(ctrl.Config.STATUS_CODES["BAD_REQUEST"], gin.H{
			"error": "Bad Request",
		})
		return
	}
	// validate payload
	var createRedirForm models.CreateRedirectForm
	if err := c.ShouldBindJSON(&createRedirForm); err != nil {
		c.Error(err)
		c.JSON(ctrl.Config.STATUS_CODES["INTERNAL_SERVER_ERROR"], "error parsing request body")
		return
	}
	// validate url
	_, err := url.ParseRequestURI(createRedirForm.Url)
	if err != nil {
		c.JSON(ctrl.Config.STATUS_CODES["BAD_REQUEST"], gin.H{
			"error": err.Error(),
		})
		return
	}
	// check if already exists
	redirect, err := ctrl.RedirService.FindByUrl(context.TODO(), createRedirForm.Url, models.REDIRECTSCOLLECTION)
	if err == nil {
		c.JSON(ctrl.Config.STATUS_CODES["OK"], redirect)
		return
	}
	// creates new redirect attributes
	createRedirForm.CreatedAt = time.Now()
	createRedirForm.UpdatedAt = time.Now()
	createRedirForm.Count = 0
	createRedirForm.Deleted = false
	// generates random string
	createRedirForm.Redirect = ctrl.RedirService.RedirectGenerator()
	inserted, err := ctrl.RedirService.Create(context.TODO(), createRedirForm, models.REDIRECTSCOLLECTION)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(ctrl.Config.STATUS_CODES["OK"], inserted)
	return
}

// short url delete method
func (ctrl *RedirCtrl) Delete(c *gin.Context) {
	if c.Param("redirect") != "api" || c.Param("url") == "" {
		c.JSON(ctrl.Config.STATUS_CODES["BAD_REQUEST"], gin.H{
			"error": "Bad Request",
		})
		return
	}
	redir := c.Param("url")
	// check if exists
	redirect, err := ctrl.RedirService.FindByRedirect(context.TODO(), redir, models.REDIRECTSCOLLECTION)
	if err != nil {
		c.JSON(ctrl.Config.STATUS_CODES["INTERNAL_SERVER_ERROR"], gin.H{
			"error": err.Error(),
		})
		return
	}
	redirect.Deleted = true
	redirDeleteForm := models.DeleteRedirectForm{
		Deleted:   redirect.Deleted,
		UpdatedAt: time.Now(),
	}
	// updates deleted field to true
	_, err = ctrl.RedirService.Update(context.TODO(), redirect.ID, redirDeleteForm, models.REDIRECTSCOLLECTION)
	if err != nil {
		c.JSON(ctrl.Config.STATUS_CODES["INTERNAL_SERVER_ERROR"], gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(ctrl.Config.STATUS_CODES["OK"], redirect)
}
