package main

import (
	"github.com/oakmound/oak/v4"
	"estacionamiento/scenes"
)

func main() {
	escenaEstacionamiento := scenes.NuevaEscenaEstacionamiento()

	escenaEstacionamiento.Comenzar()

	_ = oak.Init("escenaEstacionamiento")
}
