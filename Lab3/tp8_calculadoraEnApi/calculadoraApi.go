package tp8_calculadoraApi

import (
	"github.com/gin-gonic/gin"
	"github.com/rodri/Lab3-master/Lab3/tp8_calculadoraEnApi/controller"
)

func CalculadoraApi() {
	r := gin.Default()
	r.GET("/calc/sum", controller.Sumar)       // para usar esta funcion :http://localhost:8080/calc/sum?a=Numero&b=Numero
	r.GET("/calc/res", controller.Restar)      // para usar esta funcion :http://localhost:8080/calc/res?a=Numero&b=Numero
	r.GET("/calc/mul", controller.Multiplicar) // para usar esta funcion :http://localhost:8080/calc/mul?a=Numero&b=Numero
	r.GET("/calc/div", controller.Dividir)     // para usar esta funcion :http://localhost:8080/calc/div?a=Numero&b=Numero
	r.Run(":8080")
}
