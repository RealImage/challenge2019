from typing import Optional, Dict

import pandas as pd
from pandas import DataFrame


def check_range(input_value: int, range: str) -> bool:
    _min = int(range.split("-")[0])
    _max = int(range.split("-")[1])

    if _min <= input_value <= _max:
        return True
    return False


def get_required_cost(theatre_id: str, delivery_size: int) -> Optional[int]:
    partners = pd.read_csv(r"G:\Projects\challenge2019\partners.csv")
    min_cost: int | float = float("+inf")
    for _, partner in partners.iterrows():
        if partner.Theatre.strip() == theatre_id.strip() and check_range(input_value=delivery_size,
                                                                         range=partner["Size Slab (in GB)"]):
            cost_right_now: int = delivery_size * int(partner["Cost Per GB"])
            if cost_right_now <= partner["Minimum cost"]:
                cost_right_now = partner["Minimum cost"]

            min_cost = min(cost_right_now, min_cost)
    if min_cost == float("+inf"):
        return None
    return min_cost


if __name__ == "__main__":

    input_values = pd.read_csv(r"G:\Projects\challenge2019\input.csv", header=None)

    zipped: zip = zip(input_values[0].tolist(), input_values[1].tolist(), input_values[2].tolist())
    unpacked_values = list(map(list, zipped))

    result_list = []
    for delivery_id, size, theatre_id in unpacked_values:

        min_cost: Optional[int] = get_required_cost(theatre_id=theatre_id, delivery_size=size)
        possible: bool = True
        if not min_cost:
            possible = False

        data: Dict = {
            "delivery ID": delivery_id,
            "Delivery Possible": possible,
            "Selected Partner": "TBD",
            "Cost of Delivery": min_cost
        }

        result_list.append(data)

    result: DataFrame = pd.DataFrame(result_list)

    result.to_csv("result.csv", header=False, index=False)