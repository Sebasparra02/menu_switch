package main

import "fmt"

func main() {
	menu_principal := `
	***MENU PRINCIPAL***
	1) Cambiar su contraseña
	2) Editar perfil 
	3) Foro
	4) Finalizar programa`
	_, err := fmt.Scanf("%s", &menu_principal)
	if err != nil {
		panic(err)
	}

	menu_perfil := `
	***MENU PERFIL***
	1) Cambiar foto de perfil
	2) Cambiar nombre de usuario
	3) Editar biografia
	4) Volver al menu principal`
	_, err = fmt.Scanf("%s", &menu_perfil)
	if err != nil {
		panic(err)
	}

	menu_foro := `
	***MENU FORO***
	1) Publicar nueva entrada en el foro
	2) Editar una entrada del foro
	3) Eliminar una entrada del foro
	4) Volver al menu principal`
	_, err = fmt.Scanf("%s", &menu_foro)
	if err != nil {
		panic(err)
	}

	var seguir string

	for {
		switch menu_principal {
		case "m.p":
			fmt.Println(menu_principal)
			_, err = fmt.Scanf("%s", &menu_principal)
			if err != nil {
				panic(err)
			}

			if menu_principal == "1" {
				fmt.Println("Su contraseña a sido cambiada con exito")
			}
			fmt.Println("desea seguir navegando o desea finalizar el programa?")
			fmt.Println("si desea seguir indique SI de lo contrario NO")
			if seguir == "SI" {
				fmt.Println(menu_principal)
			} else {
				fmt.Println("te esperamos pronto")
				break
			}

		}

	}

}
