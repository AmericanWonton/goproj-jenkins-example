package jenkinsscripts

def exampleBuildApp() {
    echo 'we are inside groovy script, building an app'
}

def examplePingServer() {
    println "curl josephkeller.me".execute().text

    def resultado = new StringBuilder() //(1)
    def error     = new StringBuilder()

    def comando = "curl josephkeller.me".execute() //(2)
    comando.consumeProcessOutput(resultado, error) //(3)
    comando.waitForOrKill(1000) //(4)
    
    if (!error.toString() == ("")) {
        println "Error al ejecutar el comando"
    } else {
        println "Ejecutado correctamente"
    }
}

return this //Need this to import this groovy script into Jenkins
