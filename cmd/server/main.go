package main

import (

	"github.com/GiaBao0510/Ecommerce_golang/internal/routers"
)

func main() {
	r := routers.SetUpRouter()

	r.Run("localhost:8080") //Mặc định chạy trên localhost:8080
}