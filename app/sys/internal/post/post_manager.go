package post

import (
	"encoding/json"

	"gopkg.in/goyy/goyy.v0/app/sys/internal/menu"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

func (me *Manager) Save(p xtype.Principal, e entity.Interface) error {
	tx, err := me.DB().Begin()
	if err != nil {
		return err
	}
	err = me.Manager.Save(p, e)
	if err != nil {
		tx.Rollback()
		return err
	}
	postId := e.Get("id").(string)
	_, err = MenuMgr.Exec(deletePostMenuByPostId, postId)
	if err != nil {
		tx.Rollback()
		return err
	}
	menuIds := e.Get("menu_ids").(string)
	for _, menuId := range strings.Split(menuIds, ",") {
		if strings.IsNotBlank(menuId) {
			pm := NewMenuEntity()
			pm.SetPostId(postId)
			pm.SetMenuId(menuId)
			err = MenuMgr.Save(p, pm)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func menus(id string) (string, error) {
	menuIds := NewMenuEntities(100)
	if strings.IsNotBlank(id) {
		if err := MenuMgr.SelectList(menuIds, listMenuIdByPostId, id); err != nil {
			return "", err
		}
	}
	menus := menu.NewEntities(300)
	sDeletionEQ, _ := domain.NewSift("sDeletionEQ", "0")
	sOrdinalOA, _ := domain.NewSift("sOrdinalOA", "10")
	if err := menu.Mgr.SelectListBySift(menus, sDeletionEQ, sOrdinalOA); err != nil {
		return "", err
	}
	trees := make([]xtype.Tree, menus.Len())
	for i, m := range menus.Values() {
		trees[i].Open = true
		trees[i].Id = m.Id()
		trees[i].Name = m.Name()
		trees[i].ParentId = m.ParentId()
		for _, p := range menuIds.Values() {
			if trees[i].Id == p.MenuId() {
				trees[i].Checked = true
			}
		}
	}
	if len(trees) > 0 {
		b, err := json.Marshal(trees)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	return "", nil
}
