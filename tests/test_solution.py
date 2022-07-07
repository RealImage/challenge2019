from typing import List
import pandas as pd
from Solution import get_required_cost, check_range


def test_input_file() -> None:
    input_values = pd.read_csv("../input.csv", header=None)
    zipped: zip = zip(input_values[0].tolist(), input_values[1].tolist(), input_values[2].tolist())

    assert list(map(list, zipped))[0] == ['D1', 150, 'T1']


def test_check_custom() -> None:
    assert get_required_cost("T1", 300) == 3000
    assert get_required_cost("T2", 120) == 1800
    assert get_required_cost("T1", 450) == 13500


def test_check_output_one() -> None:
    assert get_required_cost("T1", 150) == 2000
    assert get_required_cost("T2", 325) == 3250
    assert get_required_cost("T1", 510) == 15300
    assert get_required_cost("T2", 700) is None
