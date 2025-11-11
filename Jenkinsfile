pipeline {
    agent any

    stages {
        stage('Setup Docker Network') {
            steps {
                echo "Creating Docker network if it doesn't exist..."
                bat """
                docker network inspect goapp-network >nul 2>&1 || docker network create goapp-network
                docker network connect goapp-network db-container || true
                docker network connect goapp-network app-container || true
                """
            }
        }

        stage('Run Tests') {
            steps {
                echo "Running Go tests inside existing app container..."
                bat 'docker exec app-container go test ./... -v'
            }
        }

        stage('Capture Logs') {
            steps {
                echo "Fetching logs from the existing app container..."
                bat 'mkdir logs || true'
                bat 'docker logs app-container > logs/app.log || true'
            }
        }
    }

    post {
        always {
            echo "Archiving logs..."
            archiveArtifacts artifacts: 'logs/**/*', allowEmptyArchive: true
        }
    }
}
