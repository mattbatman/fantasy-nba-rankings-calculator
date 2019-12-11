const fs = require('fs-extra');
const R = require('ramda');
const rawData = require('../data/sheets');
const convertStringsToFloats = require('./strings-to-floats');
const {
  createFantasyPointsByAmountAndCategory
} = require('./create-fantasy-points-by-amount-and-category');
const mapFantasyPointsToPlayer = require('./map-fantasy-points-to-player');
const config = require('../config');

const createRankings = R.pipe(
  convertStringsToFloats,
  createFantasyPointsByAmountAndCategory,
  mapFantasyPointsToPlayer
);

const rankings = createRankings({
  playerData: rawData,
  categories: config.statCats
});

fs.outputFile(config.ranksPath, rankings, 'utf8', error =>
  console.error(error)
);
