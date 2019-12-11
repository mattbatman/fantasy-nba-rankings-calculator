const stringsToFloats = require('../../calculator/strings-to-floats');
const { statCats } = require('../../config');

describe('stringsToFloats', () => {
  it('should convert values as strings to floats for the given keys', () => {
    const data = [
      {
        'FG%': '50',
        'FT%': '51',
        '3PM': '3.9',
        PTS: '15.5',
        REB: '2.0',
        AST: '7.0',
        STL: '1.5',
        BLK: '0.5',
        TOV: '2.5'
      }
    ];

    const expected = {
      playerData: [
        {
          'FG%': 50,
          'FT%': 51,
          '3PM': 3.9,
          PTS: 15.5,
          REB: 2,
          AST: 7,
          STL: 1.5,
          BLK: 0.5,
          TOV: 2.5
        }
      ],
      categories: statCats
    };

    const actual = stringsToFloats({ playerData: data, categories: statCats });

    expect(actual).toEqual(expected);
  });
});
