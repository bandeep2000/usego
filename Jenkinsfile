pipeline {
    agent {
        docker {
            image 'golang:latest' // Using the latest Go Docker image
            args '-u root' // Optionally, run as root if necessary
        }
    }
    
    stages {
        stage('Build') {
            steps {
                echo 'Starting the build process...'
                sh 'go version' // Checking Go version as an example command
                // Add more Go build commands here
                echo 'Build completed!'
            }
        }
        // Add more stages or steps as needed
    }
    
    // Post-build actions or additional configurations can be added here
}
