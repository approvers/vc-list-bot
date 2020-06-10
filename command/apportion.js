const vcSituation = require("./process/vcSituation");
const help = require("./process/help");

exports.apportion = (message) => {
  if (message.content === "!list") {
    vcSituation.vsMembers(message);
  }
  if (message.content === "!help") {
    help.help(message);
  }
  return;
};
