def gv //Need this to declare our groovy script into a variable under 'init'

pipeline { 
    agent any //Run this on ANY Jenkins Server
    //This is used to define environmental variables; they can be used in any stage!
    environment {
        NEW_VERSION = '1.1' //Test variable, you would get this from your code usually
        /* The 'credentails' method finds credentials stored in jenkins to pass in this file.
        It's in the 'credentials binding' plug-in you need installed. It takes the ID you made as refference.
        This is just one way of doing it...you can also use the 'withCredentials()' wrapper,
        (see the deploy section)  */
        SERVER_CREDENTIALS = credentials('test-file-cred')
        DOCKER_CREDENTIALS = credentials('dockerlogin')
        RESUME_PEM = credentials('resume-private-key')
        //GO111MODULE = 'on' //Used from Go Plugin; kind of messing up go modules
        // Ensure the desired Go version is installed
        //def root = tool type: 'go', name: 'go-1.16.4'
    }
    //Access build tools for projects
    /* tools {
        /* This is used for building tools you might need for applications.
        Rn, Jenkins only supports gradle, maven, and JDK. These tools need to be
        pre-installed in Jenkins. For example, adding this makes maven commands available */
        //maven 'Maven'
        //go 'go-1.16.4' //This needs to be what you named it in config
    //}*/
    //Used to deploy application with certain paramters given
    parameters{
        string(name: 'TEST_PARAMETER', defaultValue: '', description: "This is for running this application with THIS parameter")
        choice(name: "TEST_CHOICE_PARAMETER", choices: ['choice1', 'choice2'], description: 'This is an example choice parameter ')
        /* Use these variables for expediancy running these jobs */
        booleanParam(name: 'executeTests', defaultValue: false, description: 'Test description')
        booleanParam(name: 'runinit', defaultValue: true, description: 'Run initialization')
        booleanParam(name: 'runBuild', defaultValue: false, description: 'Run build')
        booleanParam(name: 'runDeploy', defaultValue: false, description: 'Run Deploy')
        booleanParam(name: 'runGroovy', defaultValue: true, description: 'Run Groovy')
    }
    
    //Things to execute in Jenkins
    stages{
        stage("init"){
            when {
                expression {
                    params.runinit
                }
            }
            steps{
                //writeFile(file: "./jenkinsscripts/script.groovy", text: "")
                /* This is how we load our groovy scripts into Jenkins */
                script {
                    gv = load "./jenkinsscripts/script.groovy"
                }
                /* Need to write a pem key file and folder for us to work in */
                withCredentials([
                    string(credentialsID: 'resume-private-key', variable: 'SECRET')
                ]) {
                    dir ('security'){
                        writeFile file:'resumekeypair.pem', text: $SECRET
                    }
                    sh 'sudo cat ./security/resumekeypair.pem'
                }
            }
        }
        stage("build"){
            when {
                expression {
                    params.runBuild
                }
            }
            steps{
                echo "building the golang applicaiton"
                /* USE DOUBLE QUOTES SO IT'S COMPATIBLE WITH GROOVY! */
                //echo "building version ${NEW_VERSION}"
                //sh "mvn install" //Available by adding in tools
                //sh "ls -a" 
                //sh "pwd"
            }
            post{
                always{
                    echo "Finished building golang application"
                }
                success{
                    echo "Golang app built successfully"
                }
                failure{
                    echo "Golang app build un-successfully"
                }
            }
        }
        stage("test"){
            /* This is an example when clause; this works when the expressions defined inside are true */
            when {
                expression {
                    params.executeTests
                }
            }
            steps{
                echo "Golang App starting Testing"
                sh 'go version'
                sh 'go env'
                sh 'go test ./testing/ -v'
                /*
                withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                    sh 'go version'
                    sh 'go test ./testing/ -v'
                }
                */
            }
            post{
                always{
                    echo "Finished testing golang app"
                }
                success{
                    echo "Golang App Tested Successfully"
                }
                failure {
                    echo "Golang App Failed testing"
                }
            }
        }
        stage("deploy"){
            /* This would be a good place to pass credentials to a server, for building on a dev machine, 
            or SSH into a dev machine */
            when {
                expression {
                    params.runDeploy
                }
            }
            steps{
                echo "Deploying Golang App"
                echo "Deploying vesrion ${params.TEST_PARAMETER}"
                //echo "Here is our server credentials: ${SERVER_CREDENTIALS}" //This is insecure, you get a warning
                /* You can also use this. It takes object Syntax, from Groovy.
                Passes in the Username and password you defined in Jenkins Admin.
                It then stores the Username you define in USER and password in PWD  */
                withCredentials([
                    usernamePassword(credentialsId: 'test-file-cred', usernameVariable: USER, passwordVariable: PWD)
                ]) {
                    //Here, you can run a shell script with those variables to do stuff
                    sh "echo ${USER}"
                    sh "echo ${PWD}"
                }
            }
            post{
                always{
                    echo "Finished Deploying Golang App"
                }
                success{
                    echo "Golang App Succeeded deploying"
                }
                failure{
                    echo "Golang App failed deploying"
                }
            }
        }
        stage("groovy-test"){
            options {
                timeout(time: 30, unit: 'SECONDS') //Timeout for this stage
            }
            when {
                expression {
                    params.runGroovy
                }
            }
            steps{
                /* Example using scripts within Jenkins */
                script {
                    gv.exampleBuildApp() //Print Line
                    gv.examplePingServer() //Ping server
                    /* Not sure how to do this in groovy with ssh key so I'll do it in here */
                    sh 'pwd'
                    //sh 'scp -i ./security/resumekeypair.pem ./project/main.go ubuntu@ec2-3-137-207-149.us-east-2.compute.amazonaws.com:'
                }
            }

            post{
                always{
                    echo "Finished Testing Groovy Stuff"
                }
                success{
                    echo "Succeeded Testing Groovy Stuff"
                }
                failure{
                    echo "Failed Testing Groovy Stuff"
                }
            }
        }
    }
    //Things to do AFTER Jenkins builds
    post{
        always{
            echo "========always========"
        }
        success{
            echo "========pipeline executed successfully ========"
        }
        failure{
            echo "========pipeline execution failed========"
        }
    }
}
