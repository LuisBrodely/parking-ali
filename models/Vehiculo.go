package models

import (
	"image/color"
	"sync"
	"time"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

const (
	entrada   = 185.00
	velocidad = 10
)

type Vehiculo struct {
	area    floatgeom.Rect2
	entidad *entities.Entity
	mu      sync.Mutex
}

func NuevoVehiculo(ctx *scene.Context) *Vehiculo {
	area := floatgeom.NewRect2(445, -20, 465, 0)
	spritePath := "assets/carro.png"

	sprite, _ := render.LoadSprite(spritePath)
	entidad := entities.New(ctx, entities.WithRect(area), entities.WithColor(color.RGBA{255, 100, 0, 255}), entities.WithRenderable(sprite), entities.WithDrawLayers([]int{1, 2}))

	return &Vehiculo{
		area:    area,
		entidad: entidad,
	}
}
func (a *Vehiculo) Encolar(gestor *GestorDeVehiculos) {
	for a.Y() < 145 {
		if !a.Collision("abajo", gestor.ObtenerVehiculos()) {
			a.DesplazarY(1)
			time.Sleep(velocidad * time.Millisecond)
		}
	}
}

func (a *Vehiculo) UnirsePuerta(gestor *GestorDeVehiculos) {
	for a.Y() < entrada {
		if !a.Collision("abajo", gestor.ObtenerVehiculos()) {
			a.DesplazarY(1)
			time.Sleep(velocidad * time.Millisecond)
		}
	}
}

func (a *Vehiculo) SalirPuerta(gestor *GestorDeVehiculos) {
	for a.Y() > 145 {
		if !a.Collision("arriba", gestor.ObtenerVehiculos()) {
			a.DesplazarY(-1)
			time.Sleep(velocidad * time.Millisecond)
		}
	}
}

func (a *Vehiculo) Estacionarse(spot *ParkingSpot, gestor *GestorDeVehiculos) {
	for índice := 0; índice < len(*spot.ObtenerDireccionesParaEstacionar()); índice++ {
		direcciones := *spot.ObtenerDireccionesParaEstacionar()
		if direcciones[índice].Dirección == "derecha" {
			for a.X() < direcciones[índice].Punto {
				if !a.Collision("derecha", gestor.ObtenerVehiculos()) {
					a.DesplazarX(1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		} else if direcciones[índice].Dirección == "abajo" {
			for a.Y() < direcciones[índice].Punto {
				if !a.Collision("abajo", gestor.ObtenerVehiculos()) {
					a.DesplazarY(1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		} else if direcciones[índice].Dirección == "izquierda" {
			for a.X() > direcciones[índice].Punto {
				if !a.Collision("izquierda", gestor.ObtenerVehiculos()) {
					a.DesplazarX(-1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		} else if direcciones[índice].Dirección == "arriba" {
			for a.Y() > direcciones[índice].Punto {
				if !a.Collision("arriba", gestor.ObtenerVehiculos()) {
					a.DesplazarY(-1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		}
	}
}

func (a *Vehiculo) Salir(spot *ParkingSpot, gestor *GestorDeVehiculos) {
	for índice := 0; índice < len(*spot.ObtenerDireccionesParaSalir()); índice++ {
		direcciones := *spot.ObtenerDireccionesParaSalir()
		if direcciones[índice].Dirección == "izquierda" {
			for a.X() > direcciones[índice].Punto {
				if !a.Collision("izquierda", gestor.ObtenerVehiculos()) {
					a.DesplazarX(-1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		} else if direcciones[índice].Dirección == "derecha" {
			for a.X() < direcciones[índice].Punto {
				if !a.Collision("derecha", gestor.ObtenerVehiculos()) {
					a.DesplazarX(1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		} else if direcciones[índice].Dirección == "arriba" {
			for a.Y() > direcciones[índice].Punto {
				if !a.Collision("arriba", gestor.ObtenerVehiculos()) {
					a.DesplazarY(-1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		} else if direcciones[índice].Dirección == "abajo" {
			for a.Y() < direcciones[índice].Punto {
				if !a.Collision("abajo", gestor.ObtenerVehiculos()) {
					a.DesplazarY(1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		}
	}
}

func (a *Vehiculo) DejarLugar(gestor *GestorDeVehiculos) {
	puntoX := a.X()
	for a.X() > puntoX-30 {
		if !a.Collision("izquierda", gestor.ObtenerVehiculos()) {
			a.DesplazarX(-1)
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func (a *Vehiculo) Alejarse(gestor *GestorDeVehiculos) {
	for a.Y() > -20 {
		if !a.Collision("arriba", gestor.ObtenerVehiculos()) {
			a.DesplazarY(-1)
			time.Sleep(velocidad * time.Millisecond)
		}
	}
}

func (a *Vehiculo) DesplazarY(dy float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.entidad.ShiftY(dy)
}

func (a *Vehiculo) DesplazarX(dx float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.entidad.ShiftX(dx)
}

func (a *Vehiculo) X() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.entidad.X()
}

func (a *Vehiculo) Y() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.entidad.Y()
}

func (a *Vehiculo) Eliminar() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.entidad.Destroy()
}

func (a *Vehiculo) Collision(dirección string, vehiculos []*Vehiculo) bool {
	distanciaMínima := 30.0
	for _, vehiculo := range vehiculos {
		if dirección == "izquierda" {
			if a.X() > vehiculo.X() && a.X()-vehiculo.X() < distanciaMínima && a.Y() == vehiculo.Y() {
				return true
			}
		} else if dirección == "derecha" {
			if a.X() < vehiculo.X() && vehiculo.X()-a.X() < distanciaMínima && a.Y() == vehiculo.Y() {
				return true
			}
		} else if dirección == "arriba" {
			if a.Y() > vehiculo.Y() && a.Y()-vehiculo.Y() < distanciaMínima && a.X() == vehiculo.X() {
				return true
			}
		} else if dirección == "abajo" {
			if a.Y() < vehiculo.Y() && vehiculo.Y()-a.Y() < distanciaMínima && a.X() == vehiculo.X() {
				return true
			}
		}
	}
	return false
}
