const fs = require('fs-extra');
const R = require('ramda');
const rawData = require('../data/sheets');
const convertStringsToFloats = require('./strings-to-floats');
const {
  createFantasyPointsByAmountAndCategory
} = require('./create-fantasy-points-by-amount-and-category');
const mapFantasyPointsToPlayer = require('./map-fantasy-points-to-player');
const config = require('../config');

// createRankings is a function. It accepts the parameters expected by
// the first listed function (convertStringsToFloats), executes that function, and
// passes the returned value to the next function to continue the pipepline.
const createRankings = R.pipe(
  convertStringsToFloats,
  createFantasyPointsByAmountAndCategory,
  mapFantasyPointsToPlayer
);

const rankings = createRankings({
  playerData: rawData,
  categories: config.statCats
});

// Output the final rankings data.
fs.outputFile(config.ranksPath, JSON.stringify(rankings), 'utf8', error =>
  console.error(error)
);
