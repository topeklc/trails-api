package routes



func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/")
	{
		
	}
	return router
}