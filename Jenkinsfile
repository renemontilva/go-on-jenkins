pipeline {

 agent any

 tools {
  go 'go_17'
 }
 
 stages {
   stage("test") {
        steps {
		sh "go test"
	}
   }
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
	scannerHome = tool name: 'sonar-scanner', type: 'hudson.plugins.sonar.SonarRunnerInstallation'

       }

       steps {
		withSonarQubeEnv('sonarqube') {
			sh "env"
			sh "${scannerHome}/bin/sonar-scanner"
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
