exports.paddingRight = (val, char, length) => {
  for (; val.length < length; val += char);
  return val;
};
