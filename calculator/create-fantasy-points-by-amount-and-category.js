const R = require('ramda');

const ascendingByKey = key => R.ascend(R.prop(key));
const descendingByKey = key => R.descend(R.prop(key));
const sortData = (data, category, direction) =>
  R.sort(direction(category), data);

// createFantasyPointsByAmountAndCategory :: ([]{}, []String) -> {}
// createFantasyPointsByAmountAndCategory takes an array of player JSON data
// and a list of categories to calculate the fantasy points due per category
// for each amount. It returns an object with a key for each category. Each
// category points to another object with a key for each unique amount from all
// players that points to the fantasy points due for that amount and a list
// of points used to calculate the points due.
//
// The return value might look like:
// {
//   valueToPointsData: {
//     PTS: {
//       '5': {
//         duePer: 1,
//         dueList: [ 1 ]
//       },
//       '10': {
//         duePer: 2.5,
//         dueList: [ 2, 3 ]
//       }
//     },
//     TOV: {
//       '3.5': {
//         duePer: 1.5,
//         dueList: [ 1, 2 ]
//       },
//       '1': {
//         duePer: 3,
//         dueList: [ 3 ]
//       }
//     }
//   },
//   categories: ['PTS', 'TOV'],
//   playerData: [
//     {
//       PTS: 10,
//       TOV: 3.5
//     },
//     {
//       PTS: 5,
//       TOV: 3.5
//     },
//     {
//       PTS: 10,
//       TOV: 1
//     }
//   ]
// }

const createFantasyPointsByAmountAndCategory = ({
  playerData,
  categories
}) => ({
  valueToPointsData: R.reduce(
    (acc, category) => {
      const isTov = category === 'TOV';
      const sortDirection = isTov ? descendingByKey : ascendingByKey;
      data = sortData(playerData, category, sortDirection);

      acc[category] = assignPointsForCategory({ sortedData: data, category });

      return acc;
    },
    {},
    categories
  ),
  categories,
  playerData
});

// assignPointsForCategory :: ([]{}, String) -> {}
// assignPointsForCategory takes a sorted array of player data and the category
// fantasy points to calculate, and returns an object with the amount from each
// player at that category as a key that points to an object containing the
// fantasy points due for that amount as a list and calculated total accounting
// for ties.
//
// The return value might look like:
// {
//   '1': {
//     duePer: 1,
//     dueList: [ 1 ]
//   },
//   '5': {
//     duePer: 2,
//     dueList: [ 2 ]
//   },
//   '10': {
//     duePer: 3.5,
//     dueList: [ 3, 4 ]
//   }
// };

const assignPointsForCategory = ({ sortedData, category }) =>
  sortedData.reduce((acc, playerData, i) => {
    const amount = playerData[category];
    const pointsForAmount = i + 1;
    const isTie = !R.isNil(acc[amount]);

    if (!isTie) {
      acc[amount] = {
        duePer: pointsForAmount,
        dueList: [pointsForAmount]
      };

      return acc;
    }

    acc[amount].dueList.push(pointsForAmount);
    acc[amount].duePer = R.divide(
      R.sum(acc[amount].dueList),
      R.length(acc[amount].dueList)
    );

    return acc;
  }, {});

module.exports = {
  createFantasyPointsByAmountAndCategory,
  assignPointsForCategory
};
