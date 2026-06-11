package game_tag

type ReplaceGameTagsRequest struct {
	TagIDs []uint64 `json:"tag_ids" binding:"required"`
}
