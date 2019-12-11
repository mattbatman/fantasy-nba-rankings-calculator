const mapFantasyPointsToPlayer = require('../../calculator/map-fantasy-points-to-player');

describe('mapFantasyPointsToPlayer', () => {
  it('should return an array of objects for each player with stats and fantasy points', () => {
    const playerData = [
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

    const categories = ['PTS', 'TOV'];

    const valueToPointsData = {
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
    };

    const expected = [
      {
        stats: {
          PTS: 10,
          TOV: 1
        },
        fantasy: {
          PTS: 2.5,
          TOV: 3,
          TOT: 5.5
        }
      },
      {
        stats: {
          PTS: 10,
          TOV: 3.5
        },
        fantasy: {
          PTS: 2.5,
          TOV: 1.5,
          TOT: 4
        }
      },
      {
        stats: {
          PTS: 5,
          TOV: 3.5
        },
        fantasy: {
          PTS: 1,
          TOV: 1.5,
          TOT: 2.5
        }
      }
    ];

    const actual = mapFantasyPointsToPlayer({
      playerData,
      categories,
      valueToPointsData
    });

    expect(actual).toEqual(expected);
  });
});
