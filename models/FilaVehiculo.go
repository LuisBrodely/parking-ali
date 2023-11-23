package models

import (
	"container/list"
	"sync"
)

type ColaDeVehiculos struct {
	cola  *list.List
	mutex sync.Mutex
}

func NuevaColaDeVehiculos() *ColaDeVehiculos {
	return &ColaDeVehiculos{
		cola: list.New(),
	}
}

func (ca *ColaDeVehiculos) Encolar(vehiculo *Vehiculo) {
	ca.mutex.Lock()
	defer ca.mutex.Unlock()
	ca.cola.PushBack(vehiculo)
}

func (ca *ColaDeVehiculos) Desencolar() *Vehiculo {
	ca.mutex.Lock()
	defer ca.mutex.Unlock()
	if ca.cola.Len() == 0 {
		return nil
	}
	elemento := ca.cola.Front()
	ca.cola.Remove(elemento)
	return elemento.Value.(*Vehiculo)
}

func (ca *ColaDeVehiculos) Primero() *Vehiculo {
	ca.mutex.Lock()
	defer ca.mutex.Unlock()
	if ca.cola.Len() == 0 {
		return nil
	}
	elemento := ca.cola.Front()
	return elemento.Value.(*Vehiculo)
}

func (ca *ColaDeVehiculos) Ultimo() *Vehiculo {
	ca.mutex.Lock()
	defer ca.mutex.Unlock()
	if ca.cola.Len() == 0 {
		return nil
	}
	elemento := ca.cola.Back()
	return elemento.Value.(*Vehiculo)
}

func (ca *ColaDeVehiculos) Tamaño() int {
	ca.mutex.Lock()
	defer ca.mutex.Unlock()
	return ca.cola.Len()
}
