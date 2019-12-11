const config = {
  sheets: {
    tokenPath: 'sheets/token.json',
    credentialsPath: 'sheets/credentials.json',
    id: 'can_be_found_in_url',
    outputPath: 'data/sheets.json',
    range: 'per_game!A:AE'
  },
  ranksPath: 'data/ranks.json',
  statCats: ['FG%', 'FT%', '3PM', 'PTS', 'REB', 'AST', 'STL', 'BLK', 'TOV']
};

module.exports = config;
