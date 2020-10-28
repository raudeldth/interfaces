package multimedia

import "fmt"

type ContenidoWeb struct {
    Contenidos []Multimedia
}



func (cw *ContenidoWeb) Mostrar() {
    for _, c := range cw.Contenidos {
        c.Mostrar()
        fmt.Println("")
    }
}


type Multimedia interface {
    Mostrar()
}

type Imagen struct {
    Titulo string
    Formato string
    Canales int32
}

func (i *Imagen) Mostrar() {
    fmt.Println("Titulo: ", i.Titulo)
    fmt.Println("Formato: ", i.Formato)
    fmt.Println("Canales: ", i.Canales)
}

type Audio struct {
    Titulo string
    Formato string
    Duracion float32
}

func (a *Audio) Mostrar() {
    fmt.Println("Titulo: ", a.Titulo)
    fmt.Println("Formato: ", a.Formato)
    fmt.Println("Canales: ", a.Duracion)
}

type Video struct {
    Titulo string
    Formato string
    Frames int32
}

func (v *Video) Mostrar() {
    fmt.Println("Titulo: ", v.Titulo)
    fmt.Println("Formato: ", v.Formato)
    fmt.Println("Canales: ", v.Frames)
}
