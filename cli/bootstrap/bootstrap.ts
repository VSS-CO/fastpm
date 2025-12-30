import fs from "fs";
import path from "path";

export function init() {
    const dir = path.resolve("./node_modules");
    if (!fs.existsSync(dir)) fs.mkdirSync(dir, { recursive: true });
}
