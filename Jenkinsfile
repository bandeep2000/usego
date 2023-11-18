pipeline {
    agent {
        docker { image 'golang:latest' }
    }
    
    stages {
        stage('Build') {
            steps {
                echo 'Starting the build process...'
                sh 'go version' // Checking Go version as an example command
                sh 'pwd' // Checking Go version as an example command
                sh 'ls' // Checking Go version as an example command
                sh 'printenv' // Checking Go version as an example command
                dir ('greetings/fileutils') {
                   sh 'go mod init test' // Checking Go version as an example command
                   sh 'go test .' // Checking Go version as an example command
                }
                // Add more Go build commands here
                echo 'Build completed!'
            }
        }
        // Add more stages or steps as needed
    }
    
    // Post-build actions or additional configurations can be added here
}
