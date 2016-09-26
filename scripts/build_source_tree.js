const child_process = require('child_process');
const path = require('path');
const fs = require('fs');
const rimraf = require('rimraf');

function fileExists(path) {
  try {
    let ls = fs.lstatSync(path);
    return ls.isFile();
  } catch (e) {
    return false;
  }
}

// Adds the Go search paths to a array.
function resolveGoPaths(dirs) {
  var goRoot;
  var goPath;
  try {
    goPath = child_process.execSync("go env GOPATH", {encoding: 'utf8', env: process.env}).trim() + "/src";
    dirs.push(goPath);
    dirs.push(goPath + "/github.com/gopherjs/gopherjs/compiler/natives/src")
    dirs.push(goPath + "/github.com/gopherjs/gopherjs/nosync")
  } catch (e) {
    console.log("Error determining gopath: " + e);
    console.log("Your source maps may not show up properly.")
  }
  try {
    goRoot = child_process.execSync("go env GOROOT", {encoding: 'utf8', env: process.env}).trim() + "/src";
    dirs.push(goRoot);
  } catch (e) {
    console.log("Error determining goroot: " + e);
    console.log("Your source maps may not show up properly.")
  }
}

var goPaths = [];
resolveGoPaths(goPaths);
if (!goPaths.length) {
  console.log("Unable to determine any go paths, can't build source tree.");
  process.exit(1);
}

var sourceMap;
try {
  sourceMap = JSON.parse(fs.readFileSync("./js/goterram.js.map", 'utf8'));
} catch (e) {
  console.log('Error reading source map: ' + e);
  process.exit(1);
}

if (sourceMap.sources[0].startsWith("..")) {
  console.log("Source map already transformed, checking for backup.");
  if (!fileExists('./js/goterram.js.map.bak')) {
    console.log("Backup doesn't exist, exiting.");
    process.exit(1);
  }
  try {
    sourceMap = JSON.parse(fs.readFileSync("./js/goterram.js.map.bak", 'utf8'));
  } catch (e) {
    console.log('Error reading source map backup: ' + e);
    process.exit(1);
  }
}

fs.writeFileSync("./js/goterram.js.map.bak", JSON.stringify(sourceMap), 'utf8');

// Transform some paths in the sourcemap.
for (var i = 0; i < sourceMap.sources.length; i++) {
  var smp = sourceMap.sources[i];
  // A couple special cases
  if (smp === "atomic.go") {
    sourceMap.sources[i] = "/sync/atomic/atomic.go";
    continue;
  }
  if (smp === "pool.go") {
    sourceMap.sources[i] = "/pool.go";
    continue;
  }

  if (smp.indexOf('/') !== -1 || !smp.endsWith(".go")) {
    continue;
  }
  smp = "/" + smp;
  var smpnogo = smp.substring(0, smp.length - 3);
  var smpnogoi = smpnogo.indexOf("_");
  if (smpnogoi !== -1) {
    smp = smpnogo.substring(0, smpnogoi) + "/" + smp;
  } else {
    smp = smpnogo + "/" + smp;
  }
  sourceMap.sources[i] = smp;
}

var outputDir = "./gosources";
try {
  rimraf.sync(outputDir);
} catch(e) {}
fs.mkdirSync(outputDir);

for (var i = 0; i < sourceMap.sources.length; i++) {
  var sourcep = sourceMap.sources[i];
  var found = false;
  var tried = [];
  for (var j = 0; j < goPaths.length; j++) {
    var gop = goPaths[j];
    var gosp = gop + sourcep;
    tried.push(gosp);
    if (fileExists(gosp)) {
      child_process.execSync('mkdir -p ' + path.dirname(outputDir + sourcep));
      child_process.execSync('cp ' + gosp + ' ' + outputDir + sourcep);
      found = true;
    }
  }
  if (!found) {
    console.log('[warn] unable to find source file ' + sourcep);
    console.log('[warn] tried: ' + tried.join(", "));
  }
}

// Finishing touches
for (var i = 0; i < sourceMap.sources.length; i++) {
  sourceMap.sources[i] = "../gosources" + sourceMap.sources[i];
  sourceMap.sources[i] = sourceMap.sources[i].replace("//", "/");
}
fs.writeFileSync("./js/goterram.js.map", JSON.stringify(sourceMap), 'utf8');
