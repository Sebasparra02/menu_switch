package main

import "fmt"

func main() {
	menu_principal := `
	***MENU PRINCIPAL***
	1) Cambiar su contraseña
	2) Editar perfil 
	3) Foro
	4) Finalizar programa`

	menu_perfil := `
	***MENU PERFIL***
	1) Cambiar foto de perfil
	2) Cambiar nombre de usuario
	3) Editar biografia
	4) Volver al menu principal`

	menu_foro := `
	***MENU FORO***
	1) Publicar nueva entrada en el foro
	2) Editar una entrada del foro
	3) Eliminar una entrada del foro
	4) Volver al menu principal`

	opcion_menu_principal := 0
	opcion_menu_perfil := 0
	opcio_menu_foro := 0

	for {
		switch opcion_menu_principal {
		case 0:
			fmt.Println(menu_principal)
			_, err := fmt.Scanf("%d", &opcion_menu_principal)
			if err != nil {
				panic(err)
			}

		case 1:
			fmt.Println("tu contraseña a sido cambiada con exito!!!")
			opcion_menu_principal = 0
		case 2:
			switch opcion_menu_perfil {
			case 0:
				fmt.Println(menu_perfil)
				_, err := fmt.Scanf("%d", &opcion_menu_perfil)
				if err != nil {
					panic(err)
				}

			case 1:
				fmt.Println("tu foto de perfil a sido actualizada!!")
				opcion_menu_perfil = 0

			case 2:
				fmt.Println("tu nombre de usuario se actualizo con exito!!")
				opcion_menu_perfil = 0

			case 3:
				fmt.Println("tu biografia se actualizo con exito!!")
				opcion_menu_perfil = 0

			case 4:
				opcion_menu_principal = 0
			}
		case 3:
			switch opcio_menu_foro {
			case 0:
				fmt.Println(menu_foro)
				_, err := fmt.Scanf("%d", &opcio_menu_foro)
				if err != nil {
					panic(err)
				}
			case 1:
				fmt.Println("tu nueva entrada se publico con exito")
				opcio_menu_foro = 0

			case 2:
				fmt.Println("editaste tu nueva entrada en el foro con exito !!")
				opcio_menu_foro = 0
			case 3:
				fmt.Println("eliminaste tu entrada en el foro")
				opcio_menu_foro = 0
			case 4:
				opcion_menu_principal = 0

			}
		case 4:
			fmt.Println("gracias por visitarnos!!")
		}
		if opcion_menu_principal == 4 {
			break
		}

	}
}
