package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hackthenorth2020/go-firebase/items"
)

func createItem(c *gin.Context) {
	var createItemRequest *items.Item
	err := c.Bind(&createItemRequest)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	resp, err := itemSrv.CreateItem(createItemRequest)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	c.JSON(200, &resp)
}

func readItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	resp, err := itemSrv.ReadItem(uint(id))
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	c.JSON(200, &resp)
}

func updateItem(c *gin.Context) {
	var updateItemRequest *items.Item
	err := c.Bind(&updateItemRequest)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	resp, err := itemSrv.UpdateItem(updateItemRequest)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	c.JSON(200, &resp)
}

func deleteItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	resp, err := itemSrv.DeleteItem(uint(id))
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	c.JSON(200, &resp)
}
