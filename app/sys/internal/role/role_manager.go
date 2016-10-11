package role

import (
	"encoding/json"

	"gopkg.in/goyy/goyy.v0/app/sys/internal/post"
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/data/domain"
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func (me *Manager) Save(c xhttp.Context, e entity.Interface) error {
	tx, err := me.DB().Begin()
	if err != nil {
		return err
	}
	err = me.Manager.Save(c, e)
	if err != nil {
		tx.Rollback()
		return err
	}
	roleId := e.Get("id").(string)
	_, err = PostMgr.Exec(deleteRolePostByRoleId, roleId)
	if err != nil {
		tx.Rollback()
		return err
	}
	postIds := e.Get("post_ids").(string)
	for _, postId := range strings.Split(postIds, ",") {
		if strings.IsNotBlank(postId) {
			rp := NewPostEntity()
			rp.SetRoleId(roleId)
			rp.SetPostId(postId)
			err = PostMgr.SaveAndTx(c, rp)
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

func posts(id string) (string, error) {
	postIds := NewPostEntities(100)
	if strings.IsNotBlank(id) {
		if err := PostMgr.SelectList(postIds, listPostIdByRoleId, id); err != nil {
			return "", err
		}
	}
	posts := post.NewEntities(300)
	sDeletionEQ, _ := domain.NewSift("sDeletionEQ", "0")
	sOrdinalOA, _ := domain.NewSift("sOrdinalOA", "10")
	if err := post.Mgr.SelectListBySift(posts, sDeletionEQ, sOrdinalOA); err != nil {
		return "", err
	}
	trees := make([]xtype.Tree, posts.Len())
	for i, m := range posts.Values() {
		trees[i].Open = true
		trees[i].Id = m.Id()
		trees[i].Name = m.Name()
		trees[i].ParentId = m.ParentId()
		for _, p := range postIds.Values() {
			if trees[i].Id == p.PostId() {
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
