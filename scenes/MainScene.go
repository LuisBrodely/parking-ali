package scenes

import (
	"image/color"
	"math/rand"
	"sync"
	"time"

	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"

	"estacionamiento/models"
)

var (
	lugaresDeEstacionamiento = []*models.ParkingSpot{
		models.NewParkingSpot(410, 210, 440, 240, 1, 1),
		models.NewParkingSpot(410, 255, 440, 285, 1, 2),
		models.NewParkingSpot(410, 300, 440, 330, 1, 3),
		models.NewParkingSpot(410, 345, 440, 375, 1, 4),
		models.NewParkingSpot(320, 210, 350, 240, 2, 5),
		models.NewParkingSpot(320, 255, 350, 285, 2, 6),
		models.NewParkingSpot(320, 300, 350, 330, 2, 7),
		models.NewParkingSpot(320, 345, 350, 375, 2, 8),
		models.NewParkingSpot(230, 210, 260, 240, 3, 9),
		models.NewParkingSpot(230, 255, 260, 285, 3, 10),
		models.NewParkingSpot(230, 300, 260, 330, 3, 11),
		models.NewParkingSpot(230, 345, 260, 375, 3, 12),
		models.NewParkingSpot(140, 210, 170, 240, 4, 13),
		models.NewParkingSpot(140, 255, 170, 285, 4, 14),
		models.NewParkingSpot(140, 300, 170, 330, 4, 15),
		models.NewParkingSpot(140, 345, 170, 375, 4, 16),
		models.NewParkingSpot(50, 210, 80, 240, 5, 17),
		models.NewParkingSpot(50, 255, 80, 285, 5, 18),
		models.NewParkingSpot(50, 300, 80, 330, 5, 19),
		models.NewParkingSpot(50, 345, 80, 375, 5, 20),
	}
	estacionamiento = models.NuevoEstacionamiento(lugaresDeEstacionamiento)
	colaVehiculos   = estacionamiento.GetVehicleQueue()
	mutexPuerta     sync.Mutex
	gestorVehiculos = models.NuevoGestorDeVehiculos()
)

type EscenaEstacionamiento struct {
}

func NuevaEscenaEstacionamiento() *EscenaEstacionamiento {
	return &EscenaEstacionamiento{}
}

func (ee *EscenaEstacionamiento) Comenzar() {
	esPrimeraVez := true

	_ = oak.AddScene("escenaEstacionamiento", scene.Scene{
		Start: func(ctx *scene.Context) {
			_ = ctx.Window.SetBorderless(true)
			PrepararEscena(ctx)

			event.GlobalBind(ctx, event.Enter, func(enterPayload event.EnterPayload) event.Response {
				if !esPrimeraVez {
					return 0
				}

				esPrimeraVez = false

				for i := 0; i < 100; i++ {
					go CicloAutomovil(ctx)

					time.Sleep(time.Millisecond * time.Duration(ObtenerNúmeroAleatorio(1000, 2000)))
				}

				return 0
			})
		},
	})
}

func PrepararEscena(ctx *scene.Context) {

	áreaEstacionamiento := floatgeom.NewRect2(-110, -250, 700, 1000)
	spritePath := "assets/UP.jpg"
	sprite, _ := render.LoadSprite(spritePath)

	entities.New(ctx, entities.WithRect(áreaEstacionamiento), entities.WithRenderable(sprite))

	entrada := floatgeom.NewRect2(440, 170, 500, 180)
	entities.New(ctx, entities.WithRect(entrada), entities.WithColor(color.RGBA{0, 134, 50, 255}))

	for _, lugar := range lugaresDeEstacionamiento {
		entities.New(ctx, entities.WithRect(*lugar.ObtenerÁrea()), entities.WithColor(color.RGBA{117, 0, 68, 255}))
	}
}

func CicloAutomovil(ctx *scene.Context) {
	automovil := models.NuevoVehiculo(ctx)

	gestorVehiculos.AgregarVehiculo(automovil)

	automovil.Encolar(gestorVehiculos)

	lugarDisponible := estacionamiento.ObtenerParkingSpotDisponible()

	mutexPuerta.Lock()

	automovil.UnirsePuerta(gestorVehiculos)

	mutexPuerta.Unlock()

	automovil.Estacionarse(lugarDisponible, gestorVehiculos)

	time.Sleep(time.Millisecond * time.Duration(ObtenerNúmeroAleatorio(40000, 50000)))

	automovil.DejarLugar(gestorVehiculos)

	estacionamiento.LiberarParkingSpot(lugarDisponible)

	automovil.Salir(lugarDisponible, gestorVehiculos)

	mutexPuerta.Lock()

	automovil.SalirPuerta(gestorVehiculos)

	mutexPuerta.Unlock()

	automovil.Alejarse(gestorVehiculos)

	automovil.Eliminar()

	gestorVehiculos.EliminarVehiculo(automovil)
}

func ObtenerNúmeroAleatorio(min, max int) float64 {
	origen := rand.NewSource(time.Now().UnixNano())
	generador := rand.New(origen)
	return float64(generador.Intn(max-min+1) + min)
}
