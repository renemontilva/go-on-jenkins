pipeline {

 agent any
 
 stages {
   stage("build") {
	environment {
		GO111MODULES = 'on'
	}

	steps {
         sh "apt install -y golang"
	 echo "Hello world"
	}

   }
 }

}
