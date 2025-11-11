pipeline {
    agent any

    stages {
        stage('Build Docker Image') {
            steps {
                echo "Building Docker image for Go app..."
                bat 'docker build -t goauth-app:latest .'
            }
        }

        stage('Run Tests') {
            steps {
                echo "Running Go tests inside Docker container..."
                bat 'docker run --rm goauth-app:latest go test ./... -v'
            }
        }

        stage('Run App') {
            steps {
                echo "Starting Go app container..."
                bat '''
                REM Remove old container if exists
                docker rm -f goauth-container >nul 2>&1 || exit 0

                REM Start new container on same ports as your Docker Compose
                docker run -d --name goauth-container -p 3000:3000 goauth-app:latest
                '''
            }
        }

        stage('Capture Logs') {
            steps {
                echo "Fetching logs from container..."
                bat '''
                REM Create logs folder if it doesn't exist
                if not exist logs mkdir logs

                REM Capture logs
                docker logs goauth-container > logs\\app.log 2>&1
                '''
            }
        }
    }

    post {
        always {
            echo "Archiving logs..."
            archiveArtifacts artifacts: 'logs/**/*', allowEmptyArchive: true
            echo "Cleaning up Docker container..."
            bat 'docker rm -f goauth-container >nul 2>&1 || exit 0'
        }
    }
}
