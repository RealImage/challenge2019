from typing import Optional, Dict, List, Any

import pandas as pd
from pandas import DataFrame


def check_range(input_value: int, range: str) -> bool:
    _min = int(range.split("-")[0])
    _max = int(range.split("-")[1])

    if _min <= input_value <= _max:
        return True
    return False


def get_required_cost(theatre_id: str, delivery_size: int) -> Optional[List]:
    partners: Any = pd.read_csv(r"partners.csv")
    min_cost: int | float = float("+inf")
    partner_id: str = ""
    for _, partner in partners.iterrows():
        if partner.Theatre.strip() == theatre_id.strip() and check_range(input_value=delivery_size,
                                                                         range=partner["Size Slab (in GB)"]):

            cost_right_now: int = delivery_size * int(partner["Cost Per GB"])

            if cost_right_now <= partner["Minimum cost"]:
                cost_right_now = partner["Minimum cost"]

            if cost_right_now < min_cost:
                min_cost = cost_right_now
                partner_id = partner["Partner ID"]

    if min_cost == float("+inf"):
        return None
    return [min_cost, partner_id]


if __name__ == "__main__":

    input_values: Any = pd.read_csv(r"input.csv", header=None)
    unpacked_values: Any = input_values.values.tolist()

    result_list: List = []
    for delivery_id, size, theatre_id in unpacked_values:

        min_cost: Optional[List] = get_required_cost(theatre_id=theatre_id, delivery_size=size)
        delivery_possible: bool = True
        if not min_cost:
            delivery_possible = False
            min_cost = ["''", "''"]

        data: Dict = {
            "delivery ID": delivery_id,
            "Delivery Possible": delivery_possible,
            "Selected Partner": min_cost[1],
            "Cost of Delivery": min_cost[0]
        }

        result_list.append(data)

    result: DataFrame = pd.DataFrame(result_list)

    result.to_csv("result.csv", header=False, index=False)
