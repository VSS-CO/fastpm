"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.runGo = runGo;
const child_process_1 = require("child_process");
const path_1 = __importDefault(require("path"));
function runGo(args) {
    const goBinary = path_1.default.resolve(__dirname, "../../build/fastpm");
    const result = (0, child_process_1.spawnSync)(goBinary, args, { stdio: "inherit" });
    if (result.status !== 0)
        process.exit(result.status);
}
//# sourceMappingURL=utils.js.map
