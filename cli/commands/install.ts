import { runGo } from "../utils.js";

export function install(pkg: string, version = "latest") {
    runGo(["-cmd=install", `-pkg=${pkg}`, `-ver=${version}`]);
}
