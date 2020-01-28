from __future__ import print_function
import pickle
import os.path
from googleapiclient.discovery import build
from google_auth_oauthlib.flow import InstalledAppFlow
from google.auth.transport.requests import Request
from config import spreadsheet
import json

# If modifying these scopes, delete the file token.pickle.
SCOPES = ['https://www.googleapis.com/auth/spreadsheets.readonly']
SPREADSHEET_ID = spreadsheet['id']


def fetch_data(range, output_path):
    """This is mostly a copy and paste of the Google Sheet's API
    Python quickstart.
    """
    creds = None
    # The file token.pickle stores the user's access and refresh tokens, and is
    # created automatically when the authorization flow completes for the first
    # time.
    if os.path.exists('token.pickle'):
        with open('token.pickle', 'rb') as token:
            creds = pickle.load(token)
    # If there are no (valid) credentials available, let the user log in.
    if not creds or not creds.valid:
        if creds and creds.expired and creds.refresh_token:
            creds.refresh(Request())
        else:
            flow = InstalledAppFlow.from_client_secrets_file(
                'sheets_credentials.json', SCOPES)
            creds = flow.run_local_server(port=0)
        # Save the credentials for the next run
        with open('token.pickle', 'wb') as token:
            pickle.dump(creds, token)

    service = build('sheets', 'v4', credentials=creds)

    # Call the Sheets API
    sheet = service.spreadsheets()
    result = sheet.values().get(spreadsheetId=SPREADSHEET_ID,
                                range=range).execute()
    values = result.get('values', [])

    if not values:
        print('No data found.')
    else:
        parse_and_write(values, output_path)


def parse_and_write(sheets_values, filepath):
    """Takes an array of arrays sheets_values and string filename. The first row of the array of arrays
    is expected to be the headers (first row) in the Google Sheets columns. The data is recombined
    as a dict and written as a JSON array of objects with the provided filename.
    """
    all_players = []
    keys = sheets_values[0]
    data = sheets_values[1:]

    for datum in data:
        player_data = dict(zip(keys, datum))
        all_players.append(player_data)

    try:
        with open(filepath, 'w') as outfile:
            json.dump(all_players, outfile)
    except FileNotFoundError:
        os.makedirs(spreadsheet['data_output_dir'], exist_ok=True)
        with open(filepath, 'w') as outfile:
            json.dump(all_players, outfile)
