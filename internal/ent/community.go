// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/logn-soft/logn-back/internal/ent/community"
)

// Community is the model entity for the Community schema.
type Community struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// About holds the value of the "about" field.
	About string `json:"about,omitempty"`
	// Members holds the value of the "members" field.
	Members int `json:"members,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommunityQuery when eager-loading is set.
	Edges CommunityEdges `json:"edges"`
}

// CommunityEdges holds the relations/edges for other nodes in the graph.
type CommunityEdges struct {
	// Socials holds the value of the socials edge.
	Socials []*Social `json:"socials,omitempty"`
	// Companies holds the value of the companies edge.
	Companies []*Company `json:"companies,omitempty"`
	// Areas holds the value of the areas edge.
	Areas []*Area `json:"areas,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// SocialsOrErr returns the Socials value or an error if the edge
// was not loaded in eager-loading.
func (e CommunityEdges) SocialsOrErr() ([]*Social, error) {
	if e.loadedTypes[0] {
		return e.Socials, nil
	}
	return nil, &NotLoadedError{edge: "socials"}
}

// CompaniesOrErr returns the Companies value or an error if the edge
// was not loaded in eager-loading.
func (e CommunityEdges) CompaniesOrErr() ([]*Company, error) {
	if e.loadedTypes[1] {
		return e.Companies, nil
	}
	return nil, &NotLoadedError{edge: "companies"}
}

// AreasOrErr returns the Areas value or an error if the edge
// was not loaded in eager-loading.
func (e CommunityEdges) AreasOrErr() ([]*Area, error) {
	if e.loadedTypes[2] {
		return e.Areas, nil
	}
	return nil, &NotLoadedError{edge: "areas"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e CommunityEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[3] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Community) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case community.FieldID, community.FieldMembers:
			values[i] = new(sql.NullInt64)
		case community.FieldName, community.FieldAbout:
			values[i] = new(sql.NullString)
		case community.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Community", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Community fields.
func (c *Community) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case community.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case community.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case community.FieldAbout:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field about", values[i])
			} else if value.Valid {
				c.About = value.String
			}
		case community.FieldMembers:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field members", values[i])
			} else if value.Valid {
				c.Members = int(value.Int64)
			}
		case community.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QuerySocials queries the "socials" edge of the Community entity.
func (c *Community) QuerySocials() *SocialQuery {
	return NewCommunityClient(c.config).QuerySocials(c)
}

// QueryCompanies queries the "companies" edge of the Community entity.
func (c *Community) QueryCompanies() *CompanyQuery {
	return NewCommunityClient(c.config).QueryCompanies(c)
}

// QueryAreas queries the "areas" edge of the Community entity.
func (c *Community) QueryAreas() *AreaQuery {
	return NewCommunityClient(c.config).QueryAreas(c)
}

// QueryUsers queries the "users" edge of the Community entity.
func (c *Community) QueryUsers() *UserQuery {
	return NewCommunityClient(c.config).QueryUsers(c)
}

// Update returns a builder for updating this Community.
// Note that you need to call Community.Unwrap() before calling this method if this Community
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Community) Update() *CommunityUpdateOne {
	return NewCommunityClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Community entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Community) Unwrap() *Community {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Community is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Community) String() string {
	var builder strings.Builder
	builder.WriteString("Community(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("about=")
	builder.WriteString(c.About)
	builder.WriteString(", ")
	builder.WriteString("members=")
	builder.WriteString(fmt.Sprintf("%v", c.Members))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Communities is a parsable slice of Community.
type Communities []*Community

func (c Communities) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
