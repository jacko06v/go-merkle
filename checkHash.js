const keccak256 = require("keccak256");
const { bufferToHex } = require("ethereumjs-util");

const hexExpected =
  "0x8cc628095351985de68ca79a948c500a851880ab495a54dbee7ff84c9c9195ba";


function hashObjectJson(d) {
  const intIt = BigInt(d.int_it);
  const renting = BigInt(d.renting);
  const packed = Buffer.concat([
    Buffer.from(intIt.toString(16).padStart(64, "0"), "hex"),
    Buffer.from(renting.toString(16).padStart(64, "0"), "hex"),
  ]);
  return bufferToHex(keccak256(packed));
}

let prova = hashObjectJson({
  int_it: "631359070595067391",
  renting: "3",
});

console.log("generated hash: ", prova);
console.log("hash expected: ", hexExpected);

