package main

import (
	"FP-Sanbercode-Go-Batch-41/router"
	"os"
)

func main() {

	router.StartServer().Run(":" + os.Getenv("PORT"))
}
