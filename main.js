// Response for Uptime Robot
require("dotenv").config();
const env = process.env;

// Discord bot implements
const discord = require("discord.js");
const client = new discord.Client();

client.on("ready", (message) => {
  client.user.setPresence({ game: { name: "VC状況をおしえるよ！" } });
  console.log("bot is ready!");
});

const commandApportion = require("./command/apportion");

client.on("message", (message) => {
  if (!message.author.bot) {
    commandApportion.apportion(message);
  }
  return;
});

if (env.DISCORD_BOT_TOKEN == undefined) {
  console.log("please set ENV: DISCORD_BOT_TOKEN");
  process.exit(0);
}

client.login(env.DISCORD_BOT_TOKEN);
