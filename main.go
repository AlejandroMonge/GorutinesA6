package main

import (
	"fmt"
	"time"
	"strconv"
)

// Estructura del proceso
type Proceso struct  {
	i int64
	id int
	cTerminar chan bool
}

func (p *Proceso) start(c chan bool)  {
	continuar := true
	for (continuar) {
		select {
		case <- p.cTerminar:
			continuar = false
		default:
			//fmt.Println("BUG - 1")
		}
		if( <- c ){
			if(p.id != -1){
				fmt.Printf("id %d: %d \n", p.id, p.i)
			}	
			c <- true
		}else{
			c <- false
		}
		if p.id != -1 {
			p.i = p.i + 1
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Proceso) terminarProceso(){
	p.cTerminar <- true
}
//Fin de la estructura del proceso

func main()  {

	var input string
	countId := 0
	ocultarMenu := false
	var c chan bool = make(chan bool)
	sProcesos := []Proceso {Proceso{0,-1,make(chan bool)}}
	go sProcesos[0].start(c)
	c <- false
	for{
		if(!ocultarMenu){
			fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
			fmt.Println("Agregar proceso ...... [1]")
			fmt.Println("Mostrar procesos ..... [2]")
			fmt.Println("Terminar proceso ..... [3]")
			fmt.Println("Agregar proceso ...... [0]")
			fmt.Print("Elige una opcion: ")
		}
		fmt.Scan(&input)
		if input == "1" {
			p := Proceso{0, countId, make(chan bool)}
			countId = countId + 1
			go p.start(c)
			sProcesos = append(sProcesos, p)
			fmt.Printf("Proceso %d agregado \n" , countId - 1)
		} else if input == "2" {
			b := <- c
			if b {
				c <- false
				ocultarMenu = false
			}else {
				c <- true
				ocultarMenu = true
			}
		} else if input == "3" {
			fmt.Print("Id a eliminar: ")
			fmt.Scan(&input)
			i, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(err)
			}
			sProcesos[i+1].terminarProceso()
			fmt.Printf("Proceso %d Terminado \n" , i)
		} else if input == "0" {
			break
		}
	}
}