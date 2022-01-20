package main

import "fmt"

func main() {
	fmt.Println("***MENU PRINCIPAL***")
	menu_principal := `
	1) Cambiar su contraseña
	2) Editar perfil 
	3) Foro
	4) Finalizar programa`
	_, err := fmt.Scanf("%s", &menu_principal)
	if err != nil {
		panic(err)
	}

	fmt.Println("***MENU PERFIL***")
	menu_perfil := `
	1) Cambiar foto de perfil
	2) Cambiar nombre de usuario
	3) Editar biografia
	4) Volver al menu principal`
	_, err := fmt.Scanf("%s", &menu_perfil)
	if err != nil {
		panic(err)
	}

}
