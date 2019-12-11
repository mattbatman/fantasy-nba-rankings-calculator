const fs = require('fs-extra');
const { google } = require('googleapis');
const R = require('ramda');
const config = require('../config');

// processData :: [][]String -> []{}
// processData takes sheets data (array of arrays of strings) and returns an
// array of objects (JSON) with the first array in the sheets data used as the
// keys in each object of the return value.
function processData(data) {
  // dataKeys :: []string
  // The first row of the sheets data will be the headers of the columns.
  // shift() will remove (mutate!) the first element from the sheets data and
  // store the value in dataKeys.
  const dataKeys = data.shift();

  return R.map(d => R.zipObj(dataKeys, d), data);
}

// writeData :: Record -> Void
// writeData takes a Google Auth record and writes the data from the sheet to
// a json file.
function writeData(auth) {
  const sheets = google.sheets({ version: 'v4', auth });

  sheets.spreadsheets.values.get(
    {
      spreadsheetId: config.sheets.id,
      range: config.sheets.range
    },
    (err, { data }) => {
      if (err) return console.log('The API returned an error: ' + err);

      const zippedData = processData(data.values);
      const jsonData = JSON.stringify(zippedData);

      fs.outputFile(config.sheets.outputPath, jsonData, 'utf8', error =>
        console.error(error)
      );
    }
  );
}

module.exports = writeData;
