
from dataclasses import dataclass

import sys
import os
import csv


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


def main():
    input_file_path = sys.argv[1]
    partner_file_path = f'{os.getcwd()}/partners.csv'

    partners: list[Partner] = []
    with open(partner_file_path) as csvfile:
        file_reader = csv.reader(csvfile, delimiter=',')
        header = next(file_reader, None)
        for row in file_reader:

            theatre, slab, min_cost, cost_per_gb, partner_id = row
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

            lower = int(slab.split('-')[0])
            upper = int(slab.split('-')[1])
            rate = Rates(lower=lower, upper=upper, min_cost=int(
                min_cost), cost_per_gb=int(cost_per_gb))

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
                    print()
                    th.rates.append(rate)

    print(partners)


if __name__ == '__main__':
    main()
