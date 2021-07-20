# mary_system
Aplicacion de inventario con GO Vuejs


prerrequisitos

Wails usa cgo para enlazar a los motores de renderizado nativos, por lo que se necesitan varias dependencias de plataforma, así como una instalación de Go. Los requisitos básicos son:

Ir 1.12 o superior
Bibliotecas gcc +
npm
Compatibilidad con Docker para compilación cruzada
ir
Descargar Ir desde el Página de descargas de Go. Se recomienda descargar el paquete MSI.

Asegúrese de seguir al funcionario Ir instrucciones de instalación. También deberá asegurarse de que la variable de entorno también incluye la ruta de acceso al directorio. También asegúrese de que los módulos Go estén habilitados: puede hacerlo a través del botón "Variables de entorno" en la pestaña "Avanzado" del panel de control "Sistema". Algunas versiones de Windows proporcionan este panel de control a través de la opción "Configuración avanzada del sistema" dentro del panel de control "Sistema".PATHc:\Go\bin
Reinicie el terminal y realice las siguientes comprobaciones:

Check Go está instalado de forma básica: go version
Compruebe que "c:\Go\bin" está en la variable PATH: echo %PATH%
Si utiliza Go 1.12 en lugar de 1.13, marque que la variable de entorno "GO111MODULE" está establecida en "on": .echo %GO111MODULE%
Se recomienda encarecidamente que si utiliza un terminal en Windows, que utilice el terminal git bash y no el terminal VSC. Esto se debe a que maneja las señales correctamente. Hemos encontrado que el terminal VSC es bastante defectuoso.

npm
Descargar NPM desde el Página de descargas de nodos. Lo mejor es usar la última versión, ya que es contra lo que generalmente probamos.

Ejecutar para comprobar.npm --version
Hay un problema conocido con NPX con respecto a los espacios en su ruta de perfil de usuario. Si tiene algún problema, consulte ese ticket para obtener soluciones alternativas.

GCC + Bibliotecas
Windows requiere gcc y herramientas relacionadas. La descarga recomendada es desde https://jmeubank.github.io/tdm-gcc/download/.




Instalacion de wails

La instalación es tan simple como ejecutar el siguiente comando:

* go get -u github.com/wailsapp/wails/cmd/wails


servidor 
* wails serve
Al desarrollar sus aplicaciones con wails el método preferido es por el comando serve wails serve.

Esto produce una compilación ligera mucho más rápida en modo de depuración, excluyendo los scripts de compilación, ahorrando tiempo al desarrollar el backend y también permitiendo el uso de para el desarrollo parcial del navegador de frontend

 * npm run serve


Compilar a .exe

En la raiz del proyecto ejecuta wails build y espera hasta que se genere el archivo.exe
* wails build.

Si todo salió bien, debería tener un programa compilado en su directorio local. Ejecútelo con o haga doble clic si está en windows../my-projectmyproject.exe