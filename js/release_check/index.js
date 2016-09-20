var SRError = require('@semantic-release/error');

module.exports = function(pluginConfig, config, cb) {
  var env = config.env;
  if (!env.JENKINS) {
    return cb(new SRError('not running on Jenkins, so a new version won\'t be published.', 'ENOJENKINS'));
  }

  var should_release = env.JENKINS_SHOULD_RELEASE;
  if (should_release !== "ismaster") {
    return cb(new SRError('branch isn\'t master, not publishing.', 'ENOMASTER'));
  }

  cb(null);
};
