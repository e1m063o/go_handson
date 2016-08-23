package controllers

import "github.com/revel/revel"
import "memo/app/internal/interceptors"
import "memo/app/internal/services/memo"
import (
	"strconv"
)

type Memo struct {
	*revel.Controller
	interceptors.Interceptor
}

func (c Memo) Index() revel.Result {
	return c.Render()
}

func (c Memo) Input() revel.Result {
	return c.RenderTemplate("Memo/Input.html")
}

func (c Memo) Entry(title, body string) revel.Result {
	c.Validation.Required(title)
	c.Validation.MaxSize(title, 100)
	c.Validation.Required(body)
	c.Validation.MaxSize(body, 100)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/input")
	}
	id := memo.AddMemo(title, body)
	c.Session["id"] = strconv.Itoa(id)
	return c.Redirect("/finish")
}

func (c Memo) Finish() revel.Result {
	id := c.Session["id"]
	memo := memo.GetMemo(id)
	c.RenderArgs["memo"] = memo
	return c.RenderTemplate("Memo/Finish.html")
}

func (c Memo) List() revel.Result {
	memos := memo.GetAllMemo()
	c.RenderArgs["memos"] = memos
	return c.RenderTemplate("Memo/List.html")
}

func (c Memo) Delete(id int) revel.Result {
	memo.DeleteMemo(id)
	return c.Redirect("/list")
}

func (c Memo) Show(id string) revel.Result {
	memo := memo.GetMemo(id)
	c.RenderArgs["memo"] = memo
	return c.RenderTemplate("Memo/Show.html")
}

func (c Memo) Update(id, title, body string) revel.Result {
	c.Validation.Required(title)
	c.Validation.MaxSize(title, 100)
	c.Validation.Required(body)
	c.Validation.MaxSize(body, 100)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/show/" + id)
	}
	memo.UpdateMemo(id, title, body)
	c.Session["id"] = id
	return c.Redirect("/finish")
}
