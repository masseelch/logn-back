// Code generated by ent, DO NOT EDIT.

package vacancy

import (
	"time"
)

const (
	// Label holds the string label denoting the vacancy type in the database.
	Label = "vacancy"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIsNegotiate holds the string denoting the is_negotiate field in the database.
	FieldIsNegotiate = "is_negotiate"
	// FieldMinSalary holds the string denoting the min_salary field in the database.
	FieldMinSalary = "min_salary"
	// FieldMaxSalary holds the string denoting the max_salary field in the database.
	FieldMaxSalary = "max_salary"
	// FieldIsRemote holds the string denoting the is_remote field in the database.
	FieldIsRemote = "is_remote"
	// FieldViews holds the string denoting the views field in the database.
	FieldViews = "views"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeTechnologies holds the string denoting the technologies edge name in mutations.
	EdgeTechnologies = "technologies"
	// EdgeLocations holds the string denoting the locations edge name in mutations.
	EdgeLocations = "locations"
	// EdgeAreas holds the string denoting the areas edge name in mutations.
	EdgeAreas = "areas"
	// EdgeCompanies holds the string denoting the companies edge name in mutations.
	EdgeCompanies = "companies"
	// EdgeTechnologyLevels holds the string denoting the technology_levels edge name in mutations.
	EdgeTechnologyLevels = "technology_levels"
	// Table holds the table name of the vacancy in the database.
	Table = "vacancies"
	// TechnologiesTable is the table that holds the technologies relation/edge. The primary key declared below.
	TechnologiesTable = "technology_levels"
	// TechnologiesInverseTable is the table name for the Technology entity.
	// It exists in this package in order to avoid circular dependency with the "technology" package.
	TechnologiesInverseTable = "technologies"
	// LocationsTable is the table that holds the locations relation/edge. The primary key declared below.
	LocationsTable = "location_vacancies"
	// LocationsInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationsInverseTable = "locations"
	// AreasTable is the table that holds the areas relation/edge. The primary key declared below.
	AreasTable = "area_vacancies"
	// AreasInverseTable is the table name for the Area entity.
	// It exists in this package in order to avoid circular dependency with the "area" package.
	AreasInverseTable = "areas"
	// CompaniesTable is the table that holds the companies relation/edge. The primary key declared below.
	CompaniesTable = "vacancy_companies"
	// CompaniesInverseTable is the table name for the Company entity.
	// It exists in this package in order to avoid circular dependency with the "company" package.
	CompaniesInverseTable = "companies"
	// TechnologyLevelsTable is the table that holds the technology_levels relation/edge.
	TechnologyLevelsTable = "technology_levels"
	// TechnologyLevelsInverseTable is the table name for the TechnologyLevel entity.
	// It exists in this package in order to avoid circular dependency with the "technologylevel" package.
	TechnologyLevelsInverseTable = "technology_levels"
	// TechnologyLevelsColumn is the table column denoting the technology_levels relation/edge.
	TechnologyLevelsColumn = "vacancy_id"
)

// Columns holds all SQL columns for vacancy fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldIsNegotiate,
	FieldMinSalary,
	FieldMaxSalary,
	FieldIsRemote,
	FieldViews,
	FieldCreatedAt,
}

var (
	// TechnologiesPrimaryKey and TechnologiesColumn2 are the table columns denoting the
	// primary key for the technologies relation (M2M).
	TechnologiesPrimaryKey = []string{"technology_id", "vacancy_id"}
	// LocationsPrimaryKey and LocationsColumn2 are the table columns denoting the
	// primary key for the locations relation (M2M).
	LocationsPrimaryKey = []string{"location_id", "vacancy_id"}
	// AreasPrimaryKey and AreasColumn2 are the table columns denoting the
	// primary key for the areas relation (M2M).
	AreasPrimaryKey = []string{"area_id", "vacancy_id"}
	// CompaniesPrimaryKey and CompaniesColumn2 are the table columns denoting the
	// primary key for the companies relation (M2M).
	CompaniesPrimaryKey = []string{"vacancy_id", "company_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// MinSalaryValidator is a validator for the "min_salary" field. It is called by the builders before save.
	MinSalaryValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
