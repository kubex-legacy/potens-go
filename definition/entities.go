package definition

import (
	"github.com/kubex/potens-go/i18n"
)

/**
Entity IDs will be globally found as vendor-id/app-id/entity-id
*/

// Entity Definition of a single FDL data type
type AppEntity struct {
	//ID app specific ID for this entity e.g. ticket
	ID string
	//Type 2 character entity type
	Type string
	//Name e.g. Ticket
	Name i18n.Translations
	//Plural e.g. Tickets
	Plural i18n.Translations
	//Description e.g. Support Ticket
	Description i18n.Translations
	//Edges - ourbound edges that will be created against this entity
	Edges []Edge
}

/**
Edge IDs will be globally found as vendor-id/app-id/entity-id/edge-id
*/

type Edge struct {
	ID            string                         //e.g. friend-of
	BiDirectional bool     `yaml:"bi_direction"` //Stored the reverse edge
	Meta          []string                       //e.g. information about this edge
	Label         string                         // %src% is a friend of %dst%
	Strict        bool                           //Edges should be locked into the restrictions only
	Restrictions  []string                       //Format can be [id] or [app-id/id] where app-id must be for the same vendor
}
