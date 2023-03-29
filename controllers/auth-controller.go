package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type AuthController struct {
    merkleTree *MerkleTree
}

func NewAuthController(merkleTree *MerkleTree) *AuthController {
    return &AuthController{merkleTree}
}

func (c *AuthController) Login(ctx *gin.Context) {
    // Check username and password
    username := ctx.PostForm("username")
    password := ctx.PostForm("password")
    if !checkCredentials(username, password) {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    // Generate and return token
    token := generateToken(username, c.merkleTree.Root())
    ctx.JSON(http.StatusOK, gin.H{"token": token})
}
