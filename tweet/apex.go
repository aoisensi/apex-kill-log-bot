package main

type Kill struct {
	Format string
	Rate   int
}

// 1 一人目
// 2 二人目
var kills = []Kill{
	// Weapons
	{Format: ".%[1]s ハボック %[2]s", Rate: 1000},
	{Format: ".%[1]s フラットライン %[2]s", Rate: 1000},
	{Format: ".%[1]s ヘムロック %[2]s", Rate: 1000},
	{Format: ".%[1]s R-301 %[2]s", Rate: 1000},
	{Format: ".%[1]s オルタネーター %[2]s", Rate: 1000},
	{Format: ".%[1]s プラウラー %[2]s", Rate: 1000},
	{Format: ".%[1]s R-99 %[2]s", Rate: 1000},
	{Format: ".%[1]s ボルト %[2]s", Rate: 1000},
	{Format: ".%[1]s CAR %[2]s", Rate: 1000},
	{Format: ".%[1]s ディヴォーション %[2]s", Rate: 1000},
	{Format: ".%[1]s L-スター %[2]s", Rate: 1000},
	{Format: ".%[1]s スピットファイア %[2]s", Rate: 1000},
	{Format: ".%[1]s ランページ %[2]s", Rate: 1000},
	{Format: ".%[1]s G7スカウト %[2]s", Rate: 1000},
	{Format: ".%[1]s トリプルテイク %[2]s", Rate: 1000},
	{Format: ".%[1]s 30-30 %[2]s", Rate: 1000},
	{Format: ".%[1]s ボセック %[2]s", Rate: 1000},
	{Format: ".%[1]s チャージライフル %[2]s", Rate: 1000},
	{Format: ".%[1]s ロングボウ %[2]s", Rate: 1000},
	{Format: ".%[1]s クレーバー %[2]s", Rate: 300},
	{Format: ".%[1]s センチネル %[2]s", Rate: 1000},
	{Format: ".%[1]s EVA-8 %[2]s", Rate: 1000},
	{Format: ".%[1]s マスティフ %[2]s", Rate: 1000},
	{Format: ".%[1]s モザンビーク %[2]s", Rate: 1000},
	{Format: ".%[1]s ピースキーパー %[2]s", Rate: 1000},
	{Format: ".%[1]s RE-45 %[2]s", Rate: 200},
	{Format: ".%[1]s P2020 %[2]s", Rate: 100},
	{Format: ".%[1]s ウィングマン %[2]s", Rate: 500},
	// Grenade
	{Format: ".%[1]s グレネード %[2]s", Rate: 400},
	{Format: ".%[1]s アークスター %[2]s", Rate: 500},
	{Format: ".%[1]s テルミット %[2]s", Rate: 300},
	// Melee
	{Format: ".%[1]s 格闘 %[2]s", Rate: 100},
	{Format: ".%[1]s クナイ格闘 %[2]s", Rate: 10},
	{Format: ".%[1]s レイヴンズバイト格闘 %[2]s", Rate: 10},
	{Format: ".%[1]s ショックスティック格闘 %[2]s", Rate: 10},
	{Format: ".%[1]s ボクシンググローブ格闘 %[2]s", Rate: 10},
	{Format: ".%[1]s バタフライナイフ格闘 %[2]s", Rate: 10},
	{Format: ".%[1]s トゥーマッチウィット格闘 %[2]s", Rate: 10},
	{Format: ".%[1]s デスハンマー格闘 %[2]s", Rate: 10},
	{Format: ".%[1]s ウォークラブ格闘 %[2]s", Rate: 10},
	{Format: ".%[1]s コールドスチール格闘 %[2]s", Rate: 10},
	{Format: ".%[1]s デッドマンズカーブ格闘 %[2]s", Rate: 10},
	{Format: ".%[1]s プロブレムソルバー格闘 %[2]s", Rate: 10},
	{Format: ".%[1]s エネルギーリーダー格闘 %[2]s", Rate: 10},
	{Format: ".%[1]s ビゥオンブレード格闘 %[2]s", Rate: 10},
	// Skill
	{Format: ".%[1]s スモークランチャー %[2]s", Rate: 10},
	{Format: ".%[1]s ローリングサンダー %[2]s", Rate: 300},
	{Format: ".%[1]s コースティックガス %[2]s", Rate: 50},
	{Format: ".%[1]s ワットソンフェンス %[2]s", Rate: 10},
	{Format: ".%[1]s サイレンス %[2]s", Rate: 30},
	{Format: ".%[1]s モバイルミニガン「シーラ」 %[2]s", Rate: 100},
	{Format: ".%[1]s ナックルクラスター %[2]s", Rate: 50},
	{Format: ".%[1]s マザーロード %[2]s", Rate: 50},
	{Format: ".%[1]s アークスネア %[2]s", Rate: 10},
	{Format: ".%[1]s ライオットドリル %[2]s", Rate: 30},
	{Format: ".%[1]s レッカーボール %[2]s", Rate: 30},
	// Environment
	{Format: ".[落下] %[2]s", Rate: 10},
}

var MaxRate int

func init() {
	for _, kill := range kills {
		MaxRate += kill.Rate
	}
}
