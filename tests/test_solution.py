import pandas as pd
from Solution import get_required_cost


def test_input_file() -> None:
    input_values = pd.read_csv(r"G:\Projects\challenge2019\input.csv", header=None)
    zipped: zip = zip(input_values[0].tolist(), input_values[1].tolist(), input_values[2].tolist())

    assert list(map(list, zipped))[0] == ['D1', 150, 'T1']





def test_check_min_cost() -> None:
    assert get_required_cost("T1", 150)[0] == 2000
    assert get_required_cost("T2", 325)[0] == 3250
    assert get_required_cost("T1", 510)[0] == 15300
    assert get_required_cost("T2", 700) is None


def test_check_partnerId() -> None:
    assert get_required_cost("T1", 150)[1] == "P1"
    assert get_required_cost("T2", 325)[1] == "P1"
    assert get_required_cost("T1", 510)[1] == "P3"
    assert get_required_cost("T2", 700) is None
