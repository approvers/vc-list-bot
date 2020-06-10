exports.selectEmoji = (emojis) => {
  let random,
    candidate = [];
  candidate.push(
    emojis.map((emoji) => {
      let makeEmojiStyle = "<:" + emoji.name + ":" + emoji.id + ">";
      return makeEmojiStyle;
    })
  );
  if (candidate[0].length > 0) {
    random = Math.floor(Math.random() * candidate[0].length);
    return candidate[0][random];
  } else {
    return "";
  }
};
