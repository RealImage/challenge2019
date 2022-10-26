import pandas as pd
inputFile = "input.csv"
partnersFile = "partners.csv"
outputFile = "output1.csv"

def storePartnersData(partners):
  distList = dict()
  for each in partners:
    if (each[0].strip() not in distList):
      distList[each[0].strip()] = dict()
    keys = tuple([int(val) for val in each[1].strip().split("-")])
    distList[each[0].strip()][keys] = {
        "minCost": each[2],
        "perGB": each[3],
        "partner": each[4]
    }
  return distList

def filteredMinimumCostForInputList(value,optlist):
  minKey = None
  minCost = 0
  for each in optlist:
    if (each[0]<=value and each[1]>=value):
      costTemp = max(optlist[each]["perGB"]*value,optlist[each]["minCost"])
      minCost = costTemp if minCost==0 else min(costTemp,minCost)
      minKey = each if costTemp==minCost else minKey
  if minKey:
    return [True, optlist[minKey]["partner"], minCost]
  else:
    return [False, None, None]

def calculateCostForInputList(inputs,distList):
  inputList = list()
  for each in inputs:
    temp = list()
    temp.append(each[0])
    temp += filteredMinimumCostForInputList(each[1],distList[each[2]])
    inputList.append(temp)
  return inputList

if __name__ == '__main__':
  df = pd.read_csv(partnersFile)
  partners = df.values.tolist()
  df = pd.read_csv(inputFile,header=None, names=['DeliveryId','Data','Theatre'])
  inputs = df.values.tolist()
  distList = storePartnersData(partners)
  minimumCost = calculateCostForInputList(inputs,distList)
  df = pd.DataFrame(minimumCost, columns =['DeliveryId', 'Possibility', 'Partner', 'MinimumCost'])
  df.to_csv(outputFile)
