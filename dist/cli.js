#!/usr/bin/env node
"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const child_process_1 = require("child_process");
const path_1 = __importDefault(require("path"));
const process_1 = require("process");
const [, , cmd, pkg, ver] = process_1.argv;
const goBinary = path_1.default.resolve(__dirname, ".../build/fastpm.exe");
function runGo(args) {
    const result = (0, child_process_1.spawnSync)(goBinary, args, { stdio: "inherit" });
    if (result.status !== 0)
        process.exit(result.status);
}
switch (cmd) {
    case "install":
        if (!pkg) {
            console.error("Package required");
            process.exit(1);
        }
        runGo([`-cmd=install`, `-pkg=${pkg}`, `-ver=${ver || "latest"}`]);
        break;
    case "remove":
        if (!pkg) {
            console.error("Package required");
            process.exit(1);
        }
        runGo([`-cmd=remove`, `-pkg=${pkg}`]);
        break;
    default:
        console.log("Commands: install <pkg>, remove <pkg>");
}
//# sourceMappingURL=cli.js.map
