import { runGo } from "../utils.js";

export function remove(pkg: string) {
    runGo(["-cmd=remove", `-pkg=${pkg}`]);
}
