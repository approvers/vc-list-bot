const paddingRight = require("./assets/makeBleakSpace").paddingRight;
const selectEmoji = require("./assets/function").selectEmoji;
const textLength = 8;

exports.vsMembers = (message) => {
  const guild = message.guild;
  let utterance = returnMentionContents(guild);
  message.channel.send(utterance);
};

const returnMentionContents = (guild) => {
  let emojis = guild.emojis.cache;
  let everyMembers = guild.memberCount;
  let membersStatus = countMembers(guild);
  let numberOfMember = membersStatus.numberOfMember;
  let numberOfMuteMember = membersStatus.numberOfMuteMember;

  let utterance =
    selectEmoji(emojis) + " ***限界リスト*** " + selectEmoji(emojis) + "\n";
  utterance += "```asciidoc\n= 現在の状況 =\n";
  utterance += numberOfAllMembers(everyMembers) + "\n";
  utterance += numberOfMembers(numberOfMember) + "\n";
  utterance += numberOfMute(numberOfMuteMember) + "\n";
  utterance += callRate(everyMembers, numberOfMember) + "\n";
  utterance += muteRate(numberOfMuteMember, numberOfMember) + "```\n";

  if (numberOfMember == 0) {
    utterance = "し〜ん...";
  }
  return utterance;
};

const countMembers = (guild) => {
  let numberOfMember = 0;
  let numberOfMuteMember = 0;
  guild.voiceStates.cache.map((members) => {
    numberOfMember++;
    if (members.selfMute == true) numberOfMuteMember++;
  });
  return { numberOfMember, numberOfMuteMember };
};

const numberOfAllMembers = (members) => {
  return paddingRight("鯖人数", "　", textLength) + ":: " + members + " 人";
};

const numberOfMembers = (members) => {
  return paddingRight("通話人数", "　", textLength) + ":: " + members + " 人";
};

const numberOfMute = (numberOfMuteMember) => {
  return (
    paddingRight("ミュート人数", "　", textLength) +
    ":: " +
    numberOfMuteMember +
    " 人"
  );
};

const callRate = (allMembers, numberOfMember) => {
  return (
    paddingRight("通話率", "　", 8) +
    ":: " +
    Math.round((numberOfMember / allMembers) * 1000) / 10 +
    " %"
  );
};

const muteRate = (numberOfMuteMember, numberOfMember) => {
  return (
    paddingRight("ミュート率", "　", 8) +
    ":: " +
    Math.round((numberOfMuteMember / numberOfMember) * 1000) / 10 +
    " %"
  );
};
