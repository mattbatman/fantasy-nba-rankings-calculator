# Fantasy NBA Rankings Calculator in Python

## Getting Started

To use this, you'd need to have a Google Sheet set up with statistical NBA data, a registered application in the Google developers console from which you'd need to download an OAuth 2.0 Client ID as `sheets_credentials.json`, and to copy `config_default.py` to `config.py` and adjust any settings as needed.

## Running

```
python main.py
```

## Testing

```
python -m unittest tests.test_convert_stats_to_numbers
python -m unittest tests.test_assign_points_to_categories
python -m unittest tests.test_fantasy_points_per_player
```
