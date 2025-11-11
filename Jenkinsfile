pipeline {
    agent any

    environment {
        DOCKER_HOST = 'unix:///var/run/docker.sock'
    }

    stages {
        stage('Build Docker Image') {
            steps {
                echo "Building Docker image for Go app..."
                bat 'docker build -t my-go-app:latest .'
            }
        }

        stage('Run Tests') {
            steps {
                echo "Running Go tests inside Docker container..."
                bat 'docker run --rm my-go-app:latest go test ./... -v'
            }
        }

        stage('Run App') {
            steps {
                echo "Starting Go app in Docker container..."
                bat 'docker run -d --name go-app-instance -p 8080:8080 my-go-app:latest'
            }
        }

        stage('Capture Logs') {
            steps {
                echo "Fetching logs from container..."
                bat 'mkdir -p logs'
                bat 'docker logs go-app-instance > logs/app.log || true'
            }
        }
    }

    post {
        always {
            echo "Archiving logs..."
            archiveArtifacts artifacts: 'logs/**/*', allowEmptyArchive: true
            echo "Cleaning up Docker container..."
            bat 'docker rm -f go-app-instance || true'
        }
    }
}
