package dto

type EditPostReq struct {
	Described string   `form:"described"`
	MediaDel  []string `form:"media_del"`
	ImageSort []string `form:"image_sort"`
}
