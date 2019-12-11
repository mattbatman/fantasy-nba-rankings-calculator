const R = require('ramda');

// convertStringsToFloat :: []{} -> []String -> []{}
// convertStringsToFloat takes an array of objects Data and an array Categories,
// uses the categories as keys in each object of the array to convert the
// corresponding value to a number, and returns a new array of objects with numbers
// where converted.
const convertStringsToFloat = ({ playerData, categories }) => {
  const convertedData = R.reduce(
    (acc, val) => {
      categories.forEach(v => {
        val[v] = parseFloat(val[v]);
      });
      acc.push(val);
      return acc;
    },
    [],
    playerData
  );

  return {
    playerData: convertedData,
    categories
  };
};

module.exports = convertStringsToFloat;