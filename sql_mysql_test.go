package main_test

import "testing"

func TestRevisingTheSelectQuery1(t *testing.T) {
	// select * from City where CountryCode = "USA" and Population > 100000
}

func TestRevisingTheSelectQuery2(t *testing.T) {
	// select name from City where CountryCode = "USA" and Population > 120000
}

func TestSelectAll(t *testing.T) {
	// select * from City
}

func TestSelectByID(t *testing.T) {
	// select * from City where ID = 1661
}

func TestJapaneseCitiesAttributes(t *testing.T) {
	// select * from City where CountryCode = "JPN"
}

func TestJapaneseCitiesNames(t *testing.T) {
	// select name from City where CountryCode = "JPN"
}

func TestWeatherObservationStation1(t *testing.T) {
	// select city, state from STATION
}

func TestWeatherObservationStation2(t *testing.T) {

}

func TestWeatherObservationStation3(t *testing.T) {
	// select distinct(city) from Station where id%2 = 0
}

func TestWeatherObservationStation4(t *testing.T) {
	// select count(*) - count(distinct(city)) from Station
}

func TestWeatherObservationStation5(t *testing.T) {
	// select city, length(city) from Station order by length(city), city asc  limit 1;
	// select city, length(city) from Station order by length(city) desc limit 1;
}

func TestWeatherObservationStation6(t *testing.T) {
	// select distinct(city) from Station where city regexp '^[a,e,i,o,u]'
}

func TestWeatherObservationStation7(t *testing.T) {
	// select distinct(city) from Station where city regexp '[a,e,i,o,u]$'
}

func TestWeatherObservationStation8(t *testing.T) {
	// select city from Station where city regexp '^[a,e,i,o,u]' and city regexp '[a,e,i,o,u]$'
}

func TestWeatherObservationStation9(t *testing.T) {
	// select distinct(city) from Station where city not regexp '^[a,e,i,o,u]'
}

func TestWeatherObservationStation10(t *testing.T) {
	// select distinct(city) from Station where city not regexp '[a,e,i,o,u]$'
}

func TestWeatherObservationStation11(t *testing.T) {
	// select distinct(city) from Station where city not regexp '[a,e,i,o,u]$' or city not regexp '^[a,e,i,o,u]'
}

func TestWeatherObservationStation12(t *testing.T) {
	// select distinct(city) from Station where city not regexp '[a,e,i,o,u]$' and city not regexp '^[a,e,i,o,u]'
}

func TestHigherThan75Marks(t *testing.T) {
	// select name from Station where marks > 75 order by substring(name, -3), id
}

func TestEmployeeNames(t *testing.T) {
	// select name from Employee order by name
}

func TestEmployeeSalaries(t *testing.T) {
	// select name from Employee where months < 10 and salary > 2000 order by employee_id
}

func TestTypeOfTriangle(t *testing.T) {
	// select
	// case
	//  when a+b > c and b+c > a and a+c > b then
	// 		case
	// 			when a = b and b = c         then 'Equilateral'
	// 			when a = b or  b = c or a =c then 'Isosceles'
	// 			else 'Scalene'
	// 		end
	// 	else 'Not A Triangle'
	// end
	// from Triangles
}

func TestRevisingAggregationsSum(t *testing.T) {
	// select sum(population) from City where district = 'California'
}

func TestRevisingAggregationsAvg(t *testing.T) {
	// select avg(population) from city where district = 'California'
}

func TestAveragePopulation(t *testing.T) {
	// select round(avg(population)) from City
}

func TestJapanPopulation(t *testing.T) {
	// select sum(population) from City where countrycode = 'JPN'
}

func TestPopulationDensityDifference(t *testing.T) {
	// select max(population) - min(population) from City
}

func TestTheBlunder(t *testing.T) {
	// select ceil( avg(salary) - avg( replace(salary, '0', '' ) ) ) from Employees;
}
