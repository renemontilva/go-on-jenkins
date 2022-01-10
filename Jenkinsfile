pipeline {

 agent any

 tools {
  go 'go_17'
  sonarQube 'sonarqube4'

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

       steps {
	sh './binary/.app'
       }
   }
 }

 post {
  	success {
  		archiveArtifacts(artifacts: 'binary/*', fingerprint: true, followSymlinks: false)
  	}
 }

}
