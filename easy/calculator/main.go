package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Calculator")
	win.SetDefaultSize(500, 500)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}

	// Crear 10 botones con números del 0 al 9
	for i := 0; i < 10; i++ {
		// Crear un nuevo botón con la etiqueta del número actual
		btn, err := gtk.ButtonNewWithLabel(fmt.Sprintf("%d", i))
		if err != nil {
			panic(err)
		}

		// Obtener la posición de la fila y la columna del botón
		row := i / 3
		col := i % 3

		// Agregar el botón a la cuadrícula en la fila y columna adecuadas
		grid.Attach(btn, col, row, 1, 1)
	}

	win.Add(grid)

	win.ShowAll()
	gtk.Main()
}
