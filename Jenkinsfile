pipeline {
    agent any

    environment {
        APP_CONTAINER = "go_auth_app"
        LOG_FILE = "logs.txt"
    }

    stages {
        stage('Checkout SCM') {
            steps {
                checkout scm
            }
        }

        stage('Follow App Logs') {
            steps {
                script {
                    echo "Following app logs continuously..."
                    // Keep reading the last 100 lines every 10 seconds, appending to logs.txt
                    while (true) {
                        bat "docker logs --tail 100 %APP_CONTAINER% >> %LOG_FILE% 2>&1"
                        sleep(time: 10, unit: 'SECONDS')
                    }
                }
            }
        }
    }
}
