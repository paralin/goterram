node {
  stage ("node v6") {
    sh '''
      #!/bin/bash
      set +x
      source ~/.nvm/nvm.sh
      nvm install 6
    '''
  }

  stage ("scm") {
    checkout scm
    sh './scripts/jenkins_setup_git.bash'
  }

  wrap([$class: 'AnsiColorBuildWrapper', 'colorMapName': 'XTerm']) {
    stage ("deps") {
      sh '''
        #!/bin/bash
        ./scripts/jenkins_setup_deps.bash
      '''
    }

    stage ("js") {
      sh '''
        #!/bin/bash
        source ./scripts/jenkins_env.bash
        ./scripts/build_js.bash
        ./scripts/jenkins_release.bash
        pushd js && npm run semantic-release || true && popd
      '''
    }
  }
}
