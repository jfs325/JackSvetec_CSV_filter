# csv-filter-challenge-public
# How-To
1. Copy the repository and load the project directory into an IDE.
2. use 'go run main.go' to start the program.

## Assumptions
1. The CSV file is manually created in the main function with rows with duplicate First Names, Last Names, and DOBs across the rows.
2. If the user is entering their DOB, I check to make sure their input is in an integer, but I assume they will not put 0s for year, day, or month and
they put the appropriate length for each part of the date. 
3. User can make their names lowercase or uppercase and the solution will still find the appropriate match if it exists.
4. If the use enters invalid input 5 times then I make a choice for them to continue the program.
