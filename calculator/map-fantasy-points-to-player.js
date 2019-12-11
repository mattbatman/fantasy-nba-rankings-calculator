const R = require('ramda');

// mapFantasyPointsToPlayer :: ([]{} -> []String -> {}) -> []{}
// mapFantasyPointsToPlayer takes an array of objects of stats per player, a
// list of categories to calculate for, and an object mapping categories to
// fantasy points due per amount. It returns an array of objects with keys of
// "stas" and "fantasy" for each player.
const mapFantasyPointsToPlayer = ({
  playerData,
  categories,
  valueToPointsData
}) => {
  const nestFantasyPoints = R.map(v => {
    const fantasyPointsDueForPlayer = R.reduce(
      (acc, category) => {
        acc[category] = valueToPointsData[category][v[category]].duePer;

        return acc;
      },
      {},
      categories
    );

    fantasyPointsDueForPlayer.TOT = R.sum(R.values(fantasyPointsDueForPlayer));

    return {
      stats: v,
      fantasy: fantasyPointsDueForPlayer
    };
  });

  const sortByTotal = R.sort((x, y) => y.fantasy.TOT - x.fantasy.TOT);

  return R.pipe(nestFantasyPoints, sortByTotal)(playerData);
};

module.exports = mapFantasyPointsToPlayer;
