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
        }
	
   }
   stage('Code Coverage Analyst') {
       tools {
	sonarqube 'sonarqube4'
       }
       steps {
       		withSonarQubeEnv('SonarQube') {
			sh 'sonar-scanner'
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
