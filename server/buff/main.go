package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	s, err := http.Get("https://buff.163.com/api/market/goods/sell_order?game=csgo&goods_id=912923&page_num=1&sort_by=default&mode=&allow_tradable_cooldown=1&_=1669621788895")
	if err != nil {
		return
	}
	body := s.Body

	all, err := io.ReadAll(body)
	if err != nil {
		return
	}
	//:)
	fmt.Println(string(all))
	buffVo := new(buffVo)
	json.Unmarshal(all, &buffVo)
	fmt.Println(buffVo)
	// 把结构体建起来 我这里不能复制终端

}

type buffVo struct {
	Code string `json:"code"`
	Data struct {
		FopStr     string `json:"fop_str"`
		GoodsInfos struct {
			Field1 struct {
				Appid           int         `json:"appid"`
				Description     interface{} `json:"description"`
				Game            string      `json:"game"`
				GoodsId         int         `json:"goods_id"`
				IconUrl         string      `json:"icon_url"`
				ItemId          interface{} `json:"item_id"`
				MarketHashName  string      `json:"market_hash_name"`
				MarketMinPrice  string      `json:"market_min_price"`
				Name            string      `json:"name"`
				OriginalIconUrl string      `json:"original_icon_url"`
				ShortName       string      `json:"short_name"`
				SteamPrice      string      `json:"steam_price"`
				SteamPriceCny   string      `json:"steam_price_cny"`
				Tags            struct {
					Category struct {
						Category      string `json:"category"`
						Id            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"category"`
					CategoryGroup struct {
						Category      string `json:"category"`
						Id            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"category_group"`
					Itemset struct {
						Category      string `json:"category"`
						Id            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"itemset"`
					Quality struct {
						Category      string `json:"category"`
						Id            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"quality"`
					Rarity struct {
						Category      string `json:"category"`
						Id            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"rarity"`
					Tournament struct {
						Category      string `json:"category"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"tournament"`
					Tournamentteam struct {
						Category      string `json:"category"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"tournamentteam"`
					Type struct {
						Category      string `json:"category"`
						Id            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"type"`
					Weaponcase struct {
						Category      string `json:"category"`
						Id            int    `json:"id"`
						InternalName  string `json:"internal_name"`
						LocalizedName string `json:"localized_name"`
					} `json:"weaponcase"`
				} `json:"tags"`
			} `json:"912923"`
		} `json:"goods_infos"`
		HasMarketStores struct {
			U1091463492 bool `json:"U1091463492"`
		} `json:"has_market_stores"`
		Items []struct {
			AllowBargain bool `json:"allow_bargain"`
			Appid        int  `json:"appid"`
			AssetInfo    struct {
				ActionLink          string `json:"action_link"`
				Appid               int    `json:"appid"`
				Assetid             string `json:"assetid"`
				Classid             string `json:"classid"`
				Contextid           int    `json:"contextid"`
				GoodsId             int    `json:"goods_id"`
				HasTradableCooldown bool   `json:"has_tradable_cooldown"`
				Info                struct {
					Fraudwarnings   interface{}   `json:"fraudwarnings"`
					IconUrl         string        `json:"icon_url"`
					OriginalIconUrl string        `json:"original_icon_url"`
					Stickers        []interface{} `json:"stickers"`
					TournamentTags  []interface{} `json:"tournament_tags"`
				} `json:"info"`
				Instanceid           string      `json:"instanceid"`
				Paintwear            string      `json:"paintwear"`
				TradableCooldownText string      `json:"tradable_cooldown_text"`
				TradableUnfrozenTime interface{} `json:"tradable_unfrozen_time"`
			} `json:"asset_info"`
			BackgroundImageUrl    string      `json:"background_image_url"`
			Bookmarked            bool        `json:"bookmarked"`
			CanBargain            bool        `json:"can_bargain"`
			CanUseInspectTrnUrl   bool        `json:"can_use_inspect_trn_url"`
			CannotBargainReason   string      `json:"cannot_bargain_reason"`
			CreatedAt             int         `json:"created_at"`
			Description           string      `json:"description"`
			Featured              int         `json:"featured"`
			Fee                   string      `json:"fee"`
			Game                  string      `json:"game"`
			GoodsId               int         `json:"goods_id"`
			Id                    string      `json:"id"`
			ImgSrc                string      `json:"img_src"`
			Income                string      `json:"income"`
			LowestBargainPrice    string      `json:"lowest_bargain_price"`
			Mode                  int         `json:"mode"`
			Price                 string      `json:"price"`
			RecentAverageDuration float64     `json:"recent_average_duration"`
			RecentDeliverRate     float64     `json:"recent_deliver_rate"`
			State                 int         `json:"state"`
			SupportedPayMethods   []int       `json:"supported_pay_methods"`
			TradableCooldown      interface{} `json:"tradable_cooldown"`
			UpdatedAt             int         `json:"updated_at"`
			UserId                string      `json:"user_id"`
		} `json:"items"`
		PageNum            int `json:"page_num"`
		PageSize           int `json:"page_size"`
		PreviewScreenshots struct {
		} `json:"preview_screenshots"`
		ShowPayMethodIcon bool   `json:"show_pay_method_icon"`
		SortBy            string `json:"sort_by"`
		SrcUrlBackground  string `json:"src_url_background"`
		TotalCount        int    `json:"total_count"`
		TotalPage         int    `json:"total_page"`
		UserInfos         struct {
			U1091463492 struct {
				Avatar       string      `json:"avatar"`
				AvatarSafe   string      `json:"avatar_safe"`
				IsAutoAccept bool        `json:"is_auto_accept"`
				IsPremiumVip bool        `json:"is_premium_vip"`
				Nickname     string      `json:"nickname"`
				SellerLevel  int         `json:"seller_level"`
				ShopId       string      `json:"shop_id"`
				UserId       string      `json:"user_id"`
				VTypes       interface{} `json:"v_types"`
			} `json:"U1091463492"`
		} `json:"user_infos"`
	} `json:"data"`
	Msg interface{} `json:"msg"`
}
