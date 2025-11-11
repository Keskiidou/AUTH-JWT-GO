pipeline {
    agent any

    environment {
        APP_CONTAINER = "go_auth_app"
        DB_CONTAINER  = "postgres_auth_service"
        DOCKER_NETWORK = "goapp-network"
        IMAGE_NAME = "goauth-app:latest"
    }

    stages {
        stage('Checkout SCM') {
            steps {
                checkout scm
            }
        }

        stage('Clean Old Docker') {
            steps {
                echo "Removing old Docker image and network if exist..."
                bat """
                docker rm -f %APP_CONTAINER% 2>nul || echo App container not running
                docker rmi -f %IMAGE_NAME% 2>nul || echo Image not found
                docker network rm %DOCKER_NETWORK% 2>nul || echo Network not found
                """
            }
        }

        stage('Build Docker Image') {
            steps {
                echo "Building Docker image..."
                bat "docker build -t %IMAGE_NAME% ."
            }
        }

        stage('Setup Docker Network') {
            steps {
                echo "Creating Docker network and connecting containers..."
                bat """
                docker network inspect %DOCKER_NETWORK% >nul 2>&1 || docker network create %DOCKER_NETWORK%
                docker network connect %DOCKER_NETWORK% %DB_CONTAINER% || echo DB already connected
                """
            }
        }

        stage('Run App Container') {
            steps {
                echo "Starting app container..."
                bat """
                docker run -d --name %APP_CONTAINER% --network %DOCKER_NETWORK% -p 3000:3000 %IMAGE_NAME%
                """
            }
        }

        stage('Run Tests') {
            steps {
                echo "Running Go tests inside app container..."
                bat "docker exec %APP_CONTAINER% go test ./... -v"
            }
        }

        stage('Capture Logs') {
            steps {
                echo "Capturing app logs..."
                bat "docker logs %APP_CONTAINER% > logs.txt 2>&1"
                archiveArtifacts artifacts: 'logs.txt', allowEmptyArchive: true
            }
        }
    }

    post {
        always {
            echo "Cleaning up..."
            bat "docker rm -f %APP_CONTAINER% 2>nul || echo No container to remove"
        }
    }
}
