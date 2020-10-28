package main

import (
    "fmt"
    "./multimedia"
    "bufio"
    "os"
)



func main() {
    opc := 0

    cw := multimedia.ContenidoWeb{}
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\t\tMenu de opciones\n")
        fmt.Println("1. Captura una imagen")
        fmt.Println("2. Captura un audio")
        fmt.Println("3. Captura un video")
        fmt.Println("4. Mostrar elementos")
        fmt.Println("5. Salir")
        fmt.Print("Escoge un opcion: ")
        fmt.Scan(&opc)
        fmt.Println("")

        switch opc {
            case 1:
                //imagen
                var titulo string
                var formato string
                var canales int32
                fmt.Print("Ingresa el Titulo: ")
                scanner.Scan()
                titulo = scanner.Text()
                fmt.Print("Ingresa el Formato: ")
                scanner.Scan()
                formato = scanner.Text()
                fmt.Print("Ingresa los Canales: ")
                fmt.Scan(&canales)
                img := multimedia.Imagen{titulo, formato, canales}
                cw.Contenidos = append(cw.Contenidos, &img)
            case 2:
                //Audio
                var titulo string
                var formato string
                var duracion float32
                fmt.Print("Ingresa el Titulo: ")
                scanner.Scan()
                titulo = scanner.Text()
                fmt.Print("Ingresa el Formato: ")
                scanner.Scan()
                formato = scanner.Text()
                fmt.Print("Ingresa los Duracion: ")
                fmt.Scan(&duracion)
                au := multimedia.Audio{titulo, formato, duracion}
                cw.Contenidos = append(cw.Contenidos, &au)
            case 3:
                //Video
                var titulo string
                var formato string
                var frames int32
                fmt.Print("Ingresa el Titulo: ")
                scanner.Scan()
                titulo = scanner.Text()
                fmt.Print("Ingresa el Formato: ")
                scanner.Scan()
                formato = scanner.Text()
                fmt.Print("Ingresa los Frames: ")
                fmt.Scan(&frames)
                vid := multimedia.Video{titulo, formato, frames}
                cw.Contenidos = append(cw.Contenidos, &vid)
            case 4:
                cw.Mostrar()
        }
        if opc == 5 {
            break
        }
    }

}
