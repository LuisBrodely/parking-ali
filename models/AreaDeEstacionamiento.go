package models

import (
	"image/color"

	"github.com/oakmound/oak/v4/alg/floatgeom"
)

type ParkingSpot struct {
	área                      *floatgeom.Rect2
	direccionesParaEstacionar *[]DirecciónLugarEstacionamiento
	direccionesParaSalir      *[]DirecciónLugarEstacionamiento
	número                    int
	disponible                bool
	color                     color.RGBA
}

func NewParkingSpot(x, y, x2, y2 float64, fila, número int) *ParkingSpot {
	direccionesParaEstacionar := obtenerDirecciónParaEstacionar(x, y, fila)
	direccionesParaSalir := obtenerDireccionesParaSalir()
	área := floatgeom.NewRect2(x, y, x2, y2)

	return &ParkingSpot{
		área:                      &área,
		direccionesParaEstacionar: direccionesParaEstacionar,
		direccionesParaSalir:      direccionesParaSalir,
		número:                    número,
		disponible:                true,
	}
}

func obtenerDirecciónParaEstacionar(x, y float64, fila int) *[]DirecciónLugarEstacionamiento {
	var direcciones []DirecciónLugarEstacionamiento

	if fila == 1 {
		direcciones = append(direcciones, *NuevaDirecciónLugarEstacionamiento("izquierda", 445))
	} else if fila == 2 {
		direcciones = append(direcciones, *NuevaDirecciónLugarEstacionamiento("izquierda", 355))
	} else if fila == 3 {
		direcciones = append(direcciones, *NuevaDirecciónLugarEstacionamiento("izquierda", 265))
	} else if fila == 4 {
		direcciones = append(direcciones, *NuevaDirecciónLugarEstacionamiento("izquierda", 175))
	} else if fila == 5 {
		direcciones = append(direcciones, *NuevaDirecciónLugarEstacionamiento("izquierda", 85))
	}

	direcciones = append(direcciones, *NuevaDirecciónLugarEstacionamiento("abajo", y+5))
	direcciones = append(direcciones, *NuevaDirecciónLugarEstacionamiento("izquierda", x+5))

	return &direcciones
}

func nuevaDirecciónLugarEstacionamiento(s string, i int) {
	panic("unimplemented")
}

func obtenerDireccionesParaSalir() *[]DirecciónLugarEstacionamiento {
	var direcciones []DirecciónLugarEstacionamiento

	direcciones = append(direcciones, *NuevaDirecciónLugarEstacionamiento("abajo", 420))
	direcciones = append(direcciones, *NuevaDirecciónLugarEstacionamiento("derecha", 475))
	direcciones = append(direcciones, *NuevaDirecciónLugarEstacionamiento("arriba", 185))

	return &direcciones
}

func (l *ParkingSpot) ObtenerÁrea() *floatgeom.Rect2 {
	return l.área
}

func (l *ParkingSpot) ObtenerNúmero() int {
	return l.número
}

func (l *ParkingSpot) ObtenerDireccionesParaEstacionar() *[]DirecciónLugarEstacionamiento {
	return l.direccionesParaEstacionar
}

func (l *ParkingSpot) ObtenerDireccionesParaSalir() *[]DirecciónLugarEstacionamiento {
	return l.direccionesParaSalir
}

func (l *ParkingSpot) ObtenerDisponibilidad() bool {
	return l.disponible
}

func (l *ParkingSpot) EstablecerDisponibilidad(disponible bool) {
	l.disponible = disponible
}
