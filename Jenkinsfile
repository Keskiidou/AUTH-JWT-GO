pipeline {
    agent any

    environment {
        DOCKER_NETWORK = 'goapp-network'
    }

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
                bat """
                    docker run --rm --network ${env.DOCKER_NETWORK} goauth-app:latest go test ./... -v
                """
            }
        }

        stage('Run App') {
            steps {
                echo "Starting Go app in Docker container..."
                bat """
                    docker run -d --name goauth-container --network ${env.DOCKER_NETWORK} -p 3000:3000 goauth-app:latest
                """
            }
        }

        stage('Capture Logs') {
            steps {
                echo "Fetching logs from container..."
                bat 'mkdir logs 2>nul || exit 0'
                bat 'docker logs goauth-container > logs/app.log || exit 0'
            }
        }
    }

    post {
        always {
            echo "Archiving logs..."
            archiveArtifacts artifacts: 'logs/**/*', allowEmptyArchive: true
            echo "Cleaning up Docker container..."
            bat 'docker rm -f goauth-container || exit 0'
        }
    }
}
