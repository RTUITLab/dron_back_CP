// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/0B1t322/CP-Rosseti-Back/ent"
)

// The AnswerFunc type is an adapter to allow the use of ordinary
// function as Answer mutator.
type AnswerFunc func(context.Context, *ent.AnswerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AnswerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AnswerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AnswerMutation", m)
	}
	return f(ctx, mv)
}

// The ModuleFunc type is an adapter to allow the use of ordinary
// function as Module mutator.
type ModuleFunc func(context.Context, *ent.ModuleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ModuleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ModuleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ModuleMutation", m)
	}
	return f(ctx, mv)
}

// The ModuleDependciesFunc type is an adapter to allow the use of ordinary
// function as ModuleDependcies mutator.
type ModuleDependciesFunc func(context.Context, *ent.ModuleDependciesMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ModuleDependciesFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ModuleDependciesMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ModuleDependciesMutation", m)
	}
	return f(ctx, mv)
}

// The ModuleTestFunc type is an adapter to allow the use of ordinary
// function as ModuleTest mutator.
type ModuleTestFunc func(context.Context, *ent.ModuleTestMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ModuleTestFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ModuleTestMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ModuleTestMutation", m)
	}
	return f(ctx, mv)
}

// The PractTestFunc type is an adapter to allow the use of ordinary
// function as PractTest mutator.
type PractTestFunc func(context.Context, *ent.PractTestMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PractTestFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PractTestMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PractTestMutation", m)
	}
	return f(ctx, mv)
}

// The QuestionFunc type is an adapter to allow the use of ordinary
// function as Question mutator.
type QuestionFunc func(context.Context, *ent.QuestionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f QuestionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.QuestionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.QuestionMutation", m)
	}
	return f(ctx, mv)
}

// The RoleFunc type is an adapter to allow the use of ordinary
// function as Role mutator.
type RoleFunc func(context.Context, *ent.RoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoleMutation", m)
	}
	return f(ctx, mv)
}

// The SubModuleFunc type is an adapter to allow the use of ordinary
// function as SubModule mutator.
type SubModuleFunc func(context.Context, *ent.SubModuleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SubModuleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SubModuleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SubModuleMutation", m)
	}
	return f(ctx, mv)
}

// The SubModuleTestFunc type is an adapter to allow the use of ordinary
// function as SubModuleTest mutator.
type SubModuleTestFunc func(context.Context, *ent.SubModuleTestMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SubModuleTestFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SubModuleTestMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SubModuleTestMutation", m)
	}
	return f(ctx, mv)
}

// The TaskFunc type is an adapter to allow the use of ordinary
// function as Task mutator.
type TaskFunc func(context.Context, *ent.TaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskMutation", m)
	}
	return f(ctx, mv)
}

// The TestFunc type is an adapter to allow the use of ordinary
// function as Test mutator.
type TestFunc func(context.Context, *ent.TestMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TestFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TestMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TestMutation", m)
	}
	return f(ctx, mv)
}

// The TheoreticalTestFunc type is an adapter to allow the use of ordinary
// function as TheoreticalTest mutator.
type TheoreticalTestFunc func(context.Context, *ent.TheoreticalTestMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TheoreticalTestFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TheoreticalTestMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TheoreticalTestMutation", m)
	}
	return f(ctx, mv)
}

// The TheoreticalTryFunc type is an adapter to allow the use of ordinary
// function as TheoreticalTry mutator.
type TheoreticalTryFunc func(context.Context, *ent.TheoreticalTryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TheoreticalTryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TheoreticalTryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TheoreticalTryMutation", m)
	}
	return f(ctx, mv)
}

// The TryAnswerFunc type is an adapter to allow the use of ordinary
// function as TryAnswer mutator.
type TryAnswerFunc func(context.Context, *ent.TryAnswerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TryAnswerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TryAnswerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TryAnswerMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
