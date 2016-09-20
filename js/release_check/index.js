var SRError = require('@semantic-release/error');

module.exports = function(pluginConfig, config, cb) {
  var env = config.env;
  if (!env.JENKINS) {
    return cb(new SRError('not running on Jenkins, so a new version won\'t be published.', 'ENOJENKINS'));
  }

  var branch = env.GIT_BRANCH;
  if (branch !== "origin/master") {
    return cb(new SRError('branch is ' + branch + ', not publishing.', 'ENOMASTER'));
  }

  cb(null);
};
