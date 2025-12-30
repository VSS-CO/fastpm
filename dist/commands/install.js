"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.install = install;
const utils_js_1 = require("../utils.js");
function install(pkg, version = "latest") {
    (0, utils_js_1.runGo)(["-cmd=install", `-pkg=${pkg}`, `-ver=${version}`]);
}
//# sourceMappingURL=install.js.map
