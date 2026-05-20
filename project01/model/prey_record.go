package model

// Go does not have classes or "objects" as defined by languages like Java
// A struct is strictly a composite data type—a blueprint for a block of memory
// containing a collection of named fields.
type PreyRecord struct {
	Year                string
	Species             string
	CommonName          string
	StudySite           string
	AssociatedCommunity string
	Retinol             string
}
