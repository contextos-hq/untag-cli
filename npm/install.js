const { execFileSync } = require("child_process");
const fs = require("fs");
const path = require("path");

const version = require("../package.json").version;
const platform = { darwin: "darwin", linux: "linux", win32: "windows" }[process.platform];
const arch = { x64: "amd64", arm64: "arm64" }[process.arch];

if (!platform || !arch) {
  console.error(`untag-cli does not publish a binary for ${process.platform}/${process.arch}`);
  process.exit(1);
}

const ext = platform === "windows" ? "zip" : "tar.gz";
const asset = `fl_${platform}_${arch}.${ext}`;
const url = `https://github.com/contextos-hq/untag-cli/releases/download/v${version}/${asset}`;
const dest = path.join(__dirname, "bin");
const archive = path.join(dest, asset);
const binName = platform === "windows" ? "fl.exe" : "fl";

async function main() {
  fs.mkdirSync(dest, { recursive: true });
  const response = await fetch(url);
  if (!response.ok) {
    console.error(`untag-cli: could not download ${url} (${response.status})`);
    process.exit(1);
  }
  fs.writeFileSync(archive, Buffer.from(await response.arrayBuffer()));
  const tarArgs = ext === "zip" ? ["-xf", archive, "-C", dest] : ["-xzf", archive, "-C", dest];
  execFileSync("tar", tarArgs, { stdio: "inherit" });
  fs.unlinkSync(archive);
  const bin = path.join(dest, binName);
  if (!fs.existsSync(bin)) {
    console.error(`untag-cli: ${binName} was not in ${asset}`);
    process.exit(1);
  }
  if (platform !== "windows") fs.chmodSync(bin, 0o755);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
