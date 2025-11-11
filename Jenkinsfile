pipeline {
    agent any

    environment {
        APP_CONTAINER = "go_auth_app"
    }

    stages {
        stage('Checkout SCM') {
            steps {
                checkout scm
            }
        }

        stage('Capture Logs') {
            steps {
                echo "Capturing logs from the running app container..."
                bat """
                docker logs %APP_CONTAINER% > logs.txt 2>&1
                """
                archiveArtifacts artifacts: 'logs.txt', allowEmptyArchive: true
            }
        }
    }

    post {
        always {
            echo "Pipeline finished. No containers were deleted or modified."
        }
    }
}
