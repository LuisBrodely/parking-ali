package models

import "sync"

type GestorDeVehiculos struct {
	Vehiculos []*Vehiculo
	Mutex     sync.Mutex
}

func NuevoGestorDeVehiculos() *GestorDeVehiculos {
	return &GestorDeVehiculos{
		Vehiculos: make([]*Vehiculo, 0),
	}
}

func (ga *GestorDeVehiculos) AgregarVehiculo(Vehiculo *Vehiculo) {
	ga.Mutex.Lock()
	defer ga.Mutex.Unlock()
	ga.Vehiculos = append(ga.Vehiculos, Vehiculo)
}

func (ga *GestorDeVehiculos) EliminarVehiculo(Vehiculo *Vehiculo) {
	ga.Mutex.Lock()
	defer ga.Mutex.Unlock()
	for i, a := range ga.Vehiculos {
		if a == Vehiculo {
			ga.Vehiculos = append(ga.Vehiculos[:i], ga.Vehiculos[i+1:]...)
			break
		}
	}
}

func (ga *GestorDeVehiculos) ObtenerVehiculos() []*Vehiculo {
	ga.Mutex.Lock()
	defer ga.Mutex.Unlock()
	return ga.Vehiculos
}
