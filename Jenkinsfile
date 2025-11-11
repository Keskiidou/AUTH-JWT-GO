pipeline {
    agent any

    environment {
        APP_CONTAINER = "go_auth_app"
        RELATED_CONTAINERS = "postgres_auth_service" // add more if needed
        LOG_DIR = "logs"
    }

    stages {
        stage('Checkout SCM') {
            steps {
                checkout scm
            }
        }

        stage('Build App') {
            steps {
                echo "Building the Go application..."
                bat """
                echo ===== BUILD START =====
                go build -v -o app.exe .
                echo ===== BUILD COMPLETE =====
                """
            }
        }

        stage('Run Tests') {
            steps {
                echo "Running Go unit tests..."
                bat """
                echo ===== TEST START =====
                go test ./... -v > ${LOG_DIR}/test_results.txt 2>&1
                echo ===== TEST COMPLETE =====
                """
                archiveArtifacts artifacts: "${LOG_DIR}/test_results.txt", allowEmptyArchive: true
            }
        }

        stage('Capture Application Logs') {
            steps {
                echo "Capturing logs only from the application container..."
                bat """
                mkdir ${LOG_DIR}
                echo ===== APP LOGS START =====
                docker logs %APP_CONTAINER% --tail 200 > ${LOG_DIR}/app_logs.txt 2>&1
                echo ===== APP LOGS COMPLETE =====
                """
                archiveArtifacts artifacts: "${LOG_DIR}/app_logs.txt", allowEmptyArchive: true
            }
        }

        stage('Capture Related Container Logs') {
            steps {
                echo "Capturing logs from related containers..."
                bat """
                for %%C in (%RELATED_CONTAINERS%) do (
                    echo ===== Capturing logs for %%C =====
                    docker logs %%C --tail 200 > ${LOG_DIR}/%%C_logs.txt 2>&1
                )
                """
                archiveArtifacts artifacts: "${LOG_DIR}/*_logs.txt", allowEmptyArchive: true
            }
        }

        stage('Capture Resource Stats') {
            steps {
                echo "Capturing resource stats (CPU, Memory, Disk usage)..."
                bat """
                echo ===== RESOURCE STATS =====
                docker stats --no-stream --format "table {{.Container}}\t{{.CPUPerc}}\t{{.MemUsage}}" > ${LOG_DIR}/resource_usage.txt
                echo ===== DISK USAGE =====
                docker system df > ${LOG_DIR}/disk_usage.txt
                """
                archiveArtifacts artifacts: "${LOG_DIR}/resource_usage.txt, ${LOG_DIR}/disk_usage.txt", allowEmptyArchive: true
            }
        }

        stage('Summarize Logs') {
            steps {
                echo "Summarizing captured logs..."
                bat """
                echo ===== SUMMARY START =====
                echo Captured logs:
                dir ${LOG_DIR}
                echo ===== SUMMARY END =====
                """
            }
        }
    }

    post {
        always {
            echo "Pipeline finished successfully. All logs archived for AI analysis."
        }
    }
}
