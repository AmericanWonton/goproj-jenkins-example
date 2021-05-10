package jenkinsscripts

def exampleBuildApp() {
    echo 'we are inside groovy script, building an app'
}

def examplePingServer() {

    def resultado = new StringBuilder() //(1)
    def error     = new StringBuilder()

    def comando = "curl -I josephkeller.me".execute() //(2)
    comando.consumeProcessOutput(resultado, error) //(3)
    comando.waitForOrKill(1000) //(4)

    if (!error.toString() == ("")) {
        println "Bad curl"
    } else {
        println "Successful curl"
    }
}

return this //Need this to import this groovy script into Jenkins
