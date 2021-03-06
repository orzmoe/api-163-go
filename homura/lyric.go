package homura

import (
	"encoding/json"

	"../madoka"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// LyricGroupInit 初始化搜索路由组
func LyricGroupInit(m *martini.ClassicMartini) {
	m.Group("/lyric", func(router martini.Router) {
		router.Get("/:id", func(p martini.Params, r render.Render) {
			id := p["id"]
			// 发送POST请求得到最后包含url的结果
			reqRs, reqErr := madoka.SongLyric(id)
			if reqErr != nil {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
				return
			}
			// 应该可以解析到第一层json
			var originParse map[string]interface{}
			if err := json.Unmarshal([]byte(reqRs), &originParse); err != nil || (int)(originParse["code"].(float64)) != 200 {
				r.JSON(200, map[string]interface{}{"state": false, "msg": "请求失败", "data": nil})
			} else {
				r.JSON(200, map[string]interface{}{"state": true, "msg": "success", "data": originParse})
			}
		})
	})
}
