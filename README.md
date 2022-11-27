# Construction Calendar

We bought a flat and wanted to new which parts were going to be delivered when.  The construction company gave us a number of days for each part, so I wrote a small utility that accepts a start date and then starts counting the days for each part.

Make sure to adapt both `bankholidays.txt` and `restdays.txt`, so they reflect accurate dates.    

Tranches are hardcoded for now, but if you update the `tranches`-variable with the correct slices, than you should be able to run the program with accurate values for your situation.

## How to run
1. Update `tranches` in `main.go` with correct values for your situation.
2. Update `bankholidays.txt` and `restdays.txt` with correct values for time periods that apply to your construction
3. Run the application with the following variables
   1. `startdate`: Startdate for your scenario, format `dd/MM/YYYY`.


