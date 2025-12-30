"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.remove = remove;
const utils_js_1 = require("../utils.js");
function remove(pkg) {
    (0, utils_js_1.runGo)(["-cmd=remove", `-pkg=${pkg}`]);
}
//# sourceMappingURL=remove.js.map
