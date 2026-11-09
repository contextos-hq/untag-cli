#!/usr/bin/env node
const { spawn } = require("child_process");
const path = require("path");

const binName = process.platform === "win32" ? "fl.exe" : "fl";
const bin = path.join(__dirname, "bin", binName);
const child = spawn(bin, process.argv.slice(2), { stdio: "inherit" });

child.on("exit", (code, signal) => {
  if (signal) {
    process.kill(process.pid, signal);
    return;
  }
  process.exit(code == null ? 1 : code);
});
child.on("error", (err) => {
  console.error(err.message);
  process.exit(1);
});
