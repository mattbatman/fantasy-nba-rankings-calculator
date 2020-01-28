import json
import os.path

from assign_points_to_categories import create_points_data
from config import spreadsheet, current_season, stat_categories, rankings_dir
from convert_stats_to_numbers import convert_stats_to_numbers
from fantasy_points_per_player import create_fantasy_points_per_player
from fetch_sheets_data import fetch_data


def main():
    """Fetches data from a Google Sheet, calculates player rankings, and writes
    the rankings to a JSON file.
    """
    sheets_data_file = spreadsheet['data_output_dir'] + \
        spreadsheet['data_output_file_name'] + '.json'

    fetch_data(spreadsheet['data_range'], sheets_data_file)

    with open(sheets_data_file) as data_file:
        raw_data = json.load(data_file)
        season_data = [d for d in raw_data if d['SEASON'] == current_season]
        formatted_data = convert_stats_to_numbers(season_data, stat_categories)
        points_due_per_value = create_points_data(
            formatted_data, stat_categories)
        fantasy_points = create_fantasy_points_per_player(
            formatted_data, points_due_per_value, stat_categories)

        try:
            with open('{}/{}-{}.json'.format(rankings_dir, spreadsheet['data_output_file_name'], current_season), 'w') as outfile:
                json.dump(fantasy_points, outfile)
        except FileNotFoundError:
            os.makedirs(rankings_dir, exist_ok=True)
            with open('{}/{}-{}.json'.format(rankings_dir, spreadsheet['data_output_file_name'], current_season), 'w') as outfile:
                json.dump(fantasy_points, outfile)


if __name__ == '__main__':
    main()
