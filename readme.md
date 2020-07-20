To execute the file please replace the path where the derivatives json file is stored (on line 44)

Once that is replaced, pn the terminal type python exposure_calculation.py

The expected output is a number

# This program has been written to calculate the Exposure of Default
# when the neeting set of derivatives is given
# It takes as input the derivatives json data file
# and outputs the EAD in a json format
# The document referred is  - The standardised
# approach for measuring counterparty credit risk exposures - March 2014


# EAD is calculated by the formula -
# EAD = alpha * (RC + PFE)
# where:
# alpha = 1.4,
# RC = the replacement cost calculated 
# PFE = the amount for potential future exposure calculated 
# PFE = multiplier * AddOn ^ aggregate