"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
const GetTransactions_1 = require("./GetTransactions");
const UpdateParsedTx_1 = require("./UpdateParsedTx");
function main() {
    return __awaiter(this, void 0, void 0, function* () {
        setInterval(() => __awaiter(this, void 0, void 0, function* () {
            let signatures = yield (0, GetTransactions_1.checkActivity)();
            yield (0, UpdateParsedTx_1.extractAndUpdateData)(signatures);
        }), 120000);
    });
}
main();
