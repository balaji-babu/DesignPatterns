package main

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

//Low Level Module
type Relationships struct {
	relations []Info
}

// func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
// 	result := make([]*Person, 0)

// 	for i, v := range rs.relations {
// 		if v.relationship == Parent &&
// 			v.from.name == name {
// 			result = append(result, rs.relations[i].to)
// 		}
// 	}

// 	return result
// }

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations,
		Info{parent, Parent, child})
	rs.relations = append(rs.relations,
		Info{child, Child, parent})
}

//High Level Module
type Research struct {
	// Break DIP
	relationships Relationships
}

func (r *Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" &&
			rel.relationship == Parent {
			fmt.Println("John has a child called", rel.to.name)
		}
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	// Research Module is using the internals of relationships(low level module)
	// If the lowlevel module decides to change its code in the future, it breaks the code.

	research := Research{relationships}
	research.Investigate()
}
