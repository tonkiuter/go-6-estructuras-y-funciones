package main

import (
	"./multimedia"
	"fmt"
)

func mostrarMultimedia(f ...multimedia.Multimedia) {
	//f.Mostrar()

	for _, v:= range f{
		v.Mostrar()
	}
}

func main() {
	var op int
	var tituloImg string
	var formatoImg string
	var canales int
	var tituloAud string
	var formatoAud string
	var duracion int
	var tituloVid string
	var formatoVid string
	var frames int 


	fmt.Println("Menu")
	

	fmt.Println("1) Imagen")
	fmt.Println("2) Audio")
	fmt.Println("3) Video")
	fmt.Println("4) Mostrar")

	for{
		fmt.Println("Seleccione la opcion que desea capturar")
		fmt.Scanln(&op)

		if op == 1{
			fmt.Println("Titulo: ")
			fmt.Scanln(&tituloImg)
			fmt.Println("Formato: ")
			fmt.Scanln(&formatoImg)
			fmt.Println("Canales")
			fmt.Scanln(&canales)
			

		}else if op == 2{
			fmt.Println("Titulo: ")
			fmt.Scanln(&tituloAud)
			fmt.Println("Formato: ")
			fmt.Scanln(&formatoAud)
			fmt.Println("Duracion")
			fmt.Scanln(&duracion)
		
			

		}else if op == 3{
			fmt.Println("Titulo: ")
			fmt.Scanln(&tituloVid)
			fmt.Println("Formato: ")
			fmt.Scanln(&formatoVid)
			fmt.Println("Frames")
			fmt.Scanln(&frames)

			

		}else if op == 4{
			imagen := multimedia.Imagen{tituloImg, formatoImg, canales}
			cancion := multimedia.Audio{tituloAud,formatoAud,duracion}
			video := multimedia.Video{tituloVid,formatoVid,frames}
			s := multimedia.ContenidoWeb{
				Multimedia: [] multimedia.Multimedia{
					&imagen,
					&cancion,
					&video,
				},
			}
			fmt.Println("Salida: ")
			s.Mostrar()
		}else{
			break
		} 
	}

	

	//cancion := multimedia.Audio{Titulo: "Sequentia", Formato: "Mp3", Duracion: 120}
	//video := multimedia.Video{Titulo: "Unveil", Formato: "FLAC", Frames: 60}
	//imagen := multimedia.Imagen{Titulo:"Nanji", Formato: "jpg", Canales: 2}

	//cancion.Mostrar()
	//video.Mostrar()
	//imagen.Mostrar()

	//mostrarMultimedia(&imagen,&video, &cancion)
	//mostrarMultimedia(&video)



	/*file := multimedia.ContenidoWeb{
		Multimedia: []multimedia.Multimedia{
			&cancion,
			&video,
			&imagen,
		},
	}*/


	//file.Mostrar()


}
