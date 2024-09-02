// Code generated by ent, DO NOT EDIT.

package psstrategy

import (
	"mall/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldUpdateTime, v))
}

// Owner applies equality check predicate on the "owner" field. It's identical to OwnerEQ.
func Owner(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldOwner, v))
}

// ScriptContent applies equality check predicate on the "script_content" field. It's identical to ScriptContentEQ.
func ScriptContent(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldScriptContent, v))
}

// IsDelete applies equality check predicate on the "is_delete" field. It's identical to IsDeleteEQ.
func IsDelete(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldIsDelete, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLTE(FieldUpdateTime, v))
}

// OwnerEQ applies the EQ predicate on the "owner" field.
func OwnerEQ(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldOwner, v))
}

// OwnerNEQ applies the NEQ predicate on the "owner" field.
func OwnerNEQ(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNEQ(FieldOwner, v))
}

// OwnerIn applies the In predicate on the "owner" field.
func OwnerIn(vs ...int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldIn(FieldOwner, vs...))
}

// OwnerNotIn applies the NotIn predicate on the "owner" field.
func OwnerNotIn(vs ...int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNotIn(FieldOwner, vs...))
}

// OwnerGT applies the GT predicate on the "owner" field.
func OwnerGT(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGT(FieldOwner, v))
}

// OwnerGTE applies the GTE predicate on the "owner" field.
func OwnerGTE(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGTE(FieldOwner, v))
}

// OwnerLT applies the LT predicate on the "owner" field.
func OwnerLT(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLT(FieldOwner, v))
}

// OwnerLTE applies the LTE predicate on the "owner" field.
func OwnerLTE(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLTE(FieldOwner, v))
}

// ScriptContentEQ applies the EQ predicate on the "script_content" field.
func ScriptContentEQ(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldScriptContent, v))
}

// ScriptContentNEQ applies the NEQ predicate on the "script_content" field.
func ScriptContentNEQ(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNEQ(FieldScriptContent, v))
}

// ScriptContentIn applies the In predicate on the "script_content" field.
func ScriptContentIn(vs ...string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldIn(FieldScriptContent, vs...))
}

// ScriptContentNotIn applies the NotIn predicate on the "script_content" field.
func ScriptContentNotIn(vs ...string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNotIn(FieldScriptContent, vs...))
}

// ScriptContentGT applies the GT predicate on the "script_content" field.
func ScriptContentGT(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGT(FieldScriptContent, v))
}

// ScriptContentGTE applies the GTE predicate on the "script_content" field.
func ScriptContentGTE(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGTE(FieldScriptContent, v))
}

// ScriptContentLT applies the LT predicate on the "script_content" field.
func ScriptContentLT(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLT(FieldScriptContent, v))
}

// ScriptContentLTE applies the LTE predicate on the "script_content" field.
func ScriptContentLTE(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLTE(FieldScriptContent, v))
}

// ScriptContentContains applies the Contains predicate on the "script_content" field.
func ScriptContentContains(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldContains(FieldScriptContent, v))
}

// ScriptContentHasPrefix applies the HasPrefix predicate on the "script_content" field.
func ScriptContentHasPrefix(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldHasPrefix(FieldScriptContent, v))
}

// ScriptContentHasSuffix applies the HasSuffix predicate on the "script_content" field.
func ScriptContentHasSuffix(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldHasSuffix(FieldScriptContent, v))
}

// ScriptContentEqualFold applies the EqualFold predicate on the "script_content" field.
func ScriptContentEqualFold(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEqualFold(FieldScriptContent, v))
}

// ScriptContentContainsFold applies the ContainsFold predicate on the "script_content" field.
func ScriptContentContainsFold(v string) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldContainsFold(FieldScriptContent, v))
}

// IsDeleteEQ applies the EQ predicate on the "is_delete" field.
func IsDeleteEQ(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldEQ(FieldIsDelete, v))
}

// IsDeleteNEQ applies the NEQ predicate on the "is_delete" field.
func IsDeleteNEQ(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNEQ(FieldIsDelete, v))
}

// IsDeleteIn applies the In predicate on the "is_delete" field.
func IsDeleteIn(vs ...int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldIn(FieldIsDelete, vs...))
}

// IsDeleteNotIn applies the NotIn predicate on the "is_delete" field.
func IsDeleteNotIn(vs ...int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldNotIn(FieldIsDelete, vs...))
}

// IsDeleteGT applies the GT predicate on the "is_delete" field.
func IsDeleteGT(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGT(FieldIsDelete, v))
}

// IsDeleteGTE applies the GTE predicate on the "is_delete" field.
func IsDeleteGTE(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldGTE(FieldIsDelete, v))
}

// IsDeleteLT applies the LT predicate on the "is_delete" field.
func IsDeleteLT(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLT(FieldIsDelete, v))
}

// IsDeleteLTE applies the LTE predicate on the "is_delete" field.
func IsDeleteLTE(v int) predicate.PsStrategy {
	return predicate.PsStrategy(sql.FieldLTE(FieldIsDelete, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PsStrategy) predicate.PsStrategy {
	return predicate.PsStrategy(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PsStrategy) predicate.PsStrategy {
	return predicate.PsStrategy(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PsStrategy) predicate.PsStrategy {
	return predicate.PsStrategy(sql.NotPredicates(p))
}