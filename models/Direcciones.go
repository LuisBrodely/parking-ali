package models

type DirecciónLugarEstacionamiento struct {
	Dirección string
	Punto     float64
}

func NuevaDirecciónLugarEstacionamiento(dirección string, punto float64) *DirecciónLugarEstacionamiento {
	return &DirecciónLugarEstacionamiento{
		Dirección: dirección,
		Punto:     punto,
	}
}
