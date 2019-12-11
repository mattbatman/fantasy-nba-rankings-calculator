const {
  createFantasyPointsByAmountAndCategory,
  assignPointsForCategory
} = require('../../calculator/create-fantasy-points-by-amount-and-category');

describe('assignPointsForCategory', () => {
  it('should return points due for each amount in a category', () => {
    const data = [
      {
        PTS: 1
      },
      {
        PTS: 5
      },
      {
        PTS: 10
      },
      {
        PTS: 10
      }
    ];

    const expected = {
      '1': {
        duePer: 1,
        dueList: [1]
      },
      '5': {
        duePer: 2,
        dueList: [2]
      },
      '10': {
        duePer: 3.5,
        dueList: [3, 4]
      }
    };

    const actual = assignPointsForCategory({
      sortedData: data,
      category: 'PTS'
    });

    expect(actual).toEqual(expected);
  });
});

describe('createFantasyPointsByAmountAndCategory', () => {
  it('should return points due for each category', () => {
    const data = [
      {
        PTS: 10,
        TOV: 3.5
      },
      {
        PTS: 5,
        TOV: 3.5
      },
      {
        PTS: 10,
        TOV: 1
      }
    ];

    const actual = createFantasyPointsByAmountAndCategory({
      playerData: data,
      categories: ['PTS', 'TOV']
    });

    const expected = {
      valueToPointsData: {
        PTS: {
          '5': {
            duePer: 1,
            dueList: [1]
          },
          '10': {
            duePer: 2.5,
            dueList: [2, 3]
          }
        },
        TOV: {
          '3.5': {
            duePer: 1.5,
            dueList: [1, 2]
          },
          '1': {
            duePer: 3,
            dueList: [3]
          }
        }
      },
      playerData: data,
      categories: ['PTS', 'TOV']
    };

    expect(actual).toEqual(expected);
  });
});
