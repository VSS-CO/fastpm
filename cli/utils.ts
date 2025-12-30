import { spawnSync } from "child_process";
import path from "path";

export function runGo(args: string[]) {
    const goBinary = path.resolve(__dirname, "../../build/fastpm");
    const result = spawnSync(goBinary, args, { stdio: "inherit" });
    if (result.status !== 0) process.exit(result.status!);
}
