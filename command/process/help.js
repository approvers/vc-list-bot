// コマンドのデータ
const helpData = require("./assets/helpData");
const paddingRight = require("./assets/makeBleakSpace").paddingRight;

exports.help = (message) => {
  message.channel.send(
    ":rocket: | The list of commands you have access to has been sent to your DMs."
  );
  //   DMにコマンドリストを送信
  message.author.send(returnCommandList());
};

// コマンドリストをDMに送った時に見やすい形式にする。
const returnCommandList = () => {
  let helpList = "**= Commands List =**\n";
  helpData.data.commandList.map((commandList) => {
    helpList += "**" + commandList.title + "**:\n```asciidoc\n";
    commandList.contents.map((contents) => {
      helpList += "= " + contents.subTitle + " =\n";
      contents.command.map((command, i) => {
        helpList += "!" + paddingRight(command, " ", 9);
        helpList += ":: " + contents.explanation[i] + "\n";
      });
      helpList += "```";
    });
  });
  return helpList;
};
