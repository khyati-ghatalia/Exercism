# This program has been written to calculate the Exposure of Default
# when the neeting set of derivatives is given
# It takes as input the derivatives json data file
# and outputs the EAD in a json format
# The document referred is  - The standardised
# approach for measuring counterparty credit risk exposures - March 2014


# import the json package
import json
from datetime import datetime
import math
# EAD is calculated by the formula -
# EAD = alpha * (RC + PFE)
# where:
# alpha = 1.4,
# RC = the replacement cost calculated 
# PFE = the amount for potential future exposure calculated 
# PFE = multiplier * AddOn ^ aggregate

# Define the constants
alpha = 1.4
supervisory_factor = 0.5
RC = 0
swapsList = []
residual_maturity = []
notional_amt = []
pfe = 0
exposure = 0
swaps_start = ""
swaps_end = ""
# bucket of residual maturity is 1 year, 1-5 years, 5+ years
bucket = [1,5,6]
interest = [0.5,0.5,0.5]

# The below function returns the days between 2 dates
def days_between(d1, d2):
    d1 = datetime.strptime(d1, '%Y-%m-%d')
    d2 = datetime.strptime(d2, '%Y-%m-%d')
    return abs((d2 - d1).days)

# Read the derivatives json file
# print("Started Reading Derivatives JSON file ")
with open('/Users/itkg/Desktop/RegTech/derivatives.json', 'r') as f:
    swapsDict = json.load(f)

# calculation of exposure components
for swaps in swapsDict['data']:

    RC += swaps['mtm_dirty']
    swaps_start = swaps['start_date']
    swaps_end = swaps['end_date']
    residual_maturity.append(round(days_between(swaps_start[0:10],swaps_end[0:10])/352))
    notional_amt.append(swaps['notional_amount'])

if RC < 0 : 
    RC=0
for x, y in residual_maturity,notional_amt:
   pfe+=(math.exp(-0.05*x) * y)
exposure = alpha * (RC + pfe)
print(exposure)  


