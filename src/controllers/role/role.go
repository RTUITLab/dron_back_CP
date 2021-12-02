package role

import (
	"context"

	"github.com/0B1t322/CP-Rosseti-Back/ent"
)

// This is internal controller to create role on startup in db
type RoleController struct {
	Client *ent.Client
}

func New(client *ent.Client) *RoleController {
	return &RoleController{
		Client: client,
	}
}

// Create roles student teacher and admin on startup
// if they exist do nothing
func (r RoleController) CreateRoles(ctx context.Context) error {
	roles := []string{"admin", "student", "teacher"}

	bulk := make([]*ent.RoleCreate, len(roles))
	for i, role := range roles {
		bulk[i] = r.Client.Role.Create().SetRole(role)
	}

	_, err := r.Client.Role.CreateBulk(bulk...).Save(ctx)
	if ent.IsConstraintError(err) {
		// Pass
	} else if err != nil {
		return err
	}
	
	return nil
}