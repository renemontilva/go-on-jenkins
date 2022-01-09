pipeline {

 agent any
 
 stages {
   stage("build") {
	environments {
		GO111MODULES=on
	}

	steps {
         sh "apt install -y golang"
	 echo "Hello world"
	}

   }
 }

}
