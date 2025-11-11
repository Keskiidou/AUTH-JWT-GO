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

        stage('Start Continuous Log Capture') {
            steps {
                echo "Starting continuous log capture from app container..."
                // Start log capture in the background, append to logs.txt
                bat """
                start /b cmd /c "docker logs -f %APP_CONTAINER% >> %LOG_FILE% 2>&1"
                """
            }
        }

        stage('Archive Logs Periodically') {
            steps {
                echo "Archiving logs..."
                // Wait a bit to let logs accumulate (e.g., 1 min)
                sleep(time: 60, unit: 'SECONDS')
                archiveArtifacts artifacts: LOG_FILE, allowEmptyArchive: true
            }
        }
    }

    post {
        always {
            echo "Pipeline finished. Container logs continue to be appended in the background."
        }
    }
}
