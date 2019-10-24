siatestjs = require("./siatestjs.js")



tg = siatestjs.newGroup("testgroup",1,1,1)

console.log(len(tg.Hosts()))
console.log(len(tg.Renters()))
console.log(len(tg.Miners()))