package multimedia

import "fmt"

type ContenidoWeb struct{
	Multimedia []Multimedia
}

func(mediaFile *ContenidoWeb)Mostrar(){

	for _,v:= range mediaFile.Multimedia{
		v.Mostrar()
	}
}

type Multimedia interface{
	Mostrar()
}

type Imagen struct{
	Titulo string 
	Formato string
	Canales int
}

func(img *Imagen) Mostrar(){
	fmt.Println(img.Titulo)
	fmt.Println(img.Formato)
	fmt.Println(img.Canales)
}

type Audio struct{
	Titulo string
	Formato string
	Duracion int
}

func(aud *Audio) Mostrar(){
	fmt.Println(aud.Titulo)
	fmt.Println(aud.Formato)
	fmt.Println(aud.Duracion)
}


type Video struct{
	Titulo string
	Formato string
	Frames int
}

func(vid *Video) Mostrar() {
	fmt.Println(vid.Titulo)
	fmt.Println(vid.Formato)
	fmt.Println(vid.Frames)
}