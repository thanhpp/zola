package dto

type EditPostReq struct {
	Described string   `form:"described"`
	ImageDel  []string `form:"image_del"`
	ImageSort []string `form:"image_sort"`
}
