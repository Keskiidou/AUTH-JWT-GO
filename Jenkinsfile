pipeline {
    agent any

    environment {
        APP_CONTAINER = "go_auth_app"
        DB_CONTAINER  = "postgres_auth_service"
    }

    stages {
        stage('Checkout SCM') {
            steps {
                checkout scm
            }
        }

        stage('Massive Log Collection') {
            steps {
                echo "Collecting detailed logs from containers and system for AI analysis..."
                bat '''
                mkdir logs

                echo ===== Application Logs =====
                docker logs %APP_CONTAINER% > logs\\app_logs.txt 2>&1

                echo ===== Database Logs =====
                docker logs %DB_CONTAINER% > logs\\db_logs.txt 2>&1

                echo ===== Container Runtime Info =====
                docker ps -a > logs\\container_status.txt
                docker inspect %APP_CONTAINER% > logs\\container_inspect.txt
                docker stats --no-stream > logs\\container_stats.txt
                docker events --since 1h > logs\\docker_events.txt

                echo ===== System Info =====
                systeminfo > logs\\system_info.txt
                wmic cpu get loadpercentage > logs\\cpu_usage.txt
                wmic os get freephysicalmemory > logs\\memory_usage.txt
                typeperf "\\LogicalDisk(_Total)\\% Free Space" -sc 1 > logs\\disk_usage.txt

                echo ===== Network Info =====
                docker network ls > logs\\network_list.txt
                FOR /F "tokens=*" %%i IN ('docker ps -q') DO docker inspect %%i >> logs\\network_inspect.txt
                docker exec %APP_CONTAINER% ping -n 3 %DB_CONTAINER% > logs\\network_ping.txt 2>&1

                echo ===== Security / Image Info =====
                docker images > logs\\image_inventory.txt
                docker inspect %APP_CONTAINER% | findstr "User" > logs\\container_user.txt

                echo ===== Summary =====
                echo Logs successfully collected on %DATE% %TIME% > logs\\summary.txt
                '''
                archiveArtifacts artifacts: 'logs/**', allowEmptyArchive: true
            }
        }
    }

    post {
        always {
            echo "Pipeline finished successfully. Logs archived for AI analysis. No containers were deleted or modified."
        }
    }
}
