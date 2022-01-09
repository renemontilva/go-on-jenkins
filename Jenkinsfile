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
         sh 'go build main.go'
        }
	
   }
   stage('Run Build') {

       steps {
       	sh './main' 
       }
   }

 }

}
