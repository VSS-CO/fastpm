#!/usr/bin/env node
import { spawnSync } from "child_process";
import path from "path";
import { argv } from "process";

const [,, cmd, pkg, ver] = argv;
const goBinary = path.resolve(__dirname, "../build/fastpm");

function runGo(args: string[]) {
    const result = spawnSync(goBinary, args, { stdio: "inherit" });
    if (result.status !== 0) process.exit(result.status!);
}

switch(cmd) {
    case "install":
        if (!pkg) { console.error("Package required"); process.exit(1); }
        runGo([`-cmd=install`, `-pkg=${pkg}`, `-ver=${ver || "latest"}`]);
        break;
    case "remove":
        if (!pkg) { console.error("Package required"); process.exit(1); }
        runGo([`-cmd=remove`, `-pkg=${pkg}`]);
        break;
    default:
        console.log("Commands: install <pkg>, remove <pkg>");
}
