func AuthMiddleware(cfg config.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			utils.ResponseError(c, "unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		bearer := strings.Split(token, "Bearer ")

		if len(bearer) != 2 {
			utils.ResponseError(c, "unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		token = bearer[1]

		claims, err := utils.ValidateToken(token, cfg.Secret)
		if err != nil {
			utils.ResponseError(c, "unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("id", claims.UserId)
		c.Set("role", claims.Role)
		c.Next()
	}
}