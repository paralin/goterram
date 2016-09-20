node {
  stage ("node v6") {
    sh '. /root/.bashrc && set +x && nvm install 6'
  }

  stage ("scm") {
    checkout scm
  }

  wrap([$class: 'AnsiColorBuildWrapper', 'colorMapName': 'XTerm']) {
    stage ("deps") {
      sh '''
        #!/bin/bash
        . /root/.bashrc
        git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
        glide install
      '''
    }
  }
}
