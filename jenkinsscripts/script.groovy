package jenkinsscripts

def exampleBuildApp() {
    echo 'we are inside groovy script, building an app'
}

def examplePingServer() {
    println "ls".execute().text
    "cd ./var".execute()
    println "ls".execute().text
}

return this //Need this to import this groovy script into Jenkins
