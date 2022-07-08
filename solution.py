from dataclasses import dataclass

import sys
import os
import csv

# Ideally these should be in models file
@dataclass
class Rates:
    lower: int
    upper: int
    min_cost: int
    cost_per_gb: int


@dataclass
class Theatre:
    threater_id: str
    rates: list[Rates]


@dataclass
class Partner:
    theatres: list[Theatre]
    partner_id: str


def get_parteners(partner_file_path: str) -> list[Partner]:
    """
    Process csv file to put partner object in memory

    Args:
        partner_file_path (str): location of partner csv

    Returns:
        list[Partner]: In memory representation of partners
    """
    partners = []
    with open(partner_file_path) as csvfile:
        file_reader = csv.reader(csvfile, delimiter=",")
        header = next(file_reader, None)  # skip header file
        for row in file_reader:

            theatre, slab, min_cost, cost_per_gb, partner_id = row
            # clean up data
            theatre = theatre.strip()
            partner_id = partner_id.strip()

            partner: Partner = None
            for p in partners:
                if p.partner_id == partner_id:
                    partner = p
                    break

            if not partner:
                partner = Partner(partner_id=partner_id, theatres=[])
                partners.append(partner)

            lower = int(slab.split("-")[0])
            upper = int(slab.split("-")[1])
            rate = Rates(
                lower=lower,
                upper=upper,
                min_cost=int(min_cost),
                cost_per_gb=int(cost_per_gb),
            )
            if not partner.theatres:
                theatre = Theatre(threater_id=theatre, rates=[rate])
                partner.theatres = [theatre]
            else:
                th = None
                for t in partner.theatres:

                    if t.threater_id == theatre:
                        th = t
                        break

                if not th:
                    th = Theatre(threater_id=theatre, rates=[rate])
                    partner.theatres.append(th)
                else:
                    th.rates.append(rate)
    return partners


def get_min_cost_partner(size: int, theater: str, partners: list[Partner]):
    """
    Finds partner with min cost for given size

    Args:
        size (int): size of data (in gb)
        theater (str): theather where we want to go
        partners (list[Partner]): list of partners to look at

    Returns:
        tupe(int, str): returns min cost and partner id
    """
    min_cost = sys.maxsize
    min_partner_id = None
    for partner in partners:
        # for each partners, loop over all theathers
        # if we find required theater, compute cost and see if it is min
        for p_theather in partner.theatres:
            cost = sys.maxsize
            if p_theather.threater_id == theater:

                for rate in p_theather.rates:
                    p_cost = sys.maxsize

                    if size >= rate.lower and size <= rate.upper:
                        # will it be min cost + rate*cost per gb or just rate*cost_per_gb?
                        p_cost = max(rate.cost_per_gb * size, rate.min_cost)

                    if p_cost < cost:
                        cost = p_cost

            if cost < min_cost:
                # got min cost
                min_cost = cost
                min_partner_id = partner.partner_id

    # if min_cost is still sys.maxsize then we didn't find it, mark it as none
    if min_cost == sys.maxsize:
        min_cost = None
    return (min_cost, min_partner_id)


def main():
    input_file_path = sys.argv[1]
    partner_file_path = f"{os.getcwd()}/partners.csv"
    output_file_path = f"{os.getcwd()}/output.csv"

    partners: list[Partner] = get_parteners(partner_file_path=partner_file_path)
    # both input and output files are opened together
    # maybe we can store the data in memory if size is small
    # current implementation should work fine for large sizes
    with open(input_file_path) as input_csvfile:
        input_file_reader = csv.reader(input_csvfile, delimiter=",")
        with open(
            output_file_path,
            "w",
        ) as output_csvfile:
            writer = csv.writer(output_csvfile, delimiter=",")

            for row in input_file_reader:
                delivery_id, size, theater = row
                # clean up input data for white spaces
                delivery_id = delivery_id.strip()
                size = int(size)
                theater = theater.strip()

                min_cost, partner_id = get_min_cost_partner(size, theater, partners)

                if not min_cost:
                    writer.writerow([delivery_id, "false", "", ""])
                else:
                    writer.writerow([delivery_id, "true", partner_id, min_cost])

    print(f"output written in {output_file_path}")


if __name__ == "__main__":
    main()
