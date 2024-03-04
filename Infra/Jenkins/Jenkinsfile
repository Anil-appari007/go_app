pipeline{
    agent any
    environment{
        GO_HOME = tool name: 'GO_v1.21.1', type: 'go'
        SONAR_HOME = tool name: 'SonarScanner', type: 'hudson.plugins.sonar.SonarRunnerInstallation'
        NODE_HOME = tool name: 'Node20.11.1', type: 'nodejs'
        ZAP_CONTAINER = "DAST_ZAP"
        REPORT_FOLDER = "${WORKSPACE}/zap_report"
    }
    stages{
        stage("Git_Checkout"){
            steps{
                git 'https://github.com/Anil-appari007/go_app.git'
            }
        }
        stage("Static_Code_Scan"){
            parallel{
                stage("SonarQube-Go"){
                    steps{
                        withSonarQubeEnv('SonarQube') {
                            sh '''
                                cd ${WORKSPACE}/go_backend && ${SONAR_HOME}/bin/sonar-scanner -Dsonar.projectKey=Go_App -Dsonar.sources=.
                            '''
                        }
                        // script{
                        //         def qg = waitForQualityGate()
                        //         if (qg.status != 'OK') {
                        //             error "Pipeline aborted due to quality gate failure: ${qg.status}"
                        //         }
                        // }
                    }
                }
                stage("SonarQube-React"){
                    steps{
                        withSonarQubeEnv('SonarQube') { 
                            sh '''
                                cd ${WORKSPACE}/frontend && ${SONAR_HOME}/bin/sonar-scanner -Dsonar.projectKey=ReactApp -Dsonar.sources=.
                            '''
                        }
                        // script{
                        //         def qg = waitForQualityGate() 
                        //         if (qg.status != 'OK') {
                        //             error "Pipeline aborted due to quality gate failure: ${qg.status}"
                        //         }
                        // }
                        
                    }
                }
            }
        }
        stage("Vulnerabilities_Scan"){
            parallel{
                stage("GO"){
                    steps{
                        sh '''
                            cd ${WORKSPACE}/go_backend
                            ${GO_HOME}/bin/go mod download
                            trivy fs .
                        '''
                    }
                }
                stage("React"){
                    steps{
                        sh '''
                            cd ${WORKSPACE}/frontend
                            export PATH="${NODE_HOME}/bin:$PATH"
                            ${NODE_HOME}/bin/npm ci
                            trivy fs .
                        '''
                    }
                }
            }
        }
        stage("Parallel_Builds"){
            parallel{
                stage("FrontEnd_Build"){
                    steps{
                        sh "cd ${WORKSPACE}/frontend && docker build -t cfrontend:${BUILD_ID} ."
                    }
                }
                stage("Backend_Buil"){
                    steps{
                        sh "cd ${WORKSPACE}/go_backend && docker build -t cbackend:${BUILD_ID} ."
                    }
                }
                stage("DB_Build"){
                    steps{
                        sh "cd ${WORKSPACE}/postgresql && docker build -t cpostgresql:${BUILD_ID} ."
                    }
                }
            }
        }
        stage("Deploy"){
            steps{

                withCredentials([string(credentialsId: 'POSTGRES_DB', variable: 'POSTGRES_DB')]) {
                    sh "cd ${WORKSPACE}/deployment && export IMAGE_TAG=${BUILD_ID} && export DB_PASSWORD=${POSTGRES_DB} && docker compose up -d"
                }
            }
        }
        stage("Integration_Tests"){
            steps{
                withCredentials([string(credentialsId: 'POSTGRES_DB', variable: 'POSTGRES_DB')]) {
                    sh '''
                        cd ${WORKSPACE}/go_backend
                        export DB_HOST=localhost
                        export DB_USER=postgres
                        export DB_PASSWORD=${POSTGRES_DB}
                        export DB_PORT=5432
                        export DB_NAME=inventory
                        ${GO_HOME}/bin/go test --tags=integration -coverprofile=coverage.out
                    '''
                }
            }
        }
        stage("DAST_ZAP"){
            steps{
                withCredentials([string(credentialsId: 'APP_URL', variable: 'APP_URL')]) {
                    sh '''
                        set +e
                        docker run -dit --name ${ZAP_CONTAINER} ghcr.io/zaproxy/zaproxy:stable bash
                        docker exec ${ZAP_CONTAINER} mkdir /zap/wrk
                        docker exec ${ZAP_CONTAINER} zap-baseline.py -r Zap_Report.html -t http://${APP_URL}:80
                        set -e
                        
                        mkdir ${REPORT_FOLDER}
                        docker exec ${ZAP_CONTAINER} ls -lrt /zap/wrk
                        docker cp ${ZAP_CONTAINER}:/zap/wrk/Zap_Report.html ${REPORT_FOLDER}/.
                    '''
                }
            }
        }
    }
    post{
        always{
            archiveArtifacts artifacts: 'zap_report/*.html'
            sh "docker rm -f ${ZAP_CONTAINER}"
            cleanWs()
        }
    }
}