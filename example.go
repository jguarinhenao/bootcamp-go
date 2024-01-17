currentDir, err := os.Getwd()
require.Nil(t, err)

index := strings.Index(currentDir, "desafio-go-bases")
if index == -1 {
	// La subcadena no se encontró, manejar este caso según tus necesidades
	fmt.Println("Subcadena no encontrada en la ruta.")
	return
}
err = os.Chdir(currentDir[:index+len("desafio-go-bases")])
require.Nil(t, err)

defer func() {
	err := os.Chdir(currentDir)
	require.Nil(t, err)
}()