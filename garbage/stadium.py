import sys

APrice = 15
BPrice = 12
CPrince = 9


ASeats = input('Please enter the number of Class A Seats sold:')
if ASeats < 0:
    #sys.exit() will exit if given a negative input
    print('Error: cannot have negative seats')
    sys.exit()

BSeats = input('Please enter the number of Class B Seats sold:')
if BSeats < 0:
    print('Error: cannot have negative seats')
    sys.exit()

CSeats = input('Please enter the number of Class C Seats sold:')
if CSeats < 0:
    print('Error: cannot have negative seats')
    sys.exit()

# we now have the user input for how many seats there were
# Below is the final output that is generated after the input
print('**********************************************')
print('Total number of tickets sold:',ASeats + BSeats + CSeats)
print('Income Generated:')
print('Class A: $',ASeats * APrice)
print('Class B: $',BSeats * BPrice)
print('Class C: $',CSeats * CPrice)
print('Total: ', (ASeats * APrice) + (BSeats * BPrice) + (CSeats * CPrice))
