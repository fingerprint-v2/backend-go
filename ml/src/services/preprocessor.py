import pandas as pd


class Preprocessor:
    def run(self, data: pd.DataFrame) -> pd.DataFrame:
        data = data.dropna()
        data = data.drop_duplicates()
        data = data.reset_index(drop=True)
        return data
