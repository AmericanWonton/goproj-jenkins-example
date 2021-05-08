pipeline{
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
        GO111MODULE = 'on' //Used from Go Plugin
    }
    //Access build tools for projects
    tools {
        /* This is used for building tools you might need for applications.
        Rn, Jenkins only supports gradle, maven, and JDK. These tools need to be
        pre-installed in Jenkins. For example, adding this makes maven commands available */
        //maven 'Maven'
        go 'go-1.16.4' //This needs to be what you named it in config
    }
    //Things to execute in Jenkins
    stages{
        stage("build"){
            steps{
                echo "building the golang applicaiton"
                /* USE DOUBLE QUOTES SO IT'S COMPATIBLE WITH GROOVY! */
                echo "building version ${NEW_VERSION}"
                //sh "mvn install" //Available by adding in tools
                sh "ls -a" 
                sh "pwd"
                echo "just seeing if we need to CD into anything..."
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
            steps{
                echo "Golang App starting Testing"
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
            steps{
                echo "Deploying Golang App"
                //echo "Here is our server credentials: ${SERVER_CREDENTIALS}" //This is insecure, you get a warning
                /* You can also use this. It takes object Syntax, from Groovy.
                Passes in the Username and password you defined in Jenkins Admin.
                It then stores the Username you define in USER and password in PWD  */
                withCredentials([
                    usernamePassword(credentialsId: 'test-file-cred', usernameVariable: USER, passwordVariable: PWD)
                ]) {
                    //Here, you can run a shell script with those variables to do stuff
                    sh "some script ${USER} ${PWD}"
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