pipeline {

 agent any

 tools {
  go 'go_17'
 }
 
 stages {
   stage("build") {
	environment {
		GO111MODULES = 'on'
	}
	steps {
         sh 'go build -o binary/app main.go'
  	 archiveArtifacts(artifacts: 'binary/*', fingerprint: true, followSymlinks: false)
        }
	
   }
   stage('Code Coverage Analyst') {
       environment {
	scannerHome = tool 'sonar-scanner'

       }

       steps {
	withSonarQubeEnv('sonarqube') {
		sh "echo ${SONAR_HOST_URL}"
	        
	}
       }
   }
 }

 post {
  	success {
  		archiveArtifacts(artifacts: 'binary/*', fingerprint: true, followSymlinks: false)
  	}
 }

}
