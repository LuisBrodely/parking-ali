package models

import (
	"sync"
)

type Estacionamiento struct {
	Lugares              []*ParkingSpot
	ColaVColaDeVehiculos *ColaDeVehiculos
	mu                   sync.Mutex
	CondDisponible       *sync.Cond
}

func NuevoEstacionamiento(lugares []*ParkingSpot) *Estacionamiento {
	cola := NuevaColaDeVehiculos()
	e := &Estacionamiento{
		Lugares:              lugares,
		ColaVColaDeVehiculos: cola,
	}
	e.CondDisponible = sync.NewCond(&e.mu)
	return e
}

func (e *Estacionamiento) ObtenerLugares() []*ParkingSpot {
	return e.Lugares
}

func (e *Estacionamiento) ObtenerParkingSpotDisponible() *ParkingSpot {
	e.mu.Lock()
	defer e.mu.Unlock()

	for {
		for _, lugar := range e.Lugares {
			if lugar.ObtenerDisponibilidad() {
				lugar.EstablecerDisponibilidad(false)
				return lugar
			}
		}
		e.CondDisponible.Wait()
	}
}

func (e *Estacionamiento) LiberarParkingSpot(lugar *ParkingSpot) {
	e.mu.Lock()
	defer e.mu.Unlock()

	lugar.EstablecerDisponibilidad(true)
	e.CondDisponible.Signal()
}

func (e *Estacionamiento) GetVehicleQueue() *ColaDeVehiculos {
	return e.ColaVColaDeVehiculos
}
