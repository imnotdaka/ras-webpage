package mercadopago

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/plan"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preapproval"
	"github.com/mercadopago/sdk-go/pkg/preapprovalplan"
)

type client struct {
	pap preapprovalplan.Client
	pa  preapproval.Client
}

type Client interface {
	CreatePlan(context.Context, preapprovalplan.Request) (string, error)
	CreateSuscription(ctx context.Context, req preapproval.Request) (*preapproval.Response, error)
}

func NewClient(cfg *config.Config) Client {
	return &client{
		pap: preapprovalplan.NewClient(cfg),
		pa:  preapproval.NewClient(cfg),
	}
}

func (c client) CreatePlan(ctx context.Context, req preapprovalplan.Request) (string, error) {
	res, err := c.pap.Create(ctx, req)
	// client := &http.Client{}
	// res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	fmt.Println("res:", res)
	return res.ID, nil
}

func (c client) CreateSuscription(ctx context.Context, req preapproval.Request) (*preapproval.Response, error) {
	res, err := c.pa.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

func GetAll(r plan.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		plan, err := plan.Repository.GetAll(r)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		fmt.Printf("%v+ \n", plan)
		ctx.JSON(http.StatusOK, plan)
	}
}

func Webhook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		xSignature := ctx.GetHeader("x-signature")
		xRequestId := ctx.GetHeader("x-request-id")

		// Obtener los query params
		queryParams := ctx.Request.URL.Query()

		// Extraer "data.id" de los query params
		dataID := queryParams.Get("data.id")

		// Separar el x-signature en partes
		parts := strings.Split(xSignature, ",")

		// Inicializar variables para almacenar 'ts' y 'hash'
		var ts, hash string

		// Iterar sobre los valores para obtener 'ts' y 'v1'
		for _, part := range parts {
			// Dividir cada parte en key y value
			keyValue := strings.SplitN(part, "=", 2)
			if len(keyValue) == 2 {
				key := strings.TrimSpace(keyValue[0])
				value := strings.TrimSpace(keyValue[1])
				if key == "ts" {
					ts = value
				} else if key == "v1" {
					hash = value
				}
			}
		}

		// Obtener la clave secreta de la configuración (debería estar en un archivo .env o seguro)
		secret := os.Getenv("SECRET_KEY")

		// Generar el manifest string
		manifest := fmt.Sprintf("id:%v;request-id:%v;ts:%v;", dataID, xRequestId, ts)

		// Crear la firma HMAC utilizando sha256 y la clave secreta
		hmac := hmac.New(sha256.New, []byte(secret))
		hmac.Write([]byte(manifest))

		// Obtener el resultado del hash como una cadena hexadecimal
		sha := hex.EncodeToString(hmac.Sum(nil))

		// Verificar si la firma calculada coincide con la firma recibida
		if sha != hash {
			// La verificación HMAC falló
			fmt.Println("HMAC verification failed")
			ctx.JSON(401, gin.H{"message": "HMAC verification failed"})
			return
		}
		fmt.Println("HMAC verification successfull")

		var req any

		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		fmt.Printf("%+v", req)
	}
}
